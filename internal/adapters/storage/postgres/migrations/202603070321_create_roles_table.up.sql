CREATE TABLE roles (
    id SERIAL,
    name VARCHAR(255),

    CONSTRAINT pk_roles PRIMARY KEY (id)
);

INSERT INTO roles (id, name) values (1, 'admin');
INSERT INTO roles (id, name) values (2, 'student');
INSERT INTO roles (id, name) values (3, 'teacher');