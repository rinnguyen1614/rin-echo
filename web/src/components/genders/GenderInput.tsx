import { SelectInput, SelectInputProps } from "ra-ui-materialui";
import { FC } from "react";
import data from "./data";
import { GenderField } from "./GenderField";

interface Props extends Omit<SelectInputProps, "choices" | "optionText"> {}

const GenderInput: FC<Props> = ({ source, ...props }) => {
  return (
    <SelectInput
      source={source}
      choices={data}
      optionText={(record) => <GenderField record={record} />}
      {...props}
    />
  );
};

export default GenderInput;
