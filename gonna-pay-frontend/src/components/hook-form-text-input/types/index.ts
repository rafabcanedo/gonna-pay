export type InputType = "text" | "password" | "number" | "email" | "tel";

export interface ITextInput {
  name: string;
  type: InputType;
  label: string;
  title: string;
  disabled?: boolean;
}
