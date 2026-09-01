import {
    Avatar,
    AvatarFallback,
    AvatarImage,
    AvatarBadge
} from "@/components/ui/avatar"

export function ProfileAvatar() {
    return (
        <Avatar size="lg">
            <AvatarImage
                src="https://github.com/shadcn.png"
                alt="@shadcn"
            />
            <AvatarFallback>RC</AvatarFallback>
            <AvatarBadge className="bg-green-600 dark:bg-green-800" />
        </Avatar>
    )
}
