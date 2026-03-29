-- ============================================================
-- ÁREAS - FACULTADES
-- ============================================================
CREATE TABLE departments (
    id   SERIAL,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT pk_departments PRIMARY KEY (id),
    CONSTRAINT uk_departments_name UNIQUE (name)
);