import { Identity } from "@app/types";
import { AuthProvider } from "react-admin";
import { KEY_TOKEN } from "../../constants";
import { Jwt } from "../../utils/jwt";
import { changePassword } from "./changePassword";
import { checkError } from "./checkError";
import { getIdentity, getIdentityStore } from "./getIdentity";
import { getMenus } from "./getMenus";
import { getPermissions } from "./getPermissions";
import { login } from "./login";
import { updateIdentity, updateIdentityStore } from "./updateIdentity";

const authProvider: AuthProvider = {
  // authentication
  login: async (params: any) => {
    try {
      await login(params);

      // get info of Identity from server and update to store
      var identity = await getIdentity();

      await updateIdentityStore(identity);

      return Promise.resolve();
    } catch (err) {
      return Promise.reject(err);
    }
  },
  logout: () => {
    localStorage.removeItem(KEY_TOKEN);
    return Promise.resolve();
  },
  checkError,
  checkAuth: () => {
    const token = localStorage.getItem(KEY_TOKEN);
    if (!token) {
      return Promise.reject({ message: "ra.auth.auth_check_error" });
    }

    return Jwt.isTokenExpired(token)
      ? Promise.reject({ message: "ra.auth.auth_check_error" })
      : Promise.resolve();
  },
  getIdentity: getIdentityStore,
  updateIdentity: async (params: Identity) => {
    try {
      // send to server
      var result = await updateIdentity(params);
      // update to store
      await updateIdentityStore(result);
      return Promise.resolve(result);
    } catch (err) {
      return Promise.reject(err);
    }
  },
  changePassword,
  // authorization
  getPermissions: getPermissions,
  getMenus: getMenus,
};

export default authProvider;
