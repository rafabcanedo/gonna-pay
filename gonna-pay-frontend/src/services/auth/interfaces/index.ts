export interface IUserData {
  id: string;
  name: string;
  email: string;
  phone: string;
}

export interface IAuthResponse {
  message: string;
  user?: IUserData;
}

export interface ISignInRequest {
  email: string;
  password: string;
}

export interface ISignUpRequest {
  name: string;
  email: string;
  password: string;
  phone: string;
}

export interface IUpdateProfileRequest {
  name?: string;
  email?: string;
  phone?: string;
}

export interface IUpdateProfileResponse {
  message: string;
  user: IUserData;
}

export interface IMessageResponse {
  message: string;
}
