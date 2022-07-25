import React from "react";
import Box from "@mui/material/Box";
import CheckIcon from "@mui/icons-material/Check";
import AddOutlinedIcon from "@mui/icons-material/AddOutlined";
import FormattedMessage from "../../utility/FormattedMessage";
import { useThemeContext } from "../../utility/AppContextProvider/ThemeContextProvider";
import { ThemeColorsProps } from "./ThemeColors";

interface CustomColorCellTypes {
  themeColorSet: ThemeColorsProps;
  updateThemeColors: (colorSet: ThemeColorsProps) => void;
}

const CustomColorCell: React.FC<CustomColorCellTypes> = ({
  themeColorSet,
  updateThemeColors,
}) => {
  const { theme } = useThemeContext();
  return (
    <Box
      onClick={() => {
        updateThemeColors(themeColorSet);
      }}
    >
      <Box
        sx={{
          width: "100%",
          borderRadius: 1,
          position: "relative",
          overflow: "hidden",
          cursor: "pointer",
        }}
      >
        <Box
          sx={{
            px: 2.5,
            py: 2,
            position: "relative",
            display: "flex",
            alignItems: "center",
            backgroundColor: themeColorSet.primary.main,
            color: (theme) => theme.palette.common.white,
          }}
        >
          Primary
          {theme.palette.primary.main === themeColorSet.primary.main &&
          theme.palette.secondary.main === themeColorSet.secondary.main &&
          theme.palette.mode === themeColorSet.mode ? (
            <Box
              sx={{
                ml: "auto",
                width: 20,
                height: 20,
                borderRadius: "50%",
                display: "flex",
                alignItems: "center",
                justifyContent: "center",
                overflow: "hidden",
                backgroundColor: "#fff",
                color: theme.palette.primary.main,
              }}
            >
              <CheckIcon
                sx={{
                  fontSize: 16,
                }}
              >
                <FormattedMessage id="customizer.checked" />
              </CheckIcon>
            </Box>
          ) : null}
        </Box>
        <Box
          sx={{
            p: 2.5,
            backgroundColor: themeColorSet.background.default,
            color: themeColorSet.text.primary,
          }}
        >
          <Box
            sx={{
              height: 80,
              py: 1.5,
              px: 2.5,
              mb: 2.5,
              boxShadow:
                "rgba(0, 0, 0, 0.2) 0px 2px 1px -1px, rgba(0, 0, 0, 0.14) 0px 1px 1px 0px",
              borderRadius: 1,
              backgroundColor: themeColorSet.background.paper,
              color: themeColorSet.text.primary,
              display: "flex",
              flexDirection: "column",
            }}
          >
            Paper
            <Box
              sx={{
                width: 30,
                height: 30,
                borderRadius: "50%",
                p: 1,
                backgroundColor: themeColorSet.secondary.main,
                color: (theme) => theme.palette.common.white,
                display: "flex",
                alignItems: "center",
                justifyContent: "center",
                mt: "auto",
                ml: "auto",
                "& svg": {
                  fontSize: 20,
                },
              }}
            >
              <AddOutlinedIcon />
            </Box>
          </Box>
          Background
        </Box>
      </Box>
      <Box sx={{ pt: 2, px: 3, pb: 3 }}>{themeColorSet.title}</Box>
    </Box>
  );
};

export default CustomColorCell;
