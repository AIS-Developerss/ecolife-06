'use client'

import { useState } from 'react'
import { useForm } from 'react-hook-form'
import { api } from '@/lib/api'

interface ApplicationFormProps {
  onClose: () => void
}

interface FormData {
  full_name: string
  phone: string
  address: string
  district: string
  container_id?: string
  service_type: string
}

const districts = [
  { value: 'nazranovsky', label: 'Назрановский' },
  { value: 'sunzhensky', label: 'Сунженский' },
  { value: 'malgobeksky', label: 'Малгобекский' },
  { value: 'dzheyrakhsky', label: 'Джейрахский' },
]

export default function ApplicationForm({ onClose }: ApplicationFormProps) {
  const { register, handleSubmit, formState: { errors } } = useForm<FormData>()
  const [isSubmitting, setIsSubmitting] = useState(false)
  const [isSuccess, setIsSuccess] = useState(false)

  const onSubmit = async (data: FormData) => {
    setIsSubmitting(true)
    try {
      await api.createApplication(data)
      setIsSuccess(true)
      setTimeout(() => {
        onClose()
        setIsSuccess(false)
      }, 2000)
    } catch (error) {
      console.error('Error submitting application:', error)
      alert('Ошибка при отправке заявки. Пожалуйста, попробуйте еще раз.')
    } finally {
      setIsSubmitting(false)
    }
  }

  if (isSuccess) {
    return (
      <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
        <div className="bg-white rounded-lg p-8 max-w-md mx-4">
          <h2 className="text-2xl font-bold mb-4 text-center text-green-600">
            Заявка успешно отправлена!
          </h2>
          <p className="text-center text-gray-600">
            Мы свяжемся с вами в ближайшее время.
          </p>
        </div>
      </div>
    )
  }

  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4">
      <div className="bg-white rounded-lg p-8 max-w-md w-full max-h-[90vh] overflow-y-auto">
        <div className="flex justify-between items-center mb-6">
          <h2 className="text-2xl font-bold">Оставить заявку</h2>
          <button
            onClick={onClose}
            className="text-gray-500 hover:text-gray-700"
          >
            ✕
          </button>
        </div>
        <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
          <div>
            <label className="block text-sm font-medium mb-1">
              ФИО *
            </label>
            <input
              {...register('full_name', { required: 'Обязательное поле' })}
              className="w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-primary-500"
              placeholder="Введите ФИО"
            />
            {errors.full_name && (
              <p className="text-red-500 text-sm mt-1">{errors.full_name.message}</p>
            )}
          </div>

          <div>
            <label className="block text-sm font-medium mb-1">
              Номер телефона *
            </label>
            <input
              {...register('phone', { required: 'Обязательное поле' })}
              className="w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-primary-500"
              placeholder="+7 (___) ___-__-__"
            />
            {errors.phone && (
              <p className="text-red-500 text-sm mt-1">{errors.phone.message}</p>
            )}
          </div>

          <div>
            <label className="block text-sm font-medium mb-1">
              Адрес *
            </label>
            <input
              {...register('address', { required: 'Обязательное поле' })}
              className="w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-primary-500"
              placeholder="Введите адрес"
            />
            {errors.address && (
              <p className="text-red-500 text-sm mt-1">{errors.address.message}</p>
            )}
          </div>

          <div>
            <label className="block text-sm font-medium mb-1">
              Район *
            </label>
            <select
              {...register('district', { required: 'Обязательное поле' })}
              className="w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-primary-500"
            >
              <option value="">Выберите район</option>
              {districts.map((district) => (
                <option key={district.value} value={district.value}>
                  {district.label}
                </option>
              ))}
            </select>
            {errors.district && (
              <p className="text-red-500 text-sm mt-1">{errors.district.message}</p>
            )}
          </div>

          <div>
            <label className="block text-sm font-medium mb-1">
              Тип услуги
            </label>
            <select
              {...register('service_type')}
              className="w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-primary-500"
            >
              <option value="household">Вывоз бытового мусора</option>
              <option value="building">Вывоз строительного мусора</option>
              <option value="one_time">Разовый вывоз мусора</option>
            </select>
          </div>

          <button
            type="submit"
            disabled={isSubmitting}
            className="w-full bg-primary-600 text-white py-3 rounded-lg font-semibold hover:bg-primary-700 transition disabled:opacity-50"
          >
            {isSubmitting ? 'Отправка...' : 'Отправить заявку'}
          </button>
        </form>
      </div>
    </div>
  )
}

