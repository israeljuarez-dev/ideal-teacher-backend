-- ============================================================
-- PROFESOR_CURSOS
-- ============================================================
CREATE TABLE teacher_courses (
    id          SERIAL       NOT NULL,
    teacher_id  UUID         NOT NULL,
    course_name VARCHAR(200) NOT NULL,
    created_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW(),

    CONSTRAINT pk_teacher_courses PRIMARY KEY (id),
    CONSTRAINT fk_tc_teacher FOREIGN KEY (teacher_id) REFERENCES teachers(id) ON DELETE CASCADE,
    CONSTRAINT uk_teacher_course UNIQUE (teacher_id, course_name)
);
