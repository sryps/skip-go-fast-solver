package svmrpc

import (
	"context"
	"github.com/skip-mev/go-fast-solver/shared/config"
	"sync"

	"github.com/gagliardetto/solana-go/rpc"
)

type SVMRPCClientManager interface {
	GetClient(ctx context.Context, chainID string) (SolanaRPCClient, error)
}

type svmRPCClientManager struct {
	clients map[string]SolanaRPCClient
	m       sync.Mutex
}

func NewSVMRPCClientManager() SVMRPCClientManager {
	return &svmRPCClientManager{
		clients: make(map[string]SolanaRPCClient),
	}
}

// GetClient implements SVMRPCClientManager.
func (s *svmRPCClientManager) GetClient(ctx context.Context, chainID string) (SolanaRPCClient, error) {
	s.m.Lock()
	defer s.m.Unlock()

	client, ok := s.clients[chainID]
	if ok {
		return client, nil
	}

	rpcAddr, err := config.GetConfigReader(ctx).GetRPCEndpoint(chainID)
	if err != nil {
		return nil, err
	}

	rpcConn := rpc.New(rpcAddr)

	client = NewSolanaRPCClient(rpcConn)

	s.clients[chainID] = client

	return client, nil
}
