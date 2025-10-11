import { APIResponse } from "@/types/API";

export interface HttpRequestOptions {
  body?: any;
  headers?: Record<string, string>;
}

class HttpUtils {
  private getBaseUrl(): string {
    return process.env.NEXT_PUBLIC_CICD_RUNNER_URL || 'http://localhost:8082';
  }

  private async httpRequest<T extends APIResponse<any>>(
    endpoint: string,
    method: "GET" | "POST" | "PUT" | "DELETE",
    options?: HttpRequestOptions
  ): Promise<T> {
    try {
      const baseUrl = this.getBaseUrl();
      const fullUrl = endpoint.startsWith('http') ? endpoint : `${baseUrl}${endpoint}`;

      const response = await fetch(fullUrl, {
        method: method || "GET",
        body: options?.body ? JSON.stringify(options.body) : undefined,
        headers: { "Content-Type": "application/json", ...options?.headers },
      });

      if (!response.ok) {
        return {
          success: false,
          message: `HTTP error: ${response.status}`,
          error: response.statusText,
          statusCode: response.status,
        } as T;
      }


      const responseJson = await response.json();

      return {
        ...responseJson,
        statusCode: response.status,
      } as T;

    } catch (err: any) {
      return {
        success: false,
        message: "Network or JSON parse error",
        error: err.message,
        statusCode: 500,
      } as T;
    }
  }

  async get<T extends APIResponse<any>>(endpoint: string, options?: HttpRequestOptions): Promise<T> {
    return this.httpRequest<T>(endpoint, "GET", options);
  }

  async post<T extends APIResponse<any>>(endpoint: string, options?: HttpRequestOptions): Promise<T> {
    return this.httpRequest<T>(endpoint, "POST", options);
  }

  async put<T extends APIResponse<any>>(endpoint: string, options?: HttpRequestOptions): Promise<T> {
    return this.httpRequest<T>(endpoint, "PUT", options);
  }

  async delete<T extends APIResponse<any>>(endpoint: string, options?: HttpRequestOptions): Promise<T> {
    return this.httpRequest<T>(endpoint, "DELETE", options);
  }
}

export const httpUtils = new HttpUtils();