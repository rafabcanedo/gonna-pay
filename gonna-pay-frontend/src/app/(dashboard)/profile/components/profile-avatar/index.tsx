import {
    Avatar,
    AvatarFallback,
    AvatarBadge
} from "@/components/ui/avatar"

export function ProfileAvatar() {
    return (
        <Avatar size="lg">
            <AvatarFallback>RC</AvatarFallback>
            <AvatarBadge className="bg-green-600 dark:bg-green-800" />
        </Avatar>
    )
}
