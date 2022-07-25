import { ContactObj, FolderObj, LabelObj } from "../models/apps/Contact";

export const CREATE_NEW_CONTACT = "CREATE_NEW_CONTACT";
export const DELETE_CONTACT = "DELETE_CONTACT";
export const GET_CONTACT_FOLDER_LIST = "GET_CONTACT_FOLDER_LIST";
export const GET_CONTACT_LABEL_LIST = "GET_CONTACT_LABEL_LIST";
export const GET_CONTACT_LIST = "GET_CONTACT_LIST";
export const UPDATE_CONTACT_STARRED_STATUS = "UPDATE_CONTACT_STARRED_STATUS";
export const UPDATE_CONTACT_DETAIL = "UPDATE_CONTACT_DETAIL";
export const UPDATE_CONTACT_LABEL = "UPDATE_CONTACT_LABEL";
export const TOGGLE_CONTACT_DRAWER = "TOGGLE_CONTACT_DRAWER";

export interface GetContactFolderActions {
  type: typeof GET_CONTACT_FOLDER_LIST;
  payload: FolderObj[];
}

export interface GetContactLabelActions {
  type: typeof GET_CONTACT_LABEL_LIST;
  payload: LabelObj;
}

export interface GetContactsActions {
  type: typeof GET_CONTACT_LIST;
  payload: { list: ContactObj[]; total: number };
}

export interface UpdateContactStarActions {
  type: typeof UPDATE_CONTACT_STARRED_STATUS;
  payload: { data: ContactObj[]; folderName: string };
}

export interface UpdateContactActions {
  type: typeof UPDATE_CONTACT_DETAIL;
  payload: ContactObj;
}

export interface UpdateContactLabelActions {
  type: typeof UPDATE_CONTACT_LABEL;
  payload: { data: ContactObj[]; labelName: string; labelType: number };
}

export interface CreateContactActions {
  type: typeof CREATE_NEW_CONTACT;
  payload: ContactObj;
}

export interface DeleteContactActions {
  type: typeof DELETE_CONTACT;
  payload: { list: ContactObj[]; total: number };
}

export interface ToggleContactDrawerActions {
  type: typeof TOGGLE_CONTACT_DRAWER;
}

export type ContactActions =
  | GetContactFolderActions
  | GetContactLabelActions
  | GetContactsActions
  | UpdateContactStarActions
  | UpdateContactActions
  | UpdateContactLabelActions
  | CreateContactActions
  | ToggleContactDrawerActions
  | DeleteContactActions;
