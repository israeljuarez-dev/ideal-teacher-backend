-- ============================================================
-- ROLES
-- ============================================================
CREATE TABLE roles (
    id   SMALLINT    NOT NULL,
    name VARCHAR(50) NOT NULL,

    CONSTRAINT pk_roles PRIMARY KEY (id),
    CONSTRAINT uk_roles_name UNIQUE (name)
);

INSERT INTO roles (id, name) VALUES
    (1, 'admin'),
    (2, 'user');