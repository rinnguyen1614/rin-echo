import React from "react";
import Drawer from "@mui/material/Drawer";
import Hidden from "@mui/material/Hidden";
import clsx from "clsx";
import { toggleNavCollapsed } from "../../../redux/actions";
import { useDispatch, useSelector } from "react-redux";
import AppScrollbar from "../../AppScrollbar";
import VerticalNav from "../components/VerticalNav";
import MainSidebar from "../components/MainSidebar";
import { useLayoutContext } from "../../../utility/AppContextProvider/LayoutContextProvider";
import UserInfo from "../components/UserInfo";
import { useSidebarContext } from "../../../utility/AppContextProvider/SidebarContextProvider";
import { AppState } from "../../../redux/store";

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
      <Hidden lgUp>
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
          <MainSidebar>
            <UserInfo color={sidebarTextColor} userMenus={userMenus} />
            <AppScrollbar
              sx={{
                py: 2,
                height: "calc(100vh - 70px) !important",
                borderTop: (theme: { palette: { divider: string } }) =>
                  `solid 1px ${theme.palette.divider}`,
                mt: 0.5,
              }}
            >
              <VerticalNav />
            </AppScrollbar>
          </MainSidebar>
        </Drawer>
      </Hidden>
      <Hidden lgDown>
        <MainSidebar>
          <UserInfo color={sidebarTextColor} userMenus={userMenus} />
          <AppScrollbar
            className={clsx({
              "has-footer-fixed": footer && footerType === "fixed",
            })}
            sx={{
              py: 2,
              height: "calc(100vh - 70px) !important",
              borderTop: (theme: { palette: { divider: string } }) =>
                `solid 1px ${theme.palette.divider}`,
              mt: 0.5,
              "&.has-footer-fixed": {
                height: {
                  xs: "calc(100vh - 117px) !important",
                  xl: "calc(100vh - 127px) !important",
                },
              },
            }}
          >
            <VerticalNav />
          </AppScrollbar>
        </MainSidebar>
      </Hidden>
    </>
  );
};

export default AppSidebar;
