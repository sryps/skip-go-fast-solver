package evm

import (
	"errors"
	"golang.org/x/net/context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/skip-mev/go-fast-solver/shared/signing"
)

func EthereumSignerToBindSignerFn(signer signing.Signer, chainID string) bind.SignerFn {
	return func(_ common.Address, tx *types.Transaction) (*types.Transaction, error) {
		signedTx, err := signer.Sign(context.Background(), chainID, tx)
		if err != nil {
			return nil, err
		}

		rawTx, ok := signedTx.(*types.Transaction)
		if !ok {
			return nil, errors.New("invalid transaction type")
		}

		return rawTx, nil
	}
}
