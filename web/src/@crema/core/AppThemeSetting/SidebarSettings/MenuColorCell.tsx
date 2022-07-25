import React from "react";
import Box from "@mui/material/Box";
import clsx from "clsx";
import {
  useSidebarActionsContext,
  useSidebarContext,
} from "../../../utility/AppContextProvider/SidebarContextProvider";
import {
  MenuStyle,
  NavStyle,
  ThemeMode,
} from "../../../shared/constants/AppEnums";
import { useLayoutContext } from "../../../utility/AppContextProvider/LayoutContextProvider";
import AppSelectedIcon from "../../AppSelectedIcon";
import { useThemeContext } from "../../../utility/AppContextProvider/ThemeContextProvider";
import { SidebarData } from "../../../utility/AppContextProvider/defaultConfig";

interface MenuColorCellProps {
  sidebarColors: SidebarData;
}

const MenuColorCell: React.FC<MenuColorCellProps> = ({ sidebarColors }) => {
  const {
    sidebarBgColor,
    sidebarTextColor,
    sidebarMenuSelectedBgColor,
    sidebarMenuSelectedTextColor,
    menuStyle,
  } = useSidebarContext();
  const { updateSidebarColorSet } = useSidebarActionsContext();
  const { navStyle } = useLayoutContext();
  const { theme } = useThemeContext();

  return (
    <Box
      sx={{
        border: (theme) =>
          sidebarColors.mode === ThemeMode.LIGHT
            ? `solid 2px ${theme.palette.text.disabled}`
            : `solid 2px ${sidebarColors.sidebarBgColor}`,
        borderRadius: 2,
        cursor: "pointer",
        overflow: "hidden",
        position: "relative",
      }}
      onClick={() => updateSidebarColorSet(sidebarColors)}
    >
      {navStyle === NavStyle.DEFAULT ? (
        <Box
          sx={{
            width: "100%",
            height: 60,
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
            justifyContent: "center",
            backgroundColor: sidebarColors.sidebarHeaderColor,
            borderBottom: `solid 1px ${sidebarColors.sidebarTextColor}`,
          }}
        >
          <Box
            sx={{
              width: 30,
              height: 30,
              border: `solid 1px ${sidebarColors.sidebarTextColor}`,
              borderRadius: "50%",
            }}
          />
          <Box
            sx={{
              width: 50,
              height: 4,
              marginTop: 2,
              backgroundColor: sidebarColors.sidebarTextColor,
            }}
          />
        </Box>
      ) : null}
      <Box
        sx={{
          maxWidth: 200,
          backgroundColor: sidebarColors.sidebarBgColor,
        }}
      >
        <Box
          sx={{
            width: "100%",
            minHeight: 40,
            padding: 2.5,
            py: 2.5,
            px: 4.5,
            color: sidebarColors.sidebarTextColor,
            fontSize: { xs: 12, md: 14 },
            whiteSpace: "nowrap",
          }}
        >
          Menu-1
        </Box>
        <Box
          sx={{
            width: "100%",
            minHeight: 40,
            py: 2.5,
            px: 4.5,
            color: sidebarColors.sidebarTextColor,
            fontSize: { xs: 12, md: 14 },
            whiteSpace: "nowrap",
          }}
        >
          Menu-2
        </Box>
        <Box
          className={clsx({
            "rounded-menu": menuStyle === MenuStyle.ROUNDED,
            "rounded-menu-reverse": menuStyle === MenuStyle.ROUNDED_REVERSE,
            "standard-menu": menuStyle === MenuStyle.STANDARD,
          })}
          sx={{
            width: "calc(100% - 16px)",
            minHeight: 40,
            py: 2.5,
            px: 2.5,
            mx: 2,
            borderRadius: 1,
            backgroundColor: sidebarColors.sidebarMenuSelectedBgColor,
            color: sidebarColors.sidebarMenuSelectedTextColor,
            position: "relative",
            transition: "all 0.4s ease",
            fontSize: { xs: 12, md: 14 },
            whiteSpace: "nowrap",
            "&.rounded-menu": {
              mr: 2,
              ml: 0,
              width: "calc(100% - 8px)",
              pl: 4.5,
              pr: 2.5,
              borderRadius: "0 30px 30px 0",
            },
            "&.rounded-menu-reverse": {
              ml: 2,
              mr: 0,
              width: "calc(100% - 8px)",
              pl: 2.5,
              pr: 2.5,
              borderRadius: "30px 0 0 30px",
              "&.active:after": {
                display: "none",
              },
            },
            "&.standard-menu": {
              mx: 0,
              width: "100%",
              pl: 4.5,
              pr: 2.5,
              borderRadius: 0,
              "&:after": {
                content: '""',
                position: "absolute",
                right: 0,
                top: 0,
                height: "100%",
                width: 4,
                backgroundColor: (theme) => theme.palette.primary.main,
              },
            },
          }}
        >
          Selected Menu
        </Box>
        <Box
          sx={{
            width: "100%",
            minHeight: 40,
            py: 2.5,
            px: 4.5,
            color: sidebarColors.sidebarTextColor,
          }}
        >
          Menu-4
        </Box>
      </Box>
      {sidebarColors.sidebarBgColor === sidebarBgColor &&
      sidebarColors.sidebarTextColor === sidebarTextColor &&
      sidebarColors.sidebarMenuSelectedBgColor === sidebarMenuSelectedBgColor &&
      sidebarColors.sidebarMenuSelectedTextColor ===
        sidebarMenuSelectedTextColor ? (
        <AppSelectedIcon
          isCenter={false}
          backgroundColor={
            sidebarColors.mode === ThemeMode.LIGHT
              ? theme.palette.text.primary
              : theme.palette.background.default
          }
          color={
            sidebarColors.mode === ThemeMode.LIGHT
              ? theme.palette.background.default
              : theme.palette.text.primary
          }
        />
      ) : null}
    </Box>
  );
};

export default MenuColorCell;
