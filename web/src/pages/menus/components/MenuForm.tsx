import { parse } from "query-string";
import { minLength, required, useSaveContext } from "ra-core";
import { BooleanInput, FormTab, TabbedForm, TextInput } from "ra-ui-materialui";
import { useMemo } from "react";
import { useRecordContext } from "react-admin";
import { useLocation } from "react-router";
import { FormToolbarWithRefresh } from "@app/FormToolbarWithRefresh";
import ResourceTreeInput from "../../resources/components/ResourceTreeInput";

const MenuForm = (props) => {
  const location = useLocation();
  const searchParams = parse(location.search);
  const parent_id = Number(searchParams.parent_id || 0);
  const record = useRecordContext(props);
  const value = useMemo(() => {
    let v = { ...record };
    if (record && record.resources?.length) {
      v.resource_ids = record.resources.map(({ id }) => id);
    }
    return v;
  }, [record]);
  const { saving } = useSaveContext();

  return (
    <TabbedForm
      resource="menus"
      defaultValues={{
        parent_id: parent_id,
      }}
      toolbar={<FormToolbarWithRefresh />}
      syncWithLocation={false}
      record={value}
      saving={saving}
    >
      <FormTab label={"resources.menus.tabs.general"}>
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
      </FormTab>
      <FormTab label={"resources.menus.tabs.resources"}>
        <ResourceTreeInput source="resource_ids" />
      </FormTab>
    </TabbedForm>
  );
};
export default MenuForm;
