CREATE TABLE IF NOT EXISTS services (
    id VARCHAR(255) PRIMARY KEY,
    user_profile_id VARCHAR(255) NOT NULL REFERENCES user_profiles(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    price INTEGER NOT NULL,
    own_location_price INTEGER,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS service_images (
    id VARCHAR(255) PRIMARY KEY DEFAULT new_id('service_img'),
    service_id VARCHAR(255) NOT NULL REFERENCES services(id) ON DELETE CASCADE,
    url TEXT NOT NULL,
    ordering INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
