import {
  ConnectionObj,
  FolderObj,
  LabelObj,
  MailObj,
} from "../models/apps/Mail";

export const GET_CONNECTION_LIST = "GET_CONNECTION_LIST";
export const GET_FOLDER_LIST = "GET_FOLDER_LIST";
export const GET_LABEL_LIST = "GET_LABEL_LIST";
export const COMPOSE_MAIL = "COMPOSE_MAIL";
export const GET_MAIL_DETAIL = "GET_MAIL_DETAIL";
export const GET_MAIL_LIST = "GET_MAIL_LIST";
export const NULLIFY_MAIL = "NULLIFY_MAIL";
export const TOGGLE_MAIL_DRAWER = "TOGGLE_MAIL_DRAWER";
export const UPDATE_MAIL_FOLDER = "UPDATE_MAIL_FOLDER";
export const UPDATE_MAIL_LABEL = "UPDATE_MAIL_LABEL";
export const UPDATE_STARRED_STATUS = "UPDATE_STARRED_STATUS";
export const UPDATED_MAIL_DETAIL = "UPDATED_MAIL_DETAIL";
export const CHANGE_READ_STATUS = "CHANGE_READ_STATUS";

export interface GetMailConnectionListActions {
  type: typeof GET_CONNECTION_LIST;
  payload: ConnectionObj[];
}

export interface GetMailFolderActions {
  type: typeof GET_FOLDER_LIST;
  payload: FolderObj[];
}

export interface GetMailLabelActions {
  type: typeof GET_LABEL_LIST;
  payload: LabelObj[];
}

export interface ComposeMailAction {
  type: typeof COMPOSE_MAIL;
  payload: { data: MailObj; pathname: string };
}

export interface GetMailDetailAction {
  type: typeof GET_MAIL_DETAIL;
  payload: MailObj;
}

export interface GetMailListAction {
  type: typeof GET_MAIL_LIST;
  payload: {
    list: MailObj[];
    total: number;
  };
}

export interface RemoveSelectedMailAction {
  type: typeof NULLIFY_MAIL;
}

export interface ToggleMailDrawerActions {
  type: typeof TOGGLE_MAIL_DRAWER;
}

export interface UpdateMailFolderAction {
  type: typeof UPDATE_MAIL_FOLDER;
  payload: number[];
}

export interface UpdateMailLabelAction {
  type: typeof UPDATE_MAIL_LABEL;
  payload: MailObj[];
}

export interface UpdateMailStaredAction {
  type: typeof UPDATE_STARRED_STATUS;
  payload: { data: MailObj[]; folderName: string };
}

export interface UpdateMailReadStatusAction {
  type: typeof CHANGE_READ_STATUS;
  payload: MailObj[];
}

export interface UpdateMailDetailAction {
  type: typeof UPDATED_MAIL_DETAIL;
  payload: MailObj;
}

export type MailActions =
  | GetMailConnectionListActions
  | GetMailFolderActions
  | GetMailLabelActions
  | ComposeMailAction
  | RemoveSelectedMailAction
  | GetMailDetailAction
  | GetMailListAction
  | ToggleMailDrawerActions
  | UpdateMailFolderAction
  | UpdateMailLabelAction
  | UpdateMailStaredAction
  | UpdateMailReadStatusAction
  | UpdateMailDetailAction;
