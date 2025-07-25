let brokerDataStore = {}; // store data keyed by id

import { NextResponse } from "next/server";

export async function POST(request, { params }) {
  const { id } = params; 

  const data = await request.json();
  brokerDataStore[id] = data;

  return NextResponse.json({ status: "success", id });
}

export async function GET(request, { params }) {
  const { id } = params;

  if (!brokerDataStore[id]) {
    return NextResponse.json({ error: "No data for this id" }, { status: 404 });
  }

  return NextResponse.json(brokerDataStore[id]);
}
