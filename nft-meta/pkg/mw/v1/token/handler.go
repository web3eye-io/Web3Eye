package token

import (
	"context"
	"fmt"

	constant "github.com/web3eye-io/Web3Eye/nft-meta/pkg/const"
	tokencrud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/token"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"
	tokenproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/token"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID              *uint32
	EntID           *uuid.UUID
	ChainType       *basetype.ChainType
	ChainID         *string
	Contract        *string
	TokenType       *basetype.TokenType
	TokenID         *string
	Owner           *string
	URI             *string
	URIType         *string
	ImageURL        *string
	VideoURL        *string
	Name            *string
	Description     *string
	VectorState     *tokenproto.ConvertState
	VectorID        *int64
	IPFSImageURL    *string
	ImageSnapshotID *uint32
	Remark          *string
	Reqs            []*tokencrud.Req
	Conds           *tokencrud.Conds
	Offset          int32
	Limit           int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, err
		}
	}
	return handler, nil
}

func WithID(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid id")
			}
			return nil
		}
		h.ID = u
		return nil
	}
}

func WithEntID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid entid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.EntID = &_id
		return nil
	}
}
func WithChainType(u *basetype.ChainType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid chaintype")
			}
			return nil
		}
		if _, ok := basetype.ChainType_name[int32(*u)]; !ok {
			return fmt.Errorf("invalid chaintype field")
		}
		h.ChainType = u
		return nil
	}
}
func WithChainID(u *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid chainid")
			}
			return nil
		}
		h.ChainID = u
		return nil
	}
}

func WithContract(u *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid contract")
			}
			return nil
		}
		h.Contract = u
		return nil
	}
}
func WithTokenType(u *basetype.TokenType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid tokentype")
			}
			return nil
		}
		if _, ok := basetype.TokenType_name[int32(*u)]; !ok {
			return fmt.Errorf("invalid tokentype field")
		}
		h.TokenType = u
		return nil
	}
}
func WithTokenID(u *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid tokenid")
			}
			return nil
		}
		h.TokenID = u
		return nil
	}
}
func WithOwner(u *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid owner")
			}
			return nil
		}
		h.Owner = u
		return nil
	}
}
func WithURI(u *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid uri")
			}
			return nil
		}
		h.URI = u
		return nil
	}
}
func WithURIType(u *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid uritype")
			}
			return nil
		}
		h.URIType = u
		return nil
	}
}
func WithImageURL(u *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid imageurl")
			}
			return nil
		}
		h.ImageURL = u
		return nil
	}
}
func WithVideoURL(u *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid videourl")
			}
			return nil
		}
		h.VideoURL = u
		return nil
	}
}
func WithName(u *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid name")
			}
			return nil
		}
		h.Name = u
		return nil
	}
}
func WithDescription(u *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid description")
			}
			return nil
		}
		h.Description = u
		return nil
	}
}
func WithVectorState(u *tokenproto.ConvertState, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid vectorstate")
			}
			return nil
		}
		if _, ok := tokenproto.ConvertState_name[int32(*u)]; !ok {
			return fmt.Errorf("invalid vectorstate field")
		}
		h.VectorState = u
		return nil
	}
}
func WithVectorID(u *int64, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid vectorid")
			}
			return nil
		}
		h.VectorID = u
		return nil
	}
}
func WithIPFSImageURL(u *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid ipfsimageurl")
			}
			return nil
		}
		h.IPFSImageURL = u
		return nil
	}
}
func WithImageSnapshotID(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid imagesnapshotid")
			}
			return nil
		}
		h.ImageSnapshotID = u
		return nil
	}
}

func WithRemark(u *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid remark")
			}
			return nil
		}
		h.Remark = u
		return nil
	}
}

// nolint:gocyclo
func WithReqs(reqs []*tokenproto.TokenReq, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_reqs := []*tokencrud.Req{}
		for _, req := range reqs {
			_req := &tokencrud.Req{}
			if req.EntID != nil {
				id, err := uuid.Parse(req.GetEntID())
				if err != nil {
					return err
				}
				_req.EntID = &id
			}
			if req.ChainType != nil {
				if _, ok := basetype.ChainType_name[int32(*req.ChainType)]; !ok {
					return fmt.Errorf("invalid parsestate field")
				}
				_req.ChainType = req.ChainType
			}
			if req.ChainID != nil {
				_req.ChainID = req.ChainID
			}
			if req.Contract != nil {
				_req.Contract = req.Contract
			}
			if req.TokenType != nil {
				if _, ok := basetype.TokenType_name[int32(*req.TokenType)]; !ok {
					return fmt.Errorf("invalid tokentype field")
				}
				_req.TokenType = req.TokenType
			}
			if req.TokenID != nil {
				_req.TokenID = req.TokenID
			}
			if req.Owner != nil {
				_req.Owner = req.Owner
			}
			if req.URI != nil {
				_req.URI = req.URI
			}
			if req.URIType != nil {
				_req.URIType = req.URIType
			}
			if req.ImageURL != nil {
				_req.ImageURL = req.ImageURL
			}
			if req.VideoURL != nil {
				_req.VideoURL = req.VideoURL
			}
			if req.Name != nil {
				_req.Name = req.Name
			}
			if req.Description != nil {
				_req.Description = req.Description
			}
			if req.VectorState != nil {
				if _, ok := tokenproto.ConvertState_name[int32(*req.VectorState)]; !ok {
					return fmt.Errorf("invalid vectorstate field")
				}
				_req.VectorState = req.VectorState
			}
			if req.VectorID != nil {
				_req.VectorID = req.VectorID
			}
			if req.IPFSImageURL != nil {
				_req.IPFSImageURL = req.IPFSImageURL
			}
			if req.ImageSnapshotID != nil {
				_req.ImageSnapshotID = req.ImageSnapshotID
			}
			if req.Remark != nil {
				_req.Remark = req.Remark
			}
			_reqs = append(_reqs, _req)
		}
		h.Reqs = _reqs
		return nil
	}
}

//nolint:gocyclo
func WithConds(conds *tokenproto.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &tokencrud.Conds{}
		if conds == nil {
			return nil
		}
		if conds.EntID != nil {
			id, err := uuid.Parse(conds.GetEntID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.EntID = &cruder.Cond{
				Op:  conds.GetEntID().GetOp(),
				Val: id,
			}
		}
		if conds.EntIDs != nil {
			ids := []uuid.UUID{}
			for _, id := range conds.GetEntIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.EntIDs = &cruder.Cond{
				Op:  conds.GetEntIDs().GetOp(),
				Val: ids,
			}
		}
		if conds.ChainType != nil {
			h.Conds.ChainType = &cruder.Cond{
				Op:  conds.GetChainType().GetOp(),
				Val: basetype.ChainType(conds.GetChainType().GetValue()),
			}
		}
		if conds.ChainID != nil {
			h.Conds.ChainID = &cruder.Cond{
				Op:  conds.GetChainID().GetOp(),
				Val: conds.GetChainID().GetValue(),
			}
		}
		if conds.Contract != nil {
			h.Conds.Contract = &cruder.Cond{
				Op:  conds.GetContract().GetOp(),
				Val: conds.GetContract().GetValue(),
			}
		}
		if conds.TokenType != nil {
			h.Conds.TokenType = &cruder.Cond{
				Op:  conds.GetTokenType().GetOp(),
				Val: conds.GetTokenType().GetValue(),
			}
		}
		if conds.TokenID != nil {
			h.Conds.TokenID = &cruder.Cond{
				Op:  conds.GetTokenID().GetOp(),
				Val: conds.GetTokenID().GetValue(),
			}
		}
		if conds.Owner != nil {
			h.Conds.Owner = &cruder.Cond{
				Op:  conds.GetOwner().GetOp(),
				Val: conds.GetOwner().GetValue(),
			}
		}
		if conds.URI != nil {
			h.Conds.URI = &cruder.Cond{
				Op:  conds.GetURI().GetOp(),
				Val: conds.GetURI().GetValue(),
			}
		}
		if conds.URIType != nil {
			h.Conds.URIType = &cruder.Cond{
				Op:  conds.GetURIType().GetOp(),
				Val: conds.GetURIType().GetValue(),
			}
		}
		if conds.ImageURL != nil {
			h.Conds.ImageURL = &cruder.Cond{
				Op:  conds.GetImageURL().GetOp(),
				Val: conds.GetImageURL().GetValue(),
			}
		}
		if conds.VideoURL != nil {
			h.Conds.VideoURL = &cruder.Cond{
				Op:  conds.GetVideoURL().GetOp(),
				Val: conds.GetVideoURL().GetValue(),
			}
		}
		if conds.Name != nil {
			h.Conds.Name = &cruder.Cond{
				Op:  conds.GetName().GetOp(),
				Val: conds.GetName().GetValue(),
			}
		}
		if conds.Description != nil {
			h.Conds.Description = &cruder.Cond{
				Op:  conds.GetDescription().GetOp(),
				Val: conds.GetDescription().GetValue(),
			}
		}
		if conds.VectorState != nil {
			h.Conds.VectorState = &cruder.Cond{
				Op:  conds.GetVectorState().GetOp(),
				Val: conds.GetVectorState().GetValue(),
			}
		}
		if conds.VectorID != nil {
			h.Conds.VectorID = &cruder.Cond{
				Op:  conds.GetVectorID().GetOp(),
				Val: conds.GetVectorID().GetValue(),
			}
		}
		if conds.IPFSImageURL != nil {
			h.Conds.IPFSImageURL = &cruder.Cond{
				Op:  conds.GetIPFSImageURL().GetOp(),
				Val: conds.GetIPFSImageURL().GetValue(),
			}
		}
		if conds.ImageSnapshotID != nil {
			h.Conds.ImageSnapshotID = &cruder.Cond{
				Op:  conds.GetImageSnapshotID().GetOp(),
				Val: conds.GetImageSnapshotID().GetValue(),
			}
		}
		if conds.Remark != nil {
			h.Conds.Remark = &cruder.Cond{
				Op:  conds.GetRemark().GetOp(),
				Val: conds.GetRemark().GetValue(),
			}
		}
		return nil
	}
}

func WithOffset(offset int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Offset = offset
		return nil
	}
}

func WithLimit(limit int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if limit == 0 {
			limit = constant.DefaultRowLimit
		}
		h.Limit = limit
		return nil
	}
}
