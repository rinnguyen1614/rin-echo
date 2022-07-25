import { SelectInput, SelectInputProps } from "ra-ui-materialui";
import { FC } from "react";
import data from "./data";
import { RequestMethodField } from "./RequestMethodField";

interface Props extends Omit<SelectInputProps, "choices" | "optionText"> {}

const RequestMethodInput: FC<Props> = ({ source, ...props }) => {
  return (
    <SelectInput
      source={source}
      choices={data}
      optionText={(record) => <RequestMethodField record={record} />}
      {...props}
    />
  );
};

export default RequestMethodInput;
