import { EditProps } from "ra-ui-materialui";
import React, { ReactElement } from "react";
import EditBase from "@app/auth/EditBase";
import ResourceForm from "./components/ResourceForm";
import { Resource } from "../../types/models/Resource";

interface Props extends EditProps<Resource> {}

const ResourceEdit = (props: Props): ReactElement => {
  return (
    <EditBase mutationMode="pessimistic" {...props}>
      <ResourceForm />
    </EditBase>
  );
};

export default ResourceEdit;
