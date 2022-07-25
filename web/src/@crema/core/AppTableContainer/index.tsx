import React, { ReactNode } from "react";
import TableContainer from "@mui/material/TableContainer";
import { Theme } from "@mui/material";
import { SxProps } from "@mui/system";

interface AppTableContainerProps {
  children: ReactNode;
  sxStyle?: SxProps<Theme>;
}

const AppTableContainer: React.FC<AppTableContainerProps> = ({
  children,
  sxStyle,
}) => {
  return (
    <TableContainer
      sx={{
        "& tr > th, & tr > td": {
          whiteSpace: "nowrap",
        },
        ...sxStyle,
      }}
    >
      {children}
    </TableContainer>
  );
};

export default AppTableContainer;
