import React, { ReactNode } from "react";
import { Box } from "@mui/material";

interface StandardSidebarWrapperProps {
  children: ReactNode;

  [x: string]: any;
}

const StandardSidebarWrapper: React.FC<StandardSidebarWrapperProps> = ({
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
        top: { xs: 0, lg: 71 },
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

export default StandardSidebarWrapper;
