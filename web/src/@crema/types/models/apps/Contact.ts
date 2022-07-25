export interface ContactObj {
  id: number;
  name: string;
  email: string;
  contact: string;
  company: string;
  isStarred: boolean;
  isFrequent: boolean;
  label: number;
  image: string;
  address?: string;
  facebookId?: string;
  twitterId?: string;
  notes?: string;
  website?: string;
  birthday?: string;
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
}
