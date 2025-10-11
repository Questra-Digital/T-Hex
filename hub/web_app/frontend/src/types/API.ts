export type JSONValue =
  | string
  | number
  | boolean
  | null
  | JSONValue[]
  | { [key: string]: JSONValue };

/*
  APIResponse is a generic response type for all API responses.
  It contains a success flag, a message, an error (if any), data (if any), and a status code (if any).
*/
export interface APIResponse<T = JSONValue> {
  success: boolean;
  message: string;
  error?: string;
  data?: T;
  statusCode?: number;
}