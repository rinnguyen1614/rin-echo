import { CreateProps } from "ra-ui-materialui";
import React, { ReactElement } from "react";
import CreateBase from "@app/auth/CreateBase";
import MenuForm from "./components/MenuForm";
import { Menu } from "../../types/models/Menu";

interface Props extends CreateProps<Menu> {}

const MenuCreate = (props: Props): ReactElement => {
  return (
    <CreateBase {...props}>
      <MenuForm />
    </CreateBase>
  );
};

export default MenuCreate;
