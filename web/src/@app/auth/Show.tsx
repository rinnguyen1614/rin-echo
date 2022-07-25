import { ReactElement } from "react";
import {
  ShowProps as RaShowProps,
  Show as RaShow,
  RaRecord,
} from "react-admin";
import { Actions } from "./types";
import { WithPermissions } from "./WithPermissions";

const Show = <RecordType extends RaRecord = any>(
  props: ShowProps<RecordType> & { children: ReactElement }
): ReactElement => {
  return (
    <WithPermissions action={Actions.show}>
      <RaShow {...props} />
    </WithPermissions>
  );
};

export type ShowProps<RecordType extends RaRecord = any> =
  RaShowProps<RecordType>;

export default Show;
