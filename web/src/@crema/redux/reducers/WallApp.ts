import {
  CREATE_NEW_POST,
  GET_FEED_POSTS,
  GET_WALL_DATA,
  UPDATE_POST,
  WalltActions,
} from "../../types/actions/Wall.actions";
import { PostObj, WallData } from "../../types/models/apps/Wall";

const initialState: {
  wallData: WallData | null;
  postList: PostObj[];
} = {
  wallData: null,
  postList: [],
};

const WallApp = (state = initialState, action: WalltActions) => {
  switch (action.type) {
    case GET_WALL_DATA:
      return {
        ...state,
        wallData: action.payload,
      };
    case GET_FEED_POSTS: {
      return { ...state, postList: action.payload };
    }

    case CREATE_NEW_POST: {
      return {
        ...state,
        postList: [action.payload, ...state.postList],
      };
    }

    case UPDATE_POST: {
      return {
        ...state,
        postList: state.postList.map((item) =>
          item.id === action.payload.id ? action.payload : item
        ),
      };
    }
    default: {
      return state;
    }
  }
};
export default WallApp;
