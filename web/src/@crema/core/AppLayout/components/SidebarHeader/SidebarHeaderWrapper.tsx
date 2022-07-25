import React, { ReactNode } from "react";
import { Box } from "@mui/material";
import { useSidebarContext } from "../../../../utility/AppContextProvider/SidebarContextProvider";

interface SidebarHeaderWrapperProps {
  children: ReactNode;
}

const SidebarHeaderWrapper: React.FC<SidebarHeaderWrapperProps> = ({
  children,
}) => {
  const { sidebarHeaderColor, isSidebarBgImage } = useSidebarContext();
  return (
    <Box
      sx={{
        display: "flex",
        flexDirection: "column",
        justifyContent: "center",
        backgroundColor: isSidebarBgImage ? "transparent" : sidebarHeaderColor,
        "&:hover": {
          "& .arrowIcon": {
            opacity: 1,
            visibility: "visible",
          },
        },
      }}
    >
      {children}
    </Box>
  );
};

export default SidebarHeaderWrapper;
