import { useCallback } from "react";
import { useAuthProvider } from "react-admin";
import { Identity } from "../types";

const useUpdateIdentity = (): ((identity: Identity) => Promise<Identity>) => {
  const authProvider = useAuthProvider();

  const updateIdentity = useCallback(
    (params: any = {}) => authProvider.updateIdentity(params),
    [authProvider]
  );

  return updateIdentity;
};

export default useUpdateIdentity;
