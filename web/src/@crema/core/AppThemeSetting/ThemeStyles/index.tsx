import React from "react";
import Box from "@mui/material/Box";
import ToggleButtonGroup from "@mui/material/ToggleButtonGroup";
import clsx from "clsx";
import { CustomizerItemWrapper, StyledToggleButton } from "../index.style";
import FormattedMessage from "../../../utility/FormattedMessage";
import { ThemeStyle } from "../../../shared/constants/AppEnums";
import {
  useThemeActionsContext,
  useThemeContext,
} from "../../../utility/AppContextProvider/ThemeContextProvider";

const ThemeStyles = () => {
  const { themeStyle } = useThemeContext();
  const { updateThemeStyle } = useThemeActionsContext();

  const onStyleChange = (event: any, themeStyle: string) => {
    if (themeStyle) updateThemeStyle(themeStyle);
  };
  return (
    <CustomizerItemWrapper>
      <Box component="h4" sx={{ mb: 2 }}>
        <FormattedMessage id="customizer.themeStyle" />
      </Box>
      <ToggleButtonGroup
        value={themeStyle}
        exclusive
        onChange={onStyleChange}
        aria-label="text alignment"
      >
        <StyledToggleButton
          value={ThemeStyle.MODERN}
          className={clsx({
            active: themeStyle === ThemeStyle.MODERN,
          })}
          aria-label="left aligned"
        >
          <FormattedMessage id="sidebar.pages.userList.modern" />
        </StyledToggleButton>
        <StyledToggleButton
          value={ThemeStyle.STANDARD}
          className={clsx({
            active: themeStyle === ThemeStyle.STANDARD,
          })}
          aria-label="centered"
        >
          <FormattedMessage id="sidebar.pages.userList.standard" />
        </StyledToggleButton>
      </ToggleButtonGroup>
    </CustomizerItemWrapper>
  );
};

export default ThemeStyles;
