-- ============================================================
-- PROFESOR_UNIVERSIDADES
-- ============================================================
CREATE TABLE teacher_universities (
    id              SERIAL       NOT NULL,
    teacher_id      UUID         NOT NULL,
    university_name VARCHAR(200) NOT NULL,

    CONSTRAINT pk_teacher_universities PRIMARY KEY (id),
    CONSTRAINT fk_tu_teacher FOREIGN KEY (teacher_id)  REFERENCES teachers(id) ON DELETE CASCADE,
    CONSTRAINT uk_teacher_university UNIQUE (teacher_id, university_name)            
);
