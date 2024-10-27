CREATE TABLE
    IF NOT EXISTS transfer_monitor_metadata (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        chain_id             TEXT NOT NULL UNIQUE,
        height_last_seen              BIGINT NOT NULL DEFAULT 0
    );
