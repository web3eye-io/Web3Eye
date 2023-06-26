package token

import (
	"context"
	"errors"
	"sort"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	crud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/token"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/imageconvert"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/milvusdb"
	val "github.com/web3eye-io/Web3Eye/proto/web3eye"
	nftmetanpool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/token"
	rankernpool "github.com/web3eye-io/Web3Eye/proto/web3eye/ranker/v1/token"
	converter "github.com/web3eye-io/Web3Eye/ranker/pkg/converter/v1/token"
)

func (s *Server) Search(ctx context.Context, in *rankernpool.SearchTokenRequest) (*rankernpool.SearchTokenResponse, error) {
	milvusmgr := milvusdb.NewNFTConllectionMGR()
	_vector := imageconvert.ToArrayVector(in.Vector)
	// TODO: It should be given by request
	defaultTopN := int(in.Limit)
	_scores, err := milvusmgr.Search(
		context.Background(),
		[][milvusdb.VectorDim]float32{_vector},
		defaultTopN,
	)

	if err != nil {
		return nil, err
	}
	if len(_scores) == 0 {
		return nil, errors.New("have no result")
	}

	logger.Sugar().Infof("Search Result %v", _scores)

	scores := _scores[0]
	vIDs := make([]int64, len(scores))
	for i := range scores {
		vIDs = append(vIDs, i)
	}

	conds := &nftmetanpool.Conds{
		VectorIDs: &val.Int64SliceVal{
			Op:    "in",
			Value: vIDs,
		},
	}

	rows, _, err := crud.Rows(ctx, conds, 0, len(vIDs))
	if err != nil {
		return nil, err
	}

	infos := []*rankernpool.SearchToken{}
	for i := 0; i < len(rows); i++ {
		info := converter.Ent2Grpc(rows[i])
		info.Distance = scores[info.VectorID]
		info.SiblingTokens = make([]*rankernpool.SiblingToken, 0)
		infos = append(infos, info)
	}

	sort.Slice(infos, func(i, j int) bool {
		return infos[i].Distance < infos[j].Distance
	})

	contractRecord := make(map[string]int)
	result := []*rankernpool.SearchToken{}
	for _, v := range infos {
		if _, ok := contractRecord[v.ID]; ok {
			result[contractRecord[v.ID]].SiblingTokens = append(
				result[contractRecord[v.ID]].SiblingTokens, &rankernpool.SiblingToken{
					TokenID:      v.TokenID,
					ImageURL:     v.ImageURL,
					IPFSImageURL: v.IPFSImageURL,
				})
		} else {
			result = append(result, v)
			contractRecord[v.ID] = len(result) - 1
		}
	}

	return &rankernpool.SearchTokenResponse{Infos: result, Total: int32(len(result)), Vector: in.Vector}, nil
}
