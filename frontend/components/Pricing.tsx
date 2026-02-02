'use client'

import { useEffect, useState } from 'react'
import { api } from '@/lib/api'

export default function Pricing() {
  const [tariff, setTariff] = useState<{ price: number } | null>(null)

  useEffect(() => {
    api.getCurrentTariff()
      .then((data) => setTariff(data))
      .catch(console.error)
  }, [])

  return (
    <section id="pricing" className="py-20">
      <div className="container mx-auto px-4 text-center">
        <h2 className="text-3xl md:text-4xl font-bold mb-8">
          Сколько стоит?
        </h2>
        <div className="max-w-2xl mx-auto bg-white rounded-lg shadow-md p-8">
          <h3 className="text-2xl font-semibold mb-4">
            Вывоз бытового мусора<br />два раза в неделю
          </h3>
          {tariff ? (
            <p className="text-3xl font-bold text-primary-600 mb-4">
              {tariff.price.toFixed(2)} рублей за куб.м.
            </p>
          ) : (
            <p className="text-3xl font-bold text-primary-600 mb-4">
              447,28 рублей за куб.м.
            </p>
          )}
        </div>
      </div>
    </section>
  )
}

