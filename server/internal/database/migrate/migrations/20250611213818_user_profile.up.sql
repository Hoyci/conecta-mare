CREATE TABLE IF NOT EXISTS user_profiles (
    id VARCHAR(255) PRIMARY KEY DEFAULT new_id('userprofile'),
    user_id VARCHAR(255) NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    full_name VARCHAR(255),
    category_id VARCHAR(255) REFERENCES categories(id),
    subcategory_id VARCHAR(255) REFERENCES subcategories(id),
    profile_image TEXT,
    job_description TEXT,
    phone VARCHAR(20),
    social_links JSONB,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS certifications (
    id VARCHAR(255) PRIMARY KEY DEFAULT new_id('certification'),
    user_profile_id VARCHAR(255) NOT NULL REFERENCES user_profiles(id) ON DELETE CASCADE,
    institution VARCHAR(255) NOT NULL,
    course_name VARCHAR(255) NOT NULL,
    start_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS services (
    id VARCHAR(255) PRIMARY KEY DEFAULT new_id('service'),
    user_profile_id VARCHAR(255) NOT NULL REFERENCES user_profiles(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS service_images (
    id VARCHAR(255) PRIMARY KEY DEFAULT new_id('serviceimg'),
    service_id VARCHAR(255) NOT NULL REFERENCES services(id) ON DELETE CASCADE,
    url TEXT NOT NULL,
    ordering INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
