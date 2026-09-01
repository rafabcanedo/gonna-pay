import type { InferType } from "yup";
import { addCostSchema } from "@/validations/schemas";

export type ICreateCostForm = InferType<typeof addCostSchema>;