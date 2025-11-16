import { APIResponse } from "./API";

export interface UserData {
user_id: number;
}

export interface VerifySessionResponse extends APIResponse <{ user_id: number }> {}