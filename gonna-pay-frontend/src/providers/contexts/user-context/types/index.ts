import { IUserData } from "@/services/auth";

export type UserContextType = {
  user: IUserData | null;
  isLoading: boolean;
  clearUser: () => void;
};
