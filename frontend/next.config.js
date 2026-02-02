/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  // Переменные окружения с префиксом NEXT_PUBLIC_ автоматически доступны в браузере
  // Не нужно дублировать их здесь, они уже доступны через process.env
}

module.exports = nextConfig

