"use client"

import { useUser } from "@/providers/contexts/user-context"
import { ProfileAvatar } from "./components/profile-avatar"
import { ProfileForm } from "./components/profile-form"
import { Button } from "@/components/ui/button"
import { useLogoutMutation } from "@/hooks/mutations/use-logout-mutations"
import { ProfileDelete } from "./components/profile-delete"

export default function Profile() {

    const { user, isLoading } = useUser()
    const logoutMutation = useLogoutMutation()

    return (
        <div>
            <header className="flex flex-row items-center mt-8">
                <div>
                    <ProfileAvatar />
                </div>

                <div className="ml-4">
                    <span>
                        Hey {isLoading ? "..." : user?.name}, welcome to your Profile session
                    </span>
                </div>
            </header>
            <div className="flex mt-8">
                <ProfileForm />
            </div>

            <div className="flex mt-8">
                <ProfileDelete />
            </div>

            <div className="mt-10">
                <Button
                    variant="outline"
                    onClick={() => logoutMutation.mutate()}
                    disabled={logoutMutation.isPending}
                    className="bg-red-400 text-zinc-100 hover:bg-red-300 hover:text-zinc-100"
                >
                    {logoutMutation.isPending ? "Logging out..." : "Logout"}
                </Button>
            </div>
        </div>
    )
}