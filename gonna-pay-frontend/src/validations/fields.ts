import * as yup from "yup";

const messages = {
  required: "This field is required.",
  email: "Please enter a valid email.",
  passwordMin: "Password must be at least 6 characters long.",
  nameMin: "Name must be at least 6 characters long.",
  phoneInvalid: "Phone is invalid.",
  valueInvalid: "Please enter a valid value.",
  percentInvalid: "Invalid percentage. Use multiples of 10% (e.g. 10%, 20%).",
};

export const emailField = () =>
  yup.string().email(messages.email).required(messages.required);

export const passwordField = (min = 6) =>
  yup.string().min(min, messages.passwordMin).required(messages.required);

export const nameField = (min = 6) =>
  yup.string().min(min, messages.nameMin).required(messages.required);

export const phoneField = () =>
  yup
    .string()
    .matches(/^\(?\d{2}\)?[\s-]?\d{4,5}-?\d{4}$/, messages.phoneInvalid)
    .required(messages.required);

  
export const valueField = () =>
  yup
    .string()
    .matches(/^\d+(\.\d{1,2})?$/, messages.valueInvalid)
    .required(messages.required);

export const percentField = () =>
  yup
    .string()
    .required(messages.required)
    .test(
      "is-valid-percent",
      messages.percentInvalid,
      (value) => {
        if (!value) return false;
        return /^(10|20|30|40|50|60|70|80|90|100)%$/.test(value);
      }
    );

export const optional = <T extends yup.AnySchema>(schema: T) => schema.optional();

export const trimmed = <T extends yup.StringSchema>(schema: T) =>
  schema.transform((v) => (typeof v === "string" ? v.trim() : v));
