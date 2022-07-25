import React, { ReactNode } from "react";
import ListItem from "@mui/material/ListItem";
import { Fonts, MenuStyle } from "../../../../../shared/constants/AppEnums";
import { useSidebarContext } from "../../../../../utility/AppContextProvider/SidebarContextProvider";
import clsx from "clsx";
import { alpha } from "@mui/material";

interface VerticalNavItemProps {
  children: ReactNode;
  level: number;

  [x: string]: any;
}

const VerticalNavItem: React.FC<VerticalNavItemProps> = ({
  children,
  level,
  ...rest
}) => {
  const {
    sidebarTextColor,
    sidebarMenuSelectedBgColor,
    sidebarMenuSelectedTextColor,
    menuStyle,
  } = useSidebarContext();

  return (
    <ListItem
      className={clsx("menu-vertical-item", {
        "rounded-menu": menuStyle === MenuStyle.ROUNDED,
        "rounded-menu-reverse": menuStyle === MenuStyle.ROUNDED_REVERSE,
        "standard-menu": menuStyle === MenuStyle.STANDARD,
        "curved-menu": menuStyle === MenuStyle.CURVED_MENU,
      })}
      sx={{
        height: 40,
        my: 0.25,
        cursor: "pointer",
        textDecoration: "none !important",
        mx: 2,
        width: "calc(100% - 16px)",
        pl: 22 + 33 * level + "px",
        pr: 3,
        borderRadius: 1,
        position: "relative",
        transition: "all 0.4s ease",
        whiteSpace: "nowrap",
        "& .nav-item-icon": {
          color: alpha(sidebarTextColor, 0.7),
          fontSize: 20,
          display: "block",
        },
        "& .nav-item-text": {
          color: alpha(sidebarTextColor, 0.7),
          fontWeight: Fonts.MEDIUM,
          fontSize: 14,
        },

        "& .MuiTouchRipple-root": {
          zIndex: 1,
        },
        "&.nav-item-header": {
          textTransform: "uppercase",
        },
        "&:hover, &:focus": {
          "& .nav-item-text, & .nav-item-icon, & .nav-item-icon-arrow": {
            color: sidebarTextColor,
          },
        },
        "&.active": {
          backgroundColor: sidebarMenuSelectedBgColor,
          pointerEvents: "none",
          "& .nav-item-text": {
            color: sidebarMenuSelectedTextColor + "!important",
            fontWeight: Fonts.MEDIUM,
          },
          "& .nav-item-icon": {
            color: sidebarMenuSelectedTextColor + "!important",
          },
        },
        "&.rounded-menu": {
          mr: 4,
          ml: 0,
          width: "calc(100% - 16px)",
          pl: 30 + 33 * level + "px",
          pr: 3,
          borderRadius: "0 30px 30px 0",
        },
        "&.rounded-menu-reverse": {
          ml: 4,
          mr: 0,
          width: "calc(100% - 16px)",
          pl: 14 + 33 * level + "px",
          pr: 3,
          borderRadius: "30px 0 0 30px",
        },
        "&.standard-menu": {
          mx: 0,
          width: "100%",
          pl: 30 + 33 * level + "px",
          pr: 3,
          borderRadius: 0,
          position: "relative",
          "&:after": {
            content: '""',
            position: "absolute",
            right: 0,
            top: 0,
            height: "100%",
            width: 4,
            backgroundColor: "transparent",
          },
          "&.active:after": {
            backgroundColor: (theme) => theme.palette.primary.main,
          },
        },
        "&.curved-menu": {
          ml: 4,
          mr: 0,
          width: "calc(100% - 16px)",
          pl: 14 + 33 * level + "px",
          pr: 3,
          borderRadius: "30px 0 0 30px",
          position: "relative",
          transition: "none",
          "&:before, &:after": {
            content: '""',
            position: "absolute",
            right: 0,
            zIndex: 1,
            height: 40,
            width: 40,
            backgroundColor: "transparent",
            borderRadius: "50%",
            opacity: 0,
          },
          "&:before": {
            top: -40,
            boxShadow: `30px 30px 0 10px transparent`,
          },
          "&:after": {
            bottom: -40,
            boxShadow: `30px -30px 0 10px transparent`,
          },
          "&:hover, &.active": {
            backgroundColor: sidebarMenuSelectedBgColor,
            "& .nav-item-text, & .nav-item-icon": {
              color: sidebarMenuSelectedTextColor + "!important",
            },
            "&:before": {
              boxShadow: `30px 30px 0 10px ${sidebarMenuSelectedBgColor}`,
              opacity: 1,
            },
            "&:after": {
              boxShadow: `30px -30px 0 10px ${sidebarMenuSelectedBgColor}`,
              opacity: 1,
            },
          },
          "& .MuiTouchRipple-root": {
            display: "none",
          },
        },
      }}
      {...rest}
    >
      {children}
    </ListItem>
  );
};

export default VerticalNavItem;
