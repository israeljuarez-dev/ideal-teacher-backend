CREATE TABLE reviews (
    id UUID,
    student_id UUID NOT NULL,
    teacher_id UUID NOT NULL,
    content TEXT NOT NULL,
    rating INTEGER NOT NULL,
    is_positive BOOLEAN DEFAULT true,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT pk_reviews PRIMARY KEY (id),
    CONSTRAINT fk_reviews_student FOREIGN KEY (student_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_reviews_teacher FOREIGN KEY (teacher_id) REFERENCES teachers(id) ON DELETE CASCADE,
    CONSTRAINT ck_reviews_rating CHECK (rating BETWEEN 1 AND 5),
    CONSTRAINT uk_student_teacher_review UNIQUE (student_id, teacher_id)
);