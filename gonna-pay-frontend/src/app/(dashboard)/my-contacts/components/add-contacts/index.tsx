"use client";

import { useState } from "react";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog"
import { Loader2 } from "lucide-react"
import { Button } from "@/components/ui/button";
import { ContactCategory, CreateContactInput } from "@/types";
import { HookFormTextInput } from "@/components/hook-form-text-input";
import { SelectCategory } from "../select-category";
import { addContactSchema } from "@/validations/schemas";
import { FormProvider, useForm } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import { ApiError } from "@/lib/errors/api.error";
import { useCreateContact } from "@/hooks/mutations/use-contact-mutations";

export const AddContact = () => {
  const [stepModal, setStepModal] = useState(1);
  const [openModal, setOpenModal] = useState(false);
  const [isValidating, setIsValidating] = useState(false);

  const { mutateAsync: createContact, isPending } = useCreateContact();

  const methods = useForm<CreateContactInput>({
    resolver: yupResolver(addContactSchema),
    mode: "onSubmit",
    defaultValues: {
      name: "",
      email: "",
      phone: "",
      category: ContactCategory.FRIEND,
    },
  });

  const handleContinue = async () => {
    setIsValidating(true);
    const isValid = await methods.trigger(["name", "email"]);

    if (!isValid) {
      setIsValidating(false);
      return;
    }

    await new Promise((resolve) => setTimeout(resolve, 800));
    setStepModal(2);
    setIsValidating(false);
  };

  const handleOnSubmit = async (data: CreateContactInput) => {

    try {
      await createContact(data);

      handleFinish();

    } catch (error) {
      if (error instanceof ApiError && error.status === 409) {
        methods.setError("email", {
          type: "manual",
          message: "This email is already registered",
        });
        setStepModal(1);
      }
    }
  };

  const handleFinish = () => {
    methods.reset();
    setStepModal(1);
    setOpenModal(false);
  };

  const handleCancel = () => {
    methods.reset();
    setStepModal(1);
    setOpenModal(false);
  };

  const handleBack = () => {
    setStepModal(1);
  };

  return (
    <Dialog open={openModal} onOpenChange={setOpenModal}>
      <DialogTrigger asChild>
        <Button variant="outline">Create Contact</Button>
      </DialogTrigger>

      <DialogContent>
        <FormProvider {...methods}>
          <form onSubmit={methods.handleSubmit(handleOnSubmit, (errors) => console.log("Erro de validação:", errors))}>
            {stepModal === 1 && (
              <>
                <DialogHeader>
                  <DialogTitle>Add Contact</DialogTitle>
                  <DialogDescription>
                    Fill in the form to add a contact
                  </DialogDescription>
                </DialogHeader>

                <div className="flex flex-col gap-4 py-4">
                  <HookFormTextInput
                    title="Name"
                    name="name"
                    label="John Jason"
                    type="text"
                  />
                  <HookFormTextInput
                    title="Email"
                    name="email"
                    label="john@email.com"
                    type="email"
                  />

                  <div className="flex flex-row gap-2">
                    <Button
                      className="w-1/2"
                      variant="outline"
                      type="button"
                      onClick={handleCancel}
                    >
                      Cancel
                    </Button>
                    <Button
                      className="w-1/2"
                      type="button"
                      onClick={handleContinue}
                      disabled={isValidating || isPending}
                    >
                      {isValidating ? (
                        <>
                          <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                        </>
                      ) : (
                        "Continue"
                      )}
                    </Button>
                  </div>
                </div>
              </>
            )}

            {stepModal === 2 && (
              <>
                <DialogHeader>
                  <DialogTitle>Add Contact</DialogTitle>
                  <DialogDescription>
                    Complete the contact information
                  </DialogDescription>
                </DialogHeader>

                <div className="flex flex-col gap-4 py-4">
                  <HookFormTextInput
                    title="Phone"
                    name="phone"
                    label="+55 11 997117911"
                    type="text"
                  />
                  <SelectCategory
                    value={methods.watch("category")}
                    onValueChange={(value: ContactCategory) =>
                      methods.setValue("category", value, { shouldValidate: true })
                    }
                  />

                  <div className="flex flex-row gap-2">
                    <Button
                      className="w-1/3"
                      variant="outline"
                      type="button"
                      onClick={handleBack}
                      disabled={isPending}
                    >
                      Back
                    </Button>
                    <Button
                      className="w-2/3"
                      type="submit"
                      disabled={isPending}
                    >
                      {isPending ? (
                        <>
                          <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                          Creating...
                        </>
                      ) : (
                        "Create Contact"
                      )}
                    </Button>
                  </div>
                </div>
              </>
            )}
          </form>
        </FormProvider>
      </DialogContent>
    </Dialog>
  )
}