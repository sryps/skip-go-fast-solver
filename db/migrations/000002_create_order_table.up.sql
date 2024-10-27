CREATE TABLE IF NOT EXISTS orders (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    source_chain_id      TEXT NOT NULL,
    destination_chain_id TEXT NOT NULL,
    source_chain_gateway_contract_address TEXT NOT NULL,
    sender BLOB NOT NULL,
    recipient BLOB NOT NULL,
    amount_in TEXT NOT NULL,
    amount_out TEXT NOT NULL,
    nonce              BIGINT NOT NULL,
    order_id TEXT NOT NULL,
    timeout_timestamp TIMESTAMP NOT NULL,
    order_creation_tx    TEXT NOT NULL,
    order_creation_tx_block_height BIGINT NOT NULL,
    data TEXT,

    filler TEXT,
    fill_tx   TEXT,
    refund_tx TEXT,
    order_status         TEXT NOT NULL,
    order_status_message TEXT,

    UNIQUE(source_chain_id, source_chain_gateway_contract_address, order_id)
);
