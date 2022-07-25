import React from "react";
import { Tooltip } from "recharts";

const ChartTooltip = () => (
  <Tooltip
    labelStyle={{ color: "black" }}
    contentStyle={{
      borderRadius: 3,
      borderColor: "#3AE28B",
      background: "#FFFFFF42",
    }}
    itemStyle={{ color: "#3AE28B" }}
  />
);

export default ChartTooltip;
