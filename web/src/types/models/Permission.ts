import { Record } from "@app/types";

export interface Permission extends Record {
  name: string;
  actions: string[];
}
