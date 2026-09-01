"use client";

import { useState } from "react";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Loader2 } from "lucide-react";
import { Button } from "@/components/ui/button";
import { TransactionCategory, CreateGroupInput } from "@/types";
import { HookFormTextInput } from "@/components/hook-form-text-input";
import { createGroupSchema } from "@/validations/schemas";
import { FormProvider, useForm } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import { useCreateGroup } from "@/hooks/mutations/use-group-mutations";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectLabel,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";

export const CreateGroup = () => {
  const [openModal, setOpenModal] = useState(false);

  const { mutateAsync: createGroup, isPending } = useCreateGroup();

  const methods = useForm<CreateGroupInput>({
    resolver: yupResolver(createGroupSchema),
    mode: "onSubmit",
    defaultValues: {
      name: "",
      category: TransactionCategory.OTHERS,
    },
  });

  const handleOnSubmit = async (data: CreateGroupInput) => {
    try {
      await createGroup(data);
      methods.reset();
      setOpenModal(false);
    } catch {
      // error handled by mutation hook via toast
    }
  };

  const handleCancel = () => {
    methods.reset();
    setOpenModal(false);
  };

  return (
    <Dialog open={openModal} onOpenChange={setOpenModal}>
      <DialogTrigger asChild>
        <Button variant="outline">Create Group</Button>
      </DialogTrigger>

      <DialogContent>
        <FormProvider {...methods}>
          <form onSubmit={methods.handleSubmit(handleOnSubmit)}>
            <DialogHeader>
              <DialogTitle>Create Group</DialogTitle>
              <DialogDescription>
                Fill in the form to create a group
              </DialogDescription>
            </DialogHeader>

            <div className="flex flex-col gap-4 py-4">
              <HookFormTextInput
                title="Name"
                name="name"
                label="My group name"
                type="text"
              />

              <Select
                value={methods.watch("category")}
                onValueChange={(val) =>
                  methods.setValue("category", val as TransactionCategory, {
                    shouldValidate: true,
                  })
                }
              >
                <SelectTrigger className="w-full">
                  <SelectValue placeholder="Select a category" />
                </SelectTrigger>
                <SelectContent>
                  <SelectGroup>
                    <SelectLabel>Category</SelectLabel>
                    {Object.values(TransactionCategory).map((cat) => (
                      <SelectItem key={cat} value={cat}>
                        {cat}
                      </SelectItem>
                    ))}
                  </SelectGroup>
                </SelectContent>
              </Select>

              <div className="flex flex-row gap-2">
                <Button
                  className="w-1/2"
                  variant="outline"
                  type="button"
                  onClick={handleCancel}
                  disabled={isPending}
                >
                  Cancel
                </Button>
                <Button
                  className="w-1/2"
                  type="submit"
                  disabled={isPending}
                >
                  {isPending ? (
                    <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                  ) : (
                    "Create Group"
                  )}
                </Button>
              </div>
            </div>
          </form>
        </FormProvider>
      </DialogContent>
    </Dialog>
  );
};
