import {
  CreateButton,
  ExportButton,
  FilterButton,
  TopToolbar,
} from "react-admin";

const ListActions = (props: any | undefined) => (
  <TopToolbar>
    <FilterButton />
    <CreateButton />
    <ExportButton />
  </TopToolbar>
);

export default ListActions;
