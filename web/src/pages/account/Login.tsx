import React, { useState } from "react";
import Card from "@mui/material/Card";
import GitHubIcon from "@mui/icons-material/GitHub";
import TwitterIcon from "@mui/icons-material/Twitter";
import FacebookIcon from "@mui/icons-material/Facebook";
import Button from "@mui/material/Button";
import { Checkbox, useTheme } from "@mui/material";
import Grid from "@mui/material/Grid";
import IconButton from "@mui/material/IconButton";
import Box from "@mui/material/Box";
import { grey } from "@mui/material/colors";
import { Fonts } from "@crema/shared/constants/AppEnums";
import AppAnimate from "@crema/core/AppAnimate";
import { ReactComponent as Logo } from "assets/user/login.svg";
import {
  useTranslate,
  useLogin,
  useNotify,
  HttpError,
  TextInput,
  Form,
  PasswordInput,
} from "react-admin";
import { useLocation } from "react-router-dom";

interface FormValues {
  username?: string;
  password?: string;
}

const Login = () => {
  const theme = useTheme();
  const [loading, setLoading] = useState(false);
  const translate = useTranslate();
  const notify = useNotify();
  const login = useLogin();
  const location = useLocation();

  const handleSubmit = (auth: FormValues) => {
    setLoading(true);
    login(
      auth,
      location.state ? (location.state as any).nextPathname : "/"
    ).catch((error: HttpError) => {
      setLoading(false);
      notify(error.body.error.message, {
        type: "warning",
        messageArgs: {
          _: error.body.error.message,
        },
      });
    });
  };

  const validate = (values: FormValues) => {
    const errors: FormValues = {};
    if (!values.username) {
      errors.username = translate("ra.validation.required");
    }
    if (!values.password) {
      errors.password = translate("ra.validation.required");
    }
    console.log(errors);

    return errors;
  };

  return (
    <AppAnimate animation="transition.slideUpIn" delay={200}>
      <Box
        sx={{
          pb: 6,
          py: { xl: 8 },
          display: "flex",
          flex: 1,
          flexDirection: "column",
          alignItems: "center",
          justifyContent: "center",
        }}
      >
        <Card
          sx={{
            maxWidth: 1024,
            width: "100%",
            padding: 8,
            paddingLeft: { xs: 8, md: 2 },
            overflow: "hidden",
            boxShadow:
              "0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05)",
          }}
        >
          <Grid container spacing={5}>
            <Grid
              item
              xs={12}
              md={6}
              sx={{
                width: "100%",
                height: "100%",
                textAlign: "center",
                "& svg": {
                  width: "100%",
                  height: "100%",
                  display: "inline-block",
                  paddingRight: { xs: 0, lg: 10 },
                },
              }}
            >
              <Logo fill={theme.palette.primary.main} />
            </Grid>
            <Grid
              item
              xs={12}
              md={6}
              sx={{
                textAlign: "center",
              }}
            >
              <Box
                sx={{
                  mb: { xs: 3, xl: 4 },
                  fontWeight: Fonts.BOLD,
                  fontSize: 20,
                }}
              >
                {translate("ra.auth.login")}
              </Box>

              <Form onSubmit={handleSubmit}>
                <Box sx={{ mb: { xs: 5, xl: 8 } }}>
                  <TextInput
                    source="username"
                    label={translate("ra.auth.username")}
                    variant="outlined"
                    fullWidth
                  />
                </Box>

                <Box sx={{ mb: { xs: 5, xl: 8 } }}>
                  <PasswordInput
                    source="password"
                    label={translate("ra.auth.password")}
                    variant="outlined"
                    fullWidth
                  />
                </Box>

                <Box
                  sx={{
                    mb: { xs: 3, xl: 4 },
                    display: "flex",
                    flexDirection: { xs: "column", sm: "row" },
                    alignItems: { sm: "center" },
                  }}
                >
                  <Box
                    sx={{
                      display: "flex",
                      flexDirection: "row",
                      alignItems: "center",
                    }}
                  >
                    <Box
                      sx={{
                        ml: -3,
                      }}
                    >
                      <Checkbox />
                    </Box>
                    <Box
                      component="span"
                      sx={{
                        fontSize: 14,
                      }}
                    >
                      {translate("ra.login.rememberMe")}
                    </Box>
                  </Box>
                  <Box
                    component="span"
                    sx={{
                      cursor: "pointer",
                      ml: { xs: 0, sm: "auto" },
                      mt: { xs: 2, sm: 0 },
                      color: "primary.main",
                      fontWeight: Fonts.BOLD,
                      fontSize: 14,
                    }}
                  >
                    {translate("ra.login.forgetPassword")}
                  </Box>
                </Box>
                <Button
                  variant="contained"
                  color="primary"
                  type="submit"
                  disabled={loading}
                  sx={{
                    width: "100%",
                    height: 44,
                  }}
                >
                  {translate("ra.auth.sign_in")}
                </Button>
              </Form>
              <Box
                sx={{
                  mt: { xs: 3, xl: 4 },
                  mb: 3,
                  display: "flex",
                  flexDirection: { xs: "column", sm: "row" },
                  justifyContent: { sm: "center" },
                  alignItems: { sm: "center" },
                }}
              >
                <Box
                  component="span"
                  sx={{
                    color: grey[600],
                    fontSize: 14,
                    mr: 4,
                  }}
                >
                  {translate("ra.login.orLoginWith")}
                </Box>
                <Box display="inline-block">
                  <IconButton>
                    <FacebookIcon sx={{ color: "text.primary" }} />
                  </IconButton>
                  <IconButton>
                    <GitHubIcon sx={{ color: "text.primary" }} />
                  </IconButton>
                  <IconButton>
                    <TwitterIcon sx={{ color: "text.primary" }} />
                  </IconButton>
                </Box>
              </Box>

              <Box
                sx={{
                  color: "grey.700",
                  fontSize: 14,
                  fontWeight: Fonts.BOLD,
                }}
              >
                <Box
                  component="span"
                  sx={{
                    mr: 2,
                  }}
                >
                  {translate("ra.auth.dontHaveAccount")}
                </Box>
                <Box
                  component="span"
                  sx={{
                    color: "primary.main",
                    cursor: "pointer",
                  }}
                >
                  {translate("ra.auth.signup")}
                </Box>
              </Box>
            </Grid>
          </Grid>
        </Card>
      </Box>
    </AppAnimate>
  );
};

export default Login;
