// app/api/pipeline/connection.ts
let sseWriter: WritableStreamDefaultWriter<Uint8Array> | null = null;

export const setSSEWriter = (writer: WritableStreamDefaultWriter<Uint8Array>) => {
  sseWriter = writer;
};

export const getSSEWriter = () => sseWriter;
