package token

import (
	"context"
	"errors"

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

	scores := _scores[0]
	vIDs := []int64{}
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
		infos = append(infos, info)
	}

	return &rankernpool.SearchTokenResponse{Infos: infos, Total: int32(len(infos))}, nil
}
