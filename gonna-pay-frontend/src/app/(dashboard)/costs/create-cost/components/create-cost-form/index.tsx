"use client";

import {
  Card,
  CardHeader,
  CardTitle,
  CardContent,
  CardFooter,
} from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { useForm, FormProvider, SubmitHandler } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import { HookFormTextInput } from "@/components/hook-form-text-input";
import { addCostSchema } from "@/validations/schemas";
import { TransactionCategory } from "@/types";
import { Loader2 } from "lucide-react";
import { HookFormSelect } from "@/components/hook-form-select";
import { ICreateCostForm } from "./types";
import { useRouter } from "next/navigation";
import { useCreateCost } from "@/hooks/mutations/use-cost-mutations";
import { useGroupsQuery } from "@/hooks/queries/use-group-query";

export default function CreateCostForm() {
  const router = useRouter();
  const { mutateAsync: createCost, isPending } = useCreateCost();
  const { data: groupsData } = useGroupsQuery();

  const groupOptions = (groupsData?.groups ?? []).map((group) => ({
    label: group.name,
    value: group.id,
  }));

  const methods = useForm({
    resolver: yupResolver(addCostSchema),
    mode: "onSubmit",
    defaultValues: {
      costName: "",
      totalValue: "",
      category: "" as TransactionCategory,
      groupId: "",
      ownerPercentage: "",
    },
  });

  const { handleSubmit } = methods;

  const handleOnSubmit: SubmitHandler<ICreateCostForm> = async (data) => {
    try {
      await createCost({
        costName: data.costName,
        totalValue: Number(data.totalValue),
        category: data.category,
        ...(data.groupId ? { groupId: data.groupId } : {}),
        ...(data.ownerPercentage ? { ownerPercentage: Number(data.ownerPercentage) } : {}),
      });

      router.push("/costs");
    } catch (error) {
      console.error("Error creating cost:", error);
    }
  };

  const handleCancel = () => {
    methods.reset();
    router.push("/costs");
  };

  const categoryOptions = [
    { label: "Dinner", value: TransactionCategory.DINNER },
    { label: "Lunch", value: TransactionCategory.LUNCH },
    { label: "Entertainment", value: TransactionCategory.ENTERTAINMENT },
    { label: "Travel", value: TransactionCategory.TRAVEL },
    { label: "Others", value: TransactionCategory.OTHERS },
  ];

  return (
    <div className="flex items-center justify-center min-h-screen">
      <FormProvider {...methods}>
        <form onSubmit={handleSubmit(handleOnSubmit)}>
          <Card className="w-[600px]">
            <CardHeader>
              <CardTitle className="text-center">Create Cost</CardTitle>
            </CardHeader>
            <CardContent className="space-y-4">
              <HookFormTextInput
                title="Cost Name"
                name="costName"
                label="Dinner at restaurant"
                type="text"
              />

              <HookFormSelect
                name="category"
                label="Category"
                placeholder="Select a category"
                groupLabel="Categories"
                options={categoryOptions}
              />

              <HookFormTextInput
                title="Total Value"
                name="totalValue"
                label="100.50"
                type="text"
              />

              <HookFormSelect
                name="groupId"
                label="Group (optional)"
                placeholder="Select a group"
                groupLabel="Groups"
                options={groupOptions}
              />

              <HookFormTextInput
                title="Your percentage (optional)"
                name="ownerPercentage"
                label="50"
                type="number"
              />
            </CardContent>
            <CardFooter className="flex flex-row justify-end">
              <Button
                type="button"
                className="w-1/2"
                variant="ghost"
                onClick={handleCancel}
              >
                Cancel
              </Button>
              <Button
                type="submit"
                className="w-1/2 transition-all duration-150 ease-in-out"
                variant="outline"
                disabled={isPending}
              >
                {isPending ? (
                  <>
                    <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                    Creating...
                  </>
                ) : (
                  "Add Cost"
                )}
              </Button>
            </CardFooter>
          </Card>
        </form>
      </FormProvider>
    </div>
  );
}
