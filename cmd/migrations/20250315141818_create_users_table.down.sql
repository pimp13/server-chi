DROP TABLE IF EXISTS users;
DROP INDEX IF EXISTS idx_name_email ON users (name, email);