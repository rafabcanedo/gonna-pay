import { apiCall } from "@/lib/api-client";
import type { ISignInRequest, ISignUpRequest, IAuthResponse, IUserData, IUpdateProfileRequest, IUpdateProfileResponse, IMessageResponse, IVerifyEmailRequest, IForgotPasswordRequest, IResetPasswordRequest } from './interfaces';

export type { ISignInRequest, ISignUpRequest, IAuthResponse, IUserData, IUpdateProfileRequest, IUpdateProfileResponse, IMessageResponse, IVerifyEmailRequest, IForgotPasswordRequest, IResetPasswordRequest };

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

  async verifyEmail(data: IVerifyEmailRequest): Promise<IMessageResponse> {
  return apiCall<IMessageResponse>("/auth/verify-email", {
    method: "POST",
    body: JSON.stringify(data),
  }, false);
},

async forgotPassword(data: IForgotPasswordRequest): Promise<IMessageResponse> {
  return apiCall<IMessageResponse>("/auth/forgot-password", {
    method: "POST",
    body: JSON.stringify(data),
  }, false);
},

async resetPassword(data: IResetPasswordRequest): Promise<IMessageResponse> {
  return apiCall<IMessageResponse>("/auth/reset-password", {
    method: "POST",
    body: JSON.stringify(data),
  }, false);
},
};
