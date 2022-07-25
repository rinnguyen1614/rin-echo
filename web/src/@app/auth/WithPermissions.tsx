import { Children } from "react";
import { useAuthenticated, useResourceContext } from "react-admin";
import { PermissionContextProvider } from "./PermissionContextProvider";
import { WithPermissionsProps } from "./types";
import useCanAcess from "./useCanAccess";

export const WithPermissions = (props: WithPermissionsProps) => {
  const { action, children, authParams } = props;
  const resource = useResourceContext(props);

  useAuthenticated(authParams);

  const { canAccess } = useCanAcess({ action, resource, authParams });
  const contextValue = { action, resource };

  if (Children.count(children) !== 0 && canAccess)
    return (
      <PermissionContextProvider value={contextValue}>
        {children}
      </PermissionContextProvider>
    );
  return null;
};
