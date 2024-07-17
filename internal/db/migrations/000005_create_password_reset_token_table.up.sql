CREATE TABLE IF NOT EXISTS password_reset_token (
            id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
            user_id UUID NOT NULL,
            token VARCHAR(255) NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            expires_at TIMESTAMP NOT NULL
);

CREATE INDEX idx_password_reset_token ON password_reset_token (token);