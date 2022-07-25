import { BoardObj, CardListObj } from "../models/apps/ScrumbBoard";

export const ADD_BOARD_LIST = "ADD_BOARD_LIST";
export const ADD_LIST_CARD = "ADD_LIST_CARD";
export const ADD_NEW_BOARD = "ADD_NEW_BOARD";
export const DELETE_BOARD = "DELETE_BOARD";
export const DELETE_LIST = "DELETE_LIST";
export const DELETE_LIST_CARD = "DELETE_LIST_CARD";
export const EDIT_BOARD_DETAIL = "EDIT_BOARD_DETAIL";
export const EDIT_BOARD_LIST = "EDIT_BOARD_LIST";
export const EDIT_LIST_CARD = "EDIT_LIST_CARD";
export const GET_BOARD_DETAIL = "GET_BOARD_DETAIL";
export const GET_BOARDS = "GET_BOARDS";
export const GET_MEMBER_LIST = "GET_MEMBER_LIST";
export const GET_SCRUM_LABEL_LIST = "GET_SCRUM_LABEL_LIST";

export interface AddBoardListActions {
  type: typeof ADD_BOARD_LIST;
  payload: BoardObj;
}

export interface AddListCardActions {
  type: typeof ADD_LIST_CARD;
  payload: CardListObj;
}

export interface AddNewBoardActions {
  type: typeof ADD_NEW_BOARD;
  payload: BoardObj;
}

export interface DeleteBoardActions {
  type: typeof DELETE_BOARD;
  payload: number;
}

export interface DeleteListActions {
  type: typeof DELETE_LIST;
  payload: number;
}

export interface DeleteListCardActions {
  type: typeof DELETE_LIST_CARD;
  payload: CardListObj;
}

export interface EditBoardDetailActions {
  type: typeof EDIT_BOARD_DETAIL;
  payload: BoardObj;
}

export interface EditBoardListActions {
  type: typeof EDIT_BOARD_LIST;
  payload: CardListObj;
}

export interface EditListCardActions {
  type: typeof EDIT_LIST_CARD;
  payload: CardListObj;
}

export interface GetBoardDetailActions {
  type: typeof GET_BOARD_DETAIL;
  payload: CardListObj;
}

export interface GetBoardListActions {
  type: typeof GET_BOARDS;
  payload: CardListObj;
}

export interface GetMemberListActions {
  type: typeof GET_MEMBER_LIST;
  payload: CardListObj;
}

export interface GetScrumLabelListActions {
  type: typeof GET_SCRUM_LABEL_LIST;
  payload: CardListObj;
}

export type ScrumboardActions =
  | AddBoardListActions
  | AddListCardActions
  | AddNewBoardActions
  | DeleteBoardActions
  | DeleteListActions
  | DeleteListCardActions
  | EditBoardDetailActions
  | EditBoardListActions
  | EditListCardActions
  | GetBoardDetailActions
  | GetBoardListActions
  | GetMemberListActions
  | GetScrumLabelListActions;
