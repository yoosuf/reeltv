import Header from '@/components/Header';

export default function Home() {
  return (
    <div className="flex flex-col min-h-screen bg-gradient-to-br from-purple-900 via-indigo-900 to-black">
      <Header />

      <main className="flex-1 flex items-center justify-center px-6">
        <div className="text-center max-w-3xl">
          <h2 className="text-5xl md:text-7xl font-bold text-white mb-6">
            Short Drama.
            <br />
            <span className="text-purple-400">Big Stories.</span>
          </h2>
          <p className="text-xl text-gray-300 mb-8">
            Professionally produced short-form dramatic content, curated for your viewing pleasure.
          </p>
          <div className="flex gap-4 justify-center">
            <button className="bg-purple-600 text-white px-8 py-3 rounded-lg text-lg font-semibold hover:bg-purple-700 transition-colors">
              Start Watching
            </button>
            <button className="border-2 border-white text-white px-8 py-3 rounded-lg text-lg font-semibold hover:bg-white hover:text-purple-900 transition-colors">
              Learn More
            </button>
          </div>
        </div>
      </main>

      <footer className="w-full p-6 text-center text-gray-400">
        <p>&copy; 2026 Reel TV. All rights reserved.</p>
      </footer>
    </div>
  );
}
