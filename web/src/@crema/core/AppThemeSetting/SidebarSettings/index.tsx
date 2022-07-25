import React from "react";
import Box from "@mui/material/Box";
import Switch from "@mui/material/Switch";
import CheckIcon from "@mui/icons-material/Check";
import { CustomizerItemWrapper } from "../index.style";
import {
  sidebarBgImages,
  sidebarColors,
} from "../../../services/db/navigationStyle";
import {
  useSidebarActionsContext,
  useSidebarContext,
} from "../../../utility/AppContextProvider/SidebarContextProvider";
import NavMenuStyle from "./NavMenuStyle";
import MenuColorCell from "./MenuColorCell";
import AppGrid from "../../AppGrid";

const SidebarSettings = () => {
  const { sidebarBgImage, isSidebarBgImage } = useSidebarContext();

  const { updateSidebarBgImage, setSidebarBgImage } =
    useSidebarActionsContext();

  const onToggleSidebarImage = () => {
    setSidebarBgImage(!isSidebarBgImage);
  };

  return (
    <CustomizerItemWrapper>
      <NavMenuStyle />
      <Box
        sx={{
          display: "flex",
          alignItems: "center",
          justifyContent: "center",
          mb: 2.5,
        }}
      >
        <Box component="h4">Sidebar Images</Box>
        <Box component="span" sx={{ ml: "auto" }}>
          <Switch
            className="customize-switch"
            checked={isSidebarBgImage}
            onChange={onToggleSidebarImage}
            value="checkedA"
          />
        </Box>
      </Box>
      {isSidebarBgImage ? (
        <Box
          sx={{
            display: "flex",
            alignItems: "center",
            flexWrap: "wrap",
            marginLeft: -2.5,
            marginRight: -2.5,
            mt: 5,
          }}
        >
          {sidebarBgImages.map((imagesObj) => {
            return (
              <Box
                sx={{
                  display: "flex",
                  alignItems: "center",
                  pl: 2.5,
                  pr: 2.5,
                  mb: 5,
                }}
                key={imagesObj.id}
              >
                <Box
                  sx={{
                    position: "relative",
                    cursor: "pointer",
                    " & .layout-img": {
                      width: 82,
                    },
                  }}
                  onClick={() => updateSidebarBgImage(imagesObj.id)}
                >
                  <img src={imagesObj.image} alt="nav" />
                  {sidebarBgImage === imagesObj.id ? (
                    <Box
                      sx={{
                        position: "absolute",
                        left: "50%",
                        top: "50%",
                        transform: "translate(-50%, -50%)",
                        width: 25,
                        height: 25,
                        borderRadius: "50%",
                        display: "flex",
                        alignItems: "center",
                        justifyContent: "center",
                        overflow: "hidden",
                        backgroundColor: "primary.main",
                      }}
                    >
                      <CheckIcon
                        sx={{
                          color: "white",
                          fontSize: 22,
                        }}
                      />
                    </Box>
                  ) : null}
                </Box>
              </Box>
            );
          })}
        </Box>
      ) : null}
      <Box component="h4" sx={{ mb: 3 }}>
        Sidebar Colors
      </Box>
      <AppGrid
        data={sidebarColors}
        column={2}
        itemPadding={5}
        renderRow={(colorSet, index) => (
          <MenuColorCell key={index} sidebarColors={colorSet} />
        )}
      />
    </CustomizerItemWrapper>
  );
};

export default SidebarSettings;
