'use client'

export default function Advantages() {
  const advantages = [
    {
      title: 'Качество',
      description: 'Мы постоянно совершенствуем наш сервис и стараемся каждый раз поднять планку качества',
      icon: '⭐',
    },
    {
      title: 'Поддержка',
      description: 'Мгновенный ответ операторов, которые решат любую вашу проблему',
      icon: '💬',
    },
    {
      title: 'Цена',
      description: 'Одно из лучших соотношений цены и качества на Северном Кавказе',
      icon: '💰',
    },
  ]

  return (
    <section className="py-20">
      <div className="container mx-auto px-4">
        <h2 className="text-3xl md:text-4xl font-bold text-center mb-4">
          Наши преимущества
        </h2>
        <p className="text-center text-gray-600 mb-12">
          Мы стараемся работать так, чтобы жизнь наших клиентов стала проще, а мир — чище!
        </p>
        <div className="grid md:grid-cols-3 gap-8">
          {advantages.map((advantage, index) => (
            <div key={index} className="text-center p-6">
              <div className="text-6xl mb-4">{advantage.icon}</div>
              <h3 className="text-xl font-semibold mb-3">{advantage.title}</h3>
              <p className="text-gray-600">{advantage.description}</p>
            </div>
          ))}
        </div>
      </div>
    </section>
  )
}

