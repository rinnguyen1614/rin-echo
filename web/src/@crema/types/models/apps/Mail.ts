export interface ConnectionObj {
  id: number;
  name: string;
  status: string;
  image: string;
  email: string;
  address: string;
  designation: string;
}

export interface FolderObj {
  id: number;
  name: string;
  alias: string;
  icon?: string;
}

export interface LabelObj {
  id: number;
  name: string;
  alias: string;
  color: string;
  icon?: any;
}

export interface Sender {
  id: number;
  name: string;
  email: string;
  profilePic: string;
}

export interface Message {
  messageId: number;
  description: string;
  sender: Sender;
  to: Sender[];
  cc: any[];
  bcc: any[];
  isStarred: boolean;
  sentOn: string;
}

export interface MailObj {
  id: number;
  isChecked?: boolean;
  isStarred?: boolean;
  isReplied?: boolean;
  label: LabelObj;
  sentBy?: string;
  subject: string;
  hasAttachments: boolean;
  sentAt?: string;
  messages?: Message[];
  isRead?: boolean;
  folderValue: number;
}
