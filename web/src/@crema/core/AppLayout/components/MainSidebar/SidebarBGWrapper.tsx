import React, { ReactNode } from "react";
import Box from "@mui/material/Box";
import { alpha } from "@mui/material";
import { ThemeMode } from "../../../../shared/constants/AppEnums";
import { useSidebarContext } from "../../../../utility/AppContextProvider/SidebarContextProvider";

interface SidebarBgWrapperProps {
  children: ReactNode;
}

const SidebarBgWrapper: React.FC<SidebarBgWrapperProps> = ({ children }) => {
  const {
    sidebarBgColor,
    sidebarTextColor,
    mode,
    isSidebarBgImage,
    sidebarBgImage,
  } = useSidebarContext();
  return (
    <Box
      sx={{
        position: "relative",
        height: "100%",
        width: "100%",
        overflow: "hidden",
        backgroundColor: sidebarBgColor,
        backgroundImage: isSidebarBgImage
          ? `url(/assets/images/sidebar/images/${sidebarBgImage}.png)`
          : "",
        backgroundRepeat: isSidebarBgImage ? "no-repeat" : "",
        backgroundPosition: isSidebarBgImage ? "center center" : "",
        backgroundSize: isSidebarBgImage ? "cover" : "",
        color: sidebarTextColor,
        // boxShadow: '3px 3px 4px rgba(0, 0, 0, 0.04)',
        "&:before": {
          content: '""',
          display: isSidebarBgImage ? "block" : "none",
          position: "absolute",
          left: 0,
          top: 0,
          zIndex: 1,
          width: "100%",
          height: "100%",
          backgroundColor: (theme) =>
            mode === ThemeMode.LIGHT
              ? alpha(theme.palette.common.white, 0.5)
              : alpha(theme.palette.common.black, 0.5),
        },
        "& > *": {
          position: "relative",
          zIndex: 3,
        },
      }}
    >
      {children}
    </Box>
  );
};

export default SidebarBgWrapper;
