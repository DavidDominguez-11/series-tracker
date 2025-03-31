CREATE DATABASE IF NOT EXISTS series_tracker;
USE series_tracker

CREATE TABLE IF NOT EXISTS series (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(100) NOT NULL UNIQUE,
    status ENUM('Plan to Watch', 'Watching', 'Dropped', 'Completed') NOT NULL,
    last_episode_watched INT DEFAULT 0,
    total_episodes INT DEFAULT NULL,
    ranking INT DEFAULT NULL
);

CREATE USER IF NOT EXISTS 'app_user'@'%' IDENTIFIED BY 'app_password';
GRANT ALL PRIVILEGES ON series_tracker.* TO 'app_user'@'%';
FLUSH PRIVILEGES;

