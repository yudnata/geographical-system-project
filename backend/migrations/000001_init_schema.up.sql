CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL DEFAULT '',
    password VARCHAR(255),
    role VARCHAR(50) DEFAULT 'contributor',
    sso_provider VARCHAR(50),
    sso_id VARCHAR(255),
    avatar_url TEXT,
    phone VARCHAR(50),
    institution VARCHAR(100),
    is_profile_completed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX idx_users_email ON users(email);
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    icon VARCHAR(255)
);
CREATE TABLE map_points (
    id SERIAL PRIMARY KEY,
    category_id INT REFERENCES categories(id) ON DELETE
    SET NULL,
        name VARCHAR(255) NOT NULL,
        latitude DECIMAL(10, 8) NOT NULL,
        longitude DECIMAL(11, 8) NOT NULL,
        address TEXT,
        owner_id UUID REFERENCES users(id) ON DELETE CASCADE,
        tahun_berdiri INT,
        status_kepemilikan VARCHAR(50),
        description TEXT,
        is_active BOOLEAN DEFAULT TRUE,
        status VARCHAR(50) DEFAULT 'draft',
        rejection_note TEXT,
        created_at TIMESTAMPTZ DEFAULT NOW(),
        updated_at TIMESTAMPTZ DEFAULT NOW()
);
CREATE TABLE blogs (
    id SERIAL PRIMARY KEY,
    map_point_id INT REFERENCES map_points(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    content TEXT,
    cover_photo TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);
-- Seed default categories
INSERT INTO categories (id, name, icon)
VALUES (1, 'Pura', 'temple'),
    (2, 'Air Terjun', 'waterfall'),
    (3, 'Pantai', 'beach'),
    (4, 'Gunung', 'mountain'),
    (5, 'Bukit', 'hill'),
    (6, 'Danau', 'lake'),
    (7, 'Hutan Wisata', 'forest'),
    (8, 'Desa Wisata', 'village'),
    (9, 'Museum', 'museum'),
    (10, 'Pasar Seni', 'market') ON CONFLICT (id) DO NOTHING;
-- Seed Admin Account (Password: admin123)
INSERT INTO users (
        id,
        email,
        name,
        password,
        role,
        is_profile_completed
    )
VALUES (
        uuid_generate_v4(),
        'admin@gmail.com',
        'Administrator',
        '$2a$10$vI8aWBnW3fID.ZQ4/zo1G.q1lRps.9cGLcZEiGDMVr5yUP1KUOYTa',
        'admin',
        true
    ) ON CONFLICT (email) DO NOTHING;