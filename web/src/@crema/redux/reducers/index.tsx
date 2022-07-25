import Settings from "./Setting";
import Common from "./Common";
import Dashboard from "./Dashboard";
import Ecommerce from "./Ecommerce";
import ChatApp from "./ChatApp";
import MailApp from "./MailApp";
import ScrumBoard from "./ScrumboardApp";
import ContactApp from "./ContactApp";
import WallApp from "./WallApp";
import ToDoApp from "./ToDoApp";
import UserList from "./UserList";

const reducers = {
  settings: Settings,
  dashboard: Dashboard,
  ecommerce: Ecommerce,
  common: Common,
  chatApp: ChatApp,
  mailApp: MailApp,
  contactApp: ContactApp,
  scrumboardApp: ScrumBoard,
  wallApp: WallApp,
  todoApp: ToDoApp,
  userList: UserList,
};

export default reducers;
