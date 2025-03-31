CREATE TABLE IF NOT EXISTS series (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(100) NOT NULL UNIQUE,
    status ENUM('Plan to Watch', 'Watching', 'Dropped', 'Completed') NOT NULL,
    last_episode_watched INT DEFAULT 0,
    total_episodes INT DEFAULT NULL,
    ranking INT DEFAULT NULL
);