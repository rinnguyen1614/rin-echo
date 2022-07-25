import {
  FolderObj,
  LabelObj,
  PriorityObj,
  StaffObj,
  TodoObj,
} from "../models/apps/Todo";

export const CREATE_NEW_TASK = "CREATE_NEW_TASK";
export const GET_TASK_DETAIL = "GET_TASK_DETAIL";
export const GET_TASK_LIST = "GET_TASK_LIST";
export const GET_TODO_FOLDER_LIST = "GET_TODO_FOLDER_LIST";
export const GET_TODO_LABEL_LIST = "GET_TODO_LABEL_LIST";
export const GET_TODO_PRIORITY_LIST = "GET_TODO_PRIORITY_LIST";
export const GET_TODO_STAFF_LIST = "GET_TODO_STAFF_LIST";
export const GET_TODO_STATUS_LIST = "GET_TODO_STATUS_LIST";
export const TOGGLE_TODO_DRAWER = "TOGGLE_TODO_DRAWER";
export const UPDATE_TASK_DETAIL = "UPDATE_TASK_DETAIL";
export const UPDATE_TASK_FOLDER = "UPDATE_TASK_FOLDER";
export const UPDATE_TASK_LABEL = "UPDATE_TASK_LABEL";
export const UPDATE_TASK_STARRED_STATUS = "UPDATE_TASK_STARRED_STATUS";

export interface CreateNewTaskActions {
  type: typeof CREATE_NEW_TASK;
  payload: TodoObj;
}

export interface GetTaskDetailActions {
  type: typeof GET_TASK_DETAIL;
  payload: TodoObj;
}

export interface GetTaskListActions {
  type: typeof GET_TASK_LIST;
  payload: {
    list: TodoObj[];
    total: number;
  };
}

export interface GetTodoFolderListAction {
  type: typeof GET_TODO_FOLDER_LIST;
  payload: FolderObj[];
}

export interface GetTodoLabelListAction {
  type: typeof GET_TODO_LABEL_LIST;
  payload: LabelObj[];
}

export interface GetTodoPriorityListAction {
  type: typeof GET_TODO_PRIORITY_LIST;
  payload: PriorityObj[];
}

export interface GetTodoStaffListAction {
  type: typeof GET_TODO_STAFF_LIST;
  payload: StaffObj[];
}

export interface ToggleTodoDrawerActions {
  type: typeof TOGGLE_TODO_DRAWER;
}

export interface UpdateTaskDetailAction {
  type: typeof UPDATE_TASK_DETAIL;
  payload: TodoObj;
}

export interface UpdateTodoFolderAction {
  type: typeof UPDATE_TASK_FOLDER;
  payload: { list: TodoObj[]; total: number };
}

export interface UpdateTodoLabelAction {
  type: typeof UPDATE_TASK_LABEL;
  payload: TodoObj[];
}

export interface UpdateTaskStaredAction {
  type: typeof UPDATE_TASK_STARRED_STATUS;
  payload: { data: TodoObj[]; folderName: string };
}

export interface GetTaskListStatusAction {
  type: typeof GET_TODO_STATUS_LIST;
  payload: TodoObj;
}

export type TaskActions =
  | CreateNewTaskActions
  | GetTaskDetailActions
  | GetTaskListActions
  | GetTodoFolderListAction
  | GetTodoLabelListAction
  | GetTodoPriorityListAction
  | GetTodoStaffListAction
  | ToggleTodoDrawerActions
  | UpdateTaskDetailAction
  | UpdateTodoFolderAction
  | UpdateTodoLabelAction
  | GetTaskListStatusAction
  | UpdateTaskStaredAction;
