import {
  TransformData,
  useNotify,
  useRecordContext,
  useRedirect,
  useRefresh,
  useResourceContext,
} from "ra-core";
import {
  DeleteButton,
  SaveButton,
  Toolbar,
  ToolbarProps,
} from "ra-ui-materialui";
import { useCallback } from "react";

interface Props extends ToolbarProps {
  transform?: TransformData;
}

export const FormToolbarWithRefresh = (props: Props) => {
  const { children, className, transform, ...rest } = props;
  const record = useRecordContext(props);
  const resource = useResourceContext(props);
  const notify = useNotify();
  const refresh = useRefresh();
  const redirect = useRedirect();
  const onSuccess = useCallback(
    (type?: string) => {
      switch (type) {
        case "delete":
          notify("ra.notification.deleted", {
            type: "info",
            messageArgs: { smart_count: 1 },
          });
          break;
        case "edit":
          notify("ra.notification.updated", {
            type: "info",
            messageArgs: { smart_count: 1 },
          });
          break;
        default:
          notify("ra.notification.created", {
            type: "info",
            messageArgs: { smart_count: 1 },
          });
      }
      redirect("list", resource);
      refresh();
    },
    [notify, redirect, refresh, resource]
  );

  return (
    <Toolbar sx={{ flex: 1, display: "flex", justifyContent: "space-between" }}>
      <SaveButton
        mutationOptions={{
          onSuccess: () =>
            record && typeof record.id !== "undefined"
              ? onSuccess("edit")
              : onSuccess(),
        }}
        resource={resource}
        transform={transform}
        type="button"
      />
      {record && typeof record.id !== "undefined" && (
        <DeleteButton
          mutationOptions={{ onSuccess: () => onSuccess("delete") }}
          resource={resource}
        />
      )}
    </Toolbar>
  );
};
