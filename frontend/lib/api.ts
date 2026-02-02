import axios from 'axios'

const API_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080/api'

const client = axios.create({
  baseURL: API_URL,
  headers: {
    'Content-Type': 'application/json',
  },
})

export interface Application {
  id: string
  full_name: string
  phone: string
  address: string
  district: string
  container_id?: string
  service_type: string
  status: string
  created_at: string
  updated_at: string
}

export interface Container {
  id: string
  volume: number
  price: number
  description: string
  is_active: boolean
}

export interface Benefit {
  id: string
  category: string
  description: string
  discount: number
  is_active: boolean
}

export interface Tariff {
  id: string
  price: number
  valid_from: string
  valid_to?: string
  is_active: boolean
  description: string
}

export interface CreateApplicationRequest {
  full_name: string
  phone: string
  address: string
  district: string
  container_id?: string
  service_type: string
}

export const api = {
  async createApplication(data: CreateApplicationRequest): Promise<Application> {
    const response = await client.post<Application>('/applications', data)
    return response.data
  },

  async getApplication(id: string): Promise<Application> {
    const response = await client.get<Application>(`/applications/${id}`)
    return response.data
  },

  async getContainers(): Promise<Container[]> {
    const response = await client.get<Container[]>('/containers')
    return response.data
  },

  async getContainer(id: string): Promise<Container> {
    const response = await client.get<Container>(`/containers/${id}`)
    return response.data
  },

  async getBenefits(): Promise<Benefit[]> {
    const response = await client.get<Benefit[]>('/benefits')
    return response.data
  },

  async getCurrentTariff(): Promise<Tariff> {
    const response = await client.get<Tariff>('/tariffs/current')
    return response.data
  },

  async getAllTariffs(): Promise<Tariff[]> {
    const response = await client.get<Tariff[]>('/tariffs')
    return response.data
  },
}

