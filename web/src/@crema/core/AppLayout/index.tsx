import React, { useEffect } from "react";
import {
  useLayoutActionsContext,
  useLayoutContext,
} from "../../utility/AppContextProvider/LayoutContextProvider";
import Layouts from "./Layouts";
import { useUrlSearchParams } from "use-url-search-params";
import { useSidebarActionsContext } from "../../utility/AppContextProvider/SidebarContextProvider";
import { useLocation } from "react-router-dom";
import { useDispatch } from "react-redux";
import { onNavCollapsed } from "../../redux/actions";
import { LayoutProps } from "./LayoutProps";

const AppLayout = (props: LayoutProps) => {
  const { children, userMenus } = props;
  const { navStyle } = useLayoutContext();
  const { pathname } = useLocation();
  const dispatch = useDispatch();
  const { updateNavStyle } = useLayoutActionsContext();
  const { updateMenuStyle, setSidebarBgImage } = useSidebarActionsContext();
  const AppLayout = Layouts[navStyle];
  const [params] = useUrlSearchParams({}, {});

  useEffect(() => {
    if (params.layout) updateNavStyle(params.layout as string);
    if (params.menuStyle) updateMenuStyle(params.menuStyle as string);
    if (params.sidebarImage) setSidebarBgImage(true);
  }, [params, setSidebarBgImage, updateNavStyle, updateMenuStyle]);

  useEffect(() => {
    dispatch(onNavCollapsed());
  }, [dispatch, pathname]);

  return <AppLayout userMenus={userMenus}>{children}</AppLayout>;
};

export default React.memo(AppLayout);
