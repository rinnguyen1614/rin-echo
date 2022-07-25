import { ReactElement } from "react";
import {
  ShowBase as RaShowBase,
  ShowControllerProps,
  RaRecord,
} from "react-admin";
import { Actions } from "./types";
import { WithPermissions } from "./WithPermissions";

const ShowBase = <RecordType extends RaRecord = any>({
  children,
  ...props
}: ShowControllerProps<RecordType> & { children: ReactElement }) => {
  return (
    <WithPermissions action={Actions.show}>
      <RaShowBase {...props}>{children}</RaShowBase>
    </WithPermissions>
  );
};

export default ShowBase;
