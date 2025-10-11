export type JSONValue =
  | string
  | number
  | boolean
  | null
  | JSONValue[]
  | { [key: string]: JSONValue };


export interface APIResponse<T = JSONValue> {
  success: boolean;
  message: string;
  error?: string;
  data?: T;
}