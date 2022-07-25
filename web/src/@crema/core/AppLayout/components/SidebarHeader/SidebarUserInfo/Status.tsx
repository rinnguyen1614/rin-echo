import React from "react";
import { Box } from "@mui/material";

const Status = () => {
  return (
    <Box
      sx={{
        position: "absolute",
        right: 4,
        bottom: 4,
        zIndex: 1,
        width: 15,
        height: 15,
        cursor: "pointer",
        borderRadius: "50%",
        backgroundColor: (theme) => theme.palette.success.main,
        border: (theme) => `solid 2px ${theme.palette.text.primary}`,
      }}
    />
  );
};

export default Status;
