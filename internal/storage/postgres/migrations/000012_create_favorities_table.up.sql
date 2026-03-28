-- ============================================================
-- FAVORITOS
-- ============================================================
CREATE TABLE favorites (
    user_id    UUID        NOT NULL,
    teacher_id UUID        NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT pk_favorites PRIMARY KEY (user_id, teacher_id),
    CONSTRAINT fk_favorites_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_favorites_teacher FOREIGN KEY (teacher_id) REFERENCES teachers(id) ON DELETE CASCADE
);