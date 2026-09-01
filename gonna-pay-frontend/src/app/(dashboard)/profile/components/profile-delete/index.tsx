"use client"
import { AlertDialog, AlertDialogAction, AlertDialogCancel, AlertDialogContent, AlertDialogDescription, AlertDialogFooter, AlertDialogHeader, AlertDialogTitle, AlertDialogTrigger } from "@/components/ui/alert-dialog"
import { Button } from "@/components/ui/button"
import { Card, CardHeader, CardTitle, CardContent, CardFooter } from "@/components/ui/card"
import { useDeleteAccountMutation } from "@/hooks/mutations/use-delete-account-mutation"

export const ProfileDelete = () => {
    const { mutate: deleteAccount, isPending } = useDeleteAccountMutation()

    return (
        <Card className="w-[800px] border-red-200">
            <CardHeader>
                <CardTitle className="text-center text-red-500">Delete Account</CardTitle>
            </CardHeader>
            <CardContent className="text-zinc-600">
                <p className="text-sm">
                    Deleting your account is <span className="font-semibold text-red-400">permanent</span>.
                    You will lose all your data, settings, and progress. Are you sure you want to proceed?
                </p>
            </CardContent>
            <CardFooter className="flex justify-end">
                <AlertDialog>
                    <AlertDialogTrigger asChild>
                        <Button variant="ghost" className="w-40 text-red-400 hover:bg-red-300 hover:text-stone-50">
                            Delete Account
                        </Button>
                    </AlertDialogTrigger>
                    <AlertDialogContent>
                        <AlertDialogHeader>
                            <AlertDialogTitle>Are you sure you want to delete your account?</AlertDialogTitle>
                            <AlertDialogDescription>
                                This action cannot be undone. All your data will be permanently removed.
                            </AlertDialogDescription>
                        </AlertDialogHeader>
                        <AlertDialogFooter>
                            <AlertDialogCancel>Cancel</AlertDialogCancel>
                            <AlertDialogAction
                                className="bg-red-600 hover:bg-red-700"
                                disabled={isPending}
                                onClick={() => deleteAccount()}
                            >
                                Delete
                            </AlertDialogAction>
                        </AlertDialogFooter>
                    </AlertDialogContent>
                </AlertDialog>
            </CardFooter>
        </Card>
    )
}