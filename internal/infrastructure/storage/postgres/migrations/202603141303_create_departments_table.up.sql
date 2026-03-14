CREATE TABLE departments (
    id INTEGER,
    name VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT pk_departments_id PRIMARY KEY (id),
    CONSTRAINT uk_departments_name UNIQUE (name)
);