package endpoints

import (
	"errors"
	"math/rand"
	"strings"
)

var (
	ErrEndpointExhausted = errors.New("all endpoints is peeked")
	ErrEndpointsEmpty    = errors.New("endpoints empty")
)

const (
	AddrSplitter = ","
	AddrMinLen   = 3
)

type Manager struct {
	len        int
	localAddrs []string
}

func NewManager(envStr string) (*Manager, error) {
	envStr = strings.Trim(envStr, " ")
	_walletAddrs := strings.Split(envStr, AddrSplitter)

	walletAddrs := []string{}

	for i := range _walletAddrs {
		if len(_walletAddrs[i]) > 0 {
			walletAddrs = append(walletAddrs, _walletAddrs[i])
		}
	}

	if len(walletAddrs) == 0 {
		return nil, ErrEndpointsEmpty
	}

	// TODO:probability is not equal, should use Fisher-Yates algorithm
	if len(walletAddrs) > 1 {
		rand.Shuffle(len(walletAddrs), func(i, j int) {
			walletAddrs[i], walletAddrs[j] = walletAddrs[j], walletAddrs[i]
		})
	}

	// random start
	return &Manager{
		len:        len(walletAddrs),
		localAddrs: walletAddrs,
	}, nil
}

func (m *Manager) Peek() (addr string, err error) {
	ll := len(m.localAddrs)
	if ll > 0 {
		addr = m.localAddrs[ll-1]
		m.localAddrs = m.localAddrs[0 : ll-1]
		return addr, nil
	}
	return "", ErrEndpointExhausted
}

func (m *Manager) Len() int {
	return m.len
}
