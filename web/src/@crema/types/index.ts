import { CommonActionTypes } from "./actions/Common.action";
import { SettingsActionTypes } from "./actions/Settings.action";
import { DashboardActionTypes } from "./actions/Dashboard.action";
import { EcommerceActionTypes } from "./actions/Ecommerce.action";
import { AuthActions } from "./actions/Auth.actions";
import { ChatActions } from "./actions/Chat.actions";
import { ContactActions } from "./actions/Contact.actions";
import { MailActions } from "./actions/Mail.action";
import { TaskActions } from "./actions/Todo.action";
import { WalltActions } from "./actions/Wall.actions";
import { ScrumboardActions } from "./actions/Scrumboard.actions";
import { UserListActions } from "./actions/UserList.actions";

export type AppActions =
  | CommonActionTypes
  | SettingsActionTypes
  | DashboardActionTypes
  | EcommerceActionTypes
  | AuthActions
  | ChatActions
  | MailActions
  | TaskActions
  | WalltActions
  | ScrumboardActions
  | ContactActions
  | UserListActions;
