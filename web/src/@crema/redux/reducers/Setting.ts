import { Setting } from "../../types/models/Setting";
import {
  NAV_COLLAPSED,
  ROUTE_CHANGE,
  SET_INITIAL_PATH,
  SettingsActionTypes,
  TOGGLE_NAV_COLLAPSED,
} from "../../types/actions/Settings.action";

const initialSettings: Setting = {
  navCollapsed: false,
  initialPath: "/",
};

const Settings = (state = initialSettings, action: SettingsActionTypes) => {
  switch (action.type) {
    case ROUTE_CHANGE:
      return {
        ...state,
        navCollapsed: false,
      };

    case TOGGLE_NAV_COLLAPSED:
      return {
        ...state,
        navCollapsed: !state.navCollapsed,
      };

    case SET_INITIAL_PATH:
      return {
        ...state,
        initialPath: action.initialPath,
      };

    case NAV_COLLAPSED:
      return {
        ...state,
        navCollapsed: false,
      };
    default:
      return state;
  }
};

export default Settings;
