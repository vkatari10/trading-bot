// File-scope variable to hold last payload
let technicalData: any = null;

import { NextRequest, NextResponse } from "next/server";

// POST is called by your Go backend
export async function POST(req: NextRequest) {
    const data = await req.json();
    technicalData = data;
    return NextResponse.json({ status: "success"});
} // POST

// GET is called by your Next.js frontend to retrieve the latest data
export async function GET() {
    if (technicalData === null) {
        return NextResponse.json({ error: "No data yet" }, { status: 404 });
    } // if
    return NextResponse.json(technicalData);
} // GET()
