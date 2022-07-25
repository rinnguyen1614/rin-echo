import React, { useCallback } from "react";
import Box from "@mui/material/Box";
import { Typography } from "@mui/material";
import * as yup from "yup";
import { Fonts } from "@crema/shared/constants/AppEnums";
import FormattedMessage from "@crema/utility/FormattedMessage";
import InfoForm from "./InfoForm";
import { useGetIdentity } from "react-admin";
import { yupResolver } from "@hookform/resolvers/yup";

const phoneRegExp =
  /^((\\+[1-9]{1,4}[ \\-]*)|(\\([0-9]{2,3}\\)[ \\-]*)|([0-9]{2,4})[ \\-]*)*?[0-9]{3,4}?[ \\-]*[0-9]{3,4}?$/;

const validationSchema = yup.object({
  phone: yup.string().matches(phoneRegExp, "Phone number is not valid"),
});
const Information = () => {
  const { isLoading, identity } = useGetIdentity();

  const handleSubmit = useCallback(() => {}, []);

  return (
    !isLoading && (
      <Box
        sx={{
          position: "relative",
          maxWidth: 550,
        }}
      >
        <Typography
          component="h3"
          sx={{
            fontSize: 16,
            fontWeight: Fonts.BOLD,
            mb: { xs: 3, lg: 5 },
          }}
        >
          <FormattedMessage id="common.information" />
        </Typography>
        <InfoForm
          defaultValues={identity}
          onSubmit={handleSubmit}
          resolver={yupResolver(validationSchema)}
        />
      </Box>
    )
  );
};

export default Information;
