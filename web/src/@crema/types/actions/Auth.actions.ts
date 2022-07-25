import { AuthUser } from "../models/AuthUser";

export const UPDATE_AUTH_USER = "UPDATE_AUTH_USER";
export const SET_AUTH_TOKEN = "SET_AUTH_TOKEN";
export const SIGNOUT_AUTH_SUCCESS = "SIGNOUT_AUTH_SUCCESS";

export interface SetAuthTokenActions {
  type: typeof SET_AUTH_TOKEN;
  payload: string | null;
}

export interface UpdateAuthUserActions {
  type: typeof UPDATE_AUTH_USER;
  payload: AuthUser | null;
}

export interface SignoutAuthUserActions {
  type: typeof SIGNOUT_AUTH_SUCCESS;
}

export type AuthActions =
  | UpdateAuthUserActions
  | SetAuthTokenActions
  | SignoutAuthUserActions;
