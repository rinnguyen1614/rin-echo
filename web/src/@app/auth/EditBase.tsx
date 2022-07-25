import { ReactNode } from "react";
import {
  EditBase as RaEditBase,
  EditControllerProps,
  RaRecord,
} from "react-admin";
import { Actions } from "./types";
import { WithPermissions } from "./WithPermissions";

const EditBase = <RecordType extends RaRecord = any>({
  children,
  ...props
}: EditControllerProps<RecordType> & { children: ReactNode }) => {
  return (
    <WithPermissions action={Actions.edit}>
      <RaEditBase {...props}>{children}</RaEditBase>
    </WithPermissions>
  );
};

export default EditBase;
