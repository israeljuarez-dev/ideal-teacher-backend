-- ============================================================
-- TAGS  (etiquetas: positivas / neutrales / negativas)
-- ============================================================
CREATE TABLE tags (
    id        SERIAL       NOT NULL,
    name      VARCHAR(100) NOT NULL,
    sentiment tag_sentiment NOT NULL,

    CONSTRAINT pk_tags PRIMARY KEY (id),
    CONSTRAINT uk_tags_name UNIQUE (name)
);

-- Seed según requerimientos
INSERT INTO tags (name, sentiment) VALUES
    -- Positivas (verde)
    ('amigable',                 'positive'),
    ('puntualidad',              'positive'),
    ('flexible',                 'positive'),
    ('responsable',              'positive'),
    ('práctico',                 'positive'),
    ('teórico',                  'positive'),
    ('comunicación efectiva',    'positive'),
    ('domina los temas',         'positive'),
    -- Neutrales / exigente (amarillo)
    ('exigente',                 'neutral'),
    -- Negativas (rojo)
    ('retroalimentación limitada','negative'),
    ('no domina los temas',      'negative'),
    ('poca interacción',         'negative'),
    ('explicaciones confusas',   'negative');