package contract

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	entcontract "github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/contract"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"

	"github.com/google/uuid"
)

type Req struct {
	ID          *uint32
	EntID       *uuid.UUID
	ChainType   *basetype.ChainType
	ChainID     *string
	Address     *string
	Name        *string
	Symbol      *string
	Decimals    *uint32
	Creator     *string
	BlockNum    *uint64
	TxHash      *string
	TxTime      *uint32
	ProfileURL  *string
	BaseURL     *string
	BannerURL   *string
	Description *string
	Remark      *string
}

func CreateSet(c *ent.ContractCreate, req *Req) *ent.ContractCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.ChainType != nil {
		c.SetChainType((*req.ChainType).String())
	}
	if req.ChainID != nil {
		c.SetChainID(*req.ChainID)
	}
	if req.Address != nil {
		c.SetAddress(*req.Address)
	}
	if req.Name != nil {
		c.SetName(*req.Name)
	}
	if req.Symbol != nil {
		c.SetSymbol(*req.Symbol)
	}
	if req.Decimals != nil {
		c.SetDecimals(*req.Decimals)
	}
	if req.Creator != nil {
		c.SetCreator(*req.Creator)
	}
	if req.BlockNum != nil {
		c.SetBlockNum(*req.BlockNum)
	}
	if req.TxHash != nil {
		c.SetTxHash(*req.TxHash)
	}
	if req.TxTime != nil {
		c.SetTxTime(*req.TxTime)
	}
	if req.ProfileURL != nil {
		c.SetProfileURL(*req.ProfileURL)
	}
	if req.BaseURL != nil {
		c.SetBaseURL(*req.BaseURL)
	}
	if req.BannerURL != nil {
		c.SetBannerURL(*req.BannerURL)
	}
	if req.Description != nil {
		c.SetDescription(*req.Description)
	}
	if req.Remark != nil {
		c.SetRemark(*req.Remark)
	}
	return c
}

func UpdateSet(u *ent.ContractUpdateOne, req *Req) (*ent.ContractUpdateOne, error) {
	if req.ChainType != nil {
		u.SetChainType(req.ChainType.String())
	}
	if req.Address != nil {
		u.SetAddress(*req.Address)
	}
	if req.Name != nil {
		u.SetName(*req.Name)
	}
	if req.Symbol != nil {
		u.SetSymbol(*req.Symbol)
	}
	if req.Decimals != nil {
		u.SetDecimals(*req.Decimals)
	}
	if req.Creator != nil {
		u.SetCreator(*req.Creator)
	}
	if req.BlockNum != nil {
		u.SetBlockNum(*req.BlockNum)
	}
	if req.TxHash != nil {
		u.SetTxHash(*req.TxHash)
	}
	if req.TxTime != nil {
		u.SetTxTime(*req.TxTime)
	}
	if req.ProfileURL != nil {
		u.SetProfileURL(*req.ProfileURL)
	}
	if req.BaseURL != nil {
		u.SetBaseURL(*req.BaseURL)
	}
	if req.BannerURL != nil {
		u.SetBannerURL(*req.BannerURL)
	}
	if req.Description != nil {
		u.SetDescription(*req.Description)
	}
	if req.Remark != nil {
		u.SetRemark(*req.Remark)
	}
	return u, nil
}

type Conds struct {
	EntID       *cruder.Cond
	EntIDs      *cruder.Cond
	ChainType   *cruder.Cond
	ChainID     *cruder.Cond
	Address     *cruder.Cond
	Name        *cruder.Cond
	Symbol      *cruder.Cond
	Decimals    *cruder.Cond
	Creator     *cruder.Cond
	BlockNum    *cruder.Cond
	TxHash      *cruder.Cond
	TxTime      *cruder.Cond
	ProfileURL  *cruder.Cond
	BaseURL     *cruder.Cond
	BannerURL   *cruder.Cond
	Description *cruder.Cond
	Remark      *cruder.Cond
}

func SetQueryConds(q *ent.ContractQuery, conds *Conds) (*ent.ContractQuery, error) { //nolint
	if conds.EntID != nil {
		entid, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entcontract.EntID(entid))
		default:
			return nil, fmt.Errorf("invalid entid field")
		}
	}
	if conds.EntIDs != nil {
		entids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entcontract.EntIDIn(entids...))
		default:
			return nil, fmt.Errorf("invalid entid field")
		}
	}
	if conds.ChainType != nil {
		chaintype, ok := conds.ChainType.Val.(basetype.ChainType)
		if !ok {
			return nil, fmt.Errorf("invalid chaintype")
		}
		switch conds.ChainType.Op {
		case cruder.EQ:
			q.Where(entcontract.ChainType(chaintype.String()))
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
			q.Where(entcontract.ChainID(chainid))
		default:
			return nil, fmt.Errorf("invalid chainid field")
		}
	}
	if conds.Address != nil {
		address, ok := conds.Address.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid address")
		}
		switch conds.Address.Op {
		case cruder.EQ:
			q.Where(entcontract.Address(address))
		default:
			return nil, fmt.Errorf("invalid address field")
		}
	}
	if conds.Name != nil {
		name, ok := conds.Name.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid name")
		}
		switch conds.Name.Op {
		case cruder.EQ:
			q.Where(entcontract.Name(name))
		default:
			return nil, fmt.Errorf("invalid name field")
		}
	}
	if conds.Symbol != nil {
		symbol, ok := conds.Symbol.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid symbol")
		}
		switch conds.Symbol.Op {
		case cruder.EQ:
			q.Where(entcontract.Symbol(symbol))
		default:
			return nil, fmt.Errorf("invalid symbol field")
		}
	}
	if conds.Decimals != nil {
		decimals, ok := conds.Decimals.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid decimals")
		}
		switch conds.Decimals.Op {
		case cruder.EQ:
			q.Where(entcontract.Decimals(decimals))
		default:
			return nil, fmt.Errorf("invalid decimals field")
		}
	}
	if conds.Creator != nil {
		creator, ok := conds.Creator.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid creator")
		}
		switch conds.Creator.Op {
		case cruder.EQ:
			q.Where(entcontract.Creator(creator))
		default:
			return nil, fmt.Errorf("invalid creator field")
		}
	}
	if conds.BlockNum != nil {
		blocknum, ok := conds.BlockNum.Val.(uint64)
		if !ok {
			return nil, fmt.Errorf("invalid blocknum")
		}
		switch conds.BlockNum.Op {
		case cruder.EQ:
			q.Where(entcontract.BlockNum(blocknum))
		default:
			return nil, fmt.Errorf("invalid blocknum field")
		}
	}
	if conds.TxHash != nil {
		txhash, ok := conds.TxHash.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid txhash")
		}
		switch conds.TxHash.Op {
		case cruder.EQ:
			q.Where(entcontract.TxHash(txhash))
		default:
			return nil, fmt.Errorf("invalid txhash field")
		}
	}
	if conds.TxTime != nil {
		txtime, ok := conds.TxTime.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid txtime")
		}
		switch conds.TxTime.Op {
		case cruder.EQ:
			q.Where(entcontract.TxTime(txtime))
		default:
			return nil, fmt.Errorf("invalid txtime field")
		}
	}
	if conds.ProfileURL != nil {
		profileurl, ok := conds.ProfileURL.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid profileurl")
		}
		switch conds.ProfileURL.Op {
		case cruder.EQ:
			q.Where(entcontract.ProfileURL(profileurl))
		default:
			return nil, fmt.Errorf("invalid profileurl field")
		}
	}
	if conds.BaseURL != nil {
		baseurl, ok := conds.BaseURL.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid baseurl")
		}
		switch conds.BaseURL.Op {
		case cruder.EQ:
			q.Where(entcontract.BaseURL(baseurl))
		default:
			return nil, fmt.Errorf("invalid baseurl field")
		}
	}
	if conds.BannerURL != nil {
		bannerurl, ok := conds.BannerURL.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid bannerurl")
		}
		switch conds.BannerURL.Op {
		case cruder.EQ:
			q.Where(entcontract.BannerURL(bannerurl))
		default:
			return nil, fmt.Errorf("invalid bannerurl field")
		}
	}
	if conds.Description != nil {
		description, ok := conds.Description.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid description")
		}
		switch conds.Description.Op {
		case cruder.EQ:
			q.Where(entcontract.Description(description))
		default:
			return nil, fmt.Errorf("invalid description field")
		}
	}
	if conds.Remark != nil {
		remark, ok := conds.Remark.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid remark")
		}
		switch conds.Remark.Op {
		case cruder.EQ:
			q.Where(entcontract.Remark(remark))
		default:
			return nil, fmt.Errorf("invalid remark field")
		}
	}
	return q, nil
}
