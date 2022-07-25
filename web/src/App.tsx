import { Admin } from "Admin";
import { CustomRoutes, Resource } from "react-admin";
import { BrowserRouter, Route } from "react-router-dom";
import authProvider from "./providers/authProvider";
import dataProviderFactory from "./providers/dataProvider";
import { AppContextProvider, AppStyleProvider, AppThemeProvider } from "@crema";
import Login from "pages/account/Login";
import { Provider } from "react-redux";
import configureStore from "@crema/redux/store";
import { CssBaseline } from "@mui/material";
import resources from "./pages/resources";
import menus from "./pages/menus";
import roles from "./pages/roles";
import i18nProvider from "i18n";
import UserProfile from "pages/profile/UserProfile";
import Layout from "layout/Layout";
import users from "pages/users";
import audit_logs from "pages/audit_logs";

const store = configureStore();

const App = () => {
  return (
    <AppContextProvider>
      <Provider store={store}>
        <AppThemeProvider>
          <AppStyleProvider>
            <>
              <CssBaseline />
              <BrowserRouter>
                <Admin
                  title="Rin"
                  dataProvider={dataProviderFactory(
                    process.env.REACT_APP_DATA_PROVIDER || ""
                  )}
                  authProvider={authProvider}
                  layout={Layout}
                  i18nProvider={i18nProvider}
                  loginPage={Login}
                  disableTelemetry
                >
                  <CustomRoutes>
                    <Route path="/account" element={<UserProfile />} />
                  </CustomRoutes>
                  <Resource name="admin/resources" {...resources} />
                  <Resource name="admin/menus" {...menus} />
                  <Resource name="admin/roles" {...roles} />
                  <Resource name="admin/users" {...users} />
                  <Resource name="admin/audit_logs" {...audit_logs} />
                </Admin>
              </BrowserRouter>
            </>
          </AppStyleProvider>
        </AppThemeProvider>
      </Provider>
    </AppContextProvider>
  );
};

export default App;
