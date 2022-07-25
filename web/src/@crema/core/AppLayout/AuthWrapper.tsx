import React, { ReactNode } from "react";
import Box from "@mui/material/Box";

interface AuthWrapperProps {
  children: ReactNode;
}

const AuthWrapper: React.FC<AuthWrapperProps> = ({ children }) => {
  return (
    <Box
      sx={{
        flex: 1,
        display: "flex",
        position: "relative",
        minHeight: "100vh",
        minWidth: "100%",
        backgroundColor: (theme) => theme.palette.background.default,

        "& .app-content-view": {
          flex: 1,
          display: "flex",
          flexDirection: "column",
          minHeight: "100vh",
        },
        "& .footer": {
          marginRight: 0,
          marginLeft: 0,
        },
      }}
    >
      {children}
    </Box>
  );
};

export default AuthWrapper;
