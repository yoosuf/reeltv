-- Sample seed data for Reel TV MVP
-- This file contains sample series, seasons, episodes, and test users

-- Insert sample users (passwords are hashed bcrypt)
-- Test user: email: test@example.com, password: Test123!
INSERT INTO users (uuid, email, password, name, role, status) VALUES
('550e8400-e29b-41d4-a716-446655440000', 'test@example.com', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'Test User', 'user', 'active'),
('550e8400-e29b-41d4-a716-446655440001', 'admin@example.com', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'Admin User', 'admin', 'active');

-- Insert sample series
INSERT INTO series (uuid, title, slug, description, poster_url, backdrop_url, year, rating, genre, language, status, release_date) VALUES
('550e8400-e29b-41d4-a716-446655440002', 'Breaking Bad', 'breaking-bad', 'A high school chemistry teacher diagnosed with inoperable lung cancer turns to manufacturing and selling methamphetamine.', 'https://example.com/posters/breaking-bad.jpg', 'https://example.com/backdrops/breaking-bad.jpg', 2008, 9.5, 'Drama', 'en', 'active', '2008-01-20'),
('550e8400-e29b-41d4-a716-446655440003', 'Game of Thrones', 'game-of-thrones', 'Nine noble families fight for control over the lands of Westeros.', 'https://example.com/posters/got.jpg', 'https://example.com/backdrops/got.jpg', 2011, 9.3, 'Fantasy', 'en', 'active', '2011-04-17'),
('550e8400-e29b-41d4-a716-446655440004', 'Stranger Things', 'stranger-things', 'When a young boy disappears, his mother, a police chief, and his friends must confront terrifying supernatural forces.', 'https://example.com/posters/stranger-things.jpg', 'https://example.com/backdrops/stranger-things.jpg', 2016, 8.7, 'Sci-Fi', 'en', 'active', '2016-07-15'),
('550e8400-e29b-41d4-a716-446655440005', 'The Crown', 'the-crown', 'Follows the political rivalries and romance of Queen Elizabeth II''s reign and the events that shaped the second half of the twentieth century.', 'https://example.com/posters/the-crown.jpg', 'https://example.com/backdrops/the-crown.jpg', 2016, 8.6, 'Drama', 'en', 'active', '2016-11-04');

-- Insert sample seasons for Breaking Bad
INSERT INTO seasons (uuid, series_id, season_number, title, description, poster_url, release_date, episode_count) VALUES
('550e8400-e29b-41d4-a716-446655440006', 1, 1, 'Season 1', 'First season of Breaking Bad', 'https://example.com/posters/bb-s1.jpg', '2008-01-20', 7),
('550e8400-e29b-41d4-a716-446655440007', 1, 2, 'Season 2', 'Second season of Breaking Bad', 'https://example.com/posters/bb-s2.jpg', '2009-03-08', 13),
('550e8400-e29b-41d4-a716-446655440008', 1, 3, 'Season 3', 'Third season of Breaking Bad', 'https://example.com/posters/bb-s3.jpg', '2010-03-21', 13);

-- Insert sample seasons for Game of Thrones
INSERT INTO seasons (uuid, series_id, season_number, title, description, poster_url, release_date, episode_count) VALUES
('550e8400-e29b-41d4-a716-446655440009', 2, 1, 'Season 1', 'Winter Is Coming', 'https://example.com/posters/got-s1.jpg', '2011-04-17', 10),
('550e8400-e29b-41d4-a716-446655440010', 2, 2, 'Season 2', 'The Night Lands', 'https://example.com/posters/got-s2.jpg', '2012-04-01', 10),
('550e8400-e29b-41d4-a716-446655440011', 2, 3, 'Season 3', 'Valar Dohaeris', 'https://example.com/posters/got-s3.jpg', '2013-03-31', 10);

-- Insert sample seasons for Stranger Things
INSERT INTO seasons (uuid, series_id, season_number, title, description, poster_url, release_date, episode_count) VALUES
('550e8400-e29b-41d4-a716-446655440012', 3, 1, 'Season 1', 'The Vanishing of Will Byers', 'https://example.com/posters/st-s1.jpg', '2016-07-15', 8),
('550e8400-e29b-41d4-a716-446655440013', 3, 2, 'Season 2', 'The Mind Flayer', 'https://example.com/posters/st-s2.jpg', '2017-10-27', 9),
('550e8400-e29b-41d4-a716-446655440014', 3, 3, 'Season 3', 'The Battle of Starcourt', 'https://example.com/posters/st-s3.jpg', '2019-07-04', 8);

-- Insert sample episodes for Breaking Bad Season 1
INSERT INTO episodes (uuid, season_id, episode_number, title, description, thumbnail_url, duration, video_url, release_date, is_premium) VALUES
('550e8400-e29b-41d4-a716-446655440015', 1, 1, 'Pilot', 'Walter White, a chemistry teacher, discovers he has lung cancer.', 'https://example.com/thumbnails/bb-s1e1.jpg', 2880, 'https://example.com/videos/bb-s1e1.mp4', '2008-01-20', false),
('550e8400-e29b-41d4-a716-446655440016', 1, 2, 'Cat''s in the Bag...', 'Walter and Jesse dispose of the body of Emilio.', 'https://example.com/thumbnails/bb-s1e2.jpg', 2850, 'https://example.com/videos/bb-s1e2.mp4', '2008-01-27', false),
('550e8400-e29b-41d4-a716-446655440017', 1, 3, '...And the Bag''s in the River', 'Walter and Jesse try to dispose of the RV.', 'https://example.com/thumbnails/bb-s1e3.jpg', 2910, 'https://example.com/videos/bb-s1e3.mp4', '2008-02-03', false),
('550e8400-e29b-41d4-a716-446655440018', 1, 4, 'Cancer Man', 'Walter''s family discovers he has cancer.', 'https://example.com/thumbnails/bb-s1e4.jpg', 2760, 'https://example.com/videos/bb-s1e4.mp4', '2008-02-10', false),
('550e8400-e29b-41d4-a716-446655440019', 1, 5, 'Gray Matter', 'Walter meets with his former business partners.', 'https://example.com/thumbnails/bb-s1e5.jpg', 2820, 'https://example.com/videos/bb-s1e5.mp4', '2008-02-17', false),
('550e8400-e29b-41d4-a716-446655440020', 1, 6, 'Crazy Handful of Nothin'', 'Walter and Jesse have a falling out.', 'https://example.com/thumbnails/bb-s1e6.jpg', 2790, 'https://example.com/videos/bb-s1e6.mp4', '2008-02-24', false),
('550e8400-e29b-41d4-a716-446655440021', 1, 7, 'A No-Rough-Stuff-Type Deal', 'Walter and Jesse make a new deal.', 'https://example.com/thumbnails/bb-s1e7.jpg', 2850, 'https://example.com/videos/bb-s1e7.mp4', '2008-03-02', false);

-- Insert sample episodes for Stranger Things Season 1
INSERT INTO episodes (uuid, season_id, episode_number, title, description, thumbnail_url, duration, video_url, release_date, is_premium) VALUES
('550e8400-e29b-41d4-a716-446655440022', 4, 1, 'Chapter One: The Vanishing of Will Byers', 'Will Byers goes missing.', 'https://example.com/thumbnails/st-s1e1.jpg', 3060, 'https://example.com/videos/st-s1e1.mp4', '2016-07-15', false),
('550e8400-e29b-41d4-a716-446655440023', 4, 2, 'Chapter Two: The Weirdo on Maple Street', 'Lucas, Mike, and Dustin search for Will.', 'https://example.com/thumbnails/st-s1e2.jpg', 3150, 'https://example.com/videos/st-s1e2.mp4', '2016-07-15', false),
('550e8400-e29b-41d4-a716-446655440024', 4, 3, 'Chapter Three: Holly, Jolly', 'Nancy and Jonathan search for Barb.', 'https://example.com/thumbnails/st-s1e3.jpg', 3120, 'https://example.com/videos/st-s1e3.mp4', '2016-07-15', false),
('550e8400-e29b-41d4-a716-446655440025', 4, 4, 'Chapter Four: The Body', 'The boys find Eleven.', 'https://example.com/thumbnails/st-s1e4.jpg', 3180, 'https://example.com/videos/st-s1e4.mp4', '2016-07-15', false),
('550e8400-e29b-41d4-a716-446655440026', 4, 5, 'Chapter Five: The Flea and the Acrobat', 'The boys try to communicate with Will.', 'https://example.com/thumbnails/st-s1e5.jpg', 3210, 'https://example.com/videos/st-s1e5.mp4', '2016-07-15', false),
('550e8400-e29b-41d4-a716-446655440027', 4, 6, 'Chapter Six: The Monster', 'Joyce communicates with Will.', 'https://example.com/thumbnails/st-s1e6.jpg', 3090, 'https://example.com/videos/st-s1e6.mp4', '2016-07-15', false),
('550e8400-e29b-41d4-a716-446655440028', 4, 7, 'Chapter Seven: The Bathtub', 'The boys try to help Eleven.', 'https://example.com/thumbnails/st-s1e7.jpg', 3240, 'https://example.com/videos/st-s1e7.mp4', '2016-07-15', false),
('550e8400-e29b-41d4-a716-446655440029', 4, 8, 'Chapter Eight: The Upside Down', 'The final showdown.', 'https://example.com/thumbnails/st-s1e8.jpg', 3300, 'https://example.com/videos/st-s1e8.mp4', '2016-07-15', false);

-- Update episode counts in seasons
UPDATE seasons SET episode_count = 7 WHERE id = 1;
UPDATE seasons SET episode_count = 8 WHERE id = 4;
