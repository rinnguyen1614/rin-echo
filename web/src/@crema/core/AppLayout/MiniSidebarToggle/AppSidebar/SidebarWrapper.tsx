import React, { ReactNode } from "react";
import { Box } from "@mui/material";

interface SidebarWrapperProps {
  children: ReactNode;

  [x: string]: any;
}

const SidebarWrapper: React.FC<SidebarWrapperProps> = ({
  children,
  ...rest
}) => {
  return (
    <Box
      sx={{
        width: 280,
        display: "flex",
        flexDirection: "column",
        transition: "all 0.5s ease",
        position: { xs: "relative", lg: "fixed" },
        top: 0,
        left: 0,
        zIndex: 1101,
        "& .app-sidebar": {
          position: "relative",
          top: "auto",
          left: "auto",
          width: "100%",
        },
      }}
      {...rest}
    >
      {children}
    </Box>
  );
};

export default SidebarWrapper;
