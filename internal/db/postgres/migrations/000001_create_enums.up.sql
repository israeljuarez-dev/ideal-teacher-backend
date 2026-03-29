-- ============================================================
-- TIPOS ENUM
-- ============================================================

CREATE TYPE user_status AS ENUM ('active', 'inactive', 'blocked');

CREATE TYPE tag_sentiment AS ENUM ('positive', 'neutral', 'negative');