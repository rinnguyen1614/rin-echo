import React, { ReactNode } from "react";
import ThemeContextProvider from "./ThemeContextProvider";
import LocaleContextProvider from "./LocaleContextProvide";
import LayoutContextProvider from "./LayoutContextProvider";
import SidebarContextProvider from "./SidebarContextProvider";

const AppContextProvider = ({ children }) => {
  return (
    <ThemeContextProvider>
      <LocaleContextProvider>
        <LayoutContextProvider>
          <SidebarContextProvider>{children}</SidebarContextProvider>
        </LayoutContextProvider>
      </LocaleContextProvider>
    </ThemeContextProvider>
  );
};

export default AppContextProvider;
