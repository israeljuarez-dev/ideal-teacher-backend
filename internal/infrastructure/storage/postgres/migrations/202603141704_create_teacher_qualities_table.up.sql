CREATE TYPE teacher_quality_type AS ENUM (
    'motiva_a_estudiantes',
    'respetuoso',
    'explicaciones_claras',
    'califica_estricto',
    'organizado',
    'accesible',
    'inspiring',
    'puntual',
    'domina_el_tema'
);

CREATE TABLE teacher_quality_votes (
    teacher_id UUID NOT NULL,
    quality teacher_quality_type NOT NULL,
    student_id UUID NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT pk_teacher_quality_votes PRIMARY KEY (teacher_id, quality, student_id),
    CONSTRAINT fk_quality_votes_teacher FOREIGN KEY (teacher_id) REFERENCES teachers(id) ON DELETE CASCADE,
    CONSTRAINT fk_quality_votes_student FOREIGN KEY (student_id) REFERENCES users(id) ON DELETE CASCADE
);