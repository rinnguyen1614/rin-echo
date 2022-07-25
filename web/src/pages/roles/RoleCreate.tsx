import { CreateProps } from "ra-ui-materialui";
import React, { ReactElement } from "react";
import Create from "@app/auth/Create";
import RoleForm from "./components/RoleForm";

const RoleCreate = (props: CreateProps): ReactElement => {
  return (
    <Create {...props}>
      <RoleForm />
    </Create>
  );
};

export default RoleCreate;
