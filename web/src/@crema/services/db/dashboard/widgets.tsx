import {
  blue,
  green,
  grey,
  orange,
  pink,
  purple,
  red,
} from "@mui/material/colors";
import { Widgets } from "../../../types/models/dashboards/Widgets";

const widgetsData: Widgets = {
  recentActivity: [
    {
      id: 10001,
      image: "/assets/images/avatar/A1.jpg",
      name: "Angelina Joew",
      message: "added courses to the new bucket.",
    },
    {
      id: 10002,
      image: "/assets/images/avatar/A2.jpg",
      name: "John Mathew",
      message: "like company website design.",
    },
    {
      id: 10003,
      image: "/assets/images/avatar/A3.jpg",
      name: "George Bailey",
      message: "followed your works",
    },
    {
      id: 10004,
      image: "/assets/images/avatar/A4.jpg",
      name: "Maria Lee",
      message: "liked origmi-creativity agency.",
    },
    {
      id: 10005,
      image: "/assets/images/avatar/A6.jpg",
      name: "Jacky Brothers",
      message: "invited you to join his page.",
    },
    {
      id: 10006,
      image: "/assets/images/avatar/A3.jpg",
      name: "George Bailey",
      message: "followed your works",
    },
    {
      id: 10007,
      image: "/assets/images/avatar/A1.jpg",
      name: "Angelina Joew",
      message: "added courses to the new bucket.",
    },
  ],
  categories: [
    { id: 100001, name: "Renders", isChecked: false },
    { id: 100002, name: "Graphics", isChecked: false },
    { id: 100003, name: "Buttons", isChecked: true },
    { id: 100004, name: "Patterns", isChecked: false },
    { id: 100005, name: "Icons", isChecked: false },
    { id: 100006, name: "App icons", isChecked: true },
    { id: 100007, name: "List", isChecked: false },
    { id: 100008, name: "Table", isChecked: true },
    { id: 100009, name: "Objects", isChecked: false },
    { id: 100010, name: "Design", isChecked: false },
    { id: 45434, name: "Renders", isChecked: false },
    { id: 3443, name: "Graphics", isChecked: false },
    { id: 10560003, name: "Buttons", isChecked: true },
    { id: 6776, name: "Patterns", isChecked: false },
    { id: 100056605, name: "Icons", isChecked: false },
    { id: 105560006, name: "App icons", isChecked: true },
  ],
  profile: {
    id: 10000001,
    name: "Anton Fristler",
    image: "/assets/images/avatar/A10.jpg",
    photos: 15,
    followers: 124,
    following: 17,
  },
  messages: [
    {
      id: 201,
      image: "/assets/images/avatar/A19.jpg",
      message: "Hey man! Whatsapp?",
      name: "Angelina Joew",
    },
    {
      id: 202,
      image: "/assets/images/avatar/A15.jpg",
      message: "I am fine, what about you?",
      name: "John Matthew",
    },
    {
      id: 203,
      image: "/assets/images/avatar/A21.jpg",
      message: "Call me when you are free!",
      name: "George Bailey",
    },
    {
      id: 204,
      image: "/assets/images/avatar/A25.jpg",
      message: "Send your contact details!",
      name: "Maria Lee",
    },
    {
      id: 205,
      image: "/assets/images/avatar/A19.jpg",
      message: "Hey man! Whatsapp?",
      name: "Angelina Joew",
    },
  ],
  taskList: [
    {
      id: 101,
      title: "Call Adams",
      desc: "It is a long established fact that a reader will be distracted by the readable content of a page.",
    },
    {
      id: 102,
      title: "Meeting with Matthiew",
      desc: "It is a long established fact that a reader will be distracted by the readable content of a page.",
    },
    {
      id: 103,
      title: "Team Meeting",
      desc: "It is a long established fact that a reader will be distracted by the readable content of a page.",
    },
    {
      id: 104,
      title: "Call Adams",
      desc: "It is a long established fact that a reader will be distracted by the readable content of a page.",
    },
  ],
  colorsList: [
    { id: 1, name: "Red", color: red[600], isChecked: false },
    { id: 2, name: "Blue", color: blue[600], isChecked: false },
    { id: 3, name: "Green", color: green[600], isChecked: true },
    { id: 4, name: "Purple", color: purple[800], isChecked: false },
    { id: 5, name: "Orange", color: orange[600], isChecked: true },
    { id: 6, name: "Pink", color: pink[600], isChecked: false },
    { id: 7, name: "Black", color: "black", isChecked: false },
    { id: 8, name: "Light Grey", color: grey[400], isChecked: false },
    { id: 9, name: "Red", color: red[600], isChecked: false },
  ],
  tagsList: [
    { id: 9001, label: "Primary", color: blue[600] },
    { id: 9002, label: "Secondary", color: red[600] },
    { id: 9003, label: "Error", color: green[600] },
    { id: 9004, label: "Ui Kit", color: orange[600] },
  ],
  reviewsList: [
    {
      id: 100001,
      rating: 5,
      by: "M S Brar",
      content:
        "It is a long established fact that a reader will be distracted by the readable content of a page.",
      time: "50 minutes ago",
    },
    {
      id: 100002,
      rating: 5,
      by: "Rocky Johnson",
      content: "It is a long established a reader will a page.",
      time: "3 hours ago",
    },
    {
      id: 100003,
      rating: 4,
      by: "Rahul Bajaj",
      content:
        "It is a long established fact that a reader will be distracted by the readable content of a page.",
      time: "2 hours ago",
    },
    {
      id: 100004,
      rating: 5,
      by: "Rocky Johnson",
      content: "It is a long established fact will a page.",
      time: "3 hours ago",
    },
  ],
  mateInfo: {
    facebookInfo: {
      image: "/assets/images/logo-white.png",
      name: "Crema Admin",
      id: "crema.report@gmail.com",
      desc: "It usually begins with Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. create a natural looking block of text.",
    },
    twitterInfo: {
      image: "/assets/images/logo-white.png",
      name: "Crema Admin",
      id: "crema.report",
      desc: "It usually begins with Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. create a natural looking block of text. ",
    },
  },
  formatList: [
    { id: 1001, name: "psd" },
    { id: 1002, name: "tiff" },
    { id: 1003, name: "jpeg" },
    { id: 1004, name: "gif" },
    { id: 1005, name: "png" },
    { id: 1006, name: "text" },
    { id: 1007, name: "pdf" },
    { id: 1008, name: "doc" },
  ],
  temperatures: [
    {
      id: 332332,
      day: "Mon",
      image: "/assets/images/weather/weather2.png",
    },
    {
      id: 3233232,
      day: "Tues",
      image: "/assets/images/weather/weather3.png",
    },
    {
      id: 4343443,
      day: "Wed",
      image: "/assets/images/weather/weather4.png",
    },
  ],
  cityData: [
    {
      id: 33323,
      name: "New York",
      desc: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
      image: "/assets/images/widgets-companyInfo/building1.png",
    },
    {
      id: 3332,
      name: "Sydney",
      desc: "Lorem Ipsum has been the industry's standard dummy text ever since the 1500s.",
      image: "/assets/images/widgets-companyInfo/building2.png",
    },
    {
      id: 4332,
      name: "New Delhi",
      desc: "It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages.",
      image: "/assets/images/widgets-companyInfo/building3.png",
    },
    {
      id: 5432,
      name: "Singapore",
      desc: "It has survived not only five centuries, but also the leap into remaining essentially unchanged.",
      image: "/assets/images/widgets-companyInfo/building4.png",
    },
  ],
};
export default widgetsData;
