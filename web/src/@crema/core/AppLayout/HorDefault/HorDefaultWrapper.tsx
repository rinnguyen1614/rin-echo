import React, { ReactNode } from "react";
import { Box } from "@mui/material";

interface HorDefaultWrapperProps {
  children: ReactNode;

  [x: string]: any;
}

const HorDefaultWrapper: React.FC<HorDefaultWrapperProps> = ({
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
        "& .customizerOption": {
          top: 210,
        },
      }}
      {...rest}
    >
      {children}
    </Box>
  );
};

export default HorDefaultWrapper;
