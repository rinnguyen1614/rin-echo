import { useEffect } from "react";
import { usePermissions, useSafeSetState } from "react-admin";
import { canAccess } from "./canAccess";
import { CanAccessParams } from "./types";

interface State {
  isLoading: boolean;
  canAccess?: boolean;
  error?: Error;
}

const useCanAcess = (params: CanAccessParams): State => {
  const { resource, action, authParams, ...rest } = params;
  const [state, setState] = useSafeSetState<State>({
    isLoading: true,
  });
  const { permissions, error, isLoading } = usePermissions(authParams);

  useEffect(() => {
    if (!isLoading) {
      if (!error) {
        setState({
          isLoading: false,
          canAccess: canAccess({ permissions, resource, action }),
        });
      } else {
        setState({ isLoading: false, error: error });
      }
    }
  }, [permissions, resource, action, setState, error, isLoading]);

  return state;
};

export default useCanAcess;
