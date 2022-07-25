import { CreateProps } from "ra-ui-materialui";
import React, { ReactElement } from "react";
import Create from "@app/auth/Create";
import UserForm from "./components/UserForm";

const UserCreate = (props: CreateProps): ReactElement => {
  return (
    <Create {...props}>
      <UserForm />
    </Create>
  );
};

export default UserCreate;
