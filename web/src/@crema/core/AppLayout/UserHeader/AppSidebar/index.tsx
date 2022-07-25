import React from "react";
import { toggleNavCollapsed } from "../../../../redux/actions";
import { useDispatch, useSelector } from "react-redux";
import clsx from "clsx";
import AppScrollbar from "../../../AppScrollbar";
import MainSidebar from "../../components/MainSidebar";
import Hidden from "@mui/material/Hidden";
import Drawer from "@mui/material/Drawer";
import VerticalNav from "../../components/VerticalNav";
import UserHeaderSidebarWrapper from "./UserHeaderSidebarWrapper";
import { useLayoutContext } from "../../../../utility/AppContextProvider/LayoutContextProvider";
import { AppState } from "../../../../redux/store";

interface AppSidebarProps {
  position?: "left" | "top" | "right" | "bottom";
  variant?: string;
}

const AppSidebar: React.FC<AppSidebarProps> = ({
  variant = "",
  position = "left",
}) => {
  const dispatch = useDispatch();
  // @ts-ignore
  const navCollapsed = useSelector<AppState, AppState["settings"]>(
    ({ settings }) => settings
  ).navCollapsed;
  const { footer, footerType } = useLayoutContext();

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
          <UserHeaderSidebarWrapper className="user-header-sidebar">
            <MainSidebar>
              <AppScrollbar
                sx={{
                  py: 2,
                }}
                scrollToTop={false}
              >
                <VerticalNav />
              </AppScrollbar>
            </MainSidebar>
          </UserHeaderSidebarWrapper>
        </Drawer>
      </Hidden>
      <Hidden lgDown>
        <UserHeaderSidebarWrapper className="user-header-sidebar">
          <MainSidebar>
            <AppScrollbar
              className={clsx({
                "has-footer-fixed": footer && footerType === "fixed",
              })}
              sx={{
                py: 2,
                height: "calc(100vh - 71px) !important",
                "&.has-footer-fixed": {
                  height: {
                    xs: "calc(100vh - 119px) !important",
                    xl: "calc(100vh - 131px) !important",
                  },
                },
              }}
              scrollToTop={false}
            >
              <VerticalNav />
            </AppScrollbar>
          </MainSidebar>
        </UserHeaderSidebarWrapper>
      </Hidden>
    </>
  );
};
export default AppSidebar;
