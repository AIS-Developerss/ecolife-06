'use client'

import { useEffect, useState } from 'react'
import { api } from '@/lib/api'

interface Benefit {
  id: string
  category: string
  description: string
  discount: number
}

export default function Benefits() {
  const [benefits, setBenefits] = useState<Benefit[]>([])

  useEffect(() => {
    api.getBenefits().then(setBenefits).catch(console.error)
  }, [])

  const benefitCategories: Record<string, string> = {
    large_family: 'Многодетная семья',
    disabled: 'Инвалиды',
    repressed: 'Репрессированные',
    afghan_veteran: 'Ветераны Афганской войны',
    wwii_veteran: 'Ветераны Великой Отечественной войны',
  }

  const defaultBenefits = [
    { category: 'large_family', discount: 30 },
    { category: 'disabled', discount: 30 },
    { category: 'repressed', discount: 50 },
    { category: 'afghan_veteran', discount: 100 },
    { category: 'wwii_veteran', discount: 100 },
  ]

  const displayBenefits = benefits.length > 0 ? benefits : defaultBenefits.map(b => ({
    id: b.category,
    category: b.category,
    description: benefitCategories[b.category] || b.category,
    discount: b.discount,
  }))

  return (
    <section id="benefits" className="py-20 bg-gray-50">
      <div className="container mx-auto px-4">
        <h2 className="text-3xl md:text-4xl font-bold text-center mb-4">
          Какие льготы у нас действуют?
        </h2>
        <p className="text-center text-gray-600 mb-12 max-w-3xl mx-auto">
          Государство предоставляет вам возможность вернуть частично или полностью стоимость наших услуг при условии, что договор оформлен на члена семьи, который относится к следующим категориям граждан
        </p>
        <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-6 max-w-5xl mx-auto">
          {displayBenefits.map((benefit) => (
            <div key={benefit.id} className="bg-white rounded-lg shadow-md p-6">
              <h3 className="text-xl font-semibold mb-2">
                {benefitCategories[benefit.category] || benefit.description}
              </h3>
              <p className="text-3xl font-bold text-primary-600">
                {benefit.discount}%
              </p>
              <p className="text-gray-600 text-sm mt-2">от стоимости</p>
            </div>
          ))}
        </div>
        <div className="mt-12 max-w-3xl mx-auto bg-white rounded-lg shadow-md p-6">
          <h3 className="text-xl font-semibold mb-4">Как получить возврат по акции?</h3>
          <ol className="list-decimal list-inside space-y-3 text-gray-700">
            <li>Приехать к нам в офис. Наш офис находится в городе Назрань, по улице Московская, дом 79</li>
            <li>Получить информационный лист с информацией об оплате. Наш клиентский отдел выдаст вам информационный лист за период, который вы потребуете и даст остальную интересующую вас информацию</li>
            <li>Сдать информационный лист в отдел субсидий и получить возврат средств. С информационным листом вы направляетесь в отдел субсидий и получаете возврат того процента от стоимости, который подходит вам</li>
          </ol>
        </div>
      </div>
    </section>
  )
}

