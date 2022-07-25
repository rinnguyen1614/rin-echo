import React, { ReactNode } from "react";
import { Box } from "@mui/material";

interface DefaultLayoutWrapperProps {
  children: ReactNode;

  [x: string]: any;
}

const DefaultLayoutWrapper: React.FC<DefaultLayoutWrapperProps> = ({
  children,
  ...rest
}) => {
  return (
    <Box
      sx={{
        flex: 1,
        display: "flex",
        flexDirection: "row",
        position: "relative",
        backgroundColor: (theme) => theme.palette.background.default,

        "&.appMainFixedFooter": {
          pb: { xs: 12, xl: 14.5 },
        },
        "&.appMainFixedHeader": {
          pt: { xs: 14, sm: 17.5 },
          "& .app-bar": {
            position: "fixed",
            top: 0,
            right: 0,
            zIndex: 9,
            width: {
              xs: "100%",
              lg: "calc(100% - 280px)",
            },
          },
        },
      }}
      {...rest}
    >
      {children}
    </Box>
  );
};

export default DefaultLayoutWrapper;
