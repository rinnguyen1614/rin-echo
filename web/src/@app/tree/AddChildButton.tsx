import AddIcon from "@mui/icons-material/Add";
import { Button } from "@mui/material";
import { useMemo } from "react";
import { useCreatePath, useTranslate } from "react-admin";
import { createPath as routerCreatePath } from "react-router";

export const AddChildButton = ({ id, resource }: any) => {
  const translate = useTranslate();
  const createPath = useCreatePath();
  const path = useMemo(
    () =>
      routerCreatePath({
        pathname: createPath({ resource, type: "create" }),
        search: "parent_id=" + id,
      }),
    [createPath, id, resource]
  );

  return (
    <Button startIcon={<AddIcon />} href={path} size="small">
      {translate("rin.tree.actions.add_child")}
    </Button>
  );
};
