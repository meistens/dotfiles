-- Create database
CREATE DATABASE monitoring_testing;

-- Connect to the application database
\c monitoring_testing;

-- Create a dedicated user for runtime operations (limited privileges)
CREATE USER app_user WITH PASSWORD 'app_password';

-- =============================================================================
-- RUNTIME USER (app_user) - Limited to DML operations only
-- =============================================================================

-- Grant basic connection privileges to app_user
GRANT CONNECT ON DATABASE monitoring_testing TO app_user;
GRANT USAGE ON SCHEMA public TO app_user;

-- Grant runtime privileges to app_user (DML operations only - no DDL)
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO app_user;
GRANT SELECT, USAGE ON ALL SEQUENCES IN SCHEMA public TO app_user;

-- Grant privileges for future tables and sequences (for app_user)
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT, INSERT, UPDATE, DELETE ON TABLES TO app_user;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT, USAGE ON SEQUENCES TO app_user;

-- =============================================================================
-- USAGE SUMMARY
-- =============================================================================

-- postgres (superuser) capabilities:
--   ✓ All operations including CREATE, ALTER, DROP
--   ✓ Run migrations with full privileges
--   ✓ Database administration

-- app_user capabilities:
--   ✓ SELECT, INSERT, UPDATE, DELETE on tables
--   ✓ SELECT, USAGE on sequences
--   ✗ CREATE, ALTER, DROP operations

-- Connection strings:
-- Runtime: postgresql://app_user:app_password@postgres:5432/monitoring_testing?sslmode=require
-- Migrations: postgresql://postgres:postgres@postgres:5432/monitoring_testing?sslmode=require
