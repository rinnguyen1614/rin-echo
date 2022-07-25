import React, { useState } from "react";
import Select from "@mui/material/Select";
import MenuItem from "@mui/material/MenuItem";
import { styled } from "@mui/material/styles";
import { alpha } from "@mui/material";

const SelectBox = styled(Select)(({ theme }) => {
  return {
    marginLeft: 8,
    cursor: "pointer",
    fontSize: 14,
    height: 24,
    "& .MuiSelect-select": {
      paddingLeft: 5,
      paddingTop: 1,
      paddingBottom: 3,
      color: "text.secondary",
    },
    "& .MuiSelect-icon": {
      color: "text.secondary",
    },
    "& .MuiOutlinedInput-notchedOutline": {
      borderColor: "transparent",
    },
    "&:hover": {
      "& .MuiOutlinedInput-notchedOutline": {
        borderColor: "transparent",
      },
    },
    "&.Mui-focused": {
      backgroundColor: alpha(theme.palette.common.black, 0.03),
      "& .MuiOutlinedInput-notchedOutline": {
        borderColor: "transparent",
      },
    },
  };
});

interface AppSelectProps {
  menus: any[];
  onChange: (e: any) => void;
  defaultValue: any;
  selectionKey?: string;
}

const AppSelect: React.FC<AppSelectProps> = ({
  menus = [],
  onChange,
  defaultValue = "",
  selectionKey = "",
}) => {
  const [selectionType, setSelectionType] = useState<string>(defaultValue);

  const handleSelectionType = (value: string) => {
    setSelectionType(value);
    onChange(value);
  };

  return (
    <SelectBox
      defaultValue={defaultValue}
      value={selectionType}
      onChange={(event) => handleSelectionType(event.target.value as string)}
      className="select-box"
    >
      {menus.map((menu, index) => (
        <MenuItem
          key={index}
          value={selectionKey ? menu[selectionKey] : menu}
          sx={{
            cursor: "pointer",
            p: 2,
            fontSize: 14,
          }}
        >
          {selectionKey ? menu[selectionKey] : menu}
        </MenuItem>
      ))}
    </SelectBox>
  );
};

export default AppSelect;
