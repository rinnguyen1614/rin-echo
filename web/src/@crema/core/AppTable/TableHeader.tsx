import React, { ReactNode } from "react";
import TableRow from "@mui/material/TableRow";
import PropTypes from "prop-types";

import { styled } from "@mui/material/styles";
import { Fonts } from "../../shared/constants/AppEnums";

const TableHeaderRow = styled(TableRow)(({ theme }) => {
  return {
    "& th": {
      fontSize: 14,
      padding: 8,
      fontWeight: Fonts.MEDIUM,
      color: theme.palette.text.primary,
      "&:first-of-type": {
        paddingLeft: 20,
      },
      "&:last-of-type": {
        paddingRight: 20,
      },
    },
  };
});

interface TableHeaderProps {
  children: ReactNode;

  [x: string]: any;
}

const TableHeader: React.FC<TableHeaderProps> = ({ children, ...rest }) => {
  return <TableHeaderRow {...rest}>{children}</TableHeaderRow>;
};

export default TableHeader;
TableHeader.propTypes = {
  children: PropTypes.node,
};
