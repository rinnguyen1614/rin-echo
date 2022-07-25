import { MessageType } from "../../../services/db/apps/chat/connectionList";

export interface FriendRequestObj {
  id: number;
  profilePic: string;
  name: string;
  date: any;
}

export interface ImageObj {
  id: number;
  thumb: string;
}

export interface RecentNewsObj {
  id: number;
  user: {
    name: string;
    profilePic: string;
  };
  title: string;
  desc: string;
}

export interface WhoToFollowObj {
  id: number;
  title: string;
  subTitle: string;
  profilePic: string;
}

export interface SuggestionObj {
  id: number;
  name: string;
  desc: string;
  thumb: string;
}

export interface MediaObj {
  id: number | string;
  url: string;
  mime_type: string;
}

export interface AttachmentObj {
  id: number;
  path: string;
  preview: string;
  metaData: {
    type: string;
    size: number;
  };
}

export interface CommentObj {
  id?: number;
  author: {
    name: string;
    profilePic: string;
    id: number;
  };
  liked?: boolean;
  comment: string;
  message_type: MessageType;
  date?: any;
  media?: MediaObj;
}

export interface AbutData {
  id: number;
  icon: string;
  text: string;
  linkType: string;
}

export interface SuggestTeamData {
  icon: string;
  title: string;
  subTitle: string;
  mediaImg: string;
}

export interface StoriesData {
  id: number;
  avatar: string;
  title: string;
  imgSrc: string;
}

export interface WhatsHappenData {
  id: number;
  imgSrc: string;
  title: string;
  subTitle: string;
  tag: Tag[];
}

export interface Tag {
  id: number;
  name?: string;
}

export interface WallData {
  id: number;
  name: string;
  profilePic: string;
  videoCall: {
    users: {
      id: number;
      name: string;
      profilePic: string;
    }[];
    title: string;
  };
  whatsHappen: WhatsHappenData[];
  suggestTeam: SuggestTeamData;
  stories: StoriesData[];
  about: AbutData[];
  friendRequests: FriendRequestObj[];
  photos: ImageObj[];
  recentNews: RecentNewsObj[];
  whoToFollow: WhoToFollowObj[];
  suggestions: SuggestionObj[];
}

export interface UserObj {
  name: string;
  profilePic: string;
  id: number;
}

export interface PostObj {
  id: number;
  owner: UserObj;
  date: any;
  attachments: AttachmentObj[];
  message?: string;
  liked: boolean;
  likes: number;
  shares: number;
  views: number;
  comments: CommentObj[];
  content?: string;
}
