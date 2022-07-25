import React, { ReactNode } from "react";
import { alpha, Box } from "@mui/material";
import { useSidebarContext } from "../../../../utility/AppContextProvider/SidebarContextProvider";

interface HeaderNavWrapperProps {
  children: ReactNode;
}

const HeaderNavWrapper: React.FC<HeaderNavWrapperProps> = ({ children }) => {
  const {
    sidebarBgColor,
    sidebarTextColor,
    sidebarMenuSelectedBgColor,
    sidebarMenuSelectedTextColor,
  } = useSidebarContext();
  return (
    <Box
      sx={{
        backgroundColor: sidebarBgColor,
        color: sidebarTextColor,
        py: 2.5,
        "& .navbarNav": {
          display: "flex",
          padding: 0,
        },
        "& .navItem": {
          width: "auto",
          cursor: "pointer",
          py: 1,
          px: { xs: 4, lg: 5 },
          borderRadius: 1,
          "&.active": {
            color: sidebarMenuSelectedTextColor,
            backgroundColor: alpha(sidebarMenuSelectedBgColor, 0.8),
            "& .navLinkIcon": {
              color: (theme) => theme.palette.secondary.main,
            },
          },
        },
        "& .navLinkIcon": {
          mr: 2.5,
          color: (theme) => theme.palette.common.white,
          fontSize: 20,
        },
      }}
    >
      {children}
    </Box>
  );
};

export default HeaderNavWrapper;
