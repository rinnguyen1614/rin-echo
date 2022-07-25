import HorizontalRuleIcon from "@mui/icons-material/HorizontalRule";
import { Box, Stack } from "@mui/material";
import { useGetList } from "ra-core";
import { SelectInput, SelectInputProps } from "ra-ui-materialui";
import React, { FC, ReactElement, useCallback, useMemo } from "react";

interface Props
  extends Omit<
    SelectInputProps,
    "source" | "optionText" | "optionValue" | "choices"
  > {
  source: string;
}

const OptionRenderer = (record: any) => (
  <Box key={record.id} sx={{ paddingLeft: (record.nested - 1) * 2 }}>
    <Stack direction="row" spacing={1}>
      <HorizontalRuleIcon fontSize="small" />
      <Box>{record.name}</Box>
    </Stack>
  </Box>
);

const ResourceSelectInput: FC<Props> = ({ source, ...props }): ReactElement => {
  const { data: resources, isLoading } = useGetList("admin/resources/trees", {
    filter: { select: "id,slug,name, parent_id" },
    pagination: { page: 1, perPage: 1000 },
  });

  const flatten = useCallback((node: any, parent?: any, root: any = []) => {
    root.push(node);
    node["nested"] = (parent && parent["nested"] ? parent["nested"] : 0) + 1;
    if (node.children) {
      for (var child of node.children) {
        flatten(child, node, root);
      }
    }
    return root;
  }, []);

  const resourceFlattens = useMemo(() => {
    let root: any = [];
    resources?.forEach((node: any) => flatten(node, null, root));
    return root;
  }, [resources, flatten]);

  return !isLoading ? (
    <SelectInput
      source={source}
      choices={resourceFlattens}
      isLoading={isLoading}
      optionText={(record) => <OptionRenderer {...record} />}
      {...props}
    />
  ) : (
    <></>
  );
};

export default ResourceSelectInput;
