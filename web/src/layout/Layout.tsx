import { AppLayout } from "@crema";
import { useUserMenu } from "@crema/core/AppLayout/components/UserMenuContext";
import { LayoutProps } from "@crema/core/AppLayout/LayoutProps";
import FormattedMessage from "@crema/utility/FormattedMessage";
import { MenuItem } from "@mui/material";
import React from "react";
import { useLogout } from "react-admin";
import { useNavigate } from "react-router";

const AccountMenu = React.forwardRef<any, any>((props, ref) => {
  const nagivate = useNavigate();
  const { onClose } = useUserMenu();
  return (
    <MenuItem
      ref={ref}
      onClick={(e) => {
        onClose && onClose();
        nagivate("/account");
      }}
      {...props}
    >
      <FormattedMessage id="my_account" />
    </MenuItem>
  );
});

const LogoutMenu = (props) => {
  const logout = useLogout();
  return (
    <MenuItem onClick={logout} {...props}>
      <FormattedMessage id="logout" />
    </MenuItem>
  );
};

const UserMenus = [<AccountMenu />, <LogoutMenu />];

const Layout = (props: LayoutProps) => (
  <>
    <AppLayout {...props} userMenus={UserMenus} />
  </>
);

export default Layout;
