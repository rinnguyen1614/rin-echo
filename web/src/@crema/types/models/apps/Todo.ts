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
  label: string;
  value: string;
  color: string;
}

export interface PriorityObj {
  id: number;
  name: string;
  type: number;
  color: string;
}

export interface StaffObj {
  id: number;
  name: string;
  image: string;
}

export interface StatusObj {
  id: number;
  name: string;
  type: number;
}

export interface CommentObj {
  comment: string;
  name: string;
  image: string;
  date: string;
  time?: string;
}

export interface TodoObj {
  id: number;
  title: string;
  isStarred?: boolean;
  label: LabelObj[];
  priority: PriorityObj;
  isAttachment: boolean;
  sentAt: string;
  folderValue: number;
  scheduleMobile: string;
  image: string;
  assignedTo: StaffObj;
  createdBy: {
    name: string;
    image?: string;
  };
  createdOn: string;
  startDate: string;
  status: number;
  comments: CommentObj[];
  content: string;
  isReplied?: boolean;
  isRead?: boolean;
  date?: string;
}
