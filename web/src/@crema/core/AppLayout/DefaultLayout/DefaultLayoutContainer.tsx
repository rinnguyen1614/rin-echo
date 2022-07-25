import React, { ReactNode } from "react";
import { Box } from "@mui/material";

interface DefaultLayoutContainerProps {
  children: ReactNode;

  [x: string]: any;
}

const DefaultLayoutContainer: React.FC<DefaultLayoutContainerProps> = ({
  children,
  ...rest
}) => {
  return (
    <Box
      sx={{
        minHeight: "100vh",
        display: "flex",
        flexDirection: "column",
        position: "relative",
        backgroundColor: (theme) => theme.palette.background.default,
        "&.boxedLayout": {
          maxWidth: { xl: 1480 },
          mx: { xl: "auto" },
          boxShadow: "none",
          borderLeft: "1px solid #e8e5dd",
          borderRight: "1px solid #e8e5dd",
          pt: { xl: 0 },
          "& .app-sidebar": {
            position: { xl: "sticky" },
            "& [data-simplebar]": {
              height: { xl: "calc(100vh - 70px) !important" },
            },
          },
          "& .appMainFixedHeader": {
            pt: { xl: 0 },
            "& .app-bar": {
              position: { xl: "sticky" },
              width: { xl: "100%" },
            },
          },
          "& .mainContent": {
            position: { xl: "static" },
            ml: { xl: 0 },
          },
          "& .fixed-footer": {
            position: { xl: "sticky" },
          },
          "& .appMainFixedFooter": {
            pb: { xl: 0 },
          },
        },
        "&.framedLayout": {
          padding: { xl: 5 },
          backgroundColor: (theme) => theme.palette.primary.main,

          "& .defaultLayoutWrapper": {
            borderRadius: { xl: 3 },
          },

          "& .app-sidebar": {
            position: { xl: "sticky" },
            borderTopLeftRadius: { xl: 12 },
            borderBottomLeftRadius: { xl: 12 },
            overflow: { xl: "hidden" },
            "& [data-simplebar]": {
              height: { xl: "calc(100vh - 90px) !important" },
            },
          },
          "& .app-bar": {
            borderTopRightRadius: { xl: 12 },
          },
          "& .footer": {
            borderBottomRightRadius: { xl: 12 },
          },
          "& .appMainFixedHeader": {
            pt: { xl: 0 },
            "& .app-bar": {
              position: { xl: "sticky" },
              width: { xl: "100%" },
            },
          },
          "& .mainContent": {
            position: { xl: "static" },
            ml: { xl: 0 },
          },
          "& .fixed-footer": {
            position: { xl: "sticky" },
          },
          "& .appMainFixedFooter": {
            pb: { xl: 0 },
          },
        },
      }}
      {...rest}
    >
      {children}
    </Box>
  );
};

export default DefaultLayoutContainer;
