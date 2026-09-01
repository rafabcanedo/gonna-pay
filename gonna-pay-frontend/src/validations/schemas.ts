import * as yup from "yup";
import { emailField, passwordField, nameField, phoneField, trimmed, valueField } from "./fields";
import { ContactCategory, TransactionCategory } from "@/types";

export const signInSchema = yup.object({
  email: trimmed(emailField()),
  password: passwordField(6),
});

export const signUpSchema = yup.object({
  name: trimmed(nameField(6)),
  email: trimmed(emailField()),
  password: passwordField(6),
  phone: trimmed(phoneField()),
});

export const forgotPasswordSchema = yup.object({
  email: trimmed(emailField()),
});

export const profileSchema = yup.object({
  name: trimmed(nameField(2)),
  email: trimmed(emailField()),
  phone: trimmed(phoneField()),
});

export const addContactSchema = yup.object({
  name: trimmed(nameField(2)),
  email: trimmed(emailField()),
  phone: trimmed(phoneField()),
  category: yup
    .mixed<ContactCategory>()
    .oneOf(Object.values(ContactCategory), "Invalid category")
    .required("Category is required"),
});

export const createGroupSchema = yup.object({
  name: trimmed(nameField(2)),
  category: yup
    .mixed<TransactionCategory>()
    .oneOf(Object.values(TransactionCategory), "Invalid category")
    .required("Category is required"),
});

export const addCostSchema = yup.object({
  costName: trimmed(nameField(2)),
  totalValue: trimmed(valueField()),
  category: yup
    .mixed<TransactionCategory>()
    .oneOf(Object.values(TransactionCategory), "Invalid category")
    .required("Category is required"),
  groupId: yup.string().uuid("Invalid group ID").transform((v) => v || undefined).optional(),
  ownerPercentage: yup
    .string()
    .optional()
    .test("is-percent", "Invalid percentage. Must be between 0 and 100", (val) => {
      if (!val) return true;
      const num = Number(val);
      return !isNaN(num) && num >= 0 && num <= 100;
    }),
});