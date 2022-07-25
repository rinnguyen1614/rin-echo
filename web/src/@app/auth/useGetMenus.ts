import { useCallback } from "react";
import { useAuthProvider } from "react-admin";

const getMenusWithoutProvider = () => Promise.resolve([]);

const useGetMenus = (): GetMenus => {
  const authProvider = useAuthProvider();
  const getMenus = useCallback(
    (params: any = {}) => authProvider.getMenus(params),
    [authProvider]
  );

  return authProvider ? getMenus : getMenusWithoutProvider;
};

type GetMenus = (params?: any) => Promise<any>;

export default useGetMenus;
