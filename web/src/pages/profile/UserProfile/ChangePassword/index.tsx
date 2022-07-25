import React, { useCallback, useState } from "react";
import { Box, Typography } from "@mui/material";
import FormattedMessage from "@crema/utility/FormattedMessage";
import { Fonts } from "@crema/shared/constants/AppEnums";
import ChangePasswordForm from "./ChangePasswordForm";
import * as yup from "yup";
import { yupResolver } from "@hookform/resolvers/yup";
import { useAuthProvider, useNotify } from "react-admin";

const validationSchema = yup.object({
  current_password: yup
    .string()
    .required("No password provided.")
    .min(8, "Password is too short - should be 8 chars minimum.")
    .matches(/[a-zA-Z]/, "Password can only contain Latin letters."),
  new_password: yup
    .string()
    .required("No password provided.")
    .min(8, "Password is too short - should be 8 chars minimum.")
    .matches(/[a-zA-Z]/, "Password can only contain Latin letters."),
  retype_new_password: yup
    .string()
    .oneOf([yup.ref("new_password"), null], "Passwords must match"),
});

const ChangePassword = () => {
  const authProvider = useAuthProvider();

  return (
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
        <FormattedMessage id="common.changePassword" />
      </Typography>
      <ChangePasswordForm
        onSubmit={authProvider?.changePassword}
        resolver={yupResolver(validationSchema)}
      />
    </Box>
  );
};

export default ChangePassword;
