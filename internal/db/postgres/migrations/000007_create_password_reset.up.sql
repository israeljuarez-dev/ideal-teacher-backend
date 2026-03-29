-- ============================================================
-- Tabla para "Olvidé mi contraseña"
-- ============================================================
CREATE TABLE password_reset (
    id         UUID        NOT NULL DEFAULT gen_random_uuid(),
    user_id    UUID        NOT NULL,
    code       CHAR(6)     NOT NULL,          -- código de 6 dígitos
    used       BOOLEAN     NOT NULL DEFAULT FALSE,
    expires_at TIMESTAMPTZ NOT NULL,          -- backend define TTL (ej: +15 min)
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT pk_password_reset_tokens PRIMARY KEY (id), 
    CONSTRAINT fk_prt_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT ck_prt_code_digits CHECK (code ~ '^\d{6}$')
);