"use client";

import { FC } from "react";
import Link from "next/link";
import { BadgeType } from "@/utils/badge-types";
import { getInitials } from '@/utils/get-initials'
import { IPropsGroupCards } from './interfaces'

export const GroupCards: FC<IPropsGroupCards> = ({ group }) => {
  return (
    <Link href={`/groups/details/${group.id}?name=${encodeURIComponent(group.name)}`}>
      <div className="w-full shadow rounded-xl px-5 py-4 flex flex-col gap-4 hover:shadow-md transition-shadow cursor-pointer">
        <div className="flex flex-row items-center justify-between">
          <span className="font-poppins font-semibold text-lg text-zinc-700">
            {group.name}
          </span>
          <BadgeType type={group.category} />
        </div>
        <div className="flex flex-row flex-wrap items-center gap-2">
          {(group.members ?? []).map((member) => (
            <div
              key={member.id}
              title={member.name}
              className="flex items-center justify-center w-9 h-9 rounded-full bg-zinc-100 border border-zinc-300 text-zinc-600 font-poppins text-sm font-medium cursor-default select-none"
            >
              {getInitials(member.name)}
            </div>
          ))}
        </div>
      </div>
    </Link>
  );
};
