import React from "react";
import { toggleNavCollapsed } from "../../../../redux/actions";
import { useDispatch, useSelector } from "react-redux";
import clsx from "clsx";
import AppScrollbar from "../../../AppScrollbar";
import MainSidebar from "../../components/MainSidebar";
import Box from "@mui/material/Box";
import Hidden from "@mui/material/Hidden";
import Drawer from "@mui/material/Drawer";
import VerticalNav from "../../components/VerticalNav";
import NavigateBeforeIcon from "@mui/icons-material/NavigateBefore";
import NavigateNextIcon from "@mui/icons-material/NavigateNext";
import BitBucketSidebarWrapper from "./BitBucketSidebarWrapper";
import AppSidebarContainer from "./AppSidebarContainer";
import BucketMinibar from "./BucketMinibar";
import { Typography } from "@mui/material";
import { Fonts } from "../../../../shared/constants/AppEnums";
import { AppState } from "../../../../redux/store";

interface AppSidebarProps {
  position?: "left" | "top" | "right" | "bottom";
  variant?: string;
  isCollapsed: boolean;
  setCollapsed: (collapsed: boolean) => void;
  userMenus?: React.ReactNode[];
}

const AppSidebar: React.FC<AppSidebarProps> = (props) => {
  const {
    isCollapsed,
    setCollapsed,
    variant = "",
    position = "left",
    userMenus,
  } = props;

  const dispatch = useDispatch();
  // @ts-ignore
  const navCollapsed = useSelector<AppState, AppState["settings"]>(
    ({ settings }) => settings
  ).navCollapsed;

  const handleToggleDrawer = () => {
    dispatch(toggleNavCollapsed());
  };

  const sideBarComponent = () => {
    return (
      <BitBucketSidebarWrapper className="bit-bucket-sidebar">
        <Box className="bit-bucket-sidebar-fixed">
          <Box
            className="bit-bucket-btn"
            onClick={() => setCollapsed(!isCollapsed)}
          >
            {isCollapsed ? <NavigateNextIcon /> : <NavigateBeforeIcon />}
          </Box>
          <BucketMinibar userMenus={userMenus} />
          <AppSidebarContainer className="app-sidebar-container">
            <MainSidebar>
              <Box
                sx={{
                  py: 4.5,
                  px: 7.5,
                }}
              >
                <Typography
                  sx={{
                    fontSize: 22,
                    fontWeight: Fonts.MEDIUM,
                  }}
                  component="h2"
                >
                  Crema
                </Typography>
              </Box>
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
          </AppSidebarContainer>
        </Box>
      </BitBucketSidebarWrapper>
    );
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
          {sideBarComponent()}
        </Drawer>
      </Hidden>
      <Hidden lgDown>{sideBarComponent()}</Hidden>
    </>
  );
};
export default AppSidebar;
