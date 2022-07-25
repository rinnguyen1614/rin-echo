import { Record, Tree } from "@app/types";
import { Identifier } from "react-admin";

export interface Resource extends Record {
  name: string;
  children?: Tree[];
  parent_id?: Identifier;
  slug: string;
  method: string;
  path: string;
}
