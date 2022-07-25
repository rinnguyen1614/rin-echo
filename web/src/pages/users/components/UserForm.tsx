import { Avatar, Box, Grid, Typography } from "@mui/material";
import { useEffect, useMemo, useState } from "react";
import {
  BooleanInput,
  FormTab,
  minLength,
  PasswordInput,
  required,
  TabbedForm,
  TextInput,
  useRecordContext,
} from "react-admin";
import { useDropzone } from "react-dropzone";
import { useForm } from "react-hook-form";
import EditIcon from "@mui/icons-material/Edit";
import { AvatarViewWrapper } from "../../profile/UserProfile/PersonalInfo/PersonalInfoForm";
import RoleInput from "pages/roles/components/RoleInput";
import { AppAnimate, AppGridContainer } from "@crema";

const UserForm = (props) => {
  const record = useRecordContext(props);
  const value = useMemo(() => {
    let v = { ...record };
    if (record && record.permissions?.length) {
      v.resource_ids = record.permissions.map((per) => per.resource.id);
    }

    if (record && record.user_roles?.length) {
      v.role_ids = record.user_roles.map(({ role }) => role?.id);
    }
    return v;
  }, [record]);
  const { setValue, watch } = useForm(props);
  const watchAvatarPath = watch("avatar_path", record.avatar_path);
  const [watchRandomPassword, setWatchRandomPassword] = useState(true);
  const { getRootProps, getInputProps } = useDropzone({
    accept: "image/*",
    onDrop: (acceptedFiles) => {
      setValue("avatar_path", URL.createObjectURL(acceptedFiles[0]));
    },
  });

  return (
    <AppAnimate animation="transition.slideUpIn" delay={200}>
      <AppGridContainer spacing={4}>
        <Grid item xs={12} md={12}>
          <TabbedForm resource="users" syncWithLocation={false} record={value}>
            <FormTab label={"resources.users.tabs.general"}>
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
                  <TextInput
                    source="full_name"
                    fullWidth
                    validate={(required(), minLength(5))}
                  />
                </Box>
              </Box>
              <TextInput
                autoFocus
                source="username"
                fullWidth
                validate={(required(), minLength(5))}
              />
              <TextInput
                source="email"
                fullWidth
                validate={(required(), minLength(5))}
              />
              <BooleanInput
                label="resources.users.fields.random_password"
                source="random_password"
                onChange={(event) => {
                  setWatchRandomPassword(event.target.checked);
                }}
              />
              {!watchRandomPassword && (
                <Box>
                  <PasswordInput fullWidth source="password" />
                  <PasswordInput fullWidth source="retype_password" />
                </Box>
              )}
              <BooleanInput
                label="resources.users.fields.change_password_on_next_login"
                source="change_password_on_next_login"
              />
              <BooleanInput
                label="resources.users.fields.send_activation_email"
                source="send_activation_email"
              />
              <BooleanInput
                label="resources.users.fields.active"
                source="active"
              />
              <BooleanInput
                label="resources.users.fields.lockout_enabled"
                source="lockout_enabled"
              />
            </FormTab>
            <FormTab label={"resources.users.tabs.roles"}>
              <RoleInput source="role_ids" />
            </FormTab>
          </TabbedForm>
        </Grid>
      </AppGridContainer>
    </AppAnimate>
  );
};
export default UserForm;
