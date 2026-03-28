-- ============================================================
-- PROFESORES - MAESTROS
-- ============================================================
CREATE TABLE teachers (
    id             UUID         NOT NULL DEFAULT gen_random_uuid(),
    department_id  INTEGER      NOT NULL,
    profession     VARCHAR(150),
    description    VARCHAR(500),
    experience     VARCHAR(500),
    average_rating DECIMAL(3,2) NOT NULL DEFAULT 0.00,
    total_reviews  INTEGER      NOT NULL DEFAULT 0,
    is_featured    BOOLEAN      NOT NULL DEFAULT FALSE,  -- para "Profesores destacados"
    created_at     TIMESTAMPTZ  NOT NULL DEFAULT NOW(),

    CONSTRAINT pk_teachers PRIMARY KEY (id),
    CONSTRAINT fk_teachers_department FOREIGN KEY (department_id) REFERENCES departments(id) ON DELETE RESTRICT,
    CONSTRAINT ck_teachers_rating CHECK (average_rating >= 0.00 AND average_rating <= 5.00),
    CONSTRAINT ck_teachers_total_reviews CHECK (total_reviews >= 0)
);