import React, { FC, ReactElement } from "react";
import { useGetList, CheckboxGroupInput } from "react-admin";

interface Props {
  source: string;
}

const RoleInput: FC<Props> = ({ source, ...props }): ReactElement => {
  const { data: roles, isLoading } = useGetList("admin/roles", {
    filter: { select: "id,name" },
    pagination: { perPage: 1000, page: 1 },
  });

  return (
    !isLoading && (
      <CheckboxGroupInput
        label=""
        resource="roles"
        optionText="name"
        optionValue="id"
        source={source}
        choices={roles}
        fullWidth
      />
    )
  );
};

export default RoleInput;
