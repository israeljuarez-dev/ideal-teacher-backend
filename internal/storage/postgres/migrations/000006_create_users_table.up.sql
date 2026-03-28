-- ============================================================
-- USUARIOS
-- ============================================================
CREATE TABLE users (
    id         UUID        NOT NULL DEFAULT gen_random_uuid(),
    email      VARCHAR(255) NOT NULL,
    password   TEXT        NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    last_name  VARCHAR(100) NOT NULL,
    status     user_status NOT NULL DEFAULT 'active',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT pk_users PRIMARY KEY (id),
    CONSTRAINT uk_users_email UNIQUE (email),
    CONSTRAINT ck_users_email_format CHECK (email LIKE '%@%')
);