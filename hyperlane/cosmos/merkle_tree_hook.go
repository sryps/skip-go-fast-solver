package cosmos

import (
	"context"
	"encoding/json"
	"fmt"

	cosmwasm "github.com/CosmWasm/wasmd/x/wasm/types"
)

type MerkleTreeHookQuerier struct {
	client  cosmwasm.QueryClient
	address string
}

func NewMerkleTreeHookQuerier(address string, client cosmwasm.QueryClient) *MerkleTreeHookQuerier {
	return &MerkleTreeHookQuerier{client, address}
}

type MerkleHookCountRequest struct {
	MerkleHook MerkleHookCount `json:"merkle_hook"`
}

type MerkleHookCount struct {
	Count struct{} `json:"count"`
}

func (mth *MerkleTreeHookQuerier) Count(ctx context.Context) (uint64, error) {
	req := MerkleHookCountRequest{
		MerkleHookCount{Count: struct{}{}},
	}
	data, err := json.Marshal(req)
	if err != nil {
		return 0, fmt.Errorf("marshaling get merkle hook count request: %w", err)
	}

	resp, err := mth.client.SmartContractState(ctx, &cosmwasm.QuerySmartContractStateRequest{
		Address:   mth.address,
		QueryData: data,
	})
	if err != nil {
		return 0, fmt.Errorf("querying smart contract %s for merkle hook count: %w", mth.address, err)
	}
	if resp.Data == nil {
		return 0, fmt.Errorf("got nil response when querying for merkle hook count")
	}

	type CountResponse struct {
		Count uint64 `json:"count"`
	}
	var countResponse CountResponse
	if err := json.Unmarshal(resp.Data, &countResponse); err != nil {
		return 0, fmt.Errorf("unmarshaling query bytes into formatted data: %w", err)
	}

	return countResponse.Count, nil
}
