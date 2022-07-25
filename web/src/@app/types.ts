import { RaRecord } from "ra-core";
import { Identifier, UserIdentity } from "react-admin";

export interface Identity extends UserIdentity {
  email: string;
  email_verified: boolean;
  phone: string;
  phone_vefidied: boolean;
  avatar_path: string;
  username: string;
  full_name: string;
  gender: string;
  date_of_birth: Date;
}

export type Record = RaRecord;

export interface RecordMap<RecordType extends Record = Record> {
  // Accept strings and numbers as identifiers
  [id: string]: RecordType;
  [id: number]: RecordType;
}

export interface Tree extends Record {
  name: string;
  children?: Tree[];
  parent_id?: Identifier;
  all_parent_ids: Identifier[];
  all_children_ids: Identifier[];
  [key: string]: any;
}
