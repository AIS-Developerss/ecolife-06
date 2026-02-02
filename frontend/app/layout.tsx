import type { Metadata } from 'next'
import { Inter } from 'next/font/google'
import './globals.css'

const inter = Inter({ subsets: ['cyrillic', 'latin'] })

export const metadata: Metadata = {
  title: 'Чистый мир - Региональный оператор по обращению с ТКО',
  description: 'ООО «Чистый мир» - региональный оператор по обращению с твердыми коммунальными отходами на территории Республики Ингушетия',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="ru">
      <body className={inter.className}>{children}</body>
    </html>
  )
}

