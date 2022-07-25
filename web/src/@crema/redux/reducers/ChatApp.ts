import {
  ADD_NEW_MESSAGE,
  ChatActions,
  DELETE_MESSAGE,
  DELETE_USER_MESSAGES,
  EDIT_MESSAGE,
  GET_CONNECTIONS_LIST,
  GET_USER_MESSAGES,
  SELECT_USER,
  TOGGLE_CHAT_DRAWER,
} from "../../types/actions/Chat.actions";
import { ConnectionObj, MessageObj } from "../../types/models/apps/Chat";

const initialState: {
  connectionList: ConnectionObj[];
  chatDrawer: boolean;
  userMessages: MessageObj | null;
  selectedUser: ConnectionObj | null;
} = {
  connectionList: [],
  chatDrawer: false,
  userMessages: null,
  selectedUser: null,
};

const ChatApp = (state = initialState, action: ChatActions) => {
  switch (action.type) {
    case GET_CONNECTIONS_LIST:
      return {
        ...state,
        connectionList: action.payload,
      };

    case TOGGLE_CHAT_DRAWER:
      return {
        ...state,
        chatDrawer: !state.chatDrawer,
      };

    case GET_USER_MESSAGES:
      return {
        ...state,
        userMessages: action.payload,
      };

    case ADD_NEW_MESSAGE: {
      return {
        ...state,
        connectionList: state.connectionList.map((item) =>
          item.id === action.payload.data.user.id
            ? action.payload.data.user
            : item
        ),
        userMessages: action.payload.data.userMessages,
      };
    }

    case EDIT_MESSAGE: {
      return {
        ...state,
        connectionList: state.connectionList.map((item) =>
          item.id === action.payload.data.user.id
            ? action.payload.data.user
            : item
        ),
        userMessages: action.payload.data.userMessages,
      };
    }

    case DELETE_MESSAGE: {
      return {
        ...state,
        connectionList: state.connectionList.map((item) =>
          item.id === action.payload.user.id ? action.payload.user : item
        ),
        userMessages: action.payload.userMessages,
      };
    }

    case DELETE_USER_MESSAGES: {
      return {
        ...state,
        connectionList: state.connectionList.map((item) =>
          item.id === action.payload.id ? action.payload : item
        ),
        userMessages: null,
        selectedUser: null,
      };
    }

    case SELECT_USER: {
      return {
        ...state,
        selectedUser: action.payload,
        userMessages: null,
      };
    }

    default:
      return state;
  }
};
export default ChatApp;
