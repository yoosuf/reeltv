'use client';

import { authService } from '@/lib/auth';
import Link from 'next/link';

export default function Header() {
  const isAuthenticated = authService.isAuthenticated();

  return (
    <header className="w-full p-6 bg-black/50 backdrop-blur-sm">
      <div className="max-w-7xl mx-auto flex items-center justify-between">
        <Link href="/" className="text-3xl font-bold text-white hover:text-purple-300 transition-colors">
          Reel TV
        </Link>
        <nav className="flex gap-4 items-center">
          <Link href="/catalog" className="text-white hover:text-purple-300 transition-colors">
            Catalog
          </Link>
          <Link href="/my-list" className="text-white hover:text-purple-300 transition-colors">
            My List
          </Link>
          {isAuthenticated ? (
            <>
              <button
                onClick={() => authService.logout()}
                className="text-white hover:text-purple-300 transition-colors"
              >
                Logout
              </button>
            </>
          ) : (
            <>
              <Link
                href="/login"
                className="text-white hover:text-purple-300 transition-colors"
              >
                Login
              </Link>
              <Link
                href="/register"
                className="bg-purple-600 text-white px-4 py-2 rounded-lg hover:bg-purple-700 transition-colors"
              >
                Sign Up
              </Link>
            </>
          )}
        </nav>
      </div>
    </header>
  );
}
