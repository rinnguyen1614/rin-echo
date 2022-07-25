import { ReactElement } from "react";
import {
  CreateProps as RaCreateProps,
  Create as RaCreate,
  RaRecord,
} from "react-admin";
import { Actions } from "./types";
import { WithPermissions } from "./WithPermissions";

const Create = <RecordType extends RaRecord = any>(
  props: CreateProps<RecordType> & { children: ReactElement }
): ReactElement => {
  return (
    <WithPermissions action={Actions.create}>
      <RaCreate {...props} />
    </WithPermissions>
  );
};

export type CreateProps<RecordType extends RaRecord = any> =
  RaCreateProps<RecordType>;

export default Create;
