import React from "react";
import Select from "@mui/material/Select";
import FormHelperText from "@mui/material/FormHelperText";
import { SelectProps } from "@mui/material/Select/Select";
import { AppFormFieldProps } from "./AppFormFieldProps";

const AppSelectField = (props: SelectProps & AppFormFieldProps) => {
  const {
    field,
    fieldState: { error, invalid, isTouched },
    formState: { isSubmitted },
  } = props;
  const errorText = error && isTouched ? error.message : "";
  return (
    <>
      <Select
        {...props}
        {...field}
        error={(isTouched || isSubmitted) && invalid}
      />
      {!props.disabled && (
        <FormHelperText style={{ color: "#f44336" }}>
          {errorText}
        </FormHelperText>
      )}
    </>
  );
};

export default AppSelectField;
