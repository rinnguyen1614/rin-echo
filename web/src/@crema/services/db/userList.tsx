export interface UserList {
  id: number;
  name: string;
  image: string;
  skills: string[];
  information: string;
  email: string;
  phone: string;
  website: string;
  charge: number;
  readTime: string;
  shares: string;
  retweets: string;
  topic: string;
}

const userList: UserList[] = [
  {
    id: 13223,
    name: "Asantha Powel",
    image: "/assets/images/avatar/A5.jpg",
    skills: ["React", "Javascript", "Native", "Drupal"],
    information:
      "There are many variations of passages of Lorem Ipsum available, but the majority have suffered alteration in some form, by injected humour, or randomised words which don't look even slightly believable. If you are going to use a passage of Lorem Ipsum, you need to be sure there isn't anything embarrassing hidden in the middle of text.",
    email: "asantha@example.com",
    phone: "+91324534563",
    website: "www.asantha.com",
    charge: 20,
    readTime: "2 minutes",
    shares: "5k",
    retweets: "25k",
    topic: "Job Interviews",
  },
  {
    id: 32433,
    name: "Alastair Jordan",
    image: "/assets/images/avatar/A1.jpg",
    skills: ["Java", "Javascript", "Flutter", "Drupal"],
    information:
      "There are many variations of passages of Lorem Ipsum available, but the majority have suffered alteration in some form, by injected humour, or randomised words which don't look even slightly believable. If you are going to use a passage of Lorem Ipsum, you need to be sure there isn't anything embarrassing hidden in the middle of text.",
    email: "alastair@example.com",
    phone: "+9145424363",
    website: "www.alastair.com",
    charge: 18,
    readTime: "6 minutes",
    shares: "15k",
    retweets: "235k",
    topic: "Health and Medicine",
  },
  {
    id: 54534,
    name: "Johnson Bravo",
    image: "/assets/images/avatar/A2.jpg",
    skills: ["PHP", "HTML", "Native", "Drupal"],
    information:
      "There are many variations of passages of Lorem Ipsum available, but the majority have suffered alteration in some form, by injected humour, or randomised words which don't look even slightly believable. If you are going to use a passage of Lorem Ipsum, you need to be sure there isn't anything embarrassing hidden in the middle of text.",
    email: "johnson@example.com",
    phone: "+9135432445",
    website: "www.johnson.com",
    charge: 25,
    readTime: "8 minutes",
    shares: "59k",
    retweets: "225k",
    topic: "World Economy",
  },
  {
    id: 43432,
    name: "Johana Peterson",
    image: "/assets/images/avatar/A3.jpg",
    skills: ["Wordpress", "Laravel", "Native", "CMS"],
    information:
      "There are many variations of passages of Lorem Ipsum available, but the majority have suffered alteration in some form, by injected humour, or randomised words which don't look even slightly believable. If you are going to use a passage of Lorem Ipsum, you need to be sure there isn't anything embarrassing hidden in the middle of text.",
    email: "johana@example.com",
    phone: "+9134352343",
    website: "www.johana.com",
    charge: 20,
    readTime: "5 minutes",
    shares: "52k",
    retweets: "125k",
    topic: "Study & Stress",
  },
  {
    id: 35623,
    name: "Heath Streak",
    image: "/assets/images/avatar/A4.jpg",
    skills: ["Angular", "Javascript", "Flutter", "Android"],
    information:
      "There are many variations of passages of Lorem Ipsum available, but the majority have suffered alteration in some form, by injected humour, or randomised words which don't look even slightly believable. If you are going to use a passage of Lorem Ipsum, you need to be sure there isn't anything embarrassing hidden in the middle of text.",
    email: "heath@example.com",
    phone: "+91345542435",
    website: "www.heath.com",
    charge: 22,
    readTime: "4 minutes",
    shares: "59k",
    retweets: "525k",
    topic: "Technology Advancement",
  },
  {
    id: 76312,
    name: "Sunita Gough",
    image: "/assets/images/avatar/A6.jpg",
    skills: ["Vue", "Java", "Native", "Swift"],
    information:
      "There are many variations of passages of Lorem Ipsum available, but the majority have suffered alteration in some form, by injected humour, or randomised words which don't look even slightly believable. If you are going to use a passage of Lorem Ipsum, you need to be sure there isn't anything embarrassing hidden in the middle of text.",
    email: "sunita@example.com",
    phone: "+1334534563",
    website: "www.sunita.com",
    charge: 30,
    readTime: "11 minutes",
    shares: "125k",
    retweets: "255k",
    topic: "Aeronautical Science",
  },
];
export default userList;
