import React from "react";
import CheckIcon from "@mui/icons-material/Check";
import FormattedMessage from "../../utility/FormattedMessage";
import Box from "@mui/material/Box";
import { Theme } from "@mui/material";
import { SxProps } from "@mui/system";

interface AppSelectedIconProps {
  backgroundColor?: string;
  color?: string;
  isCenter?: boolean;
}

const AppSelectedIcon: React.FC<AppSelectedIconProps> = ({
  backgroundColor,
  isCenter = true,
  color,
}) => {
  let centerStyle: SxProps<Theme> = isCenter
    ? {
        position: "absolute",
        left: "50%",
        top: "50%",
        zIndex: 1,
        transform: "translate(-50%, -50%)",
      }
    : {
        position: "absolute",
        right: 10,
        top: 10,
        zIndex: 1,
      };
  return (
    <Box
      sx={{
        width: 20,
        height: 20,
        borderRadius: "50%",
        display: "flex",
        alignItems: "center",
        justifyContent: "center",
        overflow: "hidden",
        backgroundColor: backgroundColor ? backgroundColor : "primary.main",
        color: color ? color : "primary.contrastText",
        ...centerStyle,
        "& svg": {
          fontSize: 16,
        },
      }}
    >
      <CheckIcon>
        <FormattedMessage id="customizer.checked" />
      </CheckIcon>
    </Box>
  );
};

export default AppSelectedIcon;
