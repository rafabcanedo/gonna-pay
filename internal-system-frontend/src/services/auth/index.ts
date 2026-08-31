import { apiCall } from "@/lib/api-client";
import type { ISignInRequest, ISignUpRequest, IAuthResponse, IUserData, IUpdateProfileRequest, IUpdateProfileResponse } from './interfaces';

export type { ISignInRequest, ISignUpRequest, IAuthResponse, IUserData, IUpdateProfileRequest, IUpdateProfileResponse };

export const authService = {
  async signIn(data: ISignInRequest): Promise<IAuthResponse> {
    const res = await fetch("/api/auth/signin", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(data),
    });

    const result = await res.json();

    if (!res.ok) {
      throw result;
    }

    return result as IAuthResponse;
  },

  async signUp(data: ISignUpRequest): Promise<{ message: string }> {
    return apiCall<{ message: string }>("/auth/signup", {
      method: "POST",
      body: JSON.stringify(data),
    });
  },

  async getProfile(): Promise<IUserData> {
    return apiCall<IUserData>("/auth/profile", {
      method: "GET",
    });
  },

  async updateProfile(data: IUpdateProfileRequest): Promise<IUpdateProfileResponse> {
    return apiCall<IUpdateProfileResponse>("/auth/profile", {
      method: "PUT",
      body: JSON.stringify(data),
    });
  },

  async logout(): Promise<{ message: string }> {
    return apiCall<{ message: string }>("/auth/logout", {
      method: "POST",
    });
  },

  async deleteAccount(): Promise<{ message: string }> {
    return apiCall<{ message: string }>("/auth/profile", {
      method: "DELETE",
    });
  },
};
