'use client'

import { useState, useEffect } from 'react'
import Hero from '@/components/Hero'
import About from '@/components/About'
import HowItWorks from '@/components/HowItWorks'
import Containers from '@/components/Containers'
import Advantages from '@/components/Advantages'
import Services from '@/components/Services'
import Pricing from '@/components/Pricing'
import Benefits from '@/components/Benefits'
import Contact from '@/components/Contact'
import ApplicationForm from '@/components/ApplicationForm'
import Footer from '@/components/Footer'
import Header from '@/components/Header'

export default function Home() {
  const [showForm, setShowForm] = useState(false)

  return (
    <main className="min-h-screen">
      <Header />
      <Hero onApply={() => setShowForm(true)} />
      <About />
      <HowItWorks />
      <Containers onApply={() => setShowForm(true)} />
      <Advantages />
      <Services />
      <Pricing />
      <Benefits />
      <Contact />
      <Footer />
      {showForm && <ApplicationForm onClose={() => setShowForm(false)} />}
    </main>
  )
}

