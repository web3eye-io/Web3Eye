package token

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	enttoken "github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/token"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"
	tokenproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/token"

	"github.com/google/uuid"
)

type Req struct {
	ID              *uint32
	EntID           *uuid.UUID
	ChainType       *basetype.ChainType
	ChainID         *string
	Contract        *string
	TokenType       *basetype.TokenType
	TokenID         *string
	Owner           *string
	URI             *string
	URIState        *basetype.TokenURIState
	URIType         *string
	ImageURL        *string
	VideoURL        *string
	Description     *string
	Name            *string
	VectorState     *tokenproto.ConvertState
	VectorID        *int64
	IPFSImageURL    *string
	ImageSnapshotID *uint32
	Remark          *string
}

//nolint:gocyclo
func CreateSet(c *ent.TokenCreate, req *Req) *ent.TokenCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.ChainType != nil {
		c.SetChainType(req.ChainType.String())
	}
	if req.ChainID != nil {
		c.SetChainID(*req.ChainID)
	}
	if req.Contract != nil {
		c.SetContract(*req.Contract)
	}
	if req.TokenType != nil {
		c.SetTokenType(req.TokenType.String())
	}
	if req.TokenID != nil {
		c.SetTokenID(*req.TokenID)
	}
	if req.Owner != nil {
		c.SetOwner(*req.Owner)
	}
	if req.URI != nil {
		c.SetURI(*req.URI)
	}
	if req.URIState != nil {
		c.SetURIState(req.URIState.String())
	}
	if req.URIType != nil {
		c.SetURIType(*req.URIType)
	}
	if req.ImageURL != nil {
		c.SetImageURL(*req.ImageURL)
	}
	if req.VideoURL != nil {
		c.SetVideoURL(*req.VideoURL)
	}
	if req.Description != nil {
		c.SetDescription(*req.Description)
	}
	if req.Name != nil {
		c.SetName(*req.Name)
	}
	if req.VectorState != nil {
		c.SetVectorState(req.VectorState.String())
	}
	if req.VectorID != nil {
		c.SetVectorID(*req.VectorID)
	}
	if req.IPFSImageURL != nil {
		c.SetIpfsImageURL(*req.IPFSImageURL)
	}
	if req.ImageSnapshotID != nil {
		c.SetImageSnapshotID(*req.ImageSnapshotID)
	}
	if req.Remark != nil {
		c.SetRemark(*req.Remark)
	}
	return c
}

//nolint:gocyclo
func UpdateSet(u *ent.TokenUpdateOne, req *Req) (*ent.TokenUpdateOne, error) {
	if req.ChainType != nil {
		u.SetChainType(req.ChainType.String())
	}
	if req.ChainID != nil {
		u.SetChainID(*req.ChainID)
	}
	if req.Contract != nil {
		u.SetContract(*req.Contract)
	}
	if req.TokenType != nil {
		u.SetTokenType(req.TokenType.String())
	}
	if req.TokenID != nil {
		u.SetTokenID(*req.TokenID)
	}
	if req.Owner != nil {
		u.SetOwner(*req.Owner)
	}
	if req.URI != nil {
		u.SetURI(*req.URI)
	}
	if req.URIState != nil {
		u.SetURIState(req.URIState.String())
	}
	if req.URIType != nil {
		u.SetURIType(*req.URIType)
	}
	if req.ImageURL != nil {
		u.SetImageURL(*req.ImageURL)
	}
	if req.VideoURL != nil {
		u.SetVideoURL(*req.VideoURL)
	}
	if req.Description != nil {
		u.SetDescription(*req.Description)
	}
	if req.Name != nil {
		u.SetName(*req.Name)
	}
	if req.VectorState != nil {
		u.SetVectorState(req.VectorState.String())
	}
	if req.VectorID != nil {
		u.SetVectorID(*req.VectorID)
	}
	if req.IPFSImageURL != nil {
		u.SetIpfsImageURL(*req.IPFSImageURL)
	}
	if req.ImageSnapshotID != nil {
		u.SetImageSnapshotID(*req.ImageSnapshotID)
	}
	if req.Remark != nil {
		u.SetRemark(*req.Remark)
	}
	return u, nil
}

type Conds struct {
	EntID           *cruder.Cond
	EntIDs          *cruder.Cond
	ChainType       *cruder.Cond
	ChainID         *cruder.Cond
	Contract        *cruder.Cond
	TokenType       *cruder.Cond
	TokenID         *cruder.Cond
	Owner           *cruder.Cond
	URI             *cruder.Cond
	URIState        *cruder.Cond
	URIType         *cruder.Cond
	ImageURL        *cruder.Cond
	VideoURL        *cruder.Cond
	Description     *cruder.Cond
	Name            *cruder.Cond
	VectorState     *cruder.Cond
	VectorID        *cruder.Cond
	IPFSImageURL    *cruder.Cond
	ImageSnapshotID *cruder.Cond
	Remark          *cruder.Cond
	VectorIDs       *cruder.Cond
}

func SetQueryConds(q *ent.TokenQuery, conds *Conds) (*ent.TokenQuery, error) { //nolint
	if conds.EntID != nil {
		entid, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(enttoken.EntID(entid))
		default:
			return nil, fmt.Errorf("invalid entid field")
		}
	}
	if conds.EntIDs != nil {
		entids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(enttoken.EntIDIn(entids...))
		default:
			return nil, fmt.Errorf("invalid entid field")
		}
	}
	if conds.VectorIDs != nil {
		vectorids, ok := conds.VectorIDs.Val.([]int64)
		if !ok {
			return nil, fmt.Errorf("invalid vectorids")
		}
		switch conds.VectorIDs.Op {
		case cruder.IN:
			q.Where(enttoken.VectorIDIn(vectorids...))
		default:
			return nil, fmt.Errorf("invalid vectorids field")
		}
	}
	if conds.ChainType != nil {
		chaintype, ok := conds.ChainType.Val.(basetype.ChainType)
		if !ok {
			return nil, fmt.Errorf("invalid chaintype")
		}
		switch conds.ChainType.Op {
		case cruder.EQ:
			q.Where(enttoken.ChainType(chaintype.String()))
		default:
			return nil, fmt.Errorf("invalid chaintype field")
		}
	}
	if conds.ChainID != nil {
		chainid, ok := conds.ChainID.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid chainid")
		}
		switch conds.ChainID.Op {
		case cruder.EQ:
			q.Where(enttoken.ChainID(chainid))
		default:
			return nil, fmt.Errorf("invalid chainid field")
		}
	}
	if conds.Contract != nil {
		contract, ok := conds.Contract.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid contract")
		}
		switch conds.Contract.Op {
		case cruder.EQ:
			q.Where(enttoken.Contract(contract))
		default:
			return nil, fmt.Errorf("invalid contract field")
		}
	}
	if conds.TokenType != nil {
		tokentype, ok := conds.TokenType.Val.(basetype.TokenType)
		if !ok {
			return nil, fmt.Errorf("invalid tokentype")
		}
		switch conds.TokenType.Op {
		case cruder.EQ:
			q.Where(enttoken.TokenType(tokentype.String()))
		default:
			return nil, fmt.Errorf("invalid tokentype field")
		}
	}
	if conds.TokenID != nil {
		tokenid, ok := conds.TokenID.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid tokenid")
		}
		switch conds.TokenID.Op {
		case cruder.EQ:
			q.Where(enttoken.TokenID(tokenid))
		default:
			return nil, fmt.Errorf("invalid tokenid field")
		}
	}
	if conds.Owner != nil {
		owner, ok := conds.Owner.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid owner")
		}
		switch conds.Owner.Op {
		case cruder.EQ:
			q.Where(enttoken.Owner(owner))
		default:
			return nil, fmt.Errorf("invalid owner field")
		}
	}
	if conds.URI != nil {
		uri, ok := conds.URI.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid uri")
		}
		switch conds.URI.Op {
		case cruder.EQ:
			q.Where(enttoken.URI(uri))
		default:
			return nil, fmt.Errorf("invalid uri field")
		}
	}
	if conds.URIState != nil {
		uristate, ok := conds.URIState.Val.(basetype.TokenURIState)
		if !ok {
			return nil, fmt.Errorf("invalid uristate")
		}
		switch conds.URI.Op {
		case cruder.EQ:
			q.Where(enttoken.URIState(uristate.String()))
		default:
			return nil, fmt.Errorf("invalid uristate field")
		}
	}
	if conds.URIType != nil {
		uritype, ok := conds.URIType.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid uritype")
		}
		switch conds.URIType.Op {
		case cruder.EQ:
			q.Where(enttoken.URIType(uritype))
		default:
			return nil, fmt.Errorf("invalid uritype field")
		}
	}
	if conds.ImageURL != nil {
		imageurl, ok := conds.ImageURL.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid imageurl")
		}
		switch conds.ImageURL.Op {
		case cruder.EQ:
			q.Where(enttoken.ImageURL(imageurl))
		default:
			return nil, fmt.Errorf("invalid imageurl field")
		}
	}
	if conds.VideoURL != nil {
		videourl, ok := conds.VideoURL.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid videourl")
		}
		switch conds.VideoURL.Op {
		case cruder.EQ:
			q.Where(enttoken.VideoURL(videourl))
		default:
			return nil, fmt.Errorf("invalid videourl field")
		}
	}
	if conds.Description != nil {
		description, ok := conds.Description.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid description")
		}
		switch conds.Description.Op {
		case cruder.EQ:
			q.Where(enttoken.Description(description))
		default:
			return nil, fmt.Errorf("invalid description field")
		}
	}
	if conds.Name != nil {
		name, ok := conds.Name.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid name")
		}
		switch conds.Name.Op {
		case cruder.EQ:
			q.Where(enttoken.Name(name))
		default:
			return nil, fmt.Errorf("invalid name field")
		}
	}
	if conds.VectorState != nil {
		vectorstate, ok := conds.VectorState.Val.(tokenproto.ConvertState)
		if !ok {
			return nil, fmt.Errorf("invalid vectorstate")
		}
		switch conds.VectorState.Op {
		case cruder.EQ:
			q.Where(enttoken.VectorState(vectorstate.String()))
		default:
			return nil, fmt.Errorf("invalid vectorstate field")
		}
	}
	if conds.VectorID != nil {
		vectorid, ok := conds.VectorID.Val.(int64)
		if !ok {
			return nil, fmt.Errorf("invalid vectorid")
		}
		switch conds.VectorID.Op {
		case cruder.EQ:
			q.Where(enttoken.VectorID(vectorid))
		default:
			return nil, fmt.Errorf("invalid vectorid field")
		}
	}
	if conds.IPFSImageURL != nil {
		ipfsimageurl, ok := conds.IPFSImageURL.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid ipfsimageurl")
		}
		switch conds.IPFSImageURL.Op {
		case cruder.EQ:
			q.Where(enttoken.IpfsImageURL(ipfsimageurl))
		default:
			return nil, fmt.Errorf("invalid ipfsimageurl field")
		}
	}
	if conds.ImageSnapshotID != nil {
		imagesnapshotid, ok := conds.ImageSnapshotID.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid imagesnapshotid")
		}
		switch conds.ImageSnapshotID.Op {
		case cruder.EQ:
			q.Where(enttoken.ImageSnapshotID(imagesnapshotid))
		default:
			return nil, fmt.Errorf("invalid imagesnapshotid field")
		}
	}
	if conds.Remark != nil {
		remark, ok := conds.Remark.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid remark")
		}
		switch conds.Remark.Op {
		case cruder.EQ:
			q.Where(enttoken.Remark(remark))
		default:
			return nil, fmt.Errorf("invalid remark field")
		}
	}
	return q, nil
}
