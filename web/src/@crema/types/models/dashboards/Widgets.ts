import { PaletteColorOptions } from "@mui/material";

export interface RecentActivityData {
  id: number;
  image: string;
  name: string;
  message: string;
}

export interface CategoriesData {
  id: number;
  name: string;
  isChecked: boolean;
}

export interface ProfileData {
  id: number;
  name: string;
  image: string;
  photos: number;
  followers: number;
  following: number;
}

export interface MessagesData {
  id: number;
  image: string;
  message: string;
  name: string;
}

export interface TaskListData {
  id: number;
  title: string;
  desc: string;
}

export interface ColorsList {
  id: number;
  name: string;
  color: PaletteColorOptions | string;
  isChecked: boolean;
}

export interface TagsList {
  id: number;
  label: string;
  color: string;
}

export interface ReviewsList {
  id: number;
  rating: number;
  by: string;
  content: string;
  time: string;
}

export interface SocialInfo {
  image: string;
  name: string;
  id: string;
  desc: string;
}

export interface MateInfo {
  facebookInfo: SocialInfo;
  twitterInfo: SocialInfo;
}

export interface FormatList {
  id: number;
  name: string;
}

export interface Temperatures {
  id: number;
  day: string;
  image: string;
}

export interface CityData {
  id: number;
  name: string;
  desc: string;
  image: string;
}

export interface Widgets {
  recentActivity: RecentActivityData[];
  categories: CategoriesData[];
  profile: ProfileData;
  messages: MessagesData[];
  taskList: TaskListData[];
  colorsList: ColorsList[];
  tagsList: TagsList[];
  reviewsList: ReviewsList[];
  mateInfo: MateInfo;
  formatList: FormatList[];
  temperatures: Temperatures[];
  cityData: CityData[];
}
