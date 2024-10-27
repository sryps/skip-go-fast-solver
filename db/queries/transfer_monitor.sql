-- name: InsertTransferMonitorMetadata :one
INSERT INTO transfer_monitor_metadata (chain_id, height_last_seen) VALUES (?, ?) ON CONFLICT (chain_id) DO UPDATE SET height_last_seen = excluded.height_last_seen, updated_at=CURRENT_TIMESTAMP RETURNING *;


-- name: GetTransferMonitorMetadata :one
SELECT * FROM transfer_monitor_metadata WHERE chain_id = ?;