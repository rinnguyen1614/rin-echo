import React, { useState } from "react";
import AppSidebar from "./AppSidebar";
import { AppContentView } from "../../../index";
import AppThemeSetting from "../../AppThemeSetting";
import AppHeader from "./AppHeader";
import clsx from "clsx";
import Box from "@mui/material/Box";
import MiniSidebarWrapper from "./MiniSidebarWrapper";
import AppFixedFooter from "./AppFixedFooter";
import { useLayoutContext } from "../../../utility/AppContextProvider/LayoutContextProvider";
import { LayoutType } from "../../../shared/constants/AppEnums";
import MiniSidebarContainer from "./MiniSidebarContainer";
import { LayoutProps } from "../LayoutProps";

const MiniSidebar = (props: LayoutProps) => {
  const { children, userMenus } = props;
  const [isCollapsed, setCollapsed] = useState(true);
  const { footer, layoutType, headerType, footerType } = useLayoutContext();

  return (
    <MiniSidebarContainer
      className={clsx({
        boxedLayout: layoutType === LayoutType.BOXED,
        framedLayout: layoutType === LayoutType.FRAMED,
      })}
    >
      <MiniSidebarWrapper
        className={clsx("miniSidebarWrapper", {
          "mini-sidebar-collapsed": isCollapsed,
          appMainFooter: footer && footerType === "fluid",
          appMainFixedFooter: footer && footerType === "fixed",
          appMainFixedHeader: headerType === "fixed",
        })}
      >
        <AppSidebar userMenus={userMenus} />
        <Box className="mainContent">
          <AppHeader setCollapsed={setCollapsed} isCollapsed={isCollapsed} />
          <AppContentView>{children}</AppContentView>
          <AppFixedFooter />
        </Box>
        <AppThemeSetting />
      </MiniSidebarWrapper>
    </MiniSidebarContainer>
  );
};

export default MiniSidebar;
