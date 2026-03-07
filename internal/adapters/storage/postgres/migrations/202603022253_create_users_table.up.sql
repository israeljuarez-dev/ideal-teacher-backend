CREATE TYPE user_status AS ENUM ('active', 'inactive');

CREATE TABLE users (
    id SERIAL,
    email VARCHAR(255) NOT NULL,
    password TEXT NOT NULL,
    full_name VARCHAR(200) NOT NULL,         
    status user_status NOT NULL DEFAULT 'active',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    
    CONSTRAINT pk_users PRIMARY KEY (id),
    CONSTRAINT uk_users_email UNIQUE (email)
);