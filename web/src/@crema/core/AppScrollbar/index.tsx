import React, { ReactNode } from "react";
import { styled } from "@mui/material/styles";
import SimpleBarReact from "simplebar-react";
import "simplebar/src/simplebar.css";

const StyledSimpleBarReact = styled(SimpleBarReact)(() => ({
  height: "100%",
  width: "100%",
}));

interface AppScrollbarProps {
  children: ReactNode;
  className?: string;

  [x: string]: any;
}

const AppScrollbar: React.FC<AppScrollbarProps> = (props) => {
  const { children, ...others } = props;

  return <StyledSimpleBarReact {...others}>{children}</StyledSimpleBarReact>;
};

export default AppScrollbar;
