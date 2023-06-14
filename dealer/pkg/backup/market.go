package backup

import (
	"context"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-fil-markets/storagemarket/network"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/go-state-types/builtin/v9/market"
	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p/core/peer"
	multiaddr "github.com/multiformats/go-multiaddr"
	"github.com/multiformats/go-multibase"
)

var (
	minerId = address.Address{}
)

func init() {
	minerId, _ = address.NewIDAddress(1970622)
}

func (b *backup) connectMiner(ctx context.Context) error {
	addr, err := multiaddr.NewMultiaddr("/ip4/152.32.173.11/tcp/23456/p2p/12D3KooWS26eBREdM959vDNJWyfgwsd38NMegn7KK11R9DY4EU4p")
	if err != nil {
		return err
	}

	peer, err := peer.AddrInfoFromP2pAddr(addr)
	if err != nil {
		return err
	}

	logger.Sugar().Infow(
		"Watch",
		"Connecting", peer,
	)
	if err := b.host.Connect(ctx, *peer); err != nil {
		return err
	}

	logger.Sugar().Infow(
		"Watch",
		"Connected", peer,
		"Create Deal Stream", peer,
	)
	b.stream, err = network.NewFromLibp2pHost(b.host,
		network.RetryParameters(time.Second, 5*time.Minute, 15, 5),
	).NewDealStream(ctx, peer.ID)
	if err != nil {
		return err
	}

	logger.Sugar().Infow(
		"Watch",
		"Created Deal Stream", peer,
	)

	return nil
}

func (b *backup) dealProposal(ctx context.Context, rootCid, pieceCid string, pieceSize uint64) (*market.DealProposal, error) {
	_size := abi.UnpaddedPieceSize(pieceSize)
	if err := _size.Validate(); err != nil {
		return nil, err
	}

	_rootCid, err := cid.Parse(rootCid)
	if err != nil {
		return nil, err
	}

	_pieceCid, err := cid.Parse(pieceCid)
	if err != nil {
		return nil, err
	}

	label, err := market.NewLabelFromString(_rootCid.Encode(multibase.MustNewEncoder('u')))
	if err != nil {
		return nil, err
	}

	return &market.DealProposal{
		PieceCID:             _pieceCid,
		PieceSize:            _size.Padded(),
		Client:               clientAddress,
		Provider:             minerId,
		Label:                label,
		StartEpoch:           0,    // TODO
		EndEpoch:             1700, // TODO
		StoragePricePerEpoch: big.Zero(),
		ProviderCollateral:   big.Zero(), // TODO
		ClientCollateral:     big.Zero(),
		VerifiedDeal:         true,
	}, nil
}
