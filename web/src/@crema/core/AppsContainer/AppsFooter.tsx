import React, { ReactNode } from "react";
import { Box } from "@mui/material";

interface AppsFooterProps {
  children: ReactNode;
}

const AppsFooter: React.FC<AppsFooterProps> = (props) => {
  const { children } = props;
  return (
    <Box
      sx={{
        px: 5,
        py: 2,
        borderTop: (theme) => `1px solid ${theme.palette.divider}`,
      }}
    >
      {children}
    </Box>
  );
};

export default AppsFooter;
