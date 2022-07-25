export const NAV_COLLAPSED = "NAV_COLLAPSED";
export const TOGGLE_NAV_COLLAPSED = "TOGGLE_NAV_COLLAPSED";
export const SET_INITIAL_PATH = "SET_INITIAL_PATH";
export const ROUTE_CHANGE = "@@router/LOCATION_CHANGE";

export interface ToggleNavCollapsedAction {
  type: typeof TOGGLE_NAV_COLLAPSED;
}
export interface NavCollapsedAction {
  type: typeof NAV_COLLAPSED;
}

export interface SetInitialPathAction {
  type: typeof SET_INITIAL_PATH;
  initialPath: string | undefined;
}

export interface RouteChangeAction {
  type: typeof ROUTE_CHANGE;
}

export type SettingsActionTypes =
  | ToggleNavCollapsedAction
  | SetInitialPathAction
  | NavCollapsedAction
  | RouteChangeAction;
