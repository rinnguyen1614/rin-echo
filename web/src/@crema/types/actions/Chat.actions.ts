import { ConnectionObj, MessageObj } from "../models/apps/Chat";

export const ADD_NEW_MESSAGE = "ADD_NEW_MESSAGE";
export const DELETE_MESSAGE = "DELETE_MESSAGE";
export const DELETE_USER_MESSAGES = "DELETE_USER_MESSAGES";
export const EDIT_MESSAGE = "EDIT_MESSAGE";
export const GET_CONNECTIONS_LIST = "GET_CONNECTIONS_LIST";
export const GET_USER_MESSAGES = "GET_USER_MESSAGES";
export const SELECT_USER = "SELECT_USER";
export const TOGGLE_CHAT_DRAWER = "TOGGLE_CHAT_DRAWER";

export interface AddNewMessageActions {
  type: typeof ADD_NEW_MESSAGE;
  payload: {
    data: {
      user: ConnectionObj;
      userMessages: MessageObj;
    };
  };
}

export interface DeleteMessageActions {
  type: typeof DELETE_MESSAGE;
  payload: {
    user: ConnectionObj;
    userMessages: MessageObj;
  };
}

export interface DeleteUserMessageActions {
  type: typeof DELETE_USER_MESSAGES;
  payload: ConnectionObj;
}

export interface EditMessageActions {
  type: typeof EDIT_MESSAGE;
  payload: {
    data: {
      user: ConnectionObj;
      userMessages: MessageObj;
    };
  };
}

export interface GetConnectionListActions {
  type: typeof GET_CONNECTIONS_LIST;
  payload: ConnectionObj[];
}

export interface GetUserMessageActions {
  type: typeof GET_USER_MESSAGES;
  payload: MessageObj[];
}

export interface SelectUserActions {
  type: typeof SELECT_USER;
  payload: ConnectionObj;
}

export interface ToggleChatDrawerActions {
  type: typeof TOGGLE_CHAT_DRAWER;
}

export type ChatActions =
  | AddNewMessageActions
  | DeleteMessageActions
  | DeleteUserMessageActions
  | EditMessageActions
  | GetConnectionListActions
  | GetUserMessageActions
  | SelectUserActions
  | ToggleChatDrawerActions;
