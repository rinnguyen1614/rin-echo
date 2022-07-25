import React, { useState } from "react";
import { IconButton, Theme } from "@mui/material";
import AppTooltip from "../AppTooltip";
import EmailOutlinedIcon from "@mui/icons-material/EmailOutlined";
import Drawer from "@mui/material/Drawer";
import AppMessageContent from "./AppMessageContent";
import { alpha } from "@mui/material/styles";
import Box from "@mui/material/Box";
import { SxProps } from "@mui/system";

interface AppMessagesProps {
  sxMessageContentStyle?: SxProps<Theme>;
  drawerPosition?: "left" | "top" | "right" | "bottom";
  tooltipPosition?:
    | "bottom-end"
    | "bottom-start"
    | "bottom"
    | "left-end"
    | "left-start"
    | "left"
    | "right-end"
    | "right-start"
    | "right"
    | "top-end"
    | "top-start"
    | "top";
  isMenu?: boolean;
}

const AppMessages: React.FC<AppMessagesProps> = ({
  drawerPosition = "right",
  tooltipPosition = "bottom",
  isMenu = false,
  sxMessageContentStyle = {},
}) => {
  const [showMessage, setShowMessage] = useState(false);
  return (
    <>
      {isMenu ? (
        <Box component="span" onClick={() => setShowMessage(true)}>
          Message
        </Box>
      ) : (
        <AppTooltip title="Message" placement={tooltipPosition}>
          <IconButton
            className="icon-btn"
            sx={{
              borderRadius: "50%",
              width: 40,
              height: 40,
              color: (theme) => theme.palette.text.secondary,
              backgroundColor: (theme) => theme.palette.background.default,
              border: 1,
              borderColor: "transparent",
              "&:hover, &:focus": {
                color: (theme) => theme.palette.text.primary,
                backgroundColor: (theme) =>
                  alpha(theme.palette.background.default, 0.9),
                borderColor: (theme) =>
                  alpha(theme.palette.text.secondary, 0.25),
              },
            }}
            onClick={() => setShowMessage(true)}
            size="large"
          >
            <EmailOutlinedIcon className="icon" />
          </IconButton>
        </AppTooltip>
      )}

      <Drawer
        anchor={drawerPosition}
        open={showMessage}
        onClose={() => setShowMessage(false)}
      >
        <AppMessageContent
          sxStyle={sxMessageContentStyle}
          onClose={() => setShowMessage(false)}
        />
      </Drawer>
    </>
  );
};

export default AppMessages;
