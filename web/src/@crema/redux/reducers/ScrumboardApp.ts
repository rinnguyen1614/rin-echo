import {
  ADD_BOARD_LIST,
  ADD_LIST_CARD,
  ADD_NEW_BOARD,
  DELETE_BOARD,
  DELETE_LIST,
  DELETE_LIST_CARD,
  EDIT_BOARD_DETAIL,
  EDIT_LIST_CARD,
  GET_BOARD_DETAIL,
  GET_BOARDS,
  GET_MEMBER_LIST,
  GET_SCRUM_LABEL_LIST,
  ScrumboardActions,
} from "../../types/actions/Scrumboard.actions";
import {
  BoardObj,
  LabelObj,
  MemberObj,
} from "../../types/models/apps/ScrumbBoard";

const initialState: {
  boardList: BoardObj[];
  labelList: LabelObj[];
  memberList: MemberObj[];
  boardDetail: BoardObj | null;
} = {
  boardList: [],
  labelList: [],
  memberList: [],
  boardDetail: null,
};

const ScrumBoardApp = (state = initialState, action: ScrumboardActions) => {
  switch (action.type) {
    case GET_MEMBER_LIST:
      return {
        ...state,
        memberList: action.payload,
      };

    case GET_SCRUM_LABEL_LIST:
      return {
        ...state,
        labelList: action.payload,
      };

    case GET_BOARDS:
      return {
        ...state,
        boardList: action.payload,
      };

    case GET_BOARD_DETAIL:
      return {
        ...state,
        boardDetail: action.payload,
      };

    case ADD_NEW_BOARD:
      return {
        ...state,
        boardList: state.boardList.concat(action.payload),
      };

    case EDIT_BOARD_DETAIL: {
      return {
        ...state,
        boardList: state.boardList.map((board) =>
          board.id === action.payload.id ? action.payload : board
        ),
      };
    }

    case DELETE_BOARD:
      return {
        ...state,
        boardList: state.boardList.filter(
          (board) => board.id !== action.payload
        ),
      };

    case DELETE_LIST:
      return {
        ...state,
        boardDetail: action.payload,
      };

    case ADD_BOARD_LIST:
      return {
        ...state,
        boardDetail: action.payload,
      };

    // case EDIT_BOARD_LIST:
    //   return {
    //     ...state,
    //     boardDetail: action.payload,
    //   };

    case ADD_LIST_CARD:
      return {
        ...state,
        boardDetail: action.payload,
      };

    case EDIT_LIST_CARD:
      return {
        ...state,
        boardDetail: action.payload,
      };

    case DELETE_LIST_CARD:
      return {
        ...state,
        boardDetail: action.payload,
      };

    default:
      return state;
  }
};
export default ScrumBoardApp;
