import React, { ReactNode } from "react";
import { Box } from "@mui/material";

interface StandardContainerProps {
  children: ReactNode;

  [x: string]: any;
}

const StandardContainer: React.FC<StandardContainerProps> = ({
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
          flexDirection: { xl: "column" },
          "& .standardWrapper": {
            pt: { xl: 0 },
          },
          "& .standard-sidebar": {
            position: { xl: "sticky" },
            height: { xl: "100%" },
            "& [data-simplebar]": {
              height: { xl: "calc(100vh - 141px) !important" },
            },
          },
          "& .app-bar": {
            position: { xl: "sticky" },
            width: { xl: "100%" },
          },
          "& .mainContent": {
            position: { xl: "static" },
            ml: { xl: 0 },
            width: { xl: "100%" },
            flexDirection: { xl: "row" },
            flexWrap: { xl: "wrap" },
          },
          "& .fixed-footer": {
            position: { xl: "sticky" },
          },
          "& .appMainFixedFooter": {
            pb: { xl: 0 },
            "& .standard-sidebar": {
              "& [data-simplebar]": {
                height: { xl: "calc(100vh - 188px) !important" },
              },
            },
          },
        },
        "&.framedLayout": {
          padding: { xl: 5 },
          backgroundColor: (theme) => theme.palette.primary.main,

          "& .standardWrapper": {
            borderRadius: { xl: 3 },
            pt: { xl: 0 },
          },
          "& .standard-sidebar": {
            position: { xl: "sticky" },
            height: { xl: "100%" },
            borderBottomLeftRadius: { xl: 12 },
            overflow: { xl: "hidden" },
            "& [data-simplebar]": {
              height: { xl: "calc(100vh - 161px) !important" },
            },
          },
          "& .app-bar": {
            position: { xl: "sticky" },
            width: { xl: "100%" },
            borderTopLeftRadius: { xl: 12 },
            borderTopRightRadius: { xl: 12 },
          },
          "& .footer": {
            borderBottomRightRadius: { xl: 12 },
          },
          "& .mainContent": {
            position: { xl: "static" },
            ml: { xl: 0 },
            width: { xl: "100%" },
            flexDirection: { xl: "row" },
            flexWrap: { xl: "wrap" },
          },
          "& .fixed-footer": {
            position: { xl: "sticky" },
          },
          "& .appMainFixedFooter": {
            pb: { xl: 0 },
            "& .standard-sidebar": {
              "& [data-simplebar]": {
                height: { xl: "calc(100vh - 188px) !important" },
              },
            },
          },
        },
      }}
      {...rest}
    >
      {children}
    </Box>
  );
};

export default StandardContainer;
