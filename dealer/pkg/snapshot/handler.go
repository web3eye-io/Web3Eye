package snapshot

import (
	"fmt"

	dealerpb "github.com/web3eye-io/Web3Eye/proto/web3eye/dealer/v1"
)

type Handler struct {
	CID         string
	SnapshotURI string
	Items       []*dealerpb.ContentItem
	Indexes     []uint64
}

func NewHandler(options ...func(*Handler) error) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		if err := opt(handler); err != nil {
			return nil, err
		}
	}
	return handler, nil
}

func WithCID(cid string) func(*Handler) error {
	return func(h *Handler) error {
		if cid == "" {
			return fmt.Errorf("invalid cid")
		}
		h.CID = cid
		return nil
	}
}

func WithSnapshotURI(uri string) func(*Handler) error {
	return func(h *Handler) error {
		if uri == "" {
			return fmt.Errorf("invalid uri")
		}
		h.SnapshotURI = uri
		return nil
	}
}

func WithItems(items []*dealerpb.ContentItem) func(*Handler) error {
	return func(h *Handler) error {
		if len(items) == 0 {
			return fmt.Errorf("invalid items")
		}
		for _, item := range items {
			if item.URI == "" {
				return fmt.Errorf("invalid uri")
			}
			if item.ChainType == "" {
				return fmt.Errorf("invalid chaintype")
			}
			if item.ChainID == "" {
				return fmt.Errorf("invalid chainid")
			}
			if item.Contract == "" {
				return fmt.Errorf("invalid contract")
			}
			if item.TokenID == "" {
				return fmt.Errorf("invalid uid")
			}
		}
		h.Items = items
		return nil
	}
}

func WithIndexes(indexes []uint64) func(*Handler) error {
	return func(h *Handler) error {
		h.Indexes = indexes
		return nil
	}
}
