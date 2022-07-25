import React, { ReactNode } from "react";
import ListItem from "@mui/material/ListItem";
import { Fonts } from "../../../../../shared/constants/AppEnums";
import { alpha } from "@mui/material";

interface VerticalNavGroupItemProps {
  children: ReactNode;
  sidebarTextColor: string;
  level: number;

  [x: string]: any;
}

const VerticalNavGroupItem: React.FC<VerticalNavGroupItemProps> = ({
  children,
  sidebarTextColor,
  level,
  ...rest
}) => {
  return (
    <ListItem
      sx={{
        height: 40,
        my: 0.25,
        pl: 31 + 33 * level + "px",
        pr: 3,
        color: alpha(sidebarTextColor, 0.7),
        fontWeight: Fonts.SEMI_BOLD,
        fontSize: 14,
        cursor: "pointer",
        textDecoration: "none!important",
        whiteSpace: "nowrap",
        transition: "all 0.4s ease",
        "&.nav-item-header": {
          cursor: "auto",
          transition: "all 0.4s ease",
        },
      }}
      {...rest}
    >
      {children}
    </ListItem>
  );
};

export default VerticalNavGroupItem;
