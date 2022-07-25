import React from "react";
import LinearProgress, {
  linearProgressClasses,
} from "@mui/material/LinearProgress";

interface CustomLinearProgressProps {
  pathColor: string;
  activeColor: string;
  thickness: number;
  borderRadius: number;

  [x: string]: any;
}

const CustomLinearProgress: React.FC<CustomLinearProgressProps> = ({
  pathColor = "#d6d6d6",
  activeColor = "#1a90ff",
  thickness = 5,
  borderRadius = 10,
  ...rest
}) => {
  return (
    <LinearProgress
      sx={{
        height: thickness,
        borderRadius: borderRadius,

        [`&.${linearProgressClasses.colorPrimary}`]: {
          backgroundColor: pathColor,
        },
        [`& .${linearProgressClasses.bar}`]: {
          borderRadius: pathColor,
          backgroundColor: activeColor,
        },
      }}
      {...rest}
    />
  );
};

interface AppLinearProgressProps {
  thickness?: number;
  borderRadius?: number;
  pathColor?: string;
  activeColor?: string;
  variant?: "buffer" | "determinate" | "indeterminate" | "query";

  [x: string]: any;
}

const AppLinearProgress: React.FC<AppLinearProgressProps> = ({
  thickness = 3,
  borderRadius = 10,
  pathColor = "#F5F6FA",
  activeColor = "#00905F",
  variant = "determinate",
  ...rest
}) => {
  return (
    <CustomLinearProgress
      pathColor={pathColor}
      variant={variant}
      thickness={thickness}
      borderRadius={borderRadius}
      activeColor={activeColor}
      {...rest}
    />
  );
};

export default AppLinearProgress;
