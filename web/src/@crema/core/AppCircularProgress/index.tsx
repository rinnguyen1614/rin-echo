import React, { ReactNode } from "react";
import CircularProgress, {
  circularProgressClasses,
} from "@mui/material/CircularProgress";
import Box from "@mui/material/Box";
import { Fonts } from "../../shared/constants/AppEnums";

interface AppCircularProgressProps {
  hidePercentage?: boolean;
  centerNode?: ReactNode;
  minWidth?: number;
  maxWidth?: number;
  pathColor?: string;
  activeColor?: string;
  value: number;
  thickness: number;
  valueStyle?: React.CSSProperties;

  [x: string]: any;
}

const AppCircularProgress: React.FC<AppCircularProgressProps> = ({
  value,
  centerNode,
  hidePercentage = false,
  minWidth = 160,
  maxWidth = 200,
  pathColor = "#d6d6d6",
  activeColor = "#23fa23",
  valueStyle,
  thickness,
  ...props
}) => {
  return (
    <Box
      sx={{
        position: "relative",
        minWidth: minWidth,
        maxWidth: maxWidth,
        margin: "0 auto",
      }}
    >
      <Box sx={{ position: "relative" }}>
        <CircularProgress
          variant="determinate"
          value={100}
          sx={{
            color: pathColor,
            width: "100% !important",
            height: "100% !important",
          }}
          thickness={thickness}
          {...props}
        />
        <CircularProgress
          sx={{
            color: activeColor,
            animationDuration: "550ms",
            position: "absolute",
            left: (theme) => (theme.direction === "rtl" ? -2 : 2),
            top: -2,
            width: "100% !important",
            height: "100% !important",
            [`& .${circularProgressClasses.circle}`]: {
              strokeLinecap: "round",
            },
          }}
          variant="determinate"
          value={value}
          thickness={thickness}
          {...props}
        />
      </Box>
      <Box
        sx={{
          position: "absolute",
          left: 0,
          top: 0,
          width: "100%",
          height: "100%",
          zIndex: 3,
          display: "flex",
          alignItems: "center",
          justifyContent: "center",
        }}
      >
        {centerNode}
        {hidePercentage ? null : (
          <Box
            component="p"
            sx={{
              fontSize: 30,
              fontWeight: Fonts.MEDIUM,
              color: (theme) => theme.palette.secondary.main,
              ...valueStyle,
            }}
          >
            {value}%
          </Box>
        )}
      </Box>
    </Box>
  );
};
export default AppCircularProgress;
