import React from "react";
import orange from "@mui/material/colors/orange";
import { alpha, Box } from "@mui/material";
import Avatar from "@mui/material/Avatar";
import Menu from "@mui/material/Menu";
import MenuItem from "@mui/material/MenuItem";
import { Fonts } from "../../../../../shared/constants/AppEnums";
import { useNavigate } from "react-router-dom";
import { useLogout, useGetIdentity } from "react-admin";
import { UserMenuContextProvider } from "@crema/core/AppLayout/components/UserMenuContext";

interface UserInfoProps {
  userMenus?: React.ReactNode[];
}

const UserInfo: React.FC<UserInfoProps> = ({ userMenus }) => {
  const { isLoading, identity: user } = useGetIdentity();
  const navigate = useNavigate();
  const logout = useLogout();

  const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null);

  const handleClick = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorEl(event.currentTarget);
  };

  const handleClose = React.useCallback(() => {
    setAnchorEl(null);
  }, []);

  const context = React.useMemo(
    () => ({ onClose: handleClose }),
    [handleClose]
  );

  const getUserAvatar = () => {
    if (user?.displayName) {
      return user.displayName.charAt(0).toUpperCase();
    }
    if (user?.email) {
      return user.email.charAt(0).toUpperCase();
    }
  };

  return (
    <Box
      sx={{
        py: 3,
        px: 3,
        display: "flex",
        alignItems: "center",
        justifyContent: "center",
        cursor: "pointer",
      }}
    >
      <Box onClick={handleClick}>
        {user?.avatar_path ? (
          <Avatar
            sx={{
              height: 30,
              width: 30,
              backgroundColor: orange[500],
            }}
            src={user.avatar_path}
          />
        ) : (
          <Avatar
            sx={{
              height: 30,
              width: 30,
              fontSize: 20,
              backgroundColor: orange[500],
            }}
          >
            {getUserAvatar()}
          </Avatar>
        )}
      </Box>
      <UserMenuContextProvider value={context}>
        <Menu
          id="simple-menu"
          anchorEl={anchorEl}
          keepMounted
          open={Boolean(anchorEl)}
          onClose={handleClose}
          sx={{
            py: 4,
          }}
        >
          <MenuItem
            sx={{
              backgroundColor: (theme) =>
                alpha(theme.palette.common.black, 0.08),
              px: 6,
              py: 3,
            }}
          >
            <Box
              sx={{
                mr: 3.5,
              }}
            >
              {user?.avatar_path ? (
                <Avatar
                  sx={{
                    height: 40,
                    width: 40,
                  }}
                  src={user?.avatar_path}
                />
              ) : (
                <Avatar
                  sx={{
                    height: 40,
                    width: 40,
                    fontSize: 20,
                    backgroundColor: orange[500],
                  }}
                >
                  {getUserAvatar()}
                </Avatar>
              )}
            </Box>

            <Box>
              <Box
                sx={{
                  mb: 0,
                  overflow: "hidden",
                  textOverflow: "ellipsis",
                  whiteSpace: "nowrap",
                  fontSize: 14,
                  fontWeight: Fonts.MEDIUM,
                }}
                component="span"
              >
                {user?.displayName ? user?.displayName : "Admin User "}
              </Box>
              <Box
                sx={{
                  mt: -0.5,
                  textOverflow: "ellipsis",
                  whiteSpace: "nowrap",
                  fontSize: 12,
                  color: (theme) => theme.palette.text.secondary,
                }}
              >
                {user?.roleName ? user.roleName : "System Manager "}
              </Box>
            </Box>
          </MenuItem>
          {userMenus?.length &&
            userMenus.map((menu: JSX.Element, index) =>
              React.cloneElement(menu, {
                key: index,
                sx: {
                  px: 6,
                  py: 1.5,
                },
              })
            )}
        </Menu>
      </UserMenuContextProvider>
    </Box>
  );
};

export default UserInfo;
