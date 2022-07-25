import AppAutoComplete from "./AppAutoComplete";
import React from "react";

const AppAutocompleteField = (props: any) => {
  const {
    field,
    fieldState: { error, invalid, isTouched },
    formState: { isSubmitted },
  } = props;
  const errorText = error && isTouched ? error.message : "";
  return (
    <AppAutoComplete
      {...props}
      {...field}
      helperText={!props.disabled ? errorText : ""}
      error={(isTouched || isSubmitted) && invalid}
    />
  );
};

export default AppAutocompleteField;
