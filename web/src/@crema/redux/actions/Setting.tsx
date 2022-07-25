import {
  NAV_COLLAPSED,
  SET_INITIAL_PATH,
  SettingsActionTypes,
  TOGGLE_NAV_COLLAPSED,
} from "../../types/actions/Settings.action";

export const toggleNavCollapsed = () => ({ type: TOGGLE_NAV_COLLAPSED });
export const onNavCollapsed = () => ({ type: NAV_COLLAPSED });

export const setInitialPath = (initialPath: string): SettingsActionTypes => ({
  type: SET_INITIAL_PATH,
  initialPath,
});
