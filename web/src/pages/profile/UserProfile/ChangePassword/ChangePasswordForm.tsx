import React from "react";
import { Box, Button } from "@mui/material";
import AppGridContainer from "@crema/core/AppGridContainer";
import Grid from "@mui/material/Grid";
import FormattedMessage from "@crema/utility/FormattedMessage";
import InputAdornment from "@mui/material/InputAdornment";
import IconButton from "@mui/material/IconButton";
import { Visibility, VisibilityOff } from "@mui/icons-material";
import { Form, TextInput, SaveButton, useNotify } from "react-admin";
import { useForm } from "react-hook-form";

const ChangePasswordForm = (props) => {
  const { onSubmit: onSubmitCustom } = props;
  const { reset } = useForm(props);
  const [saving, setSaving] = React.useState(false);
  const notify = useNotify();
  const [showPassword, setShowPassword] = React.useState(false);

  const [showNewPassword, setShowNewPassword] = React.useState(false);
  const [showRetypeNewPassword, setShowRetypeNewPassword] =
    React.useState(false);

  const onShowOldPassword = () => {
    setShowPassword(!showPassword);
  };

  const onDownOldPassword = (event: any) => {
    event.preventDefault();
  };

  const onShowNewPassword = () => {
    setShowNewPassword(!showNewPassword);
  };

  const onDownNewPassword = (event: any) => {
    event.preventDefault();
  };

  const onShowRetypeNewPassword = () => {
    setShowRetypeNewPassword(!showRetypeNewPassword);
  };

  const onDownRetypeNewPassword = (event: any) => {
    event.preventDefault();
  };

  const onSubmit = React.useCallback(
    async (values: any) => {
      try {
        setSaving(true);
        await onSubmitCustom(values);
        notify("rin.auth.change_password_success");
        reset();
      } catch (error) {
        notify(typeof error === "string" ? error : error.message, {
          type: "warning",
        });
      } finally {
        setSaving(false);
      }
    },
    [onSubmitCustom, notify, reset]
  );

  return (
    <Form {...props} onSubmit={onSubmit}>
      <AppGridContainer spacing={4}>
        <Grid item xs={12} md={6}>
          <TextInput
            type={showPassword ? "text" : "password"}
            variant="outlined"
            helperText=""
            source="current_password"
            InputProps={{
              endAdornment: (
                <InputAdornment position="end">
                  <IconButton
                    aria-label="toggle password visibility"
                    onClick={onShowOldPassword}
                    onMouseDown={onDownOldPassword}
                    edge="end"
                  >
                    {showPassword ? <VisibilityOff /> : <Visibility />}
                  </IconButton>
                </InputAdornment>
              ),
            }}
            margin="none"
            size="medium"
            label={"account.change_password.fields.current_password"}
          />
        </Grid>
        <Grid item xs={12} md={6} sx={{ p: "0 !important" }} />
        <Grid item xs={12} md={6}>
          <TextInput
            type={showNewPassword ? "text" : "password"}
            variant="outlined"
            helperText=""
            source="new_password"
            InputProps={{
              endAdornment: (
                <InputAdornment position="end">
                  <IconButton
                    aria-label="toggle password visibility"
                    onClick={onShowNewPassword}
                    onMouseDown={onDownNewPassword}
                    edge="end"
                  >
                    {showNewPassword ? <VisibilityOff /> : <Visibility />}
                  </IconButton>
                </InputAdornment>
              ),
            }}
            size="medium"
            margin="none"
            label={"account.change_password.fields.new_password"}
          />
        </Grid>
        <Grid item xs={12} md={6}>
          <TextInput
            type={showRetypeNewPassword ? "text" : "password"}
            variant="outlined"
            helperText=""
            source="retype_new_password"
            InputProps={{
              endAdornment: (
                <InputAdornment position="end">
                  <IconButton
                    aria-label="toggle password visibility"
                    onClick={onShowRetypeNewPassword}
                    onMouseDown={onDownRetypeNewPassword}
                    edge="end"
                  >
                    {showRetypeNewPassword ? <VisibilityOff /> : <Visibility />}
                  </IconButton>
                </InputAdornment>
              ),
            }}
            margin="none"
            size="medium"
            label={"account.change_password.fields.retype_new_password"}
          />
        </Grid>
        <Grid item xs={12} md={12}>
          <Box
            sx={{
              display: "flex",
              alignItems: "center",
            }}
          >
            <SaveButton
              sx={{
                position: "relative",
                minWidth: 100,
              }}
              color="primary"
              variant="contained"
              label="common.saveChanges"
              saving={saving}
            />
            <Button
              sx={{
                position: "relative",
                minWidth: 100,
                ml: 2.5,
              }}
              color="primary"
              variant="outlined"
              onClick={() => {
                reset();
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

export default ChangePasswordForm;
