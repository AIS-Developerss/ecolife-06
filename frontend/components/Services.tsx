'use client'

export default function Services() {
  const services = [
    {
      title: 'Вывоз твердых бытовых отходов',
      description: 'Регулярный вывоз бытового мусора два раза в неделю',
    },
    {
      title: 'Вывоз строительного мусора',
      description: 'Специализированный вывоз строительных отходов',
    },
    {
      title: 'Разовый вывоз мусора',
      description: 'Одноразовый вывоз мусора по запросу',
    },
  ]

  return (
    <section id="services" className="py-20 bg-gray-50">
      <div className="container mx-auto px-4">
        <h2 className="text-3xl md:text-4xl font-bold text-center mb-12">
          Какие услуги мы предоставляем?
        </h2>
        <div className="grid md:grid-cols-3 gap-8">
          {services.map((service, index) => (
            <div key={index} className="bg-white rounded-lg shadow-md p-6">
              <h3 className="text-xl font-semibold mb-3">{service.title}</h3>
              <p className="text-gray-600">{service.description}</p>
            </div>
          ))}
        </div>
      </div>
    </section>
  )
}

