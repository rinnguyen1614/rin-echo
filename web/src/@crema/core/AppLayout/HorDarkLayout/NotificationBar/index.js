import React from "react";
import IconButton from "@mui/material/IconButton";
import Collapse from "@mui/material/Collapse";
import { alpha, Box } from "@mui/material";
import Alert from "@mui/material/Alert";
import CloseIcon from "@mui/icons-material/Close";
import { useSidebarContext } from "../../../../utility/AppContextProvider/SidebarContextProvider";

const NotificationBar = () => {
  const [open, setOpen] = React.useState(true);
  const { sidebarBgColor, sidebarTextColor } = useSidebarContext();

  return (
    <Collapse in={open}>
      <Box
        sx={{
          borderBottom: (theme) =>
            `solid 1px ${alpha(theme.palette.common.black, 0.15)}`,
          padding: "2px 0",
          backgroundColor: sidebarBgColor,
          color: sidebarTextColor,
        }}
      >
        <Box
          sx={{
            width: "100%",
            maxWidth: { lg: 1140, xl: 1420 },
            mx: "auto",
            px: 5,
          }}
        >
          <Alert
            sx={{
              backgroundColor: "transparent !important",
              padding: 0,
              textAlign: "center",
              color: "inherit",
              "& .MuiAlert-message": {
                flex: 1,
              },
              "& .MuiAlert-action": {
                ml: 2.5,
              },
            }}
            icon={false}
            action={
              <IconButton
                aria-label="close"
                color="inherit"
                size="small"
                onClick={() => {
                  setOpen(false);
                }}
              >
                <CloseIcon fontSize="inherit" />
              </IconButton>
            }
          >
            Get flat 60% off on your first purchase
          </Alert>
        </Box>
      </Box>
    </Collapse>
  );
};

export default NotificationBar;
