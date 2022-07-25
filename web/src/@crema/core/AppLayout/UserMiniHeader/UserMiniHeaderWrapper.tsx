import React, { ReactNode } from "react";
import { Box } from "@mui/material";

interface UserMiniHeaderWrapperProps {
  children: ReactNode;

  [x: string]: any;
}

const UserMiniHeaderWrapper: React.FC<UserMiniHeaderWrapperProps> = ({
  children,
  ...rest
}) => {
  return (
    <Box
      sx={{
        flex: 1,
        display: "flex",
        flexDirection: "column",
        position: "relative",
        backgroundColor: (theme) => theme.palette.background.default,
        paddingTop: { xs: 14, sm: 17.5 },
        "& .mainContent": {
          display: "flex",
          flexDirection: "column",
          position: "relative",
          width: { xs: "100%", lg: "calc(100% - 280px)" },
          transition: "all 0.5s ease",
          ml: { lg: "280px" },
        },
        "&.mini-sidebar-collapsed": {
          "& .mainContent": {
            width: { lg: "calc(100% - 64px)" },
            ml: { lg: "64px" },
          },
          "& .user-mini-header-sidebar:not(:hover)": {
            width: { lg: 64 },
            "& .nav-item-header, & .nav-item-content, & .menu-badge, & .collapse-children, & .nav-item-icon-arrow-btn":
              {
                display: "none",
              },
            "& .nav-item-icon": {
              mr: 0,
              ml: 0.5,
            },
          },
          "& .menu-vertical-item": {
            pl: 3,
            "&.rounded-menu": {
              mx: 2,
              pl: 3,
              borderRadius: 1,
            },
            "&.rounded-menu-reverse": {
              mx: 2,
              pl: 3,
              borderRadius: 1,
            },
            "&.standard-menu": {
              mx: 2,
              width: "calc(100% - 16px)",
              pl: 3,
              borderRadius: 1,
              "&.active:after": {
                backgroundColor: "transparent",
              },
            },
            "&.curved-menu": {
              mx: 2,
              borderRadius: 1,
              "&:before, &:after": {
                display: "none",
              },
            },
          },
          "& .menu-vertical-collapse": {
            pl: 3,
            mx: 2,
            width: "calc(100% - 16px)",
          },
        },
        "&.appMainFixedFooter": {
          pb: { xs: 12, xl: 14.5 },
        },
      }}
      {...rest}
    >
      {children}
    </Box>
  );
};

export default UserMiniHeaderWrapper;
