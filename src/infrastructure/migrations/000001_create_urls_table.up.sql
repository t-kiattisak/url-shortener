CREATE TABLE IF NOT EXISTS urls (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    original TEXT NOT NULL,
    shortened TEXT NOT NULL UNIQUE,
    expiry TIMESTAMP NOT NULL DEFAULT NOW() + interval '7 days',
    created_at TIMESTAMP DEFAULT NOW()
);