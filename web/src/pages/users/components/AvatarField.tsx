import { Avatar, SxProps } from "@mui/material";
import { FieldProps, useRecordContext } from "react-admin";
import { User } from "types/models/User";
import { FILE_ENDPOINT } from "../../../constants";

interface Props extends FieldProps<User> {
  sx?: SxProps;
  size?: string;
}

const AvatarField = ({ size = "32", sx }: Props) => {
  const record = useRecordContext<User>();
  if (!record) return null;
  return (
    <Avatar
      src={`${FILE_ENDPOINT}${record.avatar_path}`}
      style={{ width: parseInt(size, 10), height: parseInt(size, 10) }}
      sx={sx}
    />
  );
};

export default AvatarField;
