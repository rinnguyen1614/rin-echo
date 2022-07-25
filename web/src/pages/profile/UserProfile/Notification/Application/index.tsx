import React from "react";
import { Box, Typography } from "@mui/material";
import { Fonts } from "@crema/shared/constants/AppEnums";
import FormattedMessage from "@crema/utility/FormattedMessage";
import FormControlLabel from "@mui/material/FormControlLabel";
import Switch from "@mui/material/Switch";
import { Activity } from "@crema/services/db/profile";

interface ActivityProps {
  application: Activity[];
}

const Application: React.FC<ActivityProps> = ({ application }) => {
  return (
    <Box sx={{ mb: { xs: 5, lg: 6 } }}>
      <Typography
        component="h3"
        sx={{
          fontSize: 16,
          fontWeight: Fonts.BOLD,
          mb: { xs: 3, lg: 4 },
        }}
      >
        <FormattedMessage id="eCommerce.application" />
      </Typography>

      {application.map((data, index) => (
        <Box key={index} sx={{ mb: 1.5 }}>
          <FormControlLabel control={<Switch />} label={data.title} />
        </Box>
      ))}
    </Box>
  );
};

export default Application;
