import { EditProps } from "ra-ui-materialui";
import React, { ReactElement } from "react";
import Edit from "@app/auth/Edit";
import UserForm from "./components/UserForm";

const UserEdit = (props: EditProps): ReactElement => {
  return (
    <Edit mutationMode="pessimistic" {...props}>
      <UserForm />
    </Edit>
  );
};

export default UserEdit;
