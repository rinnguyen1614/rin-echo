import { Box, styled } from "@mui/material";
import { useMemo } from "react";
import data from "./data";

export const RequestMethodField = ({ record, small = false }: any) => {
  const request = useMemo(
    () =>
      typeof record === "string" && record !== null
        ? data.find(({ id }) => id === record)
        : record,
    [record]
  );

  if (!request) return null;

  return (
    <Box
      component="span"
      sx={{ p: small ? 0.2 : 0.3, border: "1px solid" }}
      color={request.color}
      fontSize={small ? "8px" : "small"}
      lineHeight={2}
      borderColor={request.color}
    >
      {request.name}
    </Box>
  );
};

const StyledBox = styled(Box, {});
