import React, { ReactNode } from "react";

import { Box } from "@mui/material";

interface AppsHeaderProps {
  children: ReactNode;
}

const AppsHeader: React.FC<AppsHeaderProps> = ({ children }) => {
  return (
    <Box
      sx={{
        height: 60,
        display: "flex",
        alignItems: "center",
        borderBottom: (theme) => `1px solid ${theme.palette.divider}`,
        padding: {
          xs: "4px 10px",
          xl: "12px 10px",
        },
      }}
      className="apps-header"
    >
      {children}
    </Box>
  );
};

export default AppsHeader;
