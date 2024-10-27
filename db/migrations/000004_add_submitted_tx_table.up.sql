CREATE TABLE IF NOT EXISTS submitted_txs (
     id INTEGER PRIMARY KEY AUTOINCREMENT,
     created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
     updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

     order_id INT,
     order_settlement_id INT,
     hyperlane_transfer_id INT,

     chain_id TEXT NOT NULL,
     tx_hash         TEXT NOT NULL,
     raw_tx          TEXT NOT NULL,
     tx_type TEXT NOT NULL,
     tx_status         TEXT NOT NULL,
     tx_status_message TEXT,

     UNIQUE(chain_id, tx_hash),
     FOREIGN KEY (order_id) REFERENCES orders(id),
     FOREIGN KEY (order_settlement_id) REFERENCES order_settlements(id),
     FOREIGN KEY (hyperlane_transfer_id) REFERENCES hyperlane_transfers(id)
);