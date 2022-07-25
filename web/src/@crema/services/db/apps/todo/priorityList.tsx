import { blue, green, red } from "@mui/material/colors";
import { PriorityObj } from "../../../../types/models/apps/Todo";

const priorityList: PriorityObj[] = [
  { id: 4545554, name: "High", type: 1, color: red[500] },
  { id: 9384234, name: "Medium", type: 2, color: blue[500] },
  { id: 4354454, name: "Low", type: 3, color: green[500] },
];
export default priorityList;
