import { ReactElement } from "react";
import { PermissionContext, PermissionContextValue } from "./PermissionContext";

export const PermissionContextProvider = ({
  children,
  value,
}: {
  children: ReactElement;
  value: PermissionContextValue;
}) => (
  <PermissionContext.Provider value={value}>
    {children}
  </PermissionContext.Provider>
);
