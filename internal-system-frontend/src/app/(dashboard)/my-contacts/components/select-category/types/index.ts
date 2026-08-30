import { ContactCategory } from "@/types";

export type SelectCategoryProps = {
  value: ContactCategory | undefined;
  onValueChange: (value: ContactCategory) => void;
};
