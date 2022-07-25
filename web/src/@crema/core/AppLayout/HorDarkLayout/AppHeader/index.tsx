import React from "react";
import AppBar from "@mui/material/AppBar";
import Toolbar from "@mui/material/Toolbar";
import AppLngSwitcher from "@crema/core/AppLngSwitcher";
import Box from "@mui/material/Box";
import AppSearchBar from "@crema/core/AppSearchBar";
import Hidden from "@mui/material/Hidden";
import IconButton from "@mui/material/IconButton";
import { toggleNavCollapsed } from "../../../../redux/actions";
import MenuIcon from "@mui/icons-material/Menu";
import { useDispatch } from "react-redux";
import AppMessages from "../../../AppMessages";
import AppNotifications from "../../../AppNotifications";
import Menu from "@mui/material/Menu";
import MenuItem from "@mui/material/MenuItem";
import MoreVertIcon from "@mui/icons-material/MoreVert";
import AppTooltip from "../../../AppTooltip";
import { alpha } from "@mui/material/styles";
import NotificationBar from "../NotificationBar";
import AppLogo from "../../components/AppLogo";
import UserInfo from "../../components/UserInfo";
import HeaderNavWrapper from "./HeaderNavWrapper";
import HorizontalNav from "../../components/HorizontalNav";
import { useSidebarContext } from "../../../../utility/AppContextProvider/SidebarContextProvider";

interface AppHeaderProps {
  userMenus?: React.ReactNode[];
}

const AppHeader: React.FC<AppHeaderProps> = ({ userMenus }) => {
  const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null);
  const { sidebarBgColor, sidebarTextColor } = useSidebarContext();

  const handleClick = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorEl(event.currentTarget);
  };

  const handleClose = () => {
    setAnchorEl(null);
  };
  const dispatch = useDispatch();

  return (
    <>
      <AppBar
        position="relative"
        color="inherit"
        sx={{
          boxShadow: "none",
          width: "100%",
        }}
        className="app-bar"
      >
        <NotificationBar />
        <Toolbar
          sx={{
            boxSizing: "border-box",
            minHeight: { xs: 56, sm: 70 },
            px: { xs: 0 },
            borderBottom: (theme) =>
              `solid 1px ${alpha(theme.palette.common.black, 0.15)}`,
            backgroundColor: sidebarBgColor,
            color: sidebarTextColor,
          }}
        >
          <Box
            sx={{
              width: "100%",
              maxWidth: { lg: 1140, xl: 1420 },
              mx: "auto",
              px: 5,
              display: "flex",
              alignItems: "center",
            }}
          >
            <Hidden lgUp>
              <IconButton
                sx={{
                  marginRight: (theme) => theme.spacing(2),
                  color: sidebarTextColor,
                }}
                edge="start"
                className="menu-btn"
                color="inherit"
                aria-label="open drawer"
                onClick={() => dispatch(toggleNavCollapsed())}
                size="large"
              >
                <MenuIcon
                  sx={{
                    width: 35,
                    height: 35,
                  }}
                />
              </IconButton>
            </Hidden>

            <Box
              sx={{
                "& .app-logo": {
                  pl: 0,
                },
                "& .logo-text": {
                  display: { xs: "none", sm: "block" },
                },
              }}
            >
              <AppLogo />
            </Box>

            <Box
              sx={{
                flexGrow: 1,
              }}
            />
            <Box
              sx={{
                minHeight: 40,
                position: "relative",
                display: { xs: "none", sm: "block" },
                "& .searchRoot": {
                  position: { xs: "absolute", sm: "relative" },
                  right: { xs: 0, sm: "auto" },
                  top: { xs: 0, sm: "auto" },
                },
              }}
            >
              <AppSearchBar iconPosition="right" placeholder="Search…" />
            </Box>
            <Box sx={{ ml: 4 }}>
              <AppLngSwitcher iconOnly={true} tooltipPosition="bottom" />
            </Box>

            <Box
              sx={{
                ml: 4,
                display: "flex",
                alignItems: "center",
              }}
            >
              <Hidden smDown>
                <Box
                  sx={{
                    position: "relative",
                    display: "flex",
                    alignItems: "center",
                    marginLeft: -2,
                    marginRight: -2,
                  }}
                >
                  <Box
                    sx={{
                      px: 1.85,
                    }}
                  >
                    <AppNotifications />
                  </Box>
                  <Box
                    sx={{
                      px: 1.85,
                    }}
                  >
                    <AppMessages />
                  </Box>
                </Box>
              </Hidden>

              <Box
                sx={{
                  ml: { sm: 4 },
                  mr: { xs: 4, sm: 0 },
                  minWidth: { md: 220 },
                  "& .user-info-view": {
                    p: 0,
                  },
                  "& .user-info": {
                    display: { xs: "none", md: "block" },
                    color: "inherit",
                  },
                }}
              >
                <UserInfo userMenus={userMenus} />
              </Box>

              <Hidden smUp>
                <Box
                  sx={{
                    position: "relative",
                    display: "flex",
                    alignItems: "center",
                    marginLeft: -2,
                    marginRight: -2,
                  }}
                >
                  <Box
                    sx={{
                      px: 1.85,
                    }}
                  >
                    <AppTooltip title="More">
                      <IconButton
                        sx={{
                          borderRadius: "50%",
                          width: 40,
                          height: 40,
                          color: (theme) => theme.palette.text.secondary,
                          backgroundColor: (theme) =>
                            theme.palette.background.default,
                          border: 1,
                          borderColor: "transparent",
                          "&:hover, &:focus": {
                            color: (theme) => theme.palette.text.primary,
                            backgroundColor: (theme) =>
                              alpha(theme.palette.background.default, 0.9),
                            borderColor: (theme) =>
                              alpha(theme.palette.text.secondary, 0.25),
                          },
                        }}
                        onClick={handleClick}
                        size="large"
                      >
                        <MoreVertIcon />
                      </IconButton>
                    </AppTooltip>
                  </Box>
                </Box>
              </Hidden>
              <Menu
                id="simple-menu"
                anchorEl={anchorEl}
                keepMounted
                open={Boolean(anchorEl)}
                onClose={handleClose}
              >
                <MenuItem>
                  <AppNotifications isMenu />
                </MenuItem>
                <MenuItem>
                  <AppMessages isMenu />
                </MenuItem>
                <MenuItem>Setting</MenuItem>
              </Menu>
            </Box>
          </Box>
        </Toolbar>
        <Hidden lgDown>
          <HeaderNavWrapper>
            <Box
              sx={{
                width: "100%",
                maxWidth: { lg: 1140, xl: 1436 },
                mx: "auto",
                px: 5,
              }}
            >
              <HorizontalNav />
            </Box>
          </HeaderNavWrapper>
        </Hidden>
      </AppBar>
    </>
  );
};
export default AppHeader;
