import { ReactElement } from "react";
import {
  EditProps as RaEditProps,
  Edit as RaEdit,
  RaRecord,
} from "react-admin";
import { Actions } from "./types";
import { WithPermissions } from "./WithPermissions";

const Edit = <RecordType extends RaRecord = any>(
  props: EditProps<RecordType> & { children: ReactElement }
): ReactElement => {
  return (
    <WithPermissions action={Actions.edit}>
      <RaEdit {...props} />
    </WithPermissions>
  );
};

export type EditProps<RecordType extends RaRecord = any> =
  RaEditProps<RecordType>;

export default Edit;
