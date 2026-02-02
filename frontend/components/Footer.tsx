'use client'

export default function Footer() {
  return (
    <footer className="bg-gray-800 text-white py-8">
      <div className="container mx-auto px-4">
        <div className="grid md:grid-cols-4 gap-8">
          <div>
            <h3 className="text-xl font-bold mb-4">О компании</h3>
            <p className="text-gray-400">
              ООО «Чистый мир» — региональный оператор по обращению с ТКО
            </p>
          </div>
          <div>
            <h3 className="text-xl font-bold mb-4">Цена</h3>
            <p className="text-gray-400">447,28 руб/м³</p>
          </div>
          <div>
            <h3 className="text-xl font-bold mb-4">Услуги</h3>
            <ul className="space-y-2 text-gray-400">
              <li>Вывоз бытового мусора</li>
              <li>Вывоз строительного мусора</li>
              <li>Разовый вывоз</li>
            </ul>
          </div>
          <div>
            <h3 className="text-xl font-bold mb-4">Контакты</h3>
            <p className="text-gray-400">+7 928 799 82 82</p>
            <p className="text-gray-400">ooo_chistiy_mir@mail.ru</p>
            <p className="text-gray-400">Назрань, ул. Московская, 79</p>
          </div>
        </div>
        <div className="mt-8 pt-8 border-t border-gray-700 text-center text-gray-400">
          <p>© 2024 ООО «Чистый мир». Все права защищены.</p>
        </div>
      </div>
    </footer>
  )
}

