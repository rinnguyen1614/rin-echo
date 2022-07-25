import React from "react";
import AppGridContainer from "@crema/core/AppGridContainer";
import Grid from "@mui/material/Grid";
import TextField from "@mui/material/TextField";
import FormattedMessage from "@crema/utility/FormattedMessage";
import Box from "@mui/material/Box";
import { Button } from "@mui/material";
import DatePicker from "@mui/lab/DatePicker";
import Autocomplete from "@mui/material/Autocomplete";
import countries from "@crema/services/db/countries";
import { Form, FormProps, TextInput } from "react-admin";
import { useForm } from "react-hook-form";

const InfoForm = (props) => {
  const { defaultValues } = props;
  const { setValue } = useForm();

  return (
    <Form {...props}>
      <AppGridContainer spacing={4}>
        <Grid item xs={12} md={12}>
          <TextInput
            multiline
            source="bio"
            rows={3}
            fullWidth
            label={"common.yourBioDataHere"}
            variant="outlined"
            helperText=""
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
              label={"common.birthDate"}
              value={defaultValues.dob}
              onChange={(newValue) => {
                setValue("date_of_birth", newValue);
              }}
              renderInput={(params) => <TextField {...params} />}
            />
          </Box>
        </Grid>
        <Grid item xs={12} md={6}>
          <Autocomplete
            id="country-select-demo"
            fullWidth
            options={countries}
            autoHighlight
            onChange={(_, newValue) => {
              setValue("country", newValue);
            }}
            renderOption={(props, option) => (
              <Box
                component="li"
                sx={{ "& > img": { mr: 2, flexShrink: 0 } }}
                {...props}
              >
                <img
                  loading="lazy"
                  width="20"
                  src={`https://flagcdn.com/w20/${option.code.toLowerCase()}.png`}
                  srcSet={`https://flagcdn.com/w40/${option.code.toLowerCase()}.png 2x`}
                  alt=""
                />
                {option.label} ({option.code}) +{option.phone}
              </Box>
            )}
            renderInput={(params) => (
              <TextField
                {...params}
                label={"common.country"}
                inputProps={{
                  ...params.inputProps,
                  autoComplete: "new-password", // disable autocomplete and autofill
                }}
                variant="outlined"
                helperText=""
              />
            )}
          />
        </Grid>
        <Grid item xs={12} md={6}>
          <TextInput
            source="website"
            fullWidth
            label={"common.website"}
            variant="outlined"
            helperText=""
            margin="none"
          />
        </Grid>
        <Grid item xs={12} md={6}>
          <TextInput
            fullWidth
            source="phone"
            label={"common.phoneNumber"}
            variant="outlined"
            helperText=""
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
            >
              <FormattedMessage id="common.cancel" />
            </Button>
          </Box>
        </Grid>
      </AppGridContainer>
    </Form>
  );
};

export default InfoForm;
