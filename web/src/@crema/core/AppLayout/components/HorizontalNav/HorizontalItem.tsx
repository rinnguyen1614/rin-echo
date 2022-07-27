import React from "react";
import { Icon, ListItem, ListItemText } from "@mui/material";
import { useLocation } from "react-router-dom";
import clsx from "clsx";
import FormattedMessage from "../../../../utility/FormattedMessage";
import Box from "@mui/material/Box";
import AppNavLink from "../../../AppNavLink";
import Badge from "@mui/material/Badge";
import {
  GetMessageId,
  RouterConfig,
} from "../../../../types/models/RouterConfig";
import { useSidebarContext } from "../../../../utility/AppContextProvider/SidebarContextProvider";

interface HorizontalItemProps {
  item: RouterConfig;
  nestedLevel?: number;
  dense?: boolean;
}

const HorizontalItem: React.FC<HorizontalItemProps> = (props) => {
  const { item, dense } = props;

  const location = useLocation();
  const active = isPathInChildren(item, location.pathname);
  const { sidebarMenuSelectedBgColor, sidebarMenuSelectedTextColor } =
    useSidebarContext();

  function isPathInChildren(parent: RouterConfig, path: string) {
    if (!parent.children) {
      return false;
    }

    for (let i = 0; i < parent.children.length; i++) {
      if (parent.children[i].children) {
        if (isPathInChildren(parent.children[i], path)) {
          return true;
        }
      }

      if (
        parent.children[i].path === path ||
        path.includes(parent!.children![i].path!)
      ) {
        return true;
      }
    }

    return false;
  }

  return (
    <ListItem
      component={AppNavLink}
      to={item.path}
      activeClassName="active"
      className={clsx("navItemSubmenu", dense && "dense")}
      exact={item.exact}
      sx={{
        minHeight: 40,
        padding: "4px 12px",
        color: (theme) => theme.palette.text.primary,
        textDecoration: "none!important",
        minWidth: 160,
        "&.active": {
          backgroundColor: sidebarMenuSelectedBgColor,
          color: sidebarMenuSelectedTextColor + "!important",
          pointerEvents: "none",
          "& .list-item-text-primary": {
            color: "inherit",
          },
          "& .list-item-icon": {
            color: "inherit",
          },
        },
        "& .list-item-text": {
          padding: "0 0 0 16px",
        },
        "&.dense": {
          padding: "4px 12px",
          minHeight: 40,
          "& .list-item-text": {
            padding: "0 0 0 8px",
          },
        },
      }}
    >
      {item.icon && (
        <Icon
          sx={{
            color: active ? sidebarMenuSelectedTextColor : "action",
            mr: 3,
            fontSize: { xs: 16, xl: 18 },
          }}
        >
          {item.icon}
        </Icon>
      )}
      <ListItemText
        className="AppNavLinkTextSubmenu"
        primary={<FormattedMessage id={GetMessageId(item)} />}
      />
      {item.count && (
        <Box ml={4}>
          <Badge
            badgeContent={item.count}
            sx={{
              color: item.color,
            }}
          />
        </Box>
      )}
    </ListItem>
  );
};

export default HorizontalItem;
