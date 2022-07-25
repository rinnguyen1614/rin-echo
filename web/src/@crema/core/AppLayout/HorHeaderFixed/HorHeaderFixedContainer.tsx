import React, { ReactNode } from "react";
import { Box } from "@mui/material";
import PropsTypes from "prop-types";

interface HorHeaderFixedContainerProps {
  children: ReactNode;

  [x: string]: any;
}

const HorHeaderFixedContainer: React.FC<HorHeaderFixedContainerProps> = ({
  children,
  ...rest
}) => {
  return (
    <Box
      sx={{
        minHeight: "100vh",
        display: "flex",
        flexDirection: "column",
        position: "relative",
        backgroundColor: (theme) => theme.palette.background.default,
        pt: { xs: 14, sm: 17.5 },
        "&.boxedLayout": {
          maxWidth: { xl: 1480 },
          mx: { xl: "auto" },
          boxShadow: "none",
          borderLeft: "1px solid #e8e5dd",
          borderRight: "1px solid #e8e5dd",
          pt: { xl: 0 },

          "& .app-bar": {
            position: { xl: "sticky" },
          },
          "& .mainContent": {
            position: { xl: "static" },
          },
          "& .fixed-footer": {
            position: { xl: "sticky" },
          },
          "& .appMainFixedFooter": {
            pb: { xl: 0 },
          },
        },
        "&.framedLayout": {
          padding: { xl: 5 },
          backgroundColor: (theme) => theme.palette.primary.main,

          "& .horHeaderFixedWrapper": {
            borderRadius: { xl: 3 },
          },

          "& .app-bar": {
            borderTopLeftRadius: { xl: 12 },
            borderTopRightRadius: { xl: 12 },
            overflow: "hidden",
            position: { xl: "sticky" },
          },

          "& .footer": {
            borderBottomLeftRadius: { xl: 12 },
            borderBottomRightRadius: { xl: 12 },
          },

          "& .mainContent": {
            position: { xl: "static" },
          },
          "& .fixed-footer": {
            position: { xl: "sticky" },
          },
          "& .appMainFixedFooter": {
            pb: { xl: 0 },
          },
        },
      }}
      {...rest}
    >
      {children}
    </Box>
  );
};

export default HorHeaderFixedContainer;

HorHeaderFixedContainer.propTypes = {
  children: PropsTypes.node,
};
