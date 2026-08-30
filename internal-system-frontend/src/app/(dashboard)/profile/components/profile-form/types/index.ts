import * as yup from "yup";
import { profileSchema } from "@/validations/schemas";

export type ProfileFormValues = yup.InferType<typeof profileSchema> & {
  street?: string;
  neighborhood?: string;
  zip?: string;
};
