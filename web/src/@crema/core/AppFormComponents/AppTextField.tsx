import React from "react";
import TextField from "@mui/material/TextField";
import { TextFieldProps } from "@mui/material/TextField/TextField";
import { AppFormFieldProps } from "./AppFormFieldProps";

const AppTextField = (props: TextFieldProps & AppFormFieldProps) => {
  const {
    field,
    fieldState: { error, invalid, isTouched },
    formState: { isSubmitted },
  } = props;
  const errorText = error && isTouched ? error.message : "";
  return (
    <TextField
      {...props}
      {...field}
      helperText={errorText}
      error={(isTouched || isSubmitted) && invalid}
    />
  );
};

export default AppTextField;
