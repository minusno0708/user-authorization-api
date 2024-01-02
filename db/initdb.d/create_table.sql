CREATE DATABASE IF NOT EXISTS test_db;

USE test_db;

CREATE TABLE users
(
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    user_id VARCHAR(63),
    username VARCHAR(63),
    password VARCHAR(63)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;