ALTER TABLE users ADD COLUMN api_key VARCHAR(64) UNIQUE NOT NULL DEFAULT (
	encode(sha256(random()::text::bytes), 'hex')
);