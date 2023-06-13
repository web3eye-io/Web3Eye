package backup

import (
	"context"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/filecoin-project/go-fil-markets/storagemarket/network"
	"github.com/libp2p/go-libp2p/core/peer"
	multiaddr "github.com/multiformats/go-multiaddr"
)

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
