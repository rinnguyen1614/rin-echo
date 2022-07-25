import { EditProps } from "ra-ui-materialui";
import React, { ReactElement } from "react";
import EditBase from "@app/auth/EditBase";
import MenuForm from "./components/MenuForm";
import { Menu } from "../../types/models/Menu";

interface Props extends EditProps<Menu> {}

const MenuEdit = (props: Props): ReactElement => {
  return (
    <EditBase mutationMode="pessimistic" {...props}>
      <MenuForm />
    </EditBase>
  );
};

export default MenuEdit;
