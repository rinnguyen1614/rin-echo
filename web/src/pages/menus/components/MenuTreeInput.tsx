import React, { FC, ReactElement } from "react";
import { useGetList } from "react-admin";
import TreeInput from "@app/tree/TreeInput";

interface Props {
  source: string;
}

const MenuTreeInput: FC<Props> = ({ source, ...props }): ReactElement => {
  const { data: resources, isLoading } = useGetList("admin/menus/trees", {
    filter: { select: "id,slug,name,path, parent_id" },
    pagination: { perPage: 1000, page: 1 },
  });

  return (
    !isLoading && (
      <TreeInput
        label=""
        resource="menus"
        optionText="name"
        optionValue="id"
        source={source}
        choices={resources}
        fullWidth
      />
    )
  );
};

export default MenuTreeInput;
