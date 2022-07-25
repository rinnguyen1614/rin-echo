import React from "react";
import Box from "@mui/material/Box";
import FormControl from "@mui/material/FormControl";
import InputLabel from "@mui/material/InputLabel";
import Select from "@mui/material/Select";
import MenuItem from "@mui/material/MenuItem";
import Switch from "@mui/material/Switch";
import { FooterType } from "../../../shared/constants/AppEnums";
import {
  useLayoutActionsContext,
  useLayoutContext,
} from "../../../utility/AppContextProvider/LayoutContextProvider";
import { CustomizerItemWrapper } from "../index.style";

const ThemeFooter = () => {
  const { footer } = useLayoutContext();
  const { setFooter } = useLayoutActionsContext();
  const { footerType } = useLayoutContext();
  const { setFooterType } = useLayoutActionsContext();

  return (
    <CustomizerItemWrapper>
      <Box sx={{ display: "flex", alignItems: "center", mb: 4 }}>
        <Box component="h4">Footer</Box>
        <Box component="span" ml="auto">
          <Switch
            checked={footer}
            onChange={() => setFooter(!footer)}
            value="checkedA"
            inputProps={{ "aria-label": "secondary checkbox" }}
          />
        </Box>
      </Box>
      <FormControl
        variant="outlined"
        sx={{
          width: "100%",
        }}
      >
        <InputLabel id="select-footer">Footer Type</InputLabel>
        <Select
          sx={{
            "& .MuiOutlinedInput-input": {
              padding: "12px 32px 12px 14px",
            },
          }}
          labelId="select-footer"
          label="Footer Type"
          value={footerType}
          // labelWidth={100}
          onChange={(e) => setFooterType(e.target.value)}
        >
          <MenuItem value={FooterType.FIXED}>Fixed</MenuItem>
          <MenuItem value={FooterType.FLUID}>Fluid</MenuItem>
        </Select>
      </FormControl>
    </CustomizerItemWrapper>
  );
};

export default ThemeFooter;
