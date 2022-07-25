import { createContext } from "react";
import { CanAccessParams } from "./types";

export const PermissionContext = createContext<PermissionContextValue>({
  action: null,
  resource: null,
});

export type PermissionContextValue = CanAccessParams;
