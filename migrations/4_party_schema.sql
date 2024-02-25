CREATE TABLE parties (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
    code VARCHAR(32) NOT NULL,
    court_id BIGINT NOT NULL,
    start_time DATETIME NOT NULL,
    end_time DATETIME NOT NULL,
    FOREIGN KEY (court_id) REFERENCES courts(id)
);

CREATE TABLE party_members (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    party_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    FOREIGN KEY (party_id) REFERENCES parties(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);