

CREATE TABLE IF NOT EXISTS user_profile (
                       id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                       first_name VARCHAR(255) DEFAULT NULL,
                       last_name VARCHAR(255) DEFAULT NULL,
                       phone_number VARCHAR(255) DEFAULT NULL,
                       street VARCHAR(255) DEFAULT NULL,
                       city VARCHAR(255) DEFAULT  NULL,
                       state VARCHAR(255) DEFAULT  NULL,
                       country VARCHAR(255) DEFAULT  NULL,
                       postal_code VARCHAR(255) DEFAULT NULL,
                       user_id UUID NOT NULL,
                       created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                       FOREIGN KEY (user_id) REFERENCES users (id)
);