import { AppStyleProvider, AppThemeProvider } from "@crema";
import configureStore from "@crema/redux/store";
import AppContextProvider from "@crema/utility/AppContextProvider";
import LayoutContextProvider from "@crema/utility/AppContextProvider/LayoutContextProvider";
import LocaleContextProvider from "@crema/utility/AppContextProvider/LocaleContextProvide";
import SidebarContextProvider from "@crema/utility/AppContextProvider/SidebarContextProvider";
import ThemeContextProvider from "@crema/utility/AppContextProvider/ThemeContextProvider";
import { CssBaseline } from "@mui/material";
import Error404 from "pages/errors/Error404";
import { createElement } from "react";
import {
  AdminProps,
  AdminUIProps,
  CoreAdminContext,
  CoreAdminContextProps,
  CoreAdminUI,
  Notification,
  LoadingPage,
} from "react-admin";
import { Provider } from "react-redux";

export const Admin = (props: AdminProps) => {
  const {
    authProvider,
    basename,
    catchAll,
    children,
    dashboard,
    dataProvider,
    disableTelemetry,
    history,
    i18nProvider,
    layout,
    loading,
    loginPage,
    menu, // deprecated, use a custom layout instead
    notification,
    requireAuth,
    store,
    ready,
    theme,
    title = "React Admin",
  } = props;

  if (loginPage === true && process.env.NODE_ENV !== "production") {
    console.warn(
      "You passed true to the loginPage prop. You must either pass false to disable it or a component class to customize it"
    );
  }

  return (
    <AdminContext
      authProvider={authProvider}
      basename={basename}
      dataProvider={dataProvider}
      i18nProvider={i18nProvider}
      store={store}
      history={history}
      theme={theme}
    >
      <AdminUI
        layout={layout}
        dashboard={dashboard}
        disableTelemetry={disableTelemetry}
        menu={menu}
        catchAll={catchAll}
        title={title}
        loading={loading}
        loginPage={loginPage}
        notification={notification}
        requireAuth={requireAuth}
        ready={ready}
      >
        {children}
      </AdminUI>
    </AdminContext>
  );
};

export const AdminUI = ({ notification, ...props }: AdminUIProps) => (
  <>
    <CoreAdminUI {...props} />
    {createElement(notification)}
  </>
);

AdminUI.defaultProps = {
  notification: Notification,
  catchAll: Error404,
  loading: LoadingPage,
};

export const AdminContext = (props: CoreAdminContextProps) => {
  const { children, ...rest } = props;
  const store = configureStore();

  return (
    <CoreAdminContext {...rest}>
      <AppContextProvider>
        <Provider store={store}>
          <AppThemeProvider>
            <AppStyleProvider>
              <>
                <CssBaseline />
                {children}
              </>
            </AppStyleProvider>
          </AppThemeProvider>
        </Provider>
      </AppContextProvider>
    </CoreAdminContext>
  );
};

AdminContext.displayName = "AdminContext";
