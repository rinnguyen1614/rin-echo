import React from "react";
import Box from "@mui/material/Box";
import { Fonts } from "../../shared/constants/AppEnums";

interface BadgeProps {
  className?: string;
  count: number;
  color?: string;
}

const AppBadge: React.FC<BadgeProps> = ({ count, color = "secondary" }) => {
  if (color === "primary") {
    color = "primary.main";
  } else if (color === "v") {
    color = "secondary.main";
  }
  return (
    <Box
      sx={{
        bgcolor: color,
        padding: "0px 7px",
        fontSize: 11,
        fontWeight: Fonts.SEMI_BOLD,
        height: 20,
        minWidth: 20,
        borderRadius: 20,
        display: "flex",
        alignItems: "center",
        color: (theme) => theme.palette.secondary.contrastText,
      }}
    >
      {count}
    </Box>
  );
};

export default React.memo(AppBadge);
