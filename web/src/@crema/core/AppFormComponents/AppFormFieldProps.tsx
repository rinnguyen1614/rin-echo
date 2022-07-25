import {
  ControllerFieldState,
  ControllerRenderProps,
  UseFormStateReturn,
} from "react-hook-form";

export type AppFormFieldProps = {
  field: ControllerRenderProps;
  formState: UseFormStateReturn<Record<string, string>>;
  fieldState: ControllerFieldState;
};
