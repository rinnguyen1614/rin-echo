import { useEffect } from "react";
import { useSafeSetState } from "react-admin";
import useGetMenus from "./useGetMenus";

interface State<Menus, Error> {
  isLoading: boolean;
  menus?: Menus;
  error?: Error;
}

const emptyParams = {};

const useMenus = <Menus = any, Error = any>(
  params = emptyParams
): State<Menus, Error> => {
  const [state, setState] = useSafeSetState<State<Menus, Error>>({
    isLoading: true,
  });
  const getMenus = useGetMenus();
  useEffect(() => {
    getMenus(params)
      .then((menus) => {
        setState({ isLoading: false, menus });
      })
      .catch((error) => {
        setState({
          isLoading: false,
          error,
        });
      });
  }, [getMenus, params, setState]);
  return state;
};

export default useMenus;
