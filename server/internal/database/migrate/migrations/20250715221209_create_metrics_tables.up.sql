CREATE TABLE daily_metrics (
    user_id VARCHAR(255) NOT NULL REFERENCES users(id),
    metric_date DATE NOT NULL,
    profile_views INT DEFAULT 0,
    contact_clicks INT DEFAULT 0,
    PRIMARY KEY (user_id, metric_date)
);

CREATE TABLE reviews (
    id VARCHAR(255) PRIMARY KEY DEFAULT new_id('review'),
    user_id VARCHAR(255) NOT NULL REFERENCES users(id),
    client_user_id VARCHAR(255) NOT NULL REFERENCES users(id),
    rating SMALLINT NOT NULL CHECK (rating >= 1 AND rating <= 5),
    comment TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);
