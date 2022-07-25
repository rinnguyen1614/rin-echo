export interface LabelObj {
  id: number;
  name: string;
  type: number;
  color: string;
}

export interface MemberObj {
  id: number;
  name: string;
  image: string;
}

export interface CheckedListObj {
  id: number;
  title: string;
}

export interface AttachmentObj {
  id: number;
  file: {
    path: string;
    name: string;
    lastModified: number;
    lastModifiedDate: string;
  };
  preview: string;
}

export interface CardObj {
  id: number;
  title: string;
  attachments: AttachmentObj[];
  label: LabelObj[];
  date: any;
  comments: any[];
  desc: string;
  members: MemberObj[];
  checkedList: CheckedListObj[];
}

export interface CardListObj {
  id: number;
  name: string;
  cards: CardObj[];
}

export interface BoardObj {
  id: number;
  name: string;
  list: CardListObj[];
}
