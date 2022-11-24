CREATE DATABASE recipe_bot_db;

USE recipe_bot_db;

CREATE TABLE users(
                      id INTEGER PRIMARY KEY AUTO_INCREMENT,
                      name VARCHAR(255),
                      telegram_id INT,
                      first_name VARCHAR(255),
                      last_name VARCHAR(255),
                      chat_id INT,
                      created_at DATETIME default CURRENT_TIMESTAMP,
                      updated_at DATETIME,
                      deleted_at DATETIME
);

CREATE TABLE mails(
                      id INTEGER PRIMARY KEY AUTO_INCREMENT,
                      letter VARCHAR(255),
                      ReceivedDate DATETIME default CURRENT_TIMESTAMP
);

SELECT * FROM users;


