ALTER TABLE dealers
    ADD COLUMN created_at timestamptz NOT NULL default now(),
    ADD COLUMN updated_at timestamptz DEFAULT NULL
;

