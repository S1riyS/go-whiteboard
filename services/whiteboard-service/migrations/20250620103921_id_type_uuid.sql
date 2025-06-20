-- +goose Up
-- +goose StatementBegin

-- Add temporary column for new UUID values
ALTER TABLE whiteboard ADD COLUMN new_id UUID;

-- Generate UUIDs for all existing records
UPDATE whiteboard SET new_id = gen_random_uuid();

-- Remove primary key constraint and old ID column
ALTER TABLE whiteboard DROP CONSTRAINT whiteboard_pkey;
ALTER TABLE whiteboard DROP COLUMN id;

-- Rename new column and set it as primary key with default value
ALTER TABLE whiteboard RENAME COLUMN new_id TO id;
ALTER TABLE whiteboard ALTER COLUMN id SET NOT NULL;
ALTER TABLE whiteboard ALTER COLUMN id SET DEFAULT gen_random_uuid();
ALTER TABLE whiteboard ADD PRIMARY KEY (id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- Add temporary column for SERIAL values
ALTER TABLE whiteboard ADD COLUMN old_id SERIAL;

-- Drop primary key constraint
ALTER TABLE whiteboard DROP CONSTRAINT whiteboard_pkey;

-- Remove UUID column
ALTER TABLE whiteboard DROP COLUMN id;

-- Rename temporary column and set it as primary key
ALTER TABLE whiteboard RENAME COLUMN old_id TO id;
ALTER TABLE whiteboard ADD PRIMARY KEY (id);

-- +goose StatementEnd