import { Typography, SxProps } from "@mui/material";
import { useRecordContext, FieldProps } from "react-admin";
import { User } from "types/models/User";
import AvatarField from "./AvatarField";
import * as React from "react";

interface Props extends FieldProps<User> {
  size?: string;
  sx?: SxProps;
}

const UsernameField = (props: Props) => {
  const { size } = props;
  const record = useRecordContext<User>();
  return record ? (
    <Typography
      variant="body2"
      display="flex"
      flexWrap="nowrap"
      alignItems="center"
      component="div"
      sx={props.sx}
    >
      <AvatarField
        record={record}
        size={size}
        sx={{
          mr: 1,
          mt: -0.5,
          mb: -0.5,
        }}
      />
      {record.username}
    </Typography>
  ) : null;
};

UsernameField.defaultProps = {
  source: "username",
  label: "users.fields.username",
};

export default React.memo<Props>(UsernameField);
