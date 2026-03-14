CREATE TABLE teacher_courses (
    id SERIAL,
    teacher_id UUID NOT NULL,
    course_name VARCHAR(200) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT pk_teacher_courses PRIMARY KEY (id),
    CONSTRAINT fk_courses_teacher FOREIGN KEY (teacher_id) REFERENCES teachers(id) ON DELETE CASCADE,
    CONSTRAINT uk_teacher_course_name UNIQUE (teacher_id, course_name)
);