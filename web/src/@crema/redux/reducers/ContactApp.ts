import {
  ContactActions,
  CREATE_NEW_CONTACT,
  DELETE_CONTACT,
  GET_CONTACT_FOLDER_LIST,
  GET_CONTACT_LABEL_LIST,
  GET_CONTACT_LIST,
  TOGGLE_CONTACT_DRAWER,
  UPDATE_CONTACT_DETAIL,
  UPDATE_CONTACT_LABEL,
  UPDATE_CONTACT_STARRED_STATUS,
} from "../../types/actions/Contact.actions";
import {
  ContactObj,
  FolderObj,
  LabelObj,
} from "../../types/models/apps/Contact";

const initialState: {
  contactList: ContactObj[];
  totalContacts: number;
  contactDrawer: false;
  labelList: LabelObj[];
  folderList: FolderObj[];
  selectedContact: ContactObj | null;
} = {
  contactList: [],
  totalContacts: 0,
  contactDrawer: false,
  labelList: [],
  folderList: [],
  selectedContact: null,
};

const ContactApp = (state = initialState, action: ContactActions) => {
  switch (action.type) {
    case GET_CONTACT_LIST:
      return {
        ...state,
        contactList: action.payload.list,
        totalContacts: action.payload.total,
      };

    case GET_CONTACT_FOLDER_LIST:
      return {
        ...state,
        folderList: action.payload,
      };

    case TOGGLE_CONTACT_DRAWER:
      return {
        ...state,
        contactDrawer: !state.contactDrawer,
      };

    case GET_CONTACT_LABEL_LIST:
      return {
        ...state,
        labelList: action.payload,
      };

    case CREATE_NEW_CONTACT:
      return {
        ...state,
        contactList: [action.payload, ...state.contactList],
        totalContacts: state.totalContacts + 1,
      };

    case DELETE_CONTACT: {
      return {
        ...state,
        contactList: action.payload.list,
        totalContacts: action.payload.total,
      };
    }

    case UPDATE_CONTACT_LABEL: {
      let contactIds = action.payload.data.map((contact) => contact.id);
      const updatedList = state.contactList.map((contact) => {
        if (contactIds.includes(contact.id)) {
          return action.payload.data.find(
            (selectedContact) => selectedContact.id === contact.id
          );
        } else {
          return contact;
        }
      });
      const filteredList =
        action.payload.labelName === "label"
          ? updatedList.filter(
              (item) => item!.label !== action.payload.labelType
            )
          : updatedList;
      const total = filteredList.length;
      return {
        ...state,
        contactList: filteredList,
        totalContacts: total,
      };
    }

    case UPDATE_CONTACT_STARRED_STATUS: {
      let contactIds = action.payload.data.map((contact) => contact.id);
      const updatedList = state.contactList.map((contact) => {
        if (contactIds.includes(contact.id)) {
          return action.payload.data.find(
            (selectedContact) => selectedContact.id === contact.id
          );
        } else {
          return contact;
        }
      });
      const filteredList =
        action.payload.folderName === "starred"
          ? updatedList.filter((item) => item!.isStarred)
          : updatedList;
      const total =
        action.payload.folderName === "starred"
          ? state.totalContacts - action.payload.data.length
          : state.totalContacts;
      return {
        ...state,
        contactList: filteredList,
        totalContacts: total,
      };
    }

    case UPDATE_CONTACT_DETAIL:
      return {
        ...state,
        selectedContact: action.payload,
        contactList: state.contactList.map((contact) =>
          contact.id === action.payload.id ? action.payload : contact
        ),
      };

    default:
      return state;
  }
};
export default ContactApp;
