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
import { yupResolver } from "@hookform/resolvers/yup";
import yup from "utils/yup";

const UserForm = (props) => {
  const record = useRecordContext(props);

  const value = useMemo(() => {
    let v: any = { ...record, random_password: !record, create: !record };
    if (record && record.user_roles?.length) {
      v.role_ids = record.user_roles.map(({ role }) => role?.id);
    }
    console.log(v);
    return v;
  }, [record]);

  const { setValue, watch } = useForm(props);

  const watchAvatarPath = watch("avatar_path", record?.avatar_path);

  const [watchRandomPassword, setWatchRandomPassword] = useState(
    value.random_password
  );

  const validationSchema = yup.object({
    username: yup.string().required("Required"),
    email: yup.string().email("Invalid email format").required("Required"),
    full_name: yup.string().required("Required"),
    random_password: yup.boolean(),
    create: yup.boolean(),
    password: yup
      .string()
      .when("random_password", {
        is: false,
        then: yup.string().password(),
      })
      .when(["create", "random_password"], {
        is: (create, random_password) => create && !random_password,
        then: yup.string().required(),
      }),
    retype_password: yup.string().when("random_password", {
      is: false,
      then: yup
        .string()
        .oneOf([yup.ref("password"), null], "Passwords don't match"),
    }),
  });

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
          <TabbedForm
            resource="users"
            syncWithLocation={false}
            record={value}
            resolver={yupResolver(validationSchema)}
          >
            <FormTab label={"resources.users.tabs.general"}>
              <Box
                sx={{
                  display: "flex",
                  alignItems: "center",
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
                  <TextInput source="full_name" fullWidth />
                </Box>
              </Box>
              <TextInput
                autoFocus
                source="username"
                fullWidth
                margin="normal"
                // helperText={
                //   value.is_global_admin
                //     ? "Can not change username of the admin."
                //     : ""
                // }
                disabled={true}
              />
              <TextInput source="email" fullWidth helperText="" />
              <BooleanInput
                label="resources.users.fields.random_password"
                source="random_password"
                onChange={(event) => {
                  setWatchRandomPassword(event.target.checked);
                }}
                helperText=""
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
                helperText=""
              />
              <BooleanInput
                label="resources.users.fields.send_activation_email"
                source="send_activation_email"
                helperText=""
              />
              <BooleanInput
                label="resources.users.fields.active"
                source="active"
                helperText=""
              />
              <BooleanInput
                label="resources.users.fields.lockout_enabled"
                source="lockout_enabled"
                helperText=""
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
