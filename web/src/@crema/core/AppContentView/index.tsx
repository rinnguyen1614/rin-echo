import Box from "@mui/material/Box";
import { SxProps } from "@mui/system";
import React, { FC, ReactNode } from "react";
import { AppSuspense } from "../../index";
import AppErrorBoundary from "../AppErrorBoundary";
import AppFooter from "../AppLayout/components/AppFooter";
import AppContentViewWrapper from "./AppContentViewWrapper";

interface AppContentViewProps {
  sxStyle?: SxProps;
  children?: ReactNode;
}

const AppContentView: FC<AppContentViewProps> = ({ sxStyle, children }) => {
  return (
    <AppContentViewWrapper>
      <Box
        sx={{
          display: "flex",
          flex: 1,
          flexDirection: "column",
          p: { xs: 5, md: 7.5, xl: 12.5 },
          ...sxStyle,
        }}
        className="app-content"
      >
        <AppSuspense>
          <AppErrorBoundary>{children}</AppErrorBoundary>
        </AppSuspense>
      </Box>
      <AppFooter />
    </AppContentViewWrapper>
  );
};

export default AppContentView;
