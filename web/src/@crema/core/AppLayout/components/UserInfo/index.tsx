import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import { Box } from "@mui/material";
import Avatar from "@mui/material/Avatar";
import orange from "@mui/material/colors/orange";
import Menu from "@mui/material/Menu";
import React from "react";
import { useGetIdentity } from "react-admin";
import { Fonts } from "../../../../shared/constants/AppEnums";
import { UserMenuContextProvider } from "../UserMenuContext";

interface UserInfoProps {
  color?: string;
  userMenus?: React.ReactNode[];
}

const UserInfo: React.FC<UserInfoProps> = ({
  color = "text.secondary",
  userMenus,
}) => {
  const { isLoading, identity: user } = useGetIdentity();

  const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null);

  const handleClick = (event: any) => {
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
    <>
      <Box
        onClick={handleClick}
        sx={{
          py: 3,
          px: 3,
          display: "flex",
          alignItems: "center",
          cursor: "pointer",
        }}
        className="user-info-view"
      >
        <Box sx={{ py: 0.5 }}>
          {user?.avatar_path ? (
            <Avatar
              sx={{
                height: 40,
                width: 40,
                fontSize: 24,
                backgroundColor: orange[500],
              }}
              src={user.avatar_path}
            />
          ) : (
            <Avatar
              sx={{
                height: 40,
                width: 40,
                fontSize: 24,
                backgroundColor: orange[500],
              }}
            >
              {getUserAvatar()}
            </Avatar>
          )}
        </Box>
        <Box
          sx={{
            width: { xs: "calc(100% - 62px)", xl: "calc(100% - 72px)" },
            ml: 4,
            color: color,
          }}
          className="user-info"
        >
          <Box
            sx={{
              display: "flex",
              alignItems: "center",
              justifyContent: "space-between",
            }}
          >
            <Box
              sx={{
                mb: 0,
                overflow: "hidden",
                textOverflow: "ellipsis",
                whiteSpace: "nowrap",
                fontSize: 16,
                fontWeight: Fonts.MEDIUM,
                color: "inherit",
              }}
              component="span"
            >
              {user?.displayName ? user.displayName : "Admin User "}
            </Box>
            {userMenus?.length && (
              <Box
                sx={{
                  ml: 3,
                  color: "inherit",
                  display: "flex",
                }}
              >
                <ExpandMoreIcon />
              </Box>
            )}
          </Box>
          <Box
            sx={{
              mt: -0.5,
              textOverflow: "ellipsis",
              whiteSpace: "nowrap",
              color: "inherit",
            }}
          >
            {user?.roleName ? user.roleName : "System Manager "}
          </Box>
        </Box>
      </Box>

      {userMenus?.length && (
        <UserMenuContextProvider value={context}>
          <Menu
            id="simple-menu"
            anchorEl={anchorEl}
            keepMounted
            open={Boolean(anchorEl)}
            onClose={handleClose}
            anchorOrigin={{
              vertical: "bottom",
              horizontal: "right",
            }}
            transformOrigin={{
              vertical: "top",
              horizontal: "right",
            }}
          >
            {userMenus.map((menu: JSX.Element, index) =>
              React.cloneElement(menu, { key: index })
            )}
          </Menu>
        </UserMenuContextProvider>
      )}
    </>
  );
};

export default UserInfo;
