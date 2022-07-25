import React, { ReactNode } from "react";
import { Box } from "@mui/material";

interface BitBucketHeaderWrapperProps {
  children: ReactNode;

  [x: string]: any;
}

const BitBucketHeaderWrapper: React.FC<BitBucketHeaderWrapperProps> = ({
  children,
  ...rest
}) => {
  return (
    <Box
      sx={{
        height: { xs: 56, sm: 70 },
        display: "flex",
        alignItems: "center",
        top: 0,
        left: 0,
        right: 0,
        width: "100%",
        position: "fixed",
        px: { xs: 5, md: 7.5 },
        zIndex: 999,
        backgroundColor: (theme) => theme.palette.background.paper,
        "& .menu-btn": {
          mr: 2,
        },
        "& .menu-icon": {
          width: 35,
          height: 35,
        },
        "& .logo-text": {
          display: { xs: "none", sm: "block" },
        },
      }}
      {...rest}
    >
      {children}
    </Box>
  );
};

export default BitBucketHeaderWrapper;
