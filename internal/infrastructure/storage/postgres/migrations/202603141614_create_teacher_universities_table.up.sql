CREATE TABLE teacher_universities (
    id SERIAL,
    teacher_id UUID NOT NULL,
    university_name VARCHAR(200) NOT NULL,

    CONSTRAINT pk_teacher_universities_id PRIMARY KEY(id),
    CONSTRAINT fk_teacher_id FOREIGN KEY (teacher_id) REFERENCES teachers(id) ON DELETE CASCADE
);