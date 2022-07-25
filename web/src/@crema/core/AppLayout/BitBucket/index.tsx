import React, { useState } from "react";
import AppSidebar from "./AppSidebar";
import { AppContentView } from "../../../index";
import AppThemeSetting from "../../AppThemeSetting";
import AppHeader from "./AppHeader";
import clsx from "clsx";
import Hidden from "@mui/material/Hidden";
import Box from "@mui/material/Box";
import BitBucketWrapper from "./BitBucketWrapper";
import { LayoutType } from "../../../shared/constants/AppEnums";
import { useLayoutContext } from "../../../utility/AppContextProvider/LayoutContextProvider";
import BitBucketContainer from "./BitBucketContainer";
import { LayoutProps } from "../LayoutProps";

const BitBucket = (props: LayoutProps) => {
  const { children, userMenus } = props;
  const [isCollapsed, setCollapsed] = useState(false);
  const { layoutType } = useLayoutContext();

  return (
    <BitBucketContainer
      className={clsx({
        boxedLayout: layoutType === LayoutType.BOXED,
        framedLayout: layoutType === LayoutType.FRAMED,
      })}
    >
      <BitBucketWrapper
        className={clsx("bitBucketWrapper", {
          bitBucketCollapsed: isCollapsed,
        })}
      >
        <Hidden lgUp>
          <AppHeader />
        </Hidden>
        <AppSidebar
          isCollapsed={isCollapsed}
          setCollapsed={setCollapsed}
          userMenus={userMenus}
        />
        <Box className="mainContent">
          <AppContentView>{children}</AppContentView>
        </Box>
        <AppThemeSetting />
      </BitBucketWrapper>
    </BitBucketContainer>
  );
};

export default BitBucket;
