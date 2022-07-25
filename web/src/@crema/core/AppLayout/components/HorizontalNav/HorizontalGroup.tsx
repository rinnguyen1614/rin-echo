import React, { useState } from "react";
import {
  Grow,
  Icon,
  IconButton,
  List,
  ListItem,
  ListItemText,
  Paper,
} from "@mui/material";
import clsx from "clsx";
import { Manager, Popper, Reference } from "react-popper";
import HorizontalCollapse from "./HorizontalCollapse";
import HorizontalItem from "./HorizontalItem";
import Box from "@mui/material/Box";
import FormattedMessage from "../../../../utility/FormattedMessage";
import { Fonts } from "../../../../shared/constants/AppEnums";
import ClientOnlyPortal from "./ClientPortal";
import { useLocation } from "react-router-dom";
import { RouterConfig } from "../../../../types/models/RouterConfig";

interface HorizontalCollapseProps {
  item: RouterConfig;
  nestedLevel: number;
  dense?: number;
}

const HorizontalGroup: React.FC<HorizontalCollapseProps> = (props) => {
  const [opened, setOpened] = useState<boolean>(false);
  const { item, nestedLevel } = props;
  const location = useLocation();

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
    <Manager>
      <Reference>
        {({ ref }) => (
          <ListItem
            ref={ref}
            className={clsx(
              "navItem",
              isPathInChildren(item, location.pathname) && "active"
            )}
            onMouseEnter={() => handleToggle(true)}
            onMouseLeave={() => handleToggle(false)}
          >
            {item.icon && (
              <Icon color="action" className="navLinkIcon">
                {item.icon}
              </Icon>
            )}
            <ListItemText
              primary={<FormattedMessage id={item.slug} />}
              sx={{
                fontWeight: Fonts.MEDIUM,
              }}
            />
            {nestedLevel > 0 && (
              <IconButton
                sx={{
                  ml: 2,
                }}
                disableRipple
              >
                <Icon
                  sx={{
                    fontSize: 18,
                  }}
                  className="arrow-icon"
                >
                  keyboard_arrow_right
                </Icon>
              </IconButton>
            )}
          </ListItem>
        )}
      </Reference>
      <ClientOnlyPortal selector="#root">
        <Popper placement={nestedLevel === 0 ? "bottom-start" : "right"}>
          {({ ref, style, placement }) =>
            opened && (
              <Box
                ref={ref}
                sx={{
                  ...style,
                  boxShadow: "0 0 3px 0 rgba(0, 0, 0, 0.2)",
                  zIndex: 1110 + nestedLevel,
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
                                nestedLevel={nestedLevel}
                              />
                            )}

                            {item.type === "collapse" && (
                              <HorizontalCollapse
                                item={item}
                                nestedLevel={nestedLevel}
                              />
                            )}

                            {item.type === "item" && (
                              <HorizontalItem
                                item={item}
                                nestedLevel={nestedLevel}
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
  );
};

export default HorizontalGroup;
