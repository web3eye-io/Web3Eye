package backup

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-cbor-util"
	"github.com/filecoin-project/go-fil-markets/storagemarket"
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

var miners = map[uint64]string{
	1970622: "/ip4/152.32.173.11/tcp/23456/p2p/12D3KooWS26eBREdM959vDNJWyfgwsd38NMegn7KK11R9DY4EU4p",
	1970630: "/ip4/123.58.203.78/tcp/23456/p2p/12D3KooWHGHJiH1YuvW9BonV8YZLDpnen3JR4zNQMcMJ3gRRptrq",
	7824:    "/ip4/172.19.16.118/tcp/23456/p2p/12D3KooWQeujGARoW6BsLWjML3KwAZEHr9n8fVKAk8yzGNw2FDdK",
	5316:    "/ip4/172.19.16.117/tcp/3456/p2p/12D3KooWMJxYN71gbbv3MSKnwatUGaUq9WoAFgTC7PUycmeJa9TC",  // Calibnet miner
	1002:    "/ip4/172.23.51.211/tcp/23456/p2p/12D3KooWM6Yz39SUfqKPNbd2hjowV2Jv8BGAwVDtf9JSB65G2gVF", // Testnet miner
}

const minerID = 1002

func init() {
	minerId, _ = address.NewIDAddress(minerID)
}

func (b *backup) connectMiner(ctx context.Context) error {
	addr, err := multiaddr.NewMultiaddr(miners[minerID])
	if err != nil {
		return err
	}

	peer, err := peer.AddrInfoFromP2pAddr(addr)
	if err != nil {
		return err
	}

	if err := b.host.Connect(ctx, *peer); err != nil {
		return err
	}

	b.peer = peer

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

	start := time.Unix(1686996002, 0)
	startEpoch := abi.ChainEpoch(1000 + (time.Since(start).Seconds() / 4))

	return &market.DealProposal{
		PieceCID:             _pieceCid,
		PieceSize:            _size.Padded(),
		Client:               clientAddress,
		Provider:             minerId,
		Label:                label,
		StartEpoch:           startEpoch,          // TODO
		EndEpoch:             startEpoch + 518400, // TODO
		StoragePricePerEpoch: big.NewInt(976562),
		ProviderCollateral:   big.Zero(), // TODO
		ClientCollateral:     big.Zero(),
		VerifiedDeal:         false,
	}, nil
}

func (b *backup) sendDealProposal(ctx context.Context, proposal *market.ClientDealProposal, root string) (*cid.Cid, error) {
	_root, err := cid.Parse(root)
	if err != nil {
		return nil, err
	}

	stream, err := network.NewFromLibp2pHost(b.host,
		network.RetryParameters(time.Second, 5*time.Minute, 15, 5),
	).NewDealStream(ctx, b.peer.ID)
	if err != nil {
		return nil, err
	}
	defer stream.Close()

	logger.Sugar().Infow(
		"sendDealProposal",
		"Proposal", proposal,
		"Root", root,
		"State", "Sending",
	)

	if err = stream.WriteDealProposal(network.Proposal{
		FastRetrieval: true,
		DealProposal:  proposal,
		Piece: &storagemarket.DataRef{
			TransferType: storagemarket.TTManual,
			Root:         _root,
			PieceCid:     &proposal.Proposal.PieceCID,
			PieceSize:    proposal.Proposal.PieceSize.Unpadded(),
		},
	}); err != nil {
		return nil, fmt.Errorf("sending deal proposal failed: %v", err)
	}

	logger.Sugar().Infow(
		"sendDealProposal",
		"Proposal", proposal,
		"Root", root,
		"State", "Sent",
	)

	resp, _, err := stream.ReadDealResponse()
	if err != nil {
		return nil, fmt.Errorf("reading proposal response failed: %v", err)
	}

	ipld, err := cborutil.AsIpld(proposal)
	if err != nil {
		return nil, fmt.Errorf("serializing proposal node failed: %v", err)
	}

	if !ipld.Cid().Equals(resp.Response.Proposal) {
		return nil, fmt.Errorf("provider returned proposal cid %s but we expected %s", resp.Response.Proposal, ipld.Cid())
	}

	if resp.Response.State != storagemarket.StorageDealWaitingForData {
		return nil, fmt.Errorf("provider returned unexpected state %d for proposal %s, with message: %s", resp.Response.State, resp.Response.Proposal, resp.Response.Message)
	}

	return &resp.Response.Proposal, nil
}

func (b *backup) getDealStatus(ctx context.Context, proposal string) (*storagemarket.ProviderDealState, error) {
	_proposal, err := cid.Parse(proposal)
	if err != nil {
		return nil, err
	}

	stream, err := network.NewFromLibp2pHost(b.host,
		network.RetryParameters(time.Second, 5*time.Minute, 15, 5),
	).NewDealStatusStream(ctx, b.peer.ID)
	if err != nil {
		return nil, err
	}
	defer stream.Close()

	sig, err := b.signCid(ctx, _proposal)
	if err != nil {
		return nil, err
	}

	if err := stream.WriteDealStatusRequest(network.DealStatusRequest{
		Proposal:  _proposal,
		Signature: *sig,
	}); err != nil {
		return nil, err
	}

	resp, _, err := stream.ReadDealStatusResponse()
	if err != nil {
		return nil, err
	}

	// TODO: here we need to verify signature

	return &resp.DealState, nil
}
