'use client'

import { useEffect, useState } from 'react'
import { api } from '@/lib/api'

interface Container {
  id: string
  volume: number
  price: number
  description: string
}

interface ContainersProps {
  onApply: () => void
}

export default function Containers({ onApply }: ContainersProps) {
  const [containers, setContainers] = useState<Container[]>([])

  useEffect(() => {
    api.getContainers().then(setContainers).catch(console.error)
  }, [])

  const containerData = [
    { volume: 120, price: 3000, description: 'Если в доме живут 4-5 человек, то 120 литров окажется вполне достаточно' },
    { volume: 240, price: 4000, description: 'Такая вместительность бака будет приемлема, если в вашей семье от 5 до 10 человек' },
    { volume: 1100, price: 19000, description: 'Подходит для установки во дворах офисов, магазинов или образовательных учреждений' },
  ]

  return (
    <section className="py-20 bg-gray-50">
      <div className="container mx-auto px-4">
        <h2 className="text-3xl md:text-4xl font-bold text-center mb-12">
          Какие баки мы ставим?
        </h2>
        <div className="grid md:grid-cols-3 gap-8">
          {containerData.map((container, index) => (
            <div key={index} className="bg-white rounded-lg shadow-md p-6 text-center">
              <h3 className="text-2xl font-bold mb-2">{container.volume} литров</h3>
              <p className="text-3xl font-bold text-primary-600 mb-4">{container.price.toLocaleString()} р.</p>
              <p className="text-gray-600 mb-6">{container.description}</p>
              <button
                onClick={onApply}
                className="bg-primary-600 text-white px-6 py-2 rounded-lg hover:bg-primary-700 transition"
              >
                Оставить заявку
              </button>
            </div>
          ))}
        </div>
      </div>
    </section>
  )
}

