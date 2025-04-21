'use client';

import Image from 'next/image';
import { colors } from '@/constants/colors';
import { EmailLogin } from '@/components/Auth/EmailLogin';

export const Hero = () => {
  return (
    <div className="relative w-full bg-gradient-to-br from-[#A3D1C6] via-[#3D8D7A] to-[#3D8D7A]">
      {/* Top banner */}
      <div className="w-full bg-[#3D8D7A] text-white text-center py-3">
        <p className="text-base flex items-center justify-center space-x-4">
          <span className="text-xl">←</span>
          <span className="font-medium tracking-wide">JOIN THE 1-DAY TESTING & QA SUMMIT FEATURING 15+ SPEAKERS. REGISTER FOR FREE!</span>
          <span className="text-xl">→</span>
        </p>
      </div>

      <div className="max-w-[1440px] mx-auto px-8 xl:px-20 py-16">
        <div className="flex flex-col lg:flex-row items-center justify-between gap-16">
          {/* Left side - Text content */}
          <div className="w-full lg:w-[45%] space-y-10 pt-8">
            <h1 className="text-5xl xl:text-6xl font-bold text-black leading-tight">
              Power Your Software Testing with{' '}
              <span className="text-[#0000FF]">AI and Cloud</span>
            </h1>
            
            <p className="text-xl xl:text-2xl text-gray-800 font-medium leading-relaxed">
              Test Intelligently and ship faster. Deliver unparalled digital experiences for real world enterprises
            </p>

            <div className="pt-4">
              <EmailLogin />
            </div>

            <div className="pt-12">
              <p className="text-gray-700 text-lg font-medium mb-8">Trusted by users globally</p>
              <div className="flex flex-wrap gap-x-12 gap-y-6 items-center">
                {/* Temporarily using placeholder divs until we have the actual logos */}
                <div className="w-32 h-8 bg-gray-200 rounded"></div>
                <div className="w-28 h-8 bg-gray-200 rounded"></div>
                <div className="w-32 h-8 bg-gray-200 rounded"></div>
                <div className="w-28 h-8 bg-gray-200 rounded"></div>
                <div className="w-28 h-8 bg-gray-200 rounded"></div>
              </div>
            </div>
          </div>

          {/* Right side - App preview */}
          <div className="w-full lg:w-[50%] relative">
            <div className="relative">
              {/* Teal background shape */}
              <div className="absolute -top-8 -left-8 w-full h-full bg-[#3D8D7A] rounded-lg transform -rotate-3 opacity-50"></div>
              
              {/* Main image container */}
              <div className="relative bg-white rounded-lg shadow-xl overflow-hidden">
                <Image
                  src="/app/thex.png"
                  alt="T-Hex Application"
                  width={800}
                  height={600}
                  className="w-full h-auto"
                  priority
                />
                
                {/* Play button overlay */}
                <div className="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2">
                  <button className="bg-white/90 rounded-full p-6 shadow-xl hover:bg-white transition-all group">
                    <svg
                      className="w-16 h-16 text-[#3D8D7A] group-hover:scale-110 transition-transform"
                      fill="currentColor"
                      viewBox="0 0 20 20"
                    >
                      <path d="M10 18a8 8 0 100-16 8 8 0 000 16zM9.555 7.168A1 1 0 008 8v4a1 1 0 001.555.832l3-2a1 1 0 000-1.664l-3-2z" />
                    </svg>
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}; 