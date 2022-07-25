import { ReactElement } from "react";

export const Actions = {
  list: "get_list",
  create: "create",
  edit: "edit",
  delete: "delete",
  show: "get",
  export: "export",
};

export interface CanAccessParams {
  resource: string;
  action: string;
  authParams?: object;
  record?: object;
}

export interface WithPermissionsProps {
  action: string;
  resource?: string;
  children: ReactElement;
  authParams?: object;
  [key: string]: any;
}
