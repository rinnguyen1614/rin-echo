import { ReactElement } from "react";
import {
  List as RaList,
  ListProps as RaListProps,
  RaRecord,
} from "react-admin";
import { Actions } from "./types";
import { WithPermissions } from "./WithPermissions";

const List = <RecordType extends RaRecord = any>(
  props: ListProps<RecordType>
): ReactElement => {
  return (
    <WithPermissions action={Actions.list}>
      <RaList {...props} />
    </WithPermissions>
  );
};

export type ListProps<RecordType extends RaRecord = any> =
  RaListProps<RecordType>;

export default List;
