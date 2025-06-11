"use client";

import React, { useEffect, useState } from "react";

export default function Page() {
  const [data, setData] = useState<any>(null);
  const [loading, setLoading] = useState(true);

  let log: any[] = []

  useEffect(() => {
    async function fetchData() {
      try {
        const res = await fetch("/api/data");
        const json = await res.json();
        setData(json);
        log.push(json)
        console.log(log)
      } catch (error) {
        setData({ error: "Failed to fetch data" });
        console.log("HELLO FROM ERROR")
      } finally {
        setLoading(false);
        console.log("HELLO FROM FINALLY")
      }
    }
    fetchData();
  }, []);

  if (loading) return <p>Loading...</p>;

  if (data?.error) return <p>Error: {data.error}</p>;

  return (
    <div>
      <pre>{JSON.stringify(data, null, 2)}</pre>
    </div>
  );
}
