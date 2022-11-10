-- +goose Up
CREATE TABLE IF NOT EXISTS users (
id int NOT NULL PRIMARY KEY ,
name text NOT NULL
);

INSERT INTO users VALUES
(0, 'root'),(1, 'vojtechvitek'),(2, 'asd'),(3, 'dddd'),(4, 'vvvv'),(5, 'bbbb'),(6, 'root'),(7, 'vojtechvitek'),(8, 'asd'),(9, 'dddd'),(10, 'vvvv'),(11, 'bbbb'),
(12, 'root'),(13, 'vojtechvitek'),(14, 'asd'),(15, 'dddd'),(16, 'vvvv'),(17, 'bbbb'),(18, 'root'),(19, 'vojtechvitek'),(20, 'asd'),(21, 'dddd'),(22, 'vvvv'),(23, 'bbbb'),
(24, 'root'),(25, 'vojtechvitek'),(26, 'asd'),(27, 'dddd'),(28, 'vvvv'),(29, 'bbbb'),(30, 'root'),(31, 'vojtechvitek'),(32, 'asd'),(33, 'dddd'),(34, 'vvvv'),(35, 'bbbb');





