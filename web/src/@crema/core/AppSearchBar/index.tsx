import React from "react";
import clsx from "clsx";
import SearchIcon from "@mui/icons-material/Search";
import {
  SearchIconBox,
  SearchIconWrapper,
  SearchInputBase,
  SearchWrapper,
} from "./index.style";
import { SxProps } from "@mui/system/styleFunctionSx";
import { Theme } from "@mui/material";

interface AppSearchProps {
  iconPosition?: string;
  align?: string;
  placeholder?: string;
  overlap?: boolean;
  borderLight?: boolean;
  className?: string;
  onlyIcon?: boolean;
  disableFocus?: boolean;
  iconStyle?: SxProps<Theme>;
  sx?: SxProps<Theme>;

  [x: string]: any;
}

const AppSearch: React.FC<AppSearchProps> = ({
  placeholder,
  iconPosition = "left",
  align = "left",
  overlap = true,
  onlyIcon = false,
  disableFocus,
  iconStyle = {
    color: "grey",
  },
  sx,
  ...rest
}) => {
  return (
    <SearchWrapper sx={sx} iconPosition={iconPosition}>
      <SearchIconBox
        align={align}
        className={clsx(
          "searchRoot",
          { "hs-search": overlap },
          { "hs-disableFocus": disableFocus },
          { searchIconBox: onlyIcon }
        )}
      >
        <SearchIconWrapper
          className={clsx({
            right: iconPosition === "right",
          })}
          sx={iconStyle}
        >
          <SearchIcon />
        </SearchIconWrapper>
        <SearchInputBase
          {...rest}
          placeholder={placeholder || "Search…"}
          inputProps={{ "aria-label": "search" }}
        />
      </SearchIconBox>
    </SearchWrapper>
  );
};

export default AppSearch;
