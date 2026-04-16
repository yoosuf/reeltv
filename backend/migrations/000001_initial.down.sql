-- Rollback initial migration for Reel TV database schema

-- Drop triggers
DROP TRIGGER IF EXISTS update_episodes_updated_at ON episodes;
DROP TRIGGER IF EXISTS update_seasons_updated_at ON seasons;
DROP TRIGGER IF EXISTS update_series_updated_at ON series;
DROP TRIGGER IF EXISTS update_refresh_tokens_updated_at ON refresh_tokens;
DROP TRIGGER IF EXISTS update_users_updated_at ON users;

-- Drop function
DROP FUNCTION IF EXISTS update_updated_at_column();

-- Drop indexes
DROP INDEX IF EXISTS idx_episodes_is_premium;
DROP INDEX IF EXISTS idx_episodes_episode_number;
DROP INDEX IF EXISTS idx_episodes_season_id;
DROP INDEX IF EXISTS idx_episodes_uuid;

DROP INDEX IF EXISTS idx_seasons_season_number;
DROP INDEX IF EXISTS idx_seasons_series_id;
DROP INDEX IF EXISTS idx_seasons_uuid;

DROP INDEX IF EXISTS idx_series_status;
DROP INDEX IF EXISTS idx_series_year;
DROP INDEX IF EXISTS idx_series_genre;
DROP INDEX IF EXISTS idx_series_slug;
DROP INDEX IF EXISTS idx_series_uuid;

DROP INDEX IF EXISTS idx_refresh_tokens_expires_at;
DROP INDEX IF EXISTS idx_refresh_tokens_user_id;
DROP INDEX IF EXISTS idx_refresh_tokens_token;

DROP INDEX IF EXISTS idx_users_phone;
DROP INDEX IF EXISTS idx_users_email;
DROP INDEX IF EXISTS idx_users_uuid;

-- Drop tables (in reverse order of creation)
DROP TABLE IF EXISTS episodes;
DROP TABLE IF EXISTS seasons;
DROP TABLE IF EXISTS series;
DROP TABLE IF EXISTS refresh_tokens;
DROP TABLE IF EXISTS users;
