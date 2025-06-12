'use client'
import { useEffect, useState } from 'react'

type EnvPayload = {
  burn_time: string
  refresh_rate: string
  ticker: string
}

type LogPayload = {
  msg: string
}

type BrokerPayload = {
  account_value: number
  cash: number
  market_value: number
  stock_cost: number
  stock_qty: number
}

type DataPayload = {
  [key: string]: number | number[]
  quotes: number[]
  quotes_delta: number[]
  technicals: number[]
}

export default function Home() {
  const [envData, setEnvData] = useState<EnvPayload | null>(null)
  const [logData, setLogData] = useState<LogPayload | null>(null)
  const [brokerData, setBrokerData] = useState<BrokerPayload | null>(null)
  const [data, setData] = useState<DataPayload | null>(null)
  const [error, setError] = useState<string | null>(null)

  // Fetch /api/env once
  useEffect(() => {
    fetch('/api/env')
      .then(res => {
        if (!res.ok) throw new Error('Failed to load env data')
        return res.json()
      })
      .then(setEnvData)
      .catch(err => setError((err as Error).message))
  }, [])

  // Poll /api/log every 250ms
  useEffect(() => {
    const interval = setInterval(() => {
      fetch('/api/log')
        .then(res => {
          if (!res.ok) throw new Error('Failed to fetch log data')
          return res.json()
        })
        .then(setLogData)
        .catch(err => setError((err as Error).message))
    }, 250)
    return () => clearInterval(interval)
  }, [])

  // Poll /api/data every 2s
  useEffect(() => {
    const interval = setInterval(() => {
      fetch('/api/data')
        .then(res => {
          if (!res.ok) throw new Error('Failed to fetch data')
          return res.json()
        })
        .then(setData)
        .catch(err => setError((err as Error).message))
    }, 2000)
    return () => clearInterval(interval)
  }, [])

  // Poll /api/broker every 2s
  useEffect(() => {
    const interval = setInterval(() => {
      fetch('/api/broker')
        .then(res => {
          if (!res.ok) throw new Error('Failed to fetch broker data')
          return res.json()
        })
        .then(setBrokerData)
        .catch(err => setError((err as Error).message))
    }, 2000)
    return () => clearInterval(interval)
  }, [])

  if (error) return <div className="p-4 text-red-500">Error: {error}</div>

  return (
    <main className="p-6 space-y-4 overflow-auto h-screen">
      <h1 className="text-2xl font-bold">Trading Bot Dashboard</h1>

      <section className="bg-neutral-900 p-4 rounded">
        <h2 className="text-xl font-semibold">Environment Data (/api/env)</h2>
        <pre>{envData ? JSON.stringify(envData, null, 2) : 'Loading...'}</pre>
      </section>

      <section className="bg-neutral-900 p-4 rounded">
        <h2 className="text-xl font-semibold">Log (/api/log)</h2>
        <pre>{logData ? JSON.stringify(logData, null, 2) : 'Loading...'}</pre>
      </section>

      <section className="bg-neutral-900 p-4 rounded">
        <h2 className="text-xl font-semibold">Data (/api/data)</h2>
        <pre>{data ? JSON.stringify(data, null, 2) : 'Loading...'}</pre>
      </section>

      <section className="bg-neutral-900 p-4 rounded">
        <h2 className="text-xl font-semibold">Brokerage Information (/api/broker)</h2>
        <pre>{brokerData ? JSON.stringify(brokerData, null, 2) : 'Loading...'}</pre>
      </section>
    </main>
  )
}
