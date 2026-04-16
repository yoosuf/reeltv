// User types
export interface User {
  id: number;
  uuid: string;
  email: string;
  username?: string;
  full_name?: string;
  created_at: string;
  updated_at: string;
}

// Auth types
export interface LoginRequest {
  email: string;
  password: string;
}

export interface RegisterRequest {
  email: string;
  password: string;
  full_name?: string;
  username?: string;
}

export interface AuthResponse {
  access_token: string;
  refresh_token: string;
  user: User;
}

// Catalog types
export interface Series {
  id: number;
  uuid: string;
  title: string;
  description: string;
  genre: string;
  tags: string[];
  thumbnail_url?: string;
  release_year: number;
  rating: number;
  seasons: Season[];
  created_at: string;
  updated_at: string;
}

export interface Season {
  id: number;
  uuid: string;
  series_id: number;
  season_number: number;
  title: string;
  description?: string;
  thumbnail_url?: string;
  episodes: Episode[];
  created_at: string;
  updated_at: string;
}

export interface Episode {
  id: number;
  uuid: string;
  season_id: number;
  episode_number: number;
  title: string;
  description?: string;
  duration: number; // in seconds
  thumbnail_url?: string;
  video_url?: string;
  created_at: string;
  updated_at: string;
}

// Watch progress types
export interface WatchProgress {
  id: number;
  user_id: number;
  episode_id: number;
  progress: number; // in seconds
  completed: boolean;
  last_watched_at: string;
  created_at: string;
  updated_at: string;
}

// My List types
export interface MyListItem {
  id: number;
  user_id: number;
  series_id: number;
  added_at: string;
  series: Series;
}

// Subscription types
export interface Subscription {
  id: number;
  uuid: string;
  user_id: number;
  plan_type: string;
  status: string;
  start_date: string;
  end_date: string;
  auto_renew: boolean;
  created_at: string;
  updated_at: string;
}

// Recommendation types
export interface Recommendation {
  series: Series;
  reason: string;
}
