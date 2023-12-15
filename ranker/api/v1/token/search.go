package token

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/google/uuid"
	"github.com/web3eye-io/Web3Eye/common/ctredis"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/imageconvert"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/milvusdb"
	tokenhandler "github.com/web3eye-io/Web3Eye/nft-meta/pkg/mw/v1/token"
	transferhandler "github.com/web3eye-io/Web3Eye/nft-meta/pkg/mw/v1/transfer"
	val "github.com/web3eye-io/Web3Eye/proto/web3eye"
	nftmetanpool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/token"
	transfernpool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/transfer"
	rankernpool "github.com/web3eye-io/Web3Eye/proto/web3eye/ranker/v1/token"
	converter "github.com/web3eye-io/Web3Eye/ranker/pkg/converter/v1/token"
)

const (
	TopN               = 100
	MaxSearchN         = 10000
	ShowSiblinsNum int = 10
	StorageExpr        = time.Minute * 5
)

type SearchTokenBone struct {
	EntID         string
	SiblingEntIDs []string
	SiblingsNum   uint32
	Distance      float32
	TranserferNum int32
}

type PageBone struct {
	TokenBones  []*SearchTokenBone
	Page        uint32
	TotalPages  uint32
	TotalTokens uint32
	Limit       uint32
}

type ScoreItem struct {
	VID   int64
	Score float32
}

func (s *Server) Search(ctx context.Context, in *rankernpool.SearchTokenRequest) (*rankernpool.SearchResponse, error) {
	logger.Sugar().Info("start search")
	start := time.Now()

	// search from milvus
	scores, err := SerachFromMilvus(ctx, in.Vector, MaxSearchN)
	if err != nil {
		logger.Sugar().Errorf("search from milvus failed, %v", err)
		return nil, err
	}
	logger.Sugar().Infof("scores: %v", len(scores))

	infos, err := QueryAndCollectTokens(ctx, scores, TopN)
	if err != nil {
		logger.Sugar().Errorf("query and collect tokens failed, %v", err)
		return nil, err
	}
	logger.Sugar().Infof("infos: %v", len(infos))

	totalPages := uint32(len(infos) / int(in.Limit))
	if len(infos)%int(in.Limit) > 0 {
		totalPages += 1
	}

	storageKey := uuid.NewString()

	totalTokens := uint32(len(infos))
	for i := uint32(0); i < totalPages; i++ {
		start := i * in.Limit
		end := (i + 1) * in.Limit
		if end > totalTokens {
			end = totalTokens
		}
		pBone := &PageBone{
			TokenBones:  ToTokenBones(infos[start:end]),
			Page:        i + 1,
			TotalPages:  totalPages,
			TotalTokens: totalTokens,
			Limit:       in.Limit,
		}

		err = ctredis.Set(fmt.Sprintf("SearchToken:%v:%v", storageKey, pBone.Page), pBone, StorageExpr)
		if err != nil {
			logger.Sugar().Errorf("put the search pageBone failed, %v", err)
			return nil, err
		}
	}

	limit := int(in.Limit)
	if len(infos) < limit {
		limit = len(infos)
	}

	logger.Sugar().Infof("take %v ms to finish search", time.Since(start).Milliseconds())
	return &rankernpool.SearchResponse{
		Infos:       infos[:limit],
		StorageKey:  storageKey,
		Page:        1,
		TotalPages:  totalPages,
		TotalTokens: totalTokens,
		Limit:       in.Limit,
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
		Limit:       pBone.Limit,
	}, nil
}

func (pt *PageBone) MarshalBinary() (data []byte, err error) {
	data, err = json.Marshal(pt)
	return data, err
}

func (pt *PageBone) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, pt)
}

func SerachFromMilvus(ctx context.Context, vector []float32, topN int) (map[int64]float32, error) {
	// search from milvus
	milvusmgr := milvusdb.NewNFTConllectionMGR()
	_vector := imageconvert.ToArrayVector(vector)
	_scores, err := milvusmgr.Search(ctx, [][milvusdb.VectorDim]float32{_vector}, topN)
	if err != nil {
		return nil, err
	}

	if len(_scores) == 0 || len(_scores[0]) == 0 {
		return nil, errors.New("have no result")
	}

	return _scores[0], nil
}

func sortSroces(scores map[int64]float32) []*ScoreItem {
	topScores := make([]*ScoreItem, len(scores))
	index := int64(0)
	for k, v := range scores {
		topScores[index] = &ScoreItem{
			VID:   k,
			Score: v,
		}
		index++
	}

	sort.Slice(topScores, func(i, j int) bool {
		return topScores[i].Score < topScores[j].Score
	})
	return topScores
}

// TODO:too long,will be rewrite
//
//nolint:all
func QueryAndCollectTokens(ctx context.Context, scores map[int64]float32, topN int) ([]*rankernpool.SearchToken, error) {
	topScores := sortSroces(scores)
	result := []*rankernpool.SearchToken{}
	contractRecord := make(map[string]int)

	start, end := 0, 0
	for len(result) < topN {
		start = end
		end = start + topN
		if start >= len(topScores) {
			break
		} else if end > len(topScores) {
			end = len(topScores)
		}

		vIDs := []int64{}
		for _, v := range topScores[start:end] {
			vIDs = append(vIDs, v.VID)
		}

		conds := &nftmetanpool.Conds{
			VectorIDs: &val.Int64SliceVal{
				Op:    "in",
				Value: vIDs,
			},
		}

		h, err := tokenhandler.NewHandler(ctx,
			tokenhandler.WithConds(conds),
			tokenhandler.WithOffset(0),
			tokenhandler.WithLimit(int32(len(vIDs))),
		)
		if err != nil {
			return nil, err
		}
		rows, _, err := h.GetTokens(ctx)
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
		for _, v := range infos {
			if _, ok := contractRecord[v.Contract]; ok {
				if len(result[contractRecord[v.Contract]].SiblingTokens) >= ShowSiblinsNum {
					continue
				}
				result[contractRecord[v.Contract]].SiblingTokens = append(
					result[contractRecord[v.Contract]].SiblingTokens, &rankernpool.SiblingToken{
						EntID:        v.EntID,
						TokenID:      v.TokenID,
						ImageURL:     v.ImageURL,
						IPFSImageURL: v.IPFSImageURL,
					})
			} else {
				result = append(result, v)
				contractRecord[v.Contract] = len(result) - 1
			}
		}
	}

	// full the siblinsTokens
	for _, v := range result {
		conds := &nftmetanpool.Conds{
			ChainType: &val.Uint32Val{Op: "eq", Value: uint32(v.ChainType)},
			ChainID:   &val.StringVal{Op: "eq", Value: v.ChainID},
			Contract:  &val.StringVal{Op: "eq", Value: v.Contract},
		}

		h, err := tokenhandler.NewHandler(
			ctx,
			tokenhandler.WithConds(conds),
			tokenhandler.WithOffset(0),
			tokenhandler.WithLimit(int32(ShowSiblinsNum)),
		)
		if err != nil {
			return nil, err
		}
		tokens, num, err := h.GetTokens(ctx)
		if err != nil {
			return nil, err
		}

		// query ShowSiblinsNum+1 records,because likely to query the v-self
		if err != nil {
			return nil, err
		}
		v.SiblingsNum = uint32(num)

		for _, token := range tokens {
			if v.GetID() == token.ID {
				continue
			}
			v.SiblingTokens = append(v.SiblingTokens, &rankernpool.SiblingToken{
				EntID:        token.EntID,
				TokenID:      token.TokenID,
				ImageURL:     token.ImageURL,
				IPFSImageURL: token.IPFSImageURL,
			})
		}

		v.SiblingTokens = SliceDeduplicate(v.SiblingTokens)

		if ShowSiblinsNum < len(v.SiblingTokens) {
			v.SiblingTokens = v.SiblingTokens[:ShowSiblinsNum]
		}
	}

	// get token transfersNum
	for _, v := range result {
		conds := &transfernpool.Conds{
			Contract: &val.StringVal{Op: "eq", Value: v.GetContract()},
			TokenID:  &val.StringVal{Op: "eq", Value: v.TokenID},
		}
		h, err := transferhandler.NewHandler(ctx, transferhandler.WithConds(conds))
		if err != nil {
			logger.Sugar().Infow("QueryAndCollectTokens", "error", err)
			continue
		}
		_, num, err := h.GetTransfers(ctx)
		if err == nil {
			v.TransfersNum = int32(num)
		} else {
			logger.Sugar().Infow("QueryAndCollectTokens", "error", err)
		}
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
			EntID:         v.EntID,
			SiblingsNum:   v.SiblingsNum,
			Distance:      v.Distance,
			TranserferNum: v.TransfersNum,
		}
		siblingIDs := make([]string, len(v.SiblingTokens))
		for i, token := range v.SiblingTokens {
			siblingIDs[i] = token.EntID
		}
		bones[i].SiblingEntIDs = siblingIDs
	}
	return bones
}

func ToSearchTokens(ctx context.Context, bones []*SearchTokenBone) ([]*rankernpool.SearchToken, error) {
	tokens := make([]*rankernpool.SearchToken, len(bones))
	for i, v := range bones {
		EntIDs := []string{v.EntID}
		EntIDs = append(EntIDs, v.SiblingEntIDs...)
		conds := &nftmetanpool.Conds{
			EntIDs: &val.StringSliceVal{
				Op:    "in",
				Value: EntIDs,
			},
		}

		h, err := tokenhandler.NewHandler(
			ctx,
			tokenhandler.WithConds(conds),
			tokenhandler.WithOffset(0),
			tokenhandler.WithLimit(int32(len(EntIDs))),
		)
		if err != nil {
			return nil, err
		}

		// query from db
		rows, _, err := h.GetTokens(ctx)
		if err != nil {
			return nil, err
		}

		tokens[i] = converter.Ent2Grpc(rows[0])
		tokens[i].Distance = v.Distance
		tokens[i].SiblingsNum = v.SiblingsNum
		tokens[i].TransfersNum = v.TranserferNum
		tokens[i].SiblingTokens = make([]*rankernpool.SiblingToken, len(rows)-1)

		for j := 1; j < len(rows); j++ {
			tokens[i].SiblingTokens[j-1] = &rankernpool.SiblingToken{
				EntID:        rows[j].EntID,
				TokenID:      rows[j].TokenID,
				ImageURL:     rows[j].ImageURL,
				IPFSImageURL: rows[j].IPFSImageURL,
			}
		}
	}

	return tokens, nil
}
