import { Box } from "@mui/material";
import { get } from "lodash";
import { useMemo } from "react";
import { useRecordContext } from "react-admin";
import data from "./data";

export const GenderField = (props) => {
  const { source } = props;
  const record = useRecordContext(props);
  const item = useMemo(
    () => !record.length && data.find(({ id }) => id === get(record, source)),
    [record, source]
  );

  if (!item) return null;

  return <Box component="span">{item.name}</Box>;
};

GenderField.defaultProps = {
  source: "gender",
};
