package token

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"time"

	"github.com/google/uuid"
	"github.com/web3eye-io/Web3Eye/common/ctredis"
	"github.com/web3eye-io/Web3Eye/common/utils"
	crud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/token"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/imageconvert"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/milvusdb"
	val "github.com/web3eye-io/Web3Eye/proto/web3eye"
	nftmetanpool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/token"
	rankernpool "github.com/web3eye-io/Web3Eye/proto/web3eye/ranker/v1/token"
	converter "github.com/web3eye-io/Web3Eye/ranker/pkg/converter/v1/token"
)

const (
	TopN           = 100
	PageLimit      = 10
	ShowSiblinsNum = 10
	StorageExpr    = time.Minute * 5
)

type SearchTokenBone struct {
	ID          string
	SiblingIDs  []string
	SiblingsNum uint32
	Distance    float32
}

type PageBone struct {
	TokenBones  []*SearchTokenBone
	Page        uint32
	TotalPages  uint32
	TotalTokens uint32
	PageLimit   uint32
}

func (s *Server) Search(ctx context.Context, in *rankernpool.SearchTokenRequest) (*rankernpool.SearchResponse, error) {
	// search from milvus
	scores, err := SerachFromMilvus(ctx, in.Vector)
	if err != nil {
		return nil, err
	}

	infos, err := QueryAndCollectTokens(ctx, scores)
	if err != nil {
		return nil, err
	}

	totalPages := uint32(len(infos) / PageLimit)
	if len(infos)%PageLimit > 0 {
		totalPages += 1
	}

	storageKey := uuid.NewString()

	totalTokens := uint32(len(infos))
	for i := uint32(0); i < totalPages; i++ {
		start := i * PageLimit
		end := (i + 1) * PageLimit
		if end > totalTokens {
			end = totalTokens - 1
		}
		pBone := &PageBone{
			TokenBones:  ToTokenBones(infos[start:end]),
			Page:        uint32(i + 1),
			TotalPages:  totalPages,
			TotalTokens: totalTokens,
			PageLimit:   PageLimit,
		}

		err = ctredis.Set(fmt.Sprintf("SearchToken:%v:%v", storageKey, pBone.Page), pBone, StorageExpr)
		if err != nil {
			return nil, err
		}
	}

	return &rankernpool.SearchResponse{
		Infos:       infos[:PageLimit],
		StorageKey:  storageKey,
		Page:        1,
		TotalPages:  totalPages,
		TotalTokens: totalTokens,
		PageLimit:   PageLimit,
	}, nil
}

func (s *Server) SearchPage(ctx context.Context, in *rankernpool.SearchPageRequest) (*rankernpool.SearchResponse, error) {
	pBone := &PageBone{}
	err := ctredis.Get(fmt.Sprintf("SearchToken:%v:%v", in.StorageKey, in.Page), pBone)
	if err != nil {
		return nil, err
	}

	tokens, err := ToSearchTokens(ctx, pBone.TokenBones)
	if err != nil {
		return nil, err
	}

	return &rankernpool.SearchResponse{
		Infos:       tokens,
		StorageKey:  in.StorageKey,
		Page:        pBone.Page,
		TotalPages:  pBone.TotalPages,
		TotalTokens: pBone.TotalTokens,
		PageLimit:   pBone.PageLimit,
	}, nil
}

func (pt *PageBone) MarshalBinary() (data []byte, err error) {
	data, err = json.Marshal(pt)
	fmt.Println(utils.PrettyStruct(pt))
	return data, err
	// return json.Marshal(pt)
}

func (pt *PageBone) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, pt)
}

func SerachFromMilvus(ctx context.Context, vector []float32) (map[int64]float32, error) {
	// search from milvus
	milvusmgr := milvusdb.NewNFTConllectionMGR()
	_vector := imageconvert.ToArrayVector(vector)

	_scores, err := milvusmgr.Search(ctx, [][milvusdb.VectorDim]float32{_vector}, TopN)
	if err != nil {
		return nil, err
	}

	if len(_scores) == 0 || len(_scores[0]) == 0 {
		return nil, errors.New("have no result")
	}

	return _scores[0], nil
}

func QueryAndCollectTokens(ctx context.Context, scores map[int64]float32) ([]*rankernpool.SearchToken, error) {
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

	// query from db
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

	// collection
	contractRecord := make(map[string]int)
	result := []*rankernpool.SearchToken{}
	for _, v := range infos {
		if _, ok := contractRecord[v.Contract]; ok {
			if len(result[contractRecord[v.Contract]].SiblingTokens) >= ShowSiblinsNum {
				continue
			}
			result[contractRecord[v.Contract]].SiblingTokens = append(
				result[contractRecord[v.Contract]].SiblingTokens, &rankernpool.SiblingToken{
					ID:           v.ID,
					TokenID:      v.TokenID,
					ImageURL:     v.ImageURL,
					IPFSImageURL: v.IPFSImageURL,
				})
		} else {
			result = append(result, v)
			contractRecord[v.Contract] = len(result) - 1
		}
	}

	// count transfers and
	for _, v := range result {
		conds = &nftmetanpool.Conds{
			ChainType: &val.StringVal{
				Op:    "eq",
				Value: v.ChainType,
			},
			ChainID: &val.StringVal{
				Op:    "eq",
				Value: v.ChainID,
			},
			Contract: &val.StringVal{
				Op:    "eq",
				Value: v.Contract,
			},
		}

		limit := ShowSiblinsNum - len(v.SiblingTokens)
		if limit <= 0 {
			continue
		}

		tokens, num, err := crud.Rows(ctx, conds, 0, limit)
		if err != nil {
			return nil, err
		}
		v.SiblingsNum = uint32(num)

		for _, token := range tokens {
			v.SiblingTokens = append(v.SiblingTokens, &rankernpool.SiblingToken{
				ID:           token.ID.String(),
				TokenID:      token.TokenID,
				ImageURL:     token.ImageURL,
				IPFSImageURL: token.IpfsImageURL,
			})
		}

		v.SiblingTokens = SliceDeduplicate(v.SiblingTokens)
	}
	return result, nil
}

func SliceDeduplicate(s []*rankernpool.SiblingToken) []*rankernpool.SiblingToken {
	mapRecord := make(map[string]struct{})
	listRecord := []int{}
	for i, v := range s {
		if _, ok := mapRecord[v.TokenID]; ok {
			listRecord = append(listRecord, i)
		} else {
			mapRecord[v.TokenID] = struct{}{}
		}
	}
	recordLen := len(listRecord) - 1
	for i := range listRecord {
		s = append(s[0:listRecord[recordLen-i]], s[listRecord[recordLen-i]+1:]...)
	}
	return s
}

func ToTokenBones(infos []*rankernpool.SearchToken) []*SearchTokenBone {
	bones := make([]*SearchTokenBone, len(infos))
	for i, v := range infos {
		bones[i] = &SearchTokenBone{
			ID:          v.ID,
			SiblingsNum: v.SiblingsNum,
			Distance:    v.Distance,
		}
		siblingIDs := make([]string, len(v.SiblingTokens))
		for i, token := range v.SiblingTokens {
			siblingIDs[i] = token.ID
		}
		bones[i].SiblingIDs = siblingIDs
	}
	return bones
}

func ToSearchTokens(ctx context.Context, bones []*SearchTokenBone) ([]*rankernpool.SearchToken, error) {
	tokens := make([]*rankernpool.SearchToken, len(bones))
	for i, v := range bones {
		IDs := []string{v.ID}
		IDs = append(IDs, v.SiblingIDs...)
		conds := &nftmetanpool.Conds{
			IDs: &val.StringSliceVal{
				Op:    "in",
				Value: IDs,
			},
		}
		// query from db
		rows, _, err := crud.Rows(ctx, conds, 0, len(IDs))
		if err != nil {
			return nil, err
		}

		tokens[i] = converter.Ent2Grpc(rows[0])
		tokens[i].Distance = v.Distance
		tokens[i].SiblingsNum = v.SiblingsNum
		tokens[i].SiblingTokens = make([]*rankernpool.SiblingToken, len(rows)-1)

		for j := 1; j < len(rows); j++ {
			tokens[i].SiblingTokens[j-1] = &rankernpool.SiblingToken{
				ID:           rows[j].ID.String(),
				TokenID:      rows[j].TokenID,
				ImageURL:     rows[j].ImageURL,
				IPFSImageURL: rows[j].IpfsImageURL,
			}
		}
	}

	return tokens, nil
}
