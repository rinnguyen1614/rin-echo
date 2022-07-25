import React, { useCallback } from "react";
import PersonalInfoForm from "./PersonalInfoForm";
import PropTypes from "prop-types";
import { Box } from "@mui/material";
import * as yup from "yup";
import { useGetIdentity, useUpdate } from "react-admin";
import { yupResolver } from "@hookform/resolvers/yup";
import { AppLoader } from "@crema";
import { Identity } from "@app/types";

const validationSchema = yup.object({
  email: yup.string().email("Invalid email format").required("Required"),
});

const PersonalInfo = () => {
  const { isLoading, identity } = useGetIdentity();
  const handleSubmit = useCallback(() => {}, []);

  if (isLoading) {
    return <AppLoader />;
  }
  return (
    !isLoading && (
      <Box
        sx={{
          position: "relative",
          maxWidth: 550,
        }}
      >
        <PersonalInfoForm
          defaultValues={identity as Identity}
          onSubmit={handleSubmit}
          resolver={yupResolver(validationSchema)}
        />
      </Box>
    )
  );
};

export default PersonalInfo;

PersonalInfo.propTypes = {
  setFieldValue: PropTypes.func,
  values: PropTypes.string,
};
