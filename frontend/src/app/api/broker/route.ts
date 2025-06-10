import { NextResponse } from "next/server";

// GET will get the posted data from Go and return it to the server
export async function GET() {
    const response = await fetch("http://localhost:3000/api/broker");
    const data = await response.json();
    return NextResponse.json(data)
} // GET()