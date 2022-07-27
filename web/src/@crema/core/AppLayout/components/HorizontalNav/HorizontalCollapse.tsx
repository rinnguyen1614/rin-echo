import React, { useState } from "react";
import {
  Grow,
  Icon,
  IconButton,
  ListItem,
  ListItemText,
  Paper,
} from "@mui/material";
import { useLocation } from "react-router-dom";
import clsx from "clsx";
import List from "@mui/material/List";
import Box from "@mui/material/Box";
import { Manager, Popper, Reference } from "react-popper";
import HorizontalItem from "./HorizontalItem";
import HorizontalGroup from "./HorizontalGroup";
import FormattedMessage from "../../../../utility/FormattedMessage";
import { useThemeContext } from "../../../../utility/AppContextProvider/ThemeContextProvider";
import {
  GetMessageId,
  RouterConfig,
} from "../../../../types/models/RouterConfig";
import ClientOnlyPortal from "./ClientPortal";
import { useSidebarContext } from "../../../../utility/AppContextProvider/SidebarContextProvider";
import { AddItemButton } from "ra-ui-materialui";

interface HorizontalCollapseProps {
  item: RouterConfig;
  nestedLevel: number;
  dense?: number;
}

const HorizontalCollapse: React.FC<HorizontalCollapseProps> = (props) => {
  const [opened, setOpened] = useState<boolean>(false);
  const { theme } = useThemeContext();
  const { item, nestedLevel, dense } = props;
  const location = useLocation();
  const active = isPathInChildren(item, location.pathname);
  const { sidebarMenuSelectedBgColor, sidebarMenuSelectedTextColor } =
    useSidebarContext();

  const handleToggle = (open: boolean) => {
    setOpened(open);
  };

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
    <List
      sx={{
        py: 0,
        "& .list-item-text": {
          padding: "0 0 0 16px",
        },
      }}
      className="navbarNavSubmenu"
    >
      <Manager>
        <Reference>
          {({ ref }) => (
            <ListItem
              ref={ref}
              sx={{
                color: theme.palette.text.primary,
                padding: "0px 12px",
                "&.active, &.active:hover, &.active:focus": {
                  backgroundColor: sidebarMenuSelectedBgColor + "!important",
                  color: sidebarMenuSelectedTextColor + "!important",
                },
                "&.open": {
                  backgroundColor: "rgba(0,0,0,.08)",
                },
                "&.dense": {
                  padding: "0px 12px",
                  "& .list-item-text": {
                    padding: "0 0 0 8px",
                  },
                },
              }}
              className={clsx(
                "navItemSubmenu",
                opened && "open",
                dense && "dense",
                active && "active"
              )}
              onMouseEnter={() => handleToggle(true)}
              onMouseLeave={() => handleToggle(false)}
            >
              {item.icon && (
                <Icon
                  sx={{
                    color: active ? sidebarMenuSelectedTextColor : "action",
                    mr: 3.5,
                    fontSize: { xs: 16, xl: 18 },
                  }}
                >
                  {item.icon}
                </Icon>
              )}
              <ListItemText
                className="navLinkTextSubmenu"
                primary={<FormattedMessage id={GetMessageId(item)} />}
              />
              <Box p={0}>
                <IconButton disableRipple>
                  <Icon
                    sx={{
                      color: active ? sidebarMenuSelectedTextColor : "action",
                    }}
                  >
                    {theme.direction === "ltr"
                      ? "chevron_right"
                      : "chevron_left"}
                  </Icon>
                </IconButton>
              </Box>
            </ListItem>
          )}
        </Reference>
        <ClientOnlyPortal selector="#root">
          <Popper placement="right">
            {({ ref, style, placement }) =>
              opened && (
                <Box
                  ref={ref}
                  sx={{
                    boxShadow: "0 0 3px 0 rgba(0, 0, 0, 0.2)",
                    zIndex: 1110 + nestedLevel + 1,
                    ...style,
                    "& .popperClose": {
                      pointerEvents: "none",
                    },
                  }}
                  data-placement={placement}
                  className={clsx({
                    popperClose: !opened,
                  })}
                >
                  <Grow in={opened}>
                    <Paper
                      onMouseEnter={() => handleToggle(true)}
                      onMouseLeave={() => handleToggle(false)}
                    >
                      {item.children && (
                        <List
                          sx={{
                            px: 0,
                          }}
                        >
                          {item.children.map((item) => (
                            <React.Fragment key={item.id}>
                              {item.type === "group" && (
                                <HorizontalGroup
                                  item={item}
                                  nestedLevel={nestedLevel + 1}
                                />
                              )}

                              {item.type === "collapse" && (
                                <HorizontalCollapse
                                  item={item}
                                  nestedLevel={nestedLevel + 1}
                                />
                              )}

                              {item.type === "item" && (
                                <HorizontalItem
                                  item={item}
                                  nestedLevel={nestedLevel + 1}
                                />
                              )}
                            </React.Fragment>
                          ))}
                        </List>
                      )}
                    </Paper>
                  </Grow>
                </Box>
              )
            }
          </Popper>
        </ClientOnlyPortal>
      </Manager>
    </List>
  );
};

export default HorizontalCollapse;
