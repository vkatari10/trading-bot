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
  col_names: string[]
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
        <h2 className="text-xl font-semibold">Environment Data</h2>
        {envData ? (
          <ul className="list-disc pl-4">
            <li><strong>Burn Time:</strong> {envData.burn_time}</li>
            <li><strong>Refresh Rate:</strong> {envData.refresh_rate}</li>
            <li><strong>Ticker:</strong> {envData.ticker}</li>
          </ul>
        ) : (
          <div>Loading...</div>
        )}
      </section>

      <section className="bg-neutral-900 p-4 rounded">
        <h2 className="text-xl font-semibold">Log</h2>
        {logData ? (
          <ul className="list-disc pl-4">
            <li><strong>Message:</strong> {logData.msg}</li>
          </ul>
        ) : (
          <div>Loading...</div>
        )}
      </section>

      <section className="bg-neutral-900 p-4 rounded">
        <h2 className="text-xl font-semibold">Data</h2>
        {data ? (
          <ul className="list-disc pl-4">     
             <li><strong>Open:</strong>{data.quotes[3]} ({(data.quotes_delta[3].toFixed(2))})</li>
             <li><strong>High:</strong> {data.quotes[1]} ({(data.quotes_delta[1].toFixed(2))})</li>
             <li><strong>Low:</strong> {data.quotes[2]} ({(data.quotes_delta[2].toFixed(2))})</li>
             <li><strong>Close:</strong> {data.quotes[0]} ({(data.quotes_delta[0].toFixed(2))})</li>
             <li><strong>Volume:</strong> {data.quotes[4]} ({(data.quotes_delta[4].toFixed(2))})</li>
            {data.col_names.map((key, index) => (
              <li key={key}>
                {key}: {data.technicals[index].toFixed(2)}
              </li>
            ))}
          </ul>
        ) : (
          <div>Loading...</div>
        )}
      </section>

      <section className="bg-neutral-900 p-4 rounded">
        <h2 className="text-xl font-semibold">Brokerage Information</h2>
        {brokerData ? (
          <ul className="list-disc pl-4">
            <li><strong>Account Value:</strong> {brokerData.account_value}</li>
            <li><strong>Cash:</strong> {brokerData.cash}</li>
            <li><strong>Market Value:</strong> {brokerData.market_value}</li>
            <li><strong>Stock Cost:</strong> {brokerData.stock_cost}</li>
            <li><strong>Stock Quantity:</strong> {brokerData.stock_qty}</li>
          </ul>
        ) : (
          <div>Loading...</div>
        )}
      </section>
    </main>
  )
}