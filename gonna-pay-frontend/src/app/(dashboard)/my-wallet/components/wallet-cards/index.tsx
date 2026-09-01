"use client";

import { FC, useState } from "react";
import { Eye, EyeOff } from "lucide-react";
import { IPropsWalletCards } from './interfaces'

export const WalletCards: FC<IPropsWalletCards> = ({ title, value }) => {
  const [onSee, setOnSee] = useState(true);

  const handleOnSee = () => {
    setOnSee(!onSee);
  };

  return (
    <div className="h-20 w-64 shadow rounded-xl px-4 mt-4">
      <div className="flex flex-row itens-center justify-between">
        <span className="font-roboto font-light text-lg text-zinc-500">
          {title}
        </span>
        <button onClick={handleOnSee} className="cursor-pointer">
          {onSee ? (
            <Eye className="w-4 h-4 text-zinc-500" />
          ) : (
            <EyeOff className="w-4 h-4 text-zinc-500" />
          )}
        </button>
      </div>
      <div className="flex flex-row gap-2">
        <span className="font-poppins text-2xl text-zinc-600">$</span>
        <span className="font-poppins text-2xl text-zinc-600">
          {onSee ? `${value}` : "****"}
        </span>
      </div>
    </div>
  );
};
