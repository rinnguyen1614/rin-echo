import { PostObj, WallData } from "../models/apps/Wall";

export const CREATE_NEW_POST = "CREATE_NEW_POST";
export const GET_FEED_POSTS = "GET_FEED_POSTS";
export const GET_WALL_DATA = "GET_WALL_DATA";
export const UPDATE_POST = "UPDATE_POST";

export interface CreateNewPostActions {
  type: typeof CREATE_NEW_POST;
  payload: PostObj;
}

export interface GetFeedPostActions {
  type: typeof GET_FEED_POSTS;
  payload: PostObj[];
}

export interface GetWallDataActions {
  type: typeof GET_WALL_DATA;
  payload: WallData;
}

export interface UpdatePostActions {
  type: typeof UPDATE_POST;
  payload: PostObj;
}

export type WalltActions =
  | CreateNewPostActions
  | GetFeedPostActions
  | GetWallDataActions
  | UpdatePostActions;
