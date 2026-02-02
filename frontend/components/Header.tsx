'use client'

import Link from 'next/link'
import { useState } from 'react'

export default function Header() {
  const [isMenuOpen, setIsMenuOpen] = useState(false)

  return (
    <header className="bg-white shadow-md sticky top-0 z-50">
      <nav className="container mx-auto px-4 py-4">
        <div className="flex items-center justify-between">
          <Link href="/" className="text-2xl font-bold text-primary-600">
            Чистый мир
          </Link>
          
          <div className="hidden md:flex space-x-6">
            <Link href="#about" className="text-gray-700 hover:text-primary-600 transition">
              О компании
            </Link>
            <Link href="#services" className="text-gray-700 hover:text-primary-600 transition">
              Услуги
            </Link>
            <Link href="#pricing" className="text-gray-700 hover:text-primary-600 transition">
              Цена
            </Link>
            <Link href="#benefits" className="text-gray-700 hover:text-primary-600 transition">
              Акции
            </Link>
            <Link href="#contact" className="text-gray-700 hover:text-primary-600 transition">
              Контакты
            </Link>
            <Link href="#faq" className="text-gray-700 hover:text-primary-600 transition">
              FAQ
            </Link>
          </div>

          <button
            className="md:hidden"
            onClick={() => setIsMenuOpen(!isMenuOpen)}
          >
            <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4 6h16M4 12h16M4 18h16" />
            </svg>
          </button>
        </div>

        {isMenuOpen && (
          <div className="md:hidden mt-4 space-y-2">
            <Link href="#about" className="block text-gray-700 hover:text-primary-600">О компании</Link>
            <Link href="#services" className="block text-gray-700 hover:text-primary-600">Услуги</Link>
            <Link href="#pricing" className="block text-gray-700 hover:text-primary-600">Цена</Link>
            <Link href="#benefits" className="block text-gray-700 hover:text-primary-600">Акции</Link>
            <Link href="#contact" className="block text-gray-700 hover:text-primary-600">Контакты</Link>
            <Link href="#faq" className="block text-gray-700 hover:text-primary-600">FAQ</Link>
          </div>
        )}
      </nav>
    </header>
  )
}

