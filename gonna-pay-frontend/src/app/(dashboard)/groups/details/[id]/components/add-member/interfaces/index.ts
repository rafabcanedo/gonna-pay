import { GroupMember } from "@/types";

export interface IAddMemberProps {
  groupId: string;
  currentMembers?: GroupMember[];
}
