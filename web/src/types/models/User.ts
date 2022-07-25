import { Record } from "@app/types";

export interface User extends Record {
  username: string;
  avatar_path: string;
  full_name: string;
  email: string;
  email_verified: boolean;
  phone: string;
  phone_verified: boolean;
  gender: number;
  date_of_birth: Date;
}
