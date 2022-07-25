import React from "react";
import { Box, CircularProgress, Typography } from "@mui/material";

interface ListFooterProps {
  loading?: boolean;
  footerText: string;
}

const ListFooter: React.FC<ListFooterProps> = ({
  loading = false,
  footerText,
}) => {
  return loading ? (
    <Box
      sx={{
        width: "100%",
        display: "flex",
        color: (theme) => theme.palette.text.secondary,
        justifyContent: "center",
        p: 2,
        borderTop: 1,
        borderTopColor: (theme) => theme.palette.divider,
        boxSizing: "border-box",
      }}
    >
      <CircularProgress size={16} />
      <Box component="span" sx={{ ml: 2 }}>
        Loading...
      </Box>
    </Box>
  ) : (
    <Box
      sx={{
        p: 2.5,
        color: (theme) => theme.palette.text.secondary,
        display: "flex",
        justifyContent: "center",
      }}
    >
      <Typography>{footerText}</Typography>
    </Box>
  );
};

export default ListFooter;
