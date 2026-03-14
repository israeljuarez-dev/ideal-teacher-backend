CREATE TABLE teachers (
    id UUID,
    user_id UUID NOT NULL,
    department_id INTEGER NOT NULL,
    profession VARCHAR(150),
    description TEXT,
    experience TEXT,
    average_rating DECIMAL(3,2) DEFAULT 0.00,
    total_reviews INTEGER DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT ck_teachers_rating CHECK (average_rating >= 0 AND average_rating <= 5),
    CONSTRAINT ck_teachers_total_reviews CHECK (total_reviews >= 0),
    CONSTRAINT pk_teachers_id PRIMARY KEY (id),
    CONSTRAINT fk_teachers_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_teachers_department FOREIGN KEY (department_id) REFERENCES departments(id)
);