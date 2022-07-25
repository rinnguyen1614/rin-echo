import React, { useState } from "react";
import { SketchPicker } from "react-color";
import Box from "@mui/material/Box";

interface CustomColorPickerProps {
  title: string;
  color: string;
  onUpdateColor: (color: string) => void;
}

const CustomColorPicker: React.FC<CustomColorPickerProps> = ({
  title,
  color,
  onUpdateColor,
}) => {
  const [visible, setVisibility] = useState(false);

  return (
    <>
      <Box
        sx={{
          display: "flex",
          alignItems: "center",
          position: "relative",
          cursor: "pointer",
          marginBottom: 2.5,
          marginRight: 2.5,
        }}
        onClick={() => setVisibility(!visible)}
      >
        <Box
          sx={{
            width: 35,
            height: 35,
            mx: 1,
            my: 1,
            borderRadius: 1,
            backgroundColor: color,
          }}
        />
        <Box component="span" className="font-extrabold">
          {title}
        </Box>
      </Box>
      {visible ? (
        <Box
          sx={{
            position: "absolute",
            left: 0,
            top: 0,
            zIndex: 1,
          }}
          onClick={() => setVisibility(!visible)}
        >
          <SketchPicker
            color={color}
            onChangeComplete={(color) => onUpdateColor(color.hex)}
          />
        </Box>
      ) : null}
    </>
  );
};

export default CustomColorPicker;
