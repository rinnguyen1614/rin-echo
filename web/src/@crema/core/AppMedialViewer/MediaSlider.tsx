import React, { ReactNode } from "react";
import Box from "@mui/material/Box";

interface MediaSliderProps {
  children?: ReactNode;
}

const MediaSlider: React.FC<MediaSliderProps> = ({ children }) => {
  return (
    <Box
      sx={{
        flex: 1,
        display: "flex",
        flexDirection: "column",
        justifyContent: "center",
        padding: 5,
        "& .slick-slider": {
          pb: 0,
        },
        "& .slick-slide": {
          textAlign: "center",
          position: "relative",
          "& img": {
            width: "auto !important",
            maxHeight: "96vh",
            height: "auto !important",
          },
          "& > *": {
            position: "relative",
            zIndex: 9,
          },
        },
        "& .slick-dots": {
          bottom: 10,
        },
        "& .slick-dots li button:before, & .slick-dots li.slick-active button:before":
          {
            backgroundColor: (theme) => theme.palette.background.paper,
          },
        "& .embed-responsive": {
          position: "relative",
          display: "block",
          width: "100%",
          padding: 0,
          overflow: "hidden",
          "&:before": {
            content: '""',
            display: "block",
            paddingTop: "30%",
          },
          "& embed, & iframe, & object, & video": {
            position: "absolute",
            top: 0,
            bottom: 0,
            left: 0,
            width: "100%",
            height: "100%",
            border: 0,
            color: "white",

            "& html, & body, & pre": {
              color: "white",
            },
          },
        },
        "& .slick-next": {
          right: 0,
        },
        "& .slick-prev": {
          left: 0,
        },
      }}
    >
      {children}
    </Box>
  );
};

export default MediaSlider;
