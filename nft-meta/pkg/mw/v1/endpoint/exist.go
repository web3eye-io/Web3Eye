package endpoint

import (
	"context"

	endpointcrud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/endpoint"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	endpointent "github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/endpoint"
)

func (h *Handler) ExistEndpoint(ctx context.Context) (bool, error) {
	exist := false
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			Endpoint.
			Query().
			Where(
				endpointent.EntID(*h.EntID),
				endpointent.DeletedAt(0),
			).
			Exist(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (h *Handler) ExistEndpointConds(ctx context.Context) (bool, error) {
	exist := false
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := endpointcrud.SetQueryConds(cli.Endpoint.Query(), h.Conds)
		if err != nil {
			return err
		}
		exist, err = stm.Exist(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}
