import { Record } from "@app/types";
import { Identifier } from "react-admin";

export interface Menu extends Record {
  name: string;
  children?: Menu[];
  parent_id?: Identifier;
  slug: string;
  method: string;
  path: string;
}
