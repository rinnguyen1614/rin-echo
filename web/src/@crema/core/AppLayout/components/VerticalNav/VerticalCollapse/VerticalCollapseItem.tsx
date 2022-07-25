import React, { ReactNode } from "react";
import ListItem from "@mui/material/ListItem";
import { Fonts } from "../../../../../shared/constants/AppEnums";
import { alpha } from "@mui/material";

interface VerticalCollapseItemProps {
  children: ReactNode;
  sidebarTextColor: string;

  [x: string]: any;
}

const VerticalCollapseItem: React.FC<VerticalCollapseItemProps> = ({
  children,
  sidebarTextColor,
  ...rest
}) => {
  return (
    <ListItem
      sx={{
        height: 40,
        my: 0.25,
        pl: "31px",
        pr: 3.75,
        whiteSpace: "nowrap",
        transition: "all 0.4s ease",
        "& .nav-item-text": {
          fontWeight: Fonts.MEDIUM,
          color: alpha(sidebarTextColor, 0.7),
        },

        "& .nav-item-icon": {
          color: alpha(sidebarTextColor, 0.7),
          fontSize: 20,
          display: "block",
        },

        "& .nav-item-icon-arrow": {
          color: alpha(sidebarTextColor, 0.7),
        },

        "& .MuiIconButton-root": {
          mr: 3,
          padding: 0,
        },

        "& .MuiTouchRipple-root": {
          zIndex: 10,
        },

        "&.open, &:hover, &:focus": {
          "& .nav-item-text": {
            fontWeight: Fonts.MEDIUM,
            color: sidebarTextColor,
          },

          "& .nav-item-icon": {
            color: sidebarTextColor,
          },

          "& .nav-item-icon-arrow": {
            color: sidebarTextColor,
          },
        },
        "&:hover": {
          backgroundColor: "transparent",
        },
      }}
      {...rest}
    >
      {children}
    </ListItem>
  );
};

export default VerticalCollapseItem;
