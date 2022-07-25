import React, { useEffect } from "react";
import {
  alpha,
  Box,
  Button,
  InputAdornment,
  TextField,
  Typography,
} from "@mui/material";
import Avatar from "@mui/material/Avatar";
import AppGridContainer from "@crema/core/AppGridContainer";
import Grid from "@mui/material/Grid";
import FormattedMessage from "@crema/utility/FormattedMessage";
import { useDropzone } from "react-dropzone";
import EditIcon from "@mui/icons-material/Edit";
import { styled } from "@mui/material/styles";
import { Fonts } from "@crema/shared/constants/AppEnums";
import { Form, FormProps, RadioButtonGroupInput, TextInput } from "react-admin";
import { useForm } from "react-hook-form";
import { Identity } from "@app/types";
import { CheckCircle, HighlightOff } from "@mui/icons-material";
import AppTooltip from "@crema/core/AppTooltip";
import { DatePicker } from "@mui/lab";

export const AvatarViewWrapper = styled("div")(({ theme }) => {
  return {
    position: "relative",
    cursor: "pointer",
    "& .edit-icon": {
      position: "absolute",
      bottom: 0,
      right: 0,
      zIndex: 1,
      border: `solid 2px ${theme.palette.background.paper}`,
      backgroundColor: alpha(theme.palette.primary.main, 0.7),
      color: theme.palette.primary.contrastText,
      borderRadius: "50%",
      width: 26,
      height: 26,
      display: "none",
      alignItems: "center",
      justifyContent: "center",
      transition: "all 0.4s ease",
      cursor: "pointer",
      "& .MuiSvgIcon-root": {
        fontSize: 16,
      },
    },
    "&.dropzone": {
      outline: 0,
      "&:hover .edit-icon, &:focus .edit-icon": {
        display: "flex",
      },
    },
  };
});

interface PersonalInfoFormProps extends Omit<FormProps, "children"> {
  defaultValues: Identity;
}

const PersonalInfoForm: React.FC<PersonalInfoFormProps> = (props) => {
  const { defaultValues } = props;
  const { setValue, watch, reset } = useForm(props);
  const watchAvatarPath = watch("avatar_path", defaultValues.avatar_path);
  const watchDateOfBirth = watch("date_of_birth", defaultValues.date_of_birth);
  const { getRootProps, getInputProps } = useDropzone({
    accept: "image/*",
    onDrop: (acceptedFiles) => {
      setValue("avatar_path", URL.createObjectURL(acceptedFiles[0]));
    },
  });

  return (
    <Form {...props}>
      <Typography
        component="h3"
        sx={{
          fontSize: 16,
          fontWeight: Fonts.BOLD,
          mb: { xs: 3, lg: 4 },
        }}
      >
        <FormattedMessage id="common.personalInfo" />
      </Typography>
      <Box
        sx={{
          display: "flex",
          alignItems: "center",
          mb: { xs: 5, lg: 6 },
        }}
      >
        <AvatarViewWrapper {...getRootProps({ className: "dropzone" })}>
          <input {...getInputProps()} />
          <label htmlFor="icon-button-file">
            <Avatar
              sx={{
                width: { xs: 50, lg: 64 },
                height: { xs: 50, lg: 64 },
                cursor: "pointer",
              }}
              src={watchAvatarPath}
            />
            <Box className="edit-icon">
              <EditIcon />
            </Box>
          </label>
        </AvatarViewWrapper>
        <Box
          sx={{
            ml: 4,
          }}
        >
          <Typography
            sx={{
              fontWeight: Fonts.MEDIUM,
            }}
          >
            {defaultValues.full_name}
          </Typography>
          <Typography
            sx={{
              color: (theme) => theme.palette.text.secondary,
            }}
          >
            {defaultValues.username}
          </Typography>
        </Box>
      </Box>
      <AppGridContainer spacing={4}>
        <Grid item xs={12} md={6}>
          <TextInput
            source="full_name"
            fullWidth
            label="common.fullName"
            variant="outlined"
            helperText=""
            size="medium"
            margin="none"
          />
        </Grid>
        <Grid item xs={12} md={6}>
          <TextInput
            source="email"
            label="common.email"
            InputProps={{
              readOnly: true,
              endAdornment: (
                <InputAdornment position="end">
                  {defaultValues.email_verified ? (
                    <AppTooltip
                      title={<FormattedMessage id="tooltip.verified_email" />}
                    >
                      <CheckCircle color="success" fontSize="small" />
                    </AppTooltip>
                  ) : (
                    <AppTooltip
                      title={<FormattedMessage id="tooltip.unverified_email" />}
                    >
                      <HighlightOff color="error" fontSize="small" />
                    </AppTooltip>
                  )}
                </InputAdornment>
              ),
            }}
            fullWidth
            variant="outlined"
            helperText=""
            size="medium"
            margin="none"
          />
        </Grid>
        <Grid item xs={12} md={6}>
          <TextInput
            source="phone"
            label="common.phoneNumber"
            InputProps={{
              readOnly: true,
              endAdornment: (
                <InputAdornment position="end">
                  {defaultValues.phone_verified ? (
                    <AppTooltip
                      title={<FormattedMessage id="tooltip.verified_phone" />}
                    >
                      <CheckCircle color="success" fontSize="small" />
                    </AppTooltip>
                  ) : (
                    <AppTooltip
                      title={<FormattedMessage id="tooltip.unverified_phone" />}
                    >
                      <HighlightOff color="error" fontSize="small" />
                    </AppTooltip>
                  )}
                </InputAdornment>
              ),
            }}
            variant="outlined"
            helperText=""
            size="medium"
            fullWidth
            margin="none"
          />
        </Grid>
        <Grid item xs={12} md={6}>
          <Box
            sx={{
              position: "relative",
              "& .MuiTextField-root": {
                width: "100%",
              },
            }}
          >
            <DatePicker
              label={<FormattedMessage id="common.birthDate" />}
              value={watchDateOfBirth}
              onChange={(newValue) => {
                setValue("date_of_birth", newValue);
              }}
              renderInput={(params) => <TextField {...params} />}
            />
          </Box>
        </Grid>
        <Grid item xs={12} md={6}>
          <RadioButtonGroupInput
            label="common.gender"
            source="gender"
            variant="outlined"
            choices={[
              { id: 3, name: "Female" },
              { id: 2, name: "Male" },
              { id: 1, name: "Other" },
            ]}
            margin="none"
          />
        </Grid>
        <Grid item xs={12} md={12}>
          <Box
            sx={{
              display: "flex",
              alignItems: "center",
            }}
          >
            <Button
              sx={{
                position: "relative",
                minWidth: 100,
              }}
              color="primary"
              variant="contained"
              type="submit"
            >
              <FormattedMessage id="common.saveChanges" />
            </Button>
            <Button
              sx={{
                position: "relative",
                minWidth: 100,
                ml: 2.5,
              }}
              color="primary"
              variant="outlined"
              onClick={() => {
                reset(defaultValues);
                console.log("reset", defaultValues);
              }}
            >
              <FormattedMessage id="common.cancel" />
            </Button>
          </Box>
        </Grid>
      </AppGridContainer>
    </Form>
  );
};

export default PersonalInfoForm;
