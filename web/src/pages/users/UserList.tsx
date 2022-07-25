import { Box, Chip } from "@mui/material";
import { makeStyles } from "@mui/styles";
import classnames from "classnames";
import { FC, memo, ReactElement } from "react";
import {
  BooleanField,
  Datagrid,
  DateField,
  ListProps,
  NullableBooleanInput,
  TextField,
  TextInput,
  useTranslate,
} from "react-admin";
import ListBase from "@app/auth/ListBase";
import ListActions from "@app/ListActions";
import UsernameField from "./components/UsernameField";
import { GenderField } from "components/genders/GenderField";
import { AppAnimate } from "@crema";

const userFilters = [
  <TextInput
    label="Search"
    source="q=username:like,full_name:like"
    alwaysOn
    variant="outlined"
  />,
  <NullableBooleanInput
    label="Default"
    source="is_default"
    variant="outlined"
  />,
];

const UserList = (props: ListProps): ReactElement => {
  const translate = useTranslate();

  return (
    <AppAnimate animation="transition.slideUpIn" delay={200}>
      <ListBase
        {...props}
        actions={<ListActions />}
        filters={userFilters}
        filter={{
          select:
            "id,avatar_path, username,full_name, email, phone, email_verified, phone_verified, gender, created_at",
        }}
      >
        <Datagrid optimized rowClick="edit" size="medium">
          <UsernameField />
          <TextField source="email" />
          <BooleanField source="email_verified" />
          <TextField source="phone" />
          <BooleanField source="phone_verified" />
          <GenderField source="gender" />
          <DateField
            source="created_at"
            locales="us-US"
            showTime={true}
            label={translate("model.created_at")}
          />
        </Datagrid>
      </ListBase>
    </AppAnimate>
  );
};

export default UserList;
