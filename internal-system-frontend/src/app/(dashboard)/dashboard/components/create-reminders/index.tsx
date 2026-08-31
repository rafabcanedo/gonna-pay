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
import * as yup from "yup";
import { toast } from "sonner";
import { Button } from "@/components/ui/button";
import { HookFormTextInput } from "@/components/hook-form-text-input";
import { FormProvider, useForm } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import { SelectDate } from "../select-date";
import { ICreateReminder, CreateRemindersProps } from '../interfaces'

const createReminderSchema = yup.object().shape({
    name: yup
        .string()
        .required("Reminder name is required")
        .min(3, "Name must be at least 3 characters"),
    value: yup
        .string()
        .required("Value is required")
        .matches(/^\d+([.,]\d{2})?$/, "Invalid value format"),
    date: yup
        .date()
        .required("Date is required")
        .typeError("Please select a valid date")
        .min(new Date(), "Date must be in the future"),
});

export const CreateReminders = ({ onAddReminder }: CreateRemindersProps) => {
    const [openModal, setOpenModal] = useState(false);

    const methods = useForm<ICreateReminder>({
        resolver: yupResolver(createReminderSchema),
        mode: "onSubmit",
        defaultValues: {
            name: "",
            value: "",
            date: new Date(),
        },
    });

    const handleOnSubmit = async (data: ICreateReminder) => {
        try {
            onAddReminder(data);
            toast.success("Reminders created successfully!");
            handleFinish();
        } catch (error) {
            toast.error("Failed to create reminders.");
        }
    };

    const handleFinish = () => {
        methods.reset();
        setOpenModal(false);
    };

    return (
        <Dialog open={openModal} onOpenChange={setOpenModal}>
            <DialogTrigger asChild>
                <Button variant="outline">Create Reminder</Button>
            </DialogTrigger>

            <DialogContent>
                <FormProvider {...methods}>
                    <form onSubmit={methods.handleSubmit(handleOnSubmit)}>
                        <DialogHeader>
                            <DialogTitle>Add Reminder</DialogTitle>
                            <DialogDescription>
                                Create a new Reminder
                            </DialogDescription>
                        </DialogHeader>

                        <div className="flex flex-col gap-4 py-4">
                            <HookFormTextInput
                                title="Reminder name"
                                name="name"
                                label="Pay the Internet"
                                type="text"
                            />
                            <HookFormTextInput
                                title="Value"
                                name="value"
                                label="20,00"
                                type="text"
                            />
                            <SelectDate />
                        </div>

                        <div className="flex justify-end">
                            <Button
                                className="w-1/3"
                                variant="outline"
                                type="submit"
                            >
                                Create
                            </Button>
                        </div>
                    </form>
                </FormProvider>
            </DialogContent>
        </Dialog>
    )
}