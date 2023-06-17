package retrieve

import (
	"fmt"
)

type Handler struct {
	ChainType string
	ChainID   string
	Contract  string
	TokenID   string
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

func WithChainType(chainType string) func(*Handler) error {
	return func(h *Handler) error {
		if chainType == "" {
			return fmt.Errorf("invalid chaintype")
		}
		h.ChainType = chainType
		return nil
	}
}

func WithChainID(chainID string) func(*Handler) error {
	return func(h *Handler) error {
		if chainID == "" {
			return fmt.Errorf("invalid chainid")
		}
		h.ChainID = chainID
		return nil
	}
}

func WithContract(contract string) func(*Handler) error {
	return func(h *Handler) error {
		if contract == "" {
			return fmt.Errorf("invalid contract")
		}
		h.Contract = contract
		return nil
	}
}

func WithTokenID(tokenID string) func(*Handler) error {
	return func(h *Handler) error {
		if tokenID == "" {
			return fmt.Errorf("invalid tokenid")
		}
		h.TokenID = tokenID
		return nil
	}
}
