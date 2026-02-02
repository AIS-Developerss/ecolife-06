'use client'

export default function HowItWorks() {
  const steps = [
    {
      title: 'Заявка через сайт',
      description: 'Вы отправляете заявку через сайт, заполнив ФИО, номер телефона и адрес проживания',
      icon: '📝',
    },
    {
      title: 'Установка бака',
      description: 'В назначенный срок приезжает наш курьер и устанавливает бак',
      icon: '📦',
    },
    {
      title: 'Заключение договора',
      description: 'После установки бака наша компания заключает с вами договор на обслуживание и дальнейшее сопровождение',
      icon: '📄',
    },
  ]

  return (
    <section className="py-20">
      <div className="container mx-auto px-4">
        <h2 className="text-3xl md:text-4xl font-bold text-center mb-12">
          Как мы работаем?
        </h2>
        <div className="grid md:grid-cols-3 gap-8">
          {steps.map((step, index) => (
            <div key={index} className="text-center p-6 bg-white rounded-lg shadow-md">
              <div className="text-6xl mb-4">{step.icon}</div>
              <h3 className="text-xl font-semibold mb-3">{step.title}</h3>
              <p className="text-gray-600">{step.description}</p>
            </div>
          ))}
        </div>
      </div>
    </section>
  )
}

