import React from "react";
import TextField from "@mui/material/TextField";
import Autocomplete from "@mui/material/Autocomplete";
import CircularProgress from "@mui/material/CircularProgress";
import { Chip } from "@mui/material";
import { AutocompleteProps } from "@mui/material/Autocomplete/Autocomplete";

export interface AppAutoCompleteProps
  extends AutocompleteProps<any, any, any, any> {
  onType: (e: string) => void;
  options: any[];
  onChange: () => void;
  handleChange: (e: any) => void;
  placeholder: string;
  keyName: string;
  idField: string;
  value: any;
  name: string;
  disabled: boolean;
  multiple: boolean;
  dataLoading: boolean;
  helperText: string;
  error: boolean;
  disabledId: any[];
}

const AppAutoComplete: React.FC<AppAutoCompleteProps> = ({
  options = [],
  onType = () => {},
  keyName,
  idField = "id",
  name,
  placeholder,
  dataLoading,
  handleChange,
  disabled,
  disabledId = [],
  value,
  helperText = "",
  error,
  multiple = false,
}) => {
  const loading = !disabled && dataLoading;

  const onSelectValue = (e: any, value: any[] | any) => {
    const event = {
      target: {
        name,
        value: multiple
          ? value.map((data: any) => data?.[idField])
          : value?.[idField],
      },
    };
    if (handleChange) handleChange(event);
  };

  const getValue = () => {
    if (multiple) {
      if (value) {
        return options?.filter((option) => value.includes(option?.[idField]));
      } else {
        return [];
      }
    }
    return options?.find((option) => option?.[idField] === value) || null;
  };

  return (
    <Autocomplete
      disabled={disabled}
      multiple={multiple}
      onChange={onSelectValue}
      isOptionEqualToValue={(option, value) => {
        if (multiple) {
          return option?.[idField] === value?.[idField];
        } else {
          return option?.[idField] === value?.[idField];
        }
      }}
      getOptionLabel={(option) => option?.[keyName]}
      options={options}
      loading={loading}
      // name={name}
      value={getValue()}
      renderTags={(tagValue, getTagProps) =>
        tagValue.map((option, index) => (
          <Chip
            label={option[keyName]}
            {...getTagProps({ index })}
            disabled={disabledId.indexOf(option?.[idField]) !== -1}
          />
        ))
      }
      renderInput={(params) => (
        <TextField
          name={name}
          placeholder={placeholder}
          {...params}
          variant="outlined"
          onChange={(ev) => onType(ev.target.value)}
          InputProps={{
            ...params.InputProps,
            endAdornment: (
              <React.Fragment>
                {loading ? (
                  <CircularProgress color="inherit" size={20} />
                ) : null}
                {params.InputProps.endAdornment}
              </React.Fragment>
            ),
          }}
          helperText={helperText}
          error={error}
        />
      )}
    />
  );
};
export default AppAutoComplete;
