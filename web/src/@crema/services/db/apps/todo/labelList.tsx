import { blue, green, red } from "@mui/material/colors";
import { LabelObj } from "../../../../types/models/apps/Todo";

export const labelList: LabelObj[] = [
  {
    id: 211,
    name: "Html",
    alias: "html",
    label: "Html",
    value: "Html",
    color: red[500],
  },
  {
    id: 212,
    name: "CSS",
    alias: "css",
    label: "CSS",
    value: "CSS",
    color: blue[500],
  },
  {
    id: 213,
    name: "JQuery",
    alias: "jquery",
    label: "JQuery",
    value: "JQuery",
    color: green[500],
  },
  {
    id: 214,
    name: "Node.js",
    alias: "node",
    label: "Node.js",
    value: "Node.js",
    color: "grey.500",
  },
];

export const onGetLabel = (labelId: number): LabelObj => {
  return labelList.find((label) => label.id === labelId) as LabelObj;
};
