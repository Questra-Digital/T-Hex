import { NextRequest, NextResponse } from 'next/server';
import { PipelineEvent } from '@/types/pipeline';
import { getSSEWriter, setSSEWriter } from './connection';

interface PipelineEventRequest {
  pipeline_id: string;
  event: PipelineEvent;
}

export async function POST(request: NextRequest) {
  try {
    const body: PipelineEventRequest = await request.json();
    const { pipeline_id, event } = body;

    //Get writer globally
    const writer = getSSEWriter();
    if (!writer) {
      return NextResponse.json({ status: 500, statusText: 'No SSE Client Connected' });
    }

    // Send Pipeline_id and event to client
    const encoder = new TextEncoder();
    await writer.write(encoder.encode(`data: ${JSON.stringify({ pipeline_id, event })}\n\n`));

  } catch (error) {
    console.error('Failed to update pipeline status', error);
    return NextResponse.json({ status: 500, statusText: 'Failed to update pipeline status' });
  }
  finally {
    console.log('Pipeline status updated successfully');
    return NextResponse.json({ status: 200, statusText: 'Pipeline status updated successfully' });
  }
}

export async function GET(request: NextRequest) {
  const stream = new TransformStream();
  const writer = stream.writable.getWriter();
  const encoder = new TextEncoder();

  //Save writer gloablly
  setSSEWriter(writer);

  //Handshake: let client know connection is established
  writer.write(encoder.encode(`data: ${JSON.stringify({ status: "connected" })}\n\n`));

  return new Response(stream.readable, {
    headers: {
      'Content-Type': 'text/event-stream',
      'Cache-Control': 'no-cache',
      'Connection': 'keep-alive',
    },
  });
}