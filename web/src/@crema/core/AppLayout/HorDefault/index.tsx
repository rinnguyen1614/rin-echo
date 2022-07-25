import React from "react";
import clsx from "clsx";
import AppContentView from "@crema/core/AppContentView";
import AppFixedFooter from "./AppFixedFooter";
import AppHeader from "./AppHeader";
import AppSidebar from "./AppSidebar";
import { useLayoutContext } from "../../../utility/AppContextProvider/LayoutContextProvider";
import AppThemeSetting from "../../AppThemeSetting";
import HorDefaultWrapper from "./HorDefaultWrapper";
import MainContent from "./MainContent";
import { LayoutType } from "../../../shared/constants/AppEnums";
import HorDefaultContainer from "./HorDefaultContainer";
import { LayoutProps } from "../LayoutProps";

const HorDefault = (props: LayoutProps) => {
  const { children, userMenus } = props;
  const { footer, layoutType, footerType } = useLayoutContext();

  return (
    <HorDefaultContainer
      className={clsx({
        boxedLayout: layoutType === LayoutType.BOXED,
        framedLayout: layoutType === LayoutType.FRAMED,
      })}
    >
      <HorDefaultWrapper
        className={clsx("horDefaultWrapper", {
          appMainFooter: footer && footerType === "fluid",
          appMainFixedFooter: footer && footerType === "fixed",
        })}
      >
        <AppSidebar userMenus={userMenus} />
        <MainContent>
          <AppHeader userMenus={userMenus} />
          <AppContentView>{children}</AppContentView>
          <AppFixedFooter />
        </MainContent>
        <AppThemeSetting />
      </HorDefaultWrapper>
    </HorDefaultContainer>
  );
};

export default HorDefault;
