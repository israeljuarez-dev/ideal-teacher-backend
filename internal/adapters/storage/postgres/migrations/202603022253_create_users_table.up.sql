CREATE TYPE user_role AS ENUM ('student', 'teacher', 'admin');
CREATE TYPE user_status AS ENUM ('active', 'inactive');

CREATE TABLE users (
    id SERIAL,
    email VARCHAR(150) NOT NULL,
    password TEXT NOT NULL,
    full_name VARCHAR(200) NOT NULL,
    role user_role NOT NULL,           
    status user_status DEFAULT 'active',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    
    CONSTRAINT pk_users PRIMARY KEY (id),
    CONSTRAINT uk_users_email UNIQUE (email)
);