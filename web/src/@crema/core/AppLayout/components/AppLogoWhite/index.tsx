import React from "react";
import Hidden from "@mui/material/Hidden";
import { Box } from "@mui/material";

const AppLogoWhite = () => {
  return (
    <Box
      sx={{
        display: "flex",
        flexDirection: "row",
        cursor: "pointer",
        alignItems: "center",
      }}
    >
      <Hidden smUp>
        <img
          style={{
            height: 30,
            marginRight: 10,
          }}
          src={"/assets/images/logo-white.png"}
          alt="crema-logo"
        />
      </Hidden>
      <Hidden smDown>
        <img
          style={{
            height: 30,
            marginRight: 10,
          }}
          src={"/assets/images/logo-white-with-name.png"}
          alt="crema-logo"
        />
      </Hidden>
    </Box>
  );
};

export default AppLogoWhite;
