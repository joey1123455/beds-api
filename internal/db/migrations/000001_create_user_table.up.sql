CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TYPE user_roles AS ENUM ('CUSTOMER', 'ADMIN', 'SUPERADMIN', 'STAFF');

CREATE TABLE IF NOT EXISTS users (
                       id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                       username VARCHAR(255) DEFAULT NULL,
                       email VARCHAR(255) NOT NULL UNIQUE,
                       password VARCHAR(255) NOT NULL,
                       pin TEXT,
                       email_verified BOOLEAN NOT NULL DEFAULT FALSE,
                       verify_code VARCHAR(255),
                       code_expire_time TIMESTAMP,
                       user_role user_roles NOT NULL DEFAULT 'CUSTOMER',
                       mfa_enabled BOOLEAN NOT NULL DEFAULT FALSE,
                       created_at TIMESTAMP NOT NULL,
                       updated_at TIMESTAMP NOT NULL
);

