import { ReactNode } from "react";
import {
  CreateBase as RaCreateBase,
  CreateControllerProps,
  RaRecord,
} from "react-admin";
import { Actions } from "./types";
import { WithPermissions } from "./WithPermissions";

const CreateBase = <RecordType extends RaRecord = any>({
  children,
  ...props
}: CreateControllerProps<RecordType> & { children: ReactNode }) => {
  return (
    <WithPermissions action={Actions.create}>
      <RaCreateBase {...props}>{children}</RaCreateBase>
    </WithPermissions>
  );
};

export default CreateBase;
