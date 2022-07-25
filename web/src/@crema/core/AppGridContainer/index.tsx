import React, { ReactNode } from "react";
import Grid from "@mui/material/Grid";
import { useMediaQuery } from "@mui/material";
import { CremaTheme } from "../../types/AppContextPropsType";

interface AppGridContainerProps {
  children: ReactNode;

  [x: string]: any;
}

const AppGridContainer: React.FC<AppGridContainerProps> = ({
  children,
  ...others
}) => {
  const isMDDown = useMediaQuery((theme: CremaTheme) =>
    theme.breakpoints.down("md")
  );
  return (
    <Grid container spacing={isMDDown ? 5 : 8} {...others}>
      {children}
    </Grid>
  );
};

export default AppGridContainer;
