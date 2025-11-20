import { VerifySessionResponse } from '@/types/user';
import { httpUtils, HttpRequestOptions } from './httpUtils';
import { APIResponse } from '@/types/API';

export interface SignupRequest {
  email: string;
}

export interface VerifyOTPRequest {
  email: string;
  otp: string;
}

class AuthService {
  private getApiUrl(): string {
    return process.env.NEXT_PUBLIC_API_GATEWAY_URL || process.env.API_GATEWAY_URL || 'http://localhost:8080';
  }

  async signup(email: string): Promise<APIResponse> {
    const apiUrl = this.getApiUrl();
    console.log("apiUrl", apiUrl);
    const options: HttpRequestOptions = {
      body: { "email": email },
      credentials: 'include',
    }
    const response = await httpUtils.post<APIResponse>(`${apiUrl}/signup`, options);
    console.log("response", response);
    return response;
  }

  async verifyOTP(email: string, otp: string): Promise<APIResponse<{ user_id: number }>> {
    const apiUrl = this.getApiUrl();
    const options: HttpRequestOptions = {
      body: { "email": email, "otp": otp },
      credentials: 'include',
    }
    return httpUtils.post<APIResponse<{ user_id: number }>>(`${apiUrl}/verify-otp`, options);
  }

  validateEmail(email: string): { isValid: boolean; error?: string } {
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    
    if (!email) {
      return { isValid: false, error: 'Email is required' };
    }
    
    if (!emailRegex.test(email)) {
      return { isValid: false, error: 'Please enter a valid email address' };
    }
    
    return { isValid: true };
  }

  validateOTP(otp: string): { isValid: boolean; error?: string } {
    if (!otp) {
      return { isValid: false, error: 'OTP is required' };
    }
    
    if (otp.length !== 6) {
      return { isValid: false, error: 'OTP must be 6 digits' };
    }
    
    if (!/^\d{6}$/.test(otp)) {
      return { isValid: false, error: 'OTP must contain only numbers' };
    }
    
    return { isValid: true };
  }

  async verifySession(sessionId: string): Promise<VerifySessionResponse> {
    const apiUrl = this.getApiUrl();
    const options: HttpRequestOptions = {
      headers: {
        'Authorization': `Bearer ${sessionId}`
      },
      credentials: 'include',
    }
    const response = await httpUtils.post<APIResponse<{ user_id:  number }>>(`${apiUrl}/verify-session`, options);
    
    return response;
  }
}

export const authService = new AuthService();
