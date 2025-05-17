CREATE TABLE users (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    oauth_provider VARCHAR(50) NOT NULL,
    oauth_id VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE trustee_profiles (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT UNSIGNED NOT NULL,
    name VARCHAR(255) NOT NULL,
    position VARCHAR(255),
    work_location VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE engagements (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    citizen_id BIGINT UNSIGNED NOT NULL,
    trustee_id BIGINT UNSIGNED NOT NULL,
    category ENUM('positive', 'neutral', 'negative') NOT NULL,
    comment TEXT,
    tags JSON,
    location VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (citizen_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (trustee_id) REFERENCES trustee_profiles(id) ON DELETE CASCADE
);

CREATE INDEX idx_trustee_user_id ON trustee_profiles(user_id);
CREATE INDEX idx_engagement_trustee_id ON engagements(trustee_id);
CREATE INDEX idx_engagement_citizen_id ON engagements(citizen_id);
