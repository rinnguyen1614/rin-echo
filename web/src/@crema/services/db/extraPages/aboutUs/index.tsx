import React from "react";

import EditOutlinedIcon from "@mui/icons-material/EditOutlined";
import PhotoCameraOutlinedIcon from "@mui/icons-material/PhotoCameraOutlined";
import SearchIcon from "@mui/icons-material/Search";
import FormattedMessage from "../../../../utility/FormattedMessage";
import { blue, red, teal } from "@mui/material/colors";

export interface TeamData {
  id: number;
  name: string;
  position: string;
  image: string;
}

export interface AboutUsData {
  alias: string;
  title: any;
  avatarColor: any;
  icon: any;
  content: any;
}

export const aboutUsData: AboutUsData[] = [
  {
    alias: "branding",
    title: <FormattedMessage id="extra.branding" />,
    avatarColor: teal[600],
    icon: <EditOutlinedIcon />,
    content: <FormattedMessage id="extra.brandingContent" />,
  },
  {
    alias: "photography",
    title: <FormattedMessage id="extra.photography" />,
    avatarColor: red[500],
    icon: <PhotoCameraOutlinedIcon />,
    content: <FormattedMessage id="extra.brandingContent" />,
  },
  {
    alias: "seo",
    title: <FormattedMessage id="extra.seo" />,
    avatarColor: blue[500],
    icon: <SearchIcon />,
    content: <FormattedMessage id="extra.brandingContent" />,
  },
];

export const teamData: TeamData[] = [
  {
    id: 444,
    name: "Asantha Powel",
    position: "CEO",
    image: "/assets/images/teamImages/User1.png",
  },
  {
    id: 111,
    name: "Johna Taylor",
    position: "CTO",
    image: "/assets/images/teamImages/User4.png",
  },
  {
    id: 222,
    name: "Nick Campbell",
    position: "General Manager",
    image: "/assets/images/teamImages/User3.png",
  },
  {
    id: 333,
    name: "Johna Taylor",
    position: "CFO",
    image: "/assets/images/teamImages/User5.png",
  },
  {
    id: 555,
    name: "Ricardo Johnson",
    position: "Director",
    image: "/assets/images/teamImages/User2.png",
  },
  {
    id: 666,
    name: "Johnson Lopez",
    position: "Technical Advisor",
    image: "/assets/images/teamImages/User6.png",
  },
];
