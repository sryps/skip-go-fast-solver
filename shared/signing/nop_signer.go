package signing

import "golang.org/x/net/context"

type NopSigner struct{}

var _ Signer = (*NopSigner)(nil)

func NewNopSigner() *NopSigner {
	return &NopSigner{}
}

func (s *NopSigner) Sign(ctx context.Context, chainID string, tx Transaction) (Transaction, error) {
	return tx, nil
}

func (s *NopSigner) Address() []byte {
	return nil
}
