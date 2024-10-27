package signing

import (
	"context"
	"errors"

	"github.com/cosmos/cosmos-sdk/client"
	clienttx "github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
)

type LocalCosmosSigner struct {
	privateKey    *secp256k1.PrivKey
	signerAddress string
}

var _ Signer = (*LocalCosmosSigner)(nil)

func NewLocalCosmosSigner(privateKey *secp256k1.PrivKey, signerAddress string) *LocalCosmosSigner {
	return &LocalCosmosSigner{
		privateKey:    privateKey,
		signerAddress: signerAddress,
	}
}

func (s *LocalCosmosSigner) Sign(ctx context.Context, chainID string, tx Transaction) (Transaction, error) {
	cosmosTx, ok := tx.(*CosmosTransaction)
	if !ok {
		return nil, errors.New("unsupported transaction type")
	}

	signedTx, err := SignCosmosTx(
		ctx,
		s.privateKey,
		cosmosTx.TxConfig,
		chainID,
		cosmosTx.AccountNumber,
		cosmosTx.Sequence,
		cosmosTx.Tx,
		s.signerAddress,
	)
	if err != nil {
		return nil, err
	}
	cosmosTx.Tx = signedTx
	return cosmosTx, nil
}

func (s *LocalCosmosSigner) Address() []byte {
	return s.privateKey.PubKey().Address()
}

func SignCosmosTx(
	ctx context.Context,
	privateKey *secp256k1.PrivKey,
	txConfig client.TxConfig,
	chainID string,
	accountNumber uint64,
	sequence uint64,
	tx sdk.Tx,
	signerAddress string,
) (sdk.Tx, error) {
	txBuilder, err := txConfig.WrapTxBuilder(tx)
	if err != nil {
		return nil, err
	}

	{
		// Hack for Cosmos SDK v0.45.16 since using same version as Noble
		// https://github.com/cosmos/cosmos-sdk/blob/v0.45.16/client/tx/tx.go#L359
		if err := txBuilder.SetSignatures(signing.SignatureV2{
			PubKey: privateKey.PubKey(),
			Data: &signing.SingleSignatureData{
				SignMode:  signing.SignMode(txConfig.SignModeHandler().DefaultMode()),
				Signature: nil,
			},
			Sequence: sequence,
		}); err != nil {
			return nil, err
		}
	}
	//func SignWithPrivKey(
	//	ctx context.Context,
	//	signMode signing.SignMode, signerData authsigning.SignerData,
	//	txBuilder client.TxBuilder, priv cryptotypes.PrivKey, txConfig client.TxConfig,
	//	accSeq uint64,
	//) (signing.SignatureV2, error) {

	signature, err := clienttx.SignWithPrivKey(
		ctx,
		signing.SignMode(txConfig.SignModeHandler().DefaultMode()),
		authsigning.SignerData{
			Address:       signerAddress,
			ChainID:       chainID,
			AccountNumber: accountNumber,
			Sequence:      sequence,
			PubKey:        privateKey.PubKey(),
		},
		txBuilder,
		privateKey,
		txConfig,
		sequence,
	)
	if err != nil {
		return nil, err
	}
	err = txBuilder.SetSignatures(signature)
	if err != nil {
		return nil, err
	}

	return txBuilder.GetTx(), nil
}
