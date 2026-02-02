'use client'

export default function Contact() {
  return (
    <section id="contact" className="py-20">
      <div className="container mx-auto px-4">
        <h2 className="text-3xl md:text-4xl font-bold text-center mb-12">
          Как с нами связаться?
        </h2>
        <div className="max-w-2xl mx-auto text-center space-y-6">
          <div>
            <p className="text-xl font-semibold mb-2">Телефон</p>
            <a href="tel:+79287998282" className="text-primary-600 hover:underline">
              +7 928 799 82 82
            </a>
          </div>
          <div>
            <p className="text-xl font-semibold mb-2">Email</p>
            <a href="mailto:ooo_chistiy_mir@mail.ru" className="text-primary-600 hover:underline">
              ooo_chistiy_mir@mail.ru
            </a>
          </div>
          <div>
            <p className="text-xl font-semibold mb-2">Адрес</p>
            <p className="text-gray-700">
              Назрань, ул. Московская, 79
            </p>
          </div>
        </div>
      </div>
    </section>
  )
}

