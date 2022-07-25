import React, { ReactNode, useState } from "react";
import DeleteOutlinedIcon from "@mui/icons-material/DeleteOutlined";
import FormattedMessage from "../../utility/FormattedMessage";
import AppConfirmDialog from "../AppConfirmDialog";
import IconButton from "@mui/material/IconButton";
import AppTooltip from "../AppTooltip";
import { SxProps } from "@mui/system";
import { Theme } from "@mui/material";

interface AppsDeleteIconProps {
  deleteAction: () => void;
  deleteTitle: string | ReactNode;
  sx: SxProps<Theme>;
}

const AppsDeleteIcon: React.FC<AppsDeleteIconProps> = ({
  deleteAction,
  deleteTitle,
  sx,
}) => {
  const [isDeleteDialogOpen, setDeleteDialogOpen] = useState<boolean>(false);

  return (
    <>
      <AppTooltip title={<FormattedMessage id="common.trash" />}>
        <IconButton sx={sx} size="large">
          <DeleteOutlinedIcon onClick={() => setDeleteDialogOpen(true)} />
        </IconButton>
      </AppTooltip>
      <AppConfirmDialog
        open={isDeleteDialogOpen}
        onDeny={setDeleteDialogOpen}
        onConfirm={deleteAction}
        title={deleteTitle}
        dialogTitle={<FormattedMessage id="common.deleteItem" />}
      />
    </>
  );
};

export default AppsDeleteIcon;
