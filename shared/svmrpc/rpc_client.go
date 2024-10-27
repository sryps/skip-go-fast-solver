package svmrpc

import (
	"context"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

type SolanaRPCClient interface {
	GetTransaction(ctx context.Context, txSig string, opts *rpc.GetTransactionOpts) (out *rpc.GetTransactionResult, err error)
	GetTransactionInstructions(ctx context.Context, txSig string, opts *rpc.GetTransactionOpts) (instructions []solana.CompiledInstruction, err error)
	GetParsedTransaction(ctx context.Context, txSig string, opts *rpc.GetParsedTransactionOpts) (*rpc.GetParsedTransactionResult, error)
	GetAccountDataInto(ctx context.Context, account solana.PublicKey, inVar interface{}) (err error)
	GetAccountDataBorshInto(ctx context.Context, account solana.PublicKey, inVar interface{}) (err error)
	GetSignaturesForAddress(ctx context.Context, address string) ([]*rpc.TransactionSignature, error)
	SendEncodedTransaction(ctx context.Context, encodedTx string) (signature solana.Signature, err error)
}

type solanaRPCClient struct {
	client *rpc.Client
}

func NewSolanaRPCClient(client *rpc.Client) SolanaRPCClient {
	return &solanaRPCClient{client: client}
}

func (client *solanaRPCClient) GetTransaction(ctx context.Context, txSig string, opts *rpc.GetTransactionOpts) (out *rpc.GetTransactionResult, err error) {
	sig, err := solana.SignatureFromBase58(txSig)
	if err != nil {
		return nil, err
	}

	return client.client.GetTransaction(ctx, sig, opts)
}

func (client *solanaRPCClient) GetTransactionInstructions(ctx context.Context, txSig string, opts *rpc.GetTransactionOpts) (instructions []solana.CompiledInstruction, err error) {
	sig, err := solana.SignatureFromBase58(txSig)
	if err != nil {
		return nil, err
	}
	res, err := client.client.GetTransaction(ctx, sig, opts)
	if err != nil {
		return nil, err
	}
	tx, err := res.Transaction.GetTransaction()
	if err != nil {
		return nil, err
	}
	return tx.Message.Instructions, nil
}

func (client *solanaRPCClient) GetParsedTransaction(ctx context.Context, txSig string, opts *rpc.GetParsedTransactionOpts) (*rpc.GetParsedTransactionResult, error) {
	sig, err := solana.SignatureFromBase58(txSig)
	if err != nil {
		return nil, err
	}

	return client.client.GetParsedTransaction(ctx, sig, opts)
}

func (client *solanaRPCClient) GetAccountDataInto(ctx context.Context, account solana.PublicKey, inVar interface{}) (err error) {
	return client.client.GetAccountDataInto(ctx, account, inVar)
}

func (client *solanaRPCClient) GetAccountDataBorshInto(ctx context.Context, account solana.PublicKey, inVar interface{}) (err error) {
	return client.client.GetAccountDataBorshInto(ctx, account, inVar)
}

func (client *solanaRPCClient) GetSignaturesForAddress(ctx context.Context, address string) ([]*rpc.TransactionSignature, error) {
	return client.client.GetSignaturesForAddress(ctx, solana.MustPublicKeyFromBase58(address))
}

func (client *solanaRPCClient) SendEncodedTransaction(
	ctx context.Context,
	encodedTx string,
) (signature solana.Signature, err error) {
	return client.client.SendEncodedTransaction(ctx, encodedTx)
}
