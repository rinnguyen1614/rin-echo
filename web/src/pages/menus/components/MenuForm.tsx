import { FormToolbarWithRefresh } from "@app/FormToolbarWithRefresh";
import { parse } from "query-string";
import { minLength, required } from "ra-core";
import { BooleanInput, SimpleForm, TextInput } from "ra-ui-materialui";
import { useMemo } from "react";
import { useRecordContext } from "react-admin";
import { useLocation } from "react-router";

const MenuForm = (props) => {
  const location = useLocation();
  const searchParams = parse(location.search);
  const parent_id = Number(searchParams.parent_id || 0);

  return (
    <SimpleForm
      resource="menus"
      defaultValues={{
        parent_id: parent_id,
      }}
      toolbar={<FormToolbarWithRefresh />}
    >
      <TextInput
        autoFocus
        source="name"
        fullWidth
        validate={(required(), minLength(5))}
      />
      <TextInput
        source="slug"
        fullWidth
        validate={(required(), minLength(5))}
      />
      <TextInput source="path" fullWidth validate={minLength(6)} />
      <TextInput source="type" fullWidth validate={required()} />
      <TextInput source="title" fullWidth />
      <TextInput source="icon" fullWidth />
      <TextInput source="component" fullWidth />
      <TextInput source="description" fullWidth />
      <BooleanInput label="Hidden" source="hidden" />
    </SimpleForm>
  );
};
export default MenuForm;
