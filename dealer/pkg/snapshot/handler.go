package snapshot

import (
	"fmt"

	dealerpb "github.com/web3eye-io/Web3Eye/proto/web3eye/dealer/v1"
)

type Handler struct {
	SnapshotURI  string
	Items        []*dealerpb.ContentItem
	Indexes      []uint64
	SnapshotType dealerpb.SnapshotType
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

func WithSnapshotType(_type dealerpb.SnapshotType) func(*Handler) error {
	return func(h *Handler) error {
		switch _type {
		case dealerpb.SnapshotType_SnapshotWait:
		case dealerpb.SnapshotType_SnapshotBackup:
		default:
			return fmt.Errorf("invalid snapshot type")
		}
		h.SnapshotType = _type
		return nil
	}
}
