# MySQL ver 8.0
CREATE DATABASE word_db;
USE word_db;
CREATE TABLE word_tb (word VARCHAR(255) PRIMARY KEY, num INT NOT NULL);
CREATE USER 'user' IDENTIFIED BY 'Password@123';
GRANT ALL ON word_db.* TO 'user';
