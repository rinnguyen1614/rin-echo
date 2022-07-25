import React from "react";
import { Box, CircularProgress, Typography } from "@mui/material";

interface GridFooterProps {
  loading: boolean;
  footerText: string;
}

const GridFooter: React.FC<GridFooterProps> = ({ loading, footerText }) => {
  if (loading) {
    return (
      <Box
        sx={{
          width: "100%",
          display: "flex",
          color: (theme) => theme.palette.text.secondary,
          justifyContent: "center",
          padding: 8,
          "& .loading": {
            marginLeft: 8,
          },
        }}
      >
        <CircularProgress size={16} />
        <span className="loading">Loading...</span>
      </Box>
    );
  } else {
    return (
      <Box
        sx={{
          padding: 10,
          color: (theme) => theme.palette.text.secondary,
          display: "flex",
          justifyContent: "center",
        }}
      >
        <Typography>{footerText}</Typography>
      </Box>
    );
  }
};

export default GridFooter;
