-- Create a new user with limited privileges
CREATE USER app_user WITH PASSWORD 'app_password';

-- Create the application database (remove IF NOT EXISTS)
-- The database is already created by POSTGRES_DB environment variable
-- CREATE DATABASE your_app_db;

-- Connect to the database and grant privileges
\c monitoring_testing;

-- Grant connection privileges
GRANT CONNECT ON DATABASE monitoring_testing TO app_user;

-- Grant schema usage
GRANT USAGE ON SCHEMA public TO app_user;

-- Grant table privileges (for existing and future tables)
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO app_user;
GRANT SELECT, USAGE ON ALL SEQUENCES IN SCHEMA public TO app_user;

-- Grant privileges for future tables and sequences
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT, INSERT, UPDATE, DELETE ON TABLES TO app_user;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT, USAGE ON SEQUENCES TO app_user;

-- Optional: Grant CREATE privilege if the app needs to create tables
GRANT CREATE ON SCHEMA public TO app_user;
