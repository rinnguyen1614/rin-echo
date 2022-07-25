import { MessageType } from "../../../services/db/apps/chat/connectionList";

export interface Member {
  id: number;
  name: string;
  image: string;
  status: string;
  username?: string;
}

export interface ConnectionObj {
  id: number;
  channelId: number;
  name: string;
  image: string;
  status: string;
  username: string;
  isGroup?: boolean;
  members?: Member[];
  lastMessage?: {
    id: number;
    message: string;
    type: string;
    time: string;
  };
}

export interface MediaObj {
  id: string | number;
  url: string;
  mime_type: string;
  file_name: string;
  file_size?: number;
}

export interface MessageDataObj {
  id?: number;
  sender: number | string;
  message?: string;
  message_type: MessageType;
  time: string;
  edited?: boolean;
  media?: MediaObj[];
}

export interface MessageObj {
  channelId: number;
  messageData: MessageDataObj[];
}
