import { apiCall } from "@/lib/api-client";
import type { ISignInRequest, ISignUpRequest, IAuthResponse, IUserData, IUpdateProfileRequest, IUpdateProfileResponse, IMessageResponse } from './interfaces';

export type { ISignInRequest, ISignUpRequest, IAuthResponse, IUserData, IUpdateProfileRequest, IUpdateProfileResponse, IMessageResponse };

export const authService = {
  async signIn(data: ISignInRequest): Promise<IAuthResponse> {
    return apiCall<IAuthResponse>("/auth/signin", {
      method: "POST",
      body: JSON.stringify(data),
    }, false);
  },

  async signUp(data: ISignUpRequest): Promise<IMessageResponse> {
    return apiCall<IMessageResponse>("/auth/signup", {
      method: "POST",
      body: JSON.stringify(data),
    }, false);
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

  async logout(): Promise<IMessageResponse> {
    return apiCall<IMessageResponse>("/auth/logout", {
      method: "POST",
    });
  },

  async deleteAccount(): Promise<IMessageResponse> {
    return apiCall<IMessageResponse>("/auth/profile", {
      method: "DELETE",
    });
  },
};
