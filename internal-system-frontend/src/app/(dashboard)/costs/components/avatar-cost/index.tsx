import { Avatar, AvatarFallback } from "@/components/ui/avatar"
import { IAvatarCost } from './interfaces'

export const AvatarCost = ({ name }: IAvatarCost) => {

    const firstTwoLetters = name.trim().slice(0, 2).toUpperCase()

    return (
        <div className="flex flex-col items-center gap-2">
            <Avatar>
                <AvatarFallback>{firstTwoLetters}</AvatarFallback>
            </Avatar>
            <span className="text-sm font-medium text-center">{name}</span>
        </div>
    )
}