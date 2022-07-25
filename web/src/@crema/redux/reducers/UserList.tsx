import { UserList } from "../../services/db/userList";
import { GET_USER_LIST } from "../../types/actions/UserList.actions";
import { AppActions } from "../../types";

let initialState: { usersList: UserList[] };
initialState = {
  usersList: [],
};

const userListReducer = (state = initialState, action: AppActions) => {
  switch (action.type) {
    case GET_USER_LIST:
      return {
        ...state,
        usersList: action.payload,
      };

    default:
      return state;
  }
};
export default userListReducer;
