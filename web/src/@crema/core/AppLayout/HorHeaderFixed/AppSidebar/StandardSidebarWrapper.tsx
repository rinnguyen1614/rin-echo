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
        position: "relative",
        top: 0,
        left: 0,
        zIndex: 1101,
        "& .app-sidebar": {
          position: "relative",
          top: "auto",
          left: "auto",
          width: "100%",
          borderRight: "0 none",
        },
      }}
      {...rest}
    >
      {children}
    </Box>
  );
};

export default StandardSidebarWrapper;
