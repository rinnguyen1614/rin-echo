import { ReactElement } from "react";
import {
  ListProps as RaListProps,
  ListBase as RaListBase,
  RaRecord,
} from "react-admin";
import { Actions } from "./types";
import { WithPermissions } from "./WithPermissions";

const ListBase = <RecordType extends RaRecord = any>(
  props: ListProps<RecordType>
): ReactElement => {
  return (
    <WithPermissions action={Actions.list}>
      <RaListBase {...props} />
    </WithPermissions>
  );
};

export type ListProps<RecordType extends RaRecord = any> =
  RaListProps<RecordType>;

export default ListBase;
