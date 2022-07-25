import React from "react";
import { toggleNavCollapsed } from "../../../../redux/actions";
import { useDispatch, useSelector } from "react-redux";
import clsx from "clsx";
import AppScrollbar from "../../../AppScrollbar";
import MainSidebar from "../../components/MainSidebar";
import Hidden from "@mui/material/Hidden";
import Drawer from "@mui/material/Drawer";
import VerticalNav from "../../components/VerticalNav";
import SidebarWrapper from "./SidebarWrapper";
import { useLayoutContext } from "../../../../utility/AppContextProvider/LayoutContextProvider";
import UserInfo from "../../components/UserInfo";
import { useSidebarContext } from "../../../../utility/AppContextProvider/SidebarContextProvider";
import { AppState } from "../../../../redux/store";

interface AppSidebarProps {
  position?: "left" | "top" | "right" | "bottom";
  variant?: string;
  userMenus?: React.ReactNode[];
}

const AppSidebar: React.FC<AppSidebarProps> = ({
  variant = "",
  position = "left",
  userMenus,
}) => {
  const dispatch = useDispatch();
  // @ts-ignore
  const navCollapsed = useSelector<AppState, AppState["settings"]>(
    ({ settings }) => settings
  ).navCollapsed;
  const { footer, footerType } = useLayoutContext();

  const { sidebarTextColor } = useSidebarContext();

  const handleToggleDrawer = () => {
    dispatch(toggleNavCollapsed());
  };

  return (
    <>
      <Hidden xlUp>
        <Drawer
          anchor={position}
          open={navCollapsed}
          onClose={() => handleToggleDrawer()}
          classes={{
            root: clsx(variant),
            paper: clsx(variant),
          }}
          style={{ position: "absolute" }}
        >
          <SidebarWrapper className="mini-toggle-sidebar">
            <MainSidebar>
              <UserInfo color={sidebarTextColor} userMenus={userMenus} />
              <AppScrollbar
                sx={{
                  py: 2,
                  height: "calc(100vh - 70px) !important",
                }}
                scrollToTop={false}
              >
                <VerticalNav />
              </AppScrollbar>
            </MainSidebar>
          </SidebarWrapper>
        </Drawer>
      </Hidden>
      <Hidden lgDown>
        <SidebarWrapper className="mini-toggle-sidebar">
          <MainSidebar>
            <UserInfo color={sidebarTextColor} userMenus={userMenus} />
            <AppScrollbar
              className={clsx({
                "has-footer-fixed": footer && footerType === "fixed",
              })}
              sx={{
                py: 2,
                height: "calc(100vh - 70px) !important",
                "&.has-footer-fixed": {
                  height: "calc(100vh - 117px) !important",
                },
              }}
              scrollToTop={false}
            >
              <VerticalNav />
            </AppScrollbar>
          </MainSidebar>
        </SidebarWrapper>
      </Hidden>
    </>
  );
};
export default AppSidebar;
