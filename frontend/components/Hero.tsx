'use client'

interface HeroProps {
  onApply: () => void
}

export default function Hero({ onApply }: HeroProps) {
  return (
    <section className="bg-gradient-to-r from-primary-500 to-primary-700 text-white py-20">
      <div className="container mx-auto px-4 text-center">
        <h1 className="text-4xl md:text-6xl font-bold mb-6">
          Вывоз бытового мусора<br />два раза в неделю
        </h1>
        <p className="text-xl mb-8">
          ООО «Чистый мир» — региональный оператор по обращению с твердыми коммунальными отходами
        </p>
        <button
          onClick={onApply}
          className="bg-white text-primary-600 px-8 py-3 rounded-lg font-semibold hover:bg-gray-100 transition shadow-lg"
        >
          Оставить заявку
        </button>
      </div>
    </section>
  )
}

