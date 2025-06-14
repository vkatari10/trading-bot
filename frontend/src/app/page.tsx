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

type LogEntry = {
  msg: string
  timestamp: string
}

export default function Home() {
  const [envData, setEnvData] = useState<EnvPayload | null>(null)
  const [logHistory, setLogHistory] = useState<LogEntry[]>([])
  const [brokerData, setBrokerData] = useState<BrokerPayload | null>(null)
  const [data, setData] = useState<DataPayload | null>(null)
  const [error, setError] = useState<string | null>(null)

  // Set full-page gradient background
  useEffect(() => {
    document.body.style.background = 'linear-gradient(135deg, #444 0%, #232323 100%)'
    document.body.style.minHeight = '100vh'
    return () => {
      document.body.style.background = ''
      document.body.style.minHeight = ''
    }
  }, [])

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

  // Poll /api/log every 250ms, keep last 50 logs with timestamps
  useEffect(() => {
    const interval = setInterval(() => {
      fetch('/api/log')
        .then(res => {
          if (!res.ok) throw new Error('Failed to fetch log data')
          return res.json()
        })
        .then((log: LogPayload) => {
          setLogHistory(prev => {
            const now = new Date()
            const entry: LogEntry = {
              msg: log.msg,
              timestamp: now.toLocaleTimeString()
            }
            // Only add if different from last
            let updated = prev
            if (prev.length === 0 || prev[0].msg !== entry.msg) {
              updated = [entry, ...prev]
            }
            // Always keep at most 50
            return updated.slice(0, 50)
          })
        })
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

  return (
    <main className="main-dashboard">
      <style>{`
        .main-dashboard {
          display: flex;
          flex-direction: row;
          height: 100vh;
          color: #e0e0e0;
          font-family: 'Segoe UI', 'Arial', sans-serif;
        }
        .left-col {
          flex: 1 1 0;
          display: flex;
          flex-direction: column;
          gap: 1.5rem;
          padding: 2rem 1rem 2rem 2rem;
          min-width: 0;
        }
        .top-row {
          display: flex;
          flex-direction: row;
          gap: 1.5rem;
        }
        .data-table-container {
          max-width: 420px;
          min-width: 260px;
          flex: 1 1 0;
          display: flex;
          flex-direction: column;
        }
        .env-container {
          min-width: 220px;
          max-width: 320px;
          flex: 0 0 auto;
          display: flex;
          flex-direction: column;
        }
        .glass {
          background: rgba(40, 40, 40, 0.92);
          border: 1.5px solid #666;
          border-radius: 14px;
          box-shadow: 0 4px 32px 0 rgba(80,80,80,0.10), 0 1.5px 0 #444 inset;
          backdrop-filter: blur(8px);
          margin-bottom: 0.5rem;
        }
        .glass h2 {
          color: #bdbdbd;
          text-shadow: 0 0 8px #2228;
        }
        .data-table {
          width: 100%;
          border-collapse: collapse;
          margin-top: 0.5rem;
        }
        .data-table th, .data-table td {
          border: 1px solid #555;
          padding: 0.4em 0.8em;
          text-align: left;
        }
        .data-table th {
          background: #333;
          color: #bbb;
          font-weight: 600;
        }
        .data-table tr:nth-child(even) {
          background: #292929;
        }
        .data-table tr:nth-child(odd) {
          background: #232323;
        }
        .log-scroll {
          max-height: 80vh;
          min-height: 300px;
          overflow-y: auto;
          font-family: 'Fira Mono', 'Consolas', 'Menlo', monospace;
          font-size: 1rem;
          background: rgba(30, 30, 30, 0.97);
          border-radius: 10px;
          border: 1px solid #555;
          padding: 0.75rem 1rem;
          margin-top: 0.5rem;
          box-shadow: 0 0 12px #111a;
          scrollbar-width: thin;
          scrollbar-color: #888 #232323;
        }
        .log-scroll::-webkit-scrollbar {
          width: 8px;
        }
        .log-scroll::-webkit-scrollbar-thumb {
          background: #555;
          border-radius: 8px;
        }
        .log-scroll::-webkit-scrollbar-track {
          background: #232323;
        }
        .log-entry {
          display: flex;
          gap: 0.7em;
          align-items: baseline;
          padding-bottom: 2px;
        }
        .log-timestamp {
          color: #aaa;
          font-size: 0.93em;
          min-width: 70px;
        }
        .log-msg {
          color: #e0e0e0;
          word-break: break-word;
        }
        .glass ul li strong {
          color: #bbb;
          letter-spacing: 0.02em;
        }
        .highlight-decide { color: #ffb300; font-weight: bold; }
        .highlight-stage { color: #42a5f5; font-weight: bold; }
        .highlight-update { color: #66bb6a; font-weight: bold; }
        .highlight-quote { color: #e57373; font-weight: bold; }
        @media (max-width: 1100px) {
          .top-row {
            flex-direction: column;
            gap: 0.5rem;
          }
          .data-table-container, .env-container {
            max-width: 100%;
            min-width: 0;
          }
        }
        @media (max-width: 900px) {
          .main-dashboard {
            flex-direction: column;
          }
          .right-col {
            width: 100%;
            padding: 0 1rem 2rem 1rem;
          }
          .left-col {
            padding: 2rem 1rem;
          }
        }
      `}</style>
      <div className="left-col">
        <h1 className="text-2xl font-bold mb-2" style={{color:'#e0e0e0'}}>Trading Bot Dashboard</h1>
        <div className="top-row">
          <div className="glass p-4 data-table-container">
            <h2 className="text-xl font-semibold">Data</h2>
            {data ? (
              <table className="data-table">
                <thead>
                  <tr>
                    <th>Type</th>
                    <th>Value</th>
                    <th>Delta</th>
                  </tr>
                </thead>
                <tbody>
                  <tr>
                    <td>Open</td>
                    <td>{data.quotes[3]}</td>
                    <td>{data.quotes_delta[3].toFixed(2)}</td>
                  </tr>
                  <tr>
                    <td>High</td>
                    <td>{data.quotes[1]}</td>
                    <td>{data.quotes_delta[1].toFixed(2)}</td>
                  </tr>
                  <tr>
                    <td>Low</td>
                    <td>{data.quotes[2]}</td>
                    <td>{data.quotes_delta[2].toFixed(2)}</td>
                  </tr>
                  <tr>
                    <td>Close</td>
                    <td>{data.quotes[0]}</td>
                    <td>{data.quotes_delta[0].toFixed(2)}</td>
                  </tr>
                  <tr>
                    <td>Volume</td>
                    <td>{data.quotes[4]}</td>
                    <td>{data.quotes_delta[4].toFixed(2)}</td>
                  </tr>
                  {data.col_names.map((key, index) => (
                    <tr key={key}>
                      <td>{key}</td>
                      <td>{data.technicals[index].toFixed(2)}</td>
                      <td>-</td>
                    </tr>
                  ))}
                </tbody>
              </table>
            ) : (
              <div className="text-gray-400">Could not load data, or still inside Burn-In Stage.</div>
            )}
          </div>
          <div className="glass p-4 env-container" style={{marginLeft: '0.5rem', minWidth: 220, maxWidth: 320}}>
            <h2 className="text-xl font-semibold">Account Information</h2>
            {brokerData ? (
              <ul className="list-disc pl-4">
                <li><strong>Account Value:</strong> {brokerData.account_value}</li>
                <li><strong>Cash:</strong> {brokerData.cash}</li>
                <li><strong>Market Value:</strong> {brokerData.market_value}</li>
                <li><strong>Stock Cost:</strong> {brokerData.stock_cost}</li>
                <li><strong>Stock Quantity:</strong> {brokerData.stock_qty}</li>
              </ul>
            ) : (
              <div className="text-gray-400">Could not load brokerage info.</div>
            )}
          </div>
        </div>
        <section className="glass p-4" style={{marginTop: '1.5rem'}}>
          <h2 className="text-xl font-semibold">Environment Data</h2>
          {envData ? (
            <ul className="list-disc pl-4">
              <li><strong>Burn Time (Minutes):</strong> {envData.burn_time}</li>
              <li><strong>Refresh Rate (Seconds):</strong> {envData.refresh_rate}</li>
              <li><strong>Ticker:</strong> {envData.ticker}</li>
            </ul>
          ) : (
            <div className="text-gray-400">Could not load environment data.</div>
          )}
        </section>
      </div>
      <div className="right-col">
        <section className="glass p-4" style={{height: '100%'}}>
          <h2 className="text-xl font-semibold">Log Stream</h2>
          <div className="log-scroll">
            {logHistory.length === 0 ? (
              <div className="text-gray-400">No log data available.</div>
            ) : (
              logHistory.map((entry, idx) => (
                <div className="log-entry" key={idx}>
                  <span className="log-timestamp">{entry.timestamp}</span>
                  <span className="log-msg">{entry.msg}</span>
                </div>
              ))
            )}
          </div>
        </section>
      </div>
    </main>
  )

}