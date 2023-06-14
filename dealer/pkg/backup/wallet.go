package backup

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"strings"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-cbor-util"
	"github.com/filecoin-project/go-state-types/builtin/v9/market"
	"github.com/myxtype/filecoin-client/local"
	"github.com/myxtype/filecoin-client/types"
)

const (
	clientPrivKey = "7b2254797065223a22736563703235366b31222c22507269766174654b6579223a223836493345562b652f7973466b6d61452f4d636346565a4930444e4b582b6449446246387475546938324d3d227d"
)

var (
	clientAddress = address.Address{}
)

func init() {
	clientAddress, _ = address.NewFromString("f1a6tcxpzz6cyp5geatbogmroy6cxcrstruhnp42y")
}

func (b *backup) signDealProposal(ctx context.Context, proposal *market.DealProposal) (*market.ClientDealProposal, error) {
	data, err := hex.DecodeString(strings.TrimSpace(clientPrivKey))
	if err != nil {
		return nil, err
	}

	ki := types.KeyInfo{}
	if err := json.Unmarshal(data, &ki); err != nil {
		return nil, err
	}

	serialized, err := cborutil.Dump(proposal)
	if err != nil {
		return nil, err
	}

	sig, err := local.WalletSign(ki.Type, ki.PrivateKey, serialized)
	if err != nil {
		return nil, err
	}

	return &market.ClientDealProposal{
		Proposal:        *proposal,
		ClientSignature: *sig,
	}, nil
}
