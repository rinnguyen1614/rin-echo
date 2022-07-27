import { useMemo } from "react";
import {
  BooleanInput,
  FormTab,
  minLength,
  required,
  TabbedForm,
  TextInput,
  useRecordContext,
} from "react-admin";
import MenuTreeInput from "../../menus/components/MenuTreeInput";
import ResourceTreeInput from "../../resources/components/ResourceTreeInput";

const RoleForm = (props) => {
  const record = useRecordContext(props);
  const value = useMemo(() => {
    let v = { ...record };
    if (record && record.permissions?.length) {
      v.resource_ids = record.permissions.map((per) => per.resource.id);
    }

    if (record && record.menus?.length) {
      v.menu_ids = record.menus.map(({ id }) => id);
    }
    return v;
  }, [record]);

  return (
    <TabbedForm resource="roles" syncWithLocation={false} record={value}>
      <FormTab label={"resources.roles.tabs.general"}>
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
        <BooleanInput
          label="Default"
          source="is_default"
          helperText="Assign to new users by default."
        />
      </FormTab>
      <FormTab label={"resources.roles.tabs.permissions"}>
        <ResourceTreeInput source="resource_ids" />
      </FormTab>
      <FormTab label={"resources.roles.tabs.menus"}>
        <MenuTreeInput source="menu_ids" />
      </FormTab>
    </TabbedForm>
  );
};
export default RoleForm;
