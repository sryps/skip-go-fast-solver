// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: order_settlements.sql

package db

import (
	"context"
	"database/sql"
)

const getAllOrderSettlementsWithSettlementStatus = `-- name: GetAllOrderSettlementsWithSettlementStatus :many
SELECT id, created_at, updated_at, source_chain_id, destination_chain_id, source_chain_gateway_contract_address, amount, order_id, initiate_settlement_tx, complete_settlement_tx, settlement_status, settlement_status_message FROM order_settlements WHERE settlement_status = ?
`

func (q *Queries) GetAllOrderSettlementsWithSettlementStatus(ctx context.Context, settlementStatus string) ([]OrderSettlement, error) {
	rows, err := q.db.QueryContext(ctx, getAllOrderSettlementsWithSettlementStatus, settlementStatus)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []OrderSettlement
	for rows.Next() {
		var i OrderSettlement
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.SourceChainID,
			&i.DestinationChainID,
			&i.SourceChainGatewayContractAddress,
			&i.Amount,
			&i.OrderID,
			&i.InitiateSettlementTx,
			&i.CompleteSettlementTx,
			&i.SettlementStatus,
			&i.SettlementStatusMessage,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertOrderSettlement = `-- name: InsertOrderSettlement :one
INSERT INTO order_settlements (
    source_chain_id,
    destination_chain_id,
    source_chain_gateway_contract_address,
    amount,
    order_id,
    settlement_status
) VALUES (?, ?, ?, ?, ?, ?) ON CONFLICT DO NOTHING RETURNING id, created_at, updated_at, source_chain_id, destination_chain_id, source_chain_gateway_contract_address, amount, order_id, initiate_settlement_tx, complete_settlement_tx, settlement_status, settlement_status_message
`

type InsertOrderSettlementParams struct {
	SourceChainID                     string
	DestinationChainID                string
	SourceChainGatewayContractAddress string
	Amount                            string
	OrderID                           string
	SettlementStatus                  string
}

func (q *Queries) InsertOrderSettlement(ctx context.Context, arg InsertOrderSettlementParams) (OrderSettlement, error) {
	row := q.db.QueryRowContext(ctx, insertOrderSettlement,
		arg.SourceChainID,
		arg.DestinationChainID,
		arg.SourceChainGatewayContractAddress,
		arg.Amount,
		arg.OrderID,
		arg.SettlementStatus,
	)
	var i OrderSettlement
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.SourceChainID,
		&i.DestinationChainID,
		&i.SourceChainGatewayContractAddress,
		&i.Amount,
		&i.OrderID,
		&i.InitiateSettlementTx,
		&i.CompleteSettlementTx,
		&i.SettlementStatus,
		&i.SettlementStatusMessage,
	)
	return i, err
}

const setCompleteSettlementTx = `-- name: SetCompleteSettlementTx :one
UPDATE order_settlements
SET updated_at=CURRENT_TIMESTAMP, complete_settlement_tx = ?
WHERE source_chain_id = ? AND order_id = ? AND source_chain_gateway_contract_address = ?
    RETURNING id, created_at, updated_at, source_chain_id, destination_chain_id, source_chain_gateway_contract_address, amount, order_id, initiate_settlement_tx, complete_settlement_tx, settlement_status, settlement_status_message
`

type SetCompleteSettlementTxParams struct {
	CompleteSettlementTx              sql.NullString
	SourceChainID                     string
	OrderID                           string
	SourceChainGatewayContractAddress string
}

func (q *Queries) SetCompleteSettlementTx(ctx context.Context, arg SetCompleteSettlementTxParams) (OrderSettlement, error) {
	row := q.db.QueryRowContext(ctx, setCompleteSettlementTx,
		arg.CompleteSettlementTx,
		arg.SourceChainID,
		arg.OrderID,
		arg.SourceChainGatewayContractAddress,
	)
	var i OrderSettlement
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.SourceChainID,
		&i.DestinationChainID,
		&i.SourceChainGatewayContractAddress,
		&i.Amount,
		&i.OrderID,
		&i.InitiateSettlementTx,
		&i.CompleteSettlementTx,
		&i.SettlementStatus,
		&i.SettlementStatusMessage,
	)
	return i, err
}

const setInitiateSettlementTx = `-- name: SetInitiateSettlementTx :one
UPDATE order_settlements
SET updated_at=CURRENT_TIMESTAMP, initiate_settlement_tx = ?
WHERE source_chain_id = ? AND order_id = ? AND source_chain_gateway_contract_address = ?
    RETURNING id, created_at, updated_at, source_chain_id, destination_chain_id, source_chain_gateway_contract_address, amount, order_id, initiate_settlement_tx, complete_settlement_tx, settlement_status, settlement_status_message
`

type SetInitiateSettlementTxParams struct {
	InitiateSettlementTx              sql.NullString
	SourceChainID                     string
	OrderID                           string
	SourceChainGatewayContractAddress string
}

func (q *Queries) SetInitiateSettlementTx(ctx context.Context, arg SetInitiateSettlementTxParams) (OrderSettlement, error) {
	row := q.db.QueryRowContext(ctx, setInitiateSettlementTx,
		arg.InitiateSettlementTx,
		arg.SourceChainID,
		arg.OrderID,
		arg.SourceChainGatewayContractAddress,
	)
	var i OrderSettlement
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.SourceChainID,
		&i.DestinationChainID,
		&i.SourceChainGatewayContractAddress,
		&i.Amount,
		&i.OrderID,
		&i.InitiateSettlementTx,
		&i.CompleteSettlementTx,
		&i.SettlementStatus,
		&i.SettlementStatusMessage,
	)
	return i, err
}

const setSettlementStatus = `-- name: SetSettlementStatus :one
UPDATE order_settlements
SET updated_at=CURRENT_TIMESTAMP, settlement_status = ?, settlement_status_message = ?
WHERE source_chain_id = ? AND order_id = ? AND source_chain_gateway_contract_address = ?
    RETURNING id, created_at, updated_at, source_chain_id, destination_chain_id, source_chain_gateway_contract_address, amount, order_id, initiate_settlement_tx, complete_settlement_tx, settlement_status, settlement_status_message
`

type SetSettlementStatusParams struct {
	SettlementStatus                  string
	SettlementStatusMessage           sql.NullString
	SourceChainID                     string
	OrderID                           string
	SourceChainGatewayContractAddress string
}

func (q *Queries) SetSettlementStatus(ctx context.Context, arg SetSettlementStatusParams) (OrderSettlement, error) {
	row := q.db.QueryRowContext(ctx, setSettlementStatus,
		arg.SettlementStatus,
		arg.SettlementStatusMessage,
		arg.SourceChainID,
		arg.OrderID,
		arg.SourceChainGatewayContractAddress,
	)
	var i OrderSettlement
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.SourceChainID,
		&i.DestinationChainID,
		&i.SourceChainGatewayContractAddress,
		&i.Amount,
		&i.OrderID,
		&i.InitiateSettlementTx,
		&i.CompleteSettlementTx,
		&i.SettlementStatus,
		&i.SettlementStatusMessage,
	)
	return i, err
}
