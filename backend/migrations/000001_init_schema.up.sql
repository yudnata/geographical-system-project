CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL DEFAULT '',
    password VARCHAR(255),
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
CREATE TABLE object_types (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    icon VARCHAR(255)
);
CREATE TABLE geo_points (
    id SERIAL PRIMARY KEY,
    type_id INT REFERENCES object_types(id) ON DELETE
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
        created_at TIMESTAMPTZ DEFAULT NOW(),
        updated_at TIMESTAMPTZ DEFAULT NOW()
);
-- Seed default object types
INSERT INTO object_types (id, name, icon)
VALUES (1, 'Rumah', 'home'),
    (2, 'Kantor', 'building'),
    (3, 'Rumah Sakit', 'hospital'),
    (4, 'Sekolah', 'graduation-cap'),
    (5, 'Tempat Ibadah', 'church'),
    (6, 'Objek Wisata', 'map-pin'),
    (7, 'Kuliner', 'utensils'),
    (8, 'Mall', 'shopping-bag'),
    (9, 'Kantor Polisi', 'shield'),
    (10, 'Pemadam Kebakaran', 'flame'),
    (11, 'Taman', 'tree') ON CONFLICT (id) DO NOTHING;