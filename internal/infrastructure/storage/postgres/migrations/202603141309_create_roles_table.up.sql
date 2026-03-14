CREATE TABLE roles (
    id INTEGER,
    name VARCHAR(255),

    CONSTRAINT pk_roles PRIMARY KEY (id),
    CONSTRAINT uk_roles_name UNIQUE (name)
);

INSERT INTO roles (id, name) VALUES (1, 'admin'), (2, 'student'), (3, 'teacher');