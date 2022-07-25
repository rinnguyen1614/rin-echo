import React from "react";
import Box from "@mui/material/Box";
import { CustomizerItemWrapper } from "../index.style";
import FormattedMessage from "../../../utility/FormattedMessage";
import { layoutTypes } from "../../../services/db/navigationStyle";
import {
  useLayoutActionsContext,
  useLayoutContext,
} from "../../../utility/AppContextProvider/LayoutContextProvider";
import AppSelectedIcon from "../../AppSelectedIcon";

const LayoutTypes = () => {
  const { updateLayoutType } = useLayoutActionsContext();
  const { layoutType } = useLayoutContext();

  const onLayoutChange = (layoutType: string) => {
    updateLayoutType(layoutType);
  };
  return (
    <CustomizerItemWrapper pb={1}>
      <Box component="h4" sx={{ mb: 3 }}>
        <FormattedMessage id="customizer.layoutTypes" />
      </Box>
      <Box
        sx={{
          display: "flex",
          alignItems: "center",
          flexWrap: "wrap",
          marginLeft: "-10px",
          marginRight: "-10px",
        }}
      >
        {layoutTypes.map((layout) => {
          return (
            <Box
              sx={{
                paddingLeft: 2.5,
                paddingRight: 2.5,
                minWidth: 84,
                maxHeight: 48,
              }}
              key={layout.id}
            >
              <Box
                sx={{
                  position: "relative",
                  cursor: "pointer",
                }}
                onClick={() => onLayoutChange(layout.alias)}
              >
                <img src={layout.image} alt="nav" />
                {layoutType === layout.alias ? <AppSelectedIcon /> : null}
              </Box>
            </Box>
          );
        })}
      </Box>
    </CustomizerItemWrapper>
  );
};

export default LayoutTypes;
