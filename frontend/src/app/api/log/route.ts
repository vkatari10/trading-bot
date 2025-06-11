// File-scope variable to hold last payload
let loggingData: any = null;

import { NextRequest, NextResponse } from "next/server";

// Go Puts data here
export async function POST(req: NextRequest) {
    const data = await req.json();
    loggingData = data;
    return NextResponse.json({ status: "success"});
} // POST

// GET is called by page.tsx
export async function GET() {
    if (loggingData === null) {
        return NextResponse.json({ error: "No data yet" }, { status: 404 });
    } // if
    return NextResponse.json(loggingData);
} // GET()
