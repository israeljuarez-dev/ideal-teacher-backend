-- ============================================================
-- RESEÑA_TAGS (etiquetas que un estudiante pone en su reseña)
-- ============================================================
CREATE TABLE review_tags (
    review_id  UUID    NOT NULL,
    tag_id     INTEGER NOT NULL,
    teacher_id UUID    NOT NULL,  

    CONSTRAINT pk_review_tags PRIMARY KEY (review_id, tag_id),
    CONSTRAINT fk_rt_review FOREIGN KEY (review_id) REFERENCES reviews(id) ON DELETE CASCADE,
    CONSTRAINT fk_rt_tag FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE RESTRICT,
    CONSTRAINT fk_rt_teacher FOREIGN KEY (teacher_id) REFERENCES teachers(id) ON DELETE CASCADE
);
