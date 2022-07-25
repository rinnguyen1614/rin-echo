import { Box } from "@mui/material";
import { get } from "lodash";
import { useRecordContext } from "react-admin";

export const StatusField = (props) => {
  const { source } = props;
  const record = useRecordContext(props);
  const value = get(record, source);
  const getStatusColor = () =>
    value >= 200 && value < 300 ? "#43C888" : "#F84E4E";

  return (
    <Box
      sx={{
        color: getStatusColor(),
        backgroundColor: getStatusColor() + "44",
        padding: "3px 5px",
        borderRadius: 1,
        fontSize: "small",
        display: "inline-block",
      }}
    >
      {value}
    </Box>
  );
};

StatusField.defaultProps = {
  source: "status_code",
};
