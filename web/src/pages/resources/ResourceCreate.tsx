import { CreateProps } from "ra-ui-materialui";
import React, { ReactElement } from "react";
import CreateBase from "@app/auth/CreateBase";
import ResourceForm from "./components/ResourceForm";
import { Resource } from "../../types/models/Resource";

interface Props extends CreateProps<Resource> {}

const ResourceCreate = (props: Props): ReactElement => {
  return (
    <CreateBase {...props}>
      <ResourceForm />
    </CreateBase>
  );
};

export default ResourceCreate;
