import React, { ReactNode, useState } from "react";
import { Card } from "@mui/material";
import CardHeader from "@mui/material/CardHeader";
import CardContent from "@mui/material/CardContent";
import CodeIcon from "@mui/icons-material/Code";
import Highlight, { defaultProps } from "prism-react-renderer";
import IconButton from "@mui/material/IconButton";
import Collapse from "@mui/material/Collapse";
import AppScrollbar from "../AppScrollbar";
import { highlightTheme } from "./highlightTheme";
import Box from "@mui/material/Box";
import AppAnimate from "../AppAnimate";
import { Fonts } from "../../shared/constants/AppEnums";

interface AppComponentCardProps {
  title: ReactNode;
  component: any;
  source?: any;
  maxHeight?: number;
  description?: ReactNode;
  noScrollbar?: boolean;
}

const AppComponentCard: React.FC<AppComponentCardProps> = ({
  title,
  maxHeight,
  description,
  component: Component,
  source,
  noScrollbar,
}) => {
  const [viewSource, setToggleViewSource] = useState(false);
  const [animation, setAnimation] = useState(false);

  return (
    <AppAnimate animation="transition.slideUpIn" delay={200}>
      <Card>
        <CardHeader
          sx={{
            py: 4,
            pb: 1,
            px: 5,
            display: "flex",
            alignItems: "center",
            minHeight: 50,
            boxSizing: "border-box",
            "& .MuiTypography-h5": {
              fontSize: 14,
              fontWeight: Fonts.BOLD,
              marginBottom: 0.25,
            },
          }}
          title={title}
          subheader={description}
          root={{
            subheader: {
              fontSize: 13,
            },
          }}
          action={
            source ? (
              <Box>
                <IconButton
                  aria-label="view code"
                  onClick={() => {
                    if (animation) {
                      setAnimation(!animation);
                      setTimeout(() => setToggleViewSource(!viewSource), 400);
                    } else {
                      setAnimation(!animation);
                      setToggleViewSource(!viewSource);
                    }
                  }}
                  size="large"
                >
                  <CodeIcon />
                </IconButton>
              </Box>
            ) : null
          }
        />

        <CardContent sx={{ px: 4, pt: 0 }}>
          <Collapse in={animation}>
            {viewSource ? (
              <AppScrollbar
                sx={{
                  borderRadius: 3,
                  background: "#333333",
                }}
              >
                <Highlight
                  {...defaultProps}
                  code={source}
                  language="jsx"
                  theme={highlightTheme}
                >
                  {({ style, tokens, getLineProps, getTokenProps }) => (
                    <pre
                      style={{
                        ...style,
                        maxHeight: 500,
                        borderRadius: 8,
                        padding: 12,
                      }}
                    >
                      {tokens.map((line, i) => (
                        <Box
                          key={"line-" + i}
                          {...getLineProps({ line, key: i })}
                        >
                          {line.map((token, key) => (
                            <span
                              key={"token-" + key}
                              {...getTokenProps({ token, key })}
                            />
                          ))}
                        </Box>
                      ))}
                    </pre>
                  )}
                </Highlight>
              </AppScrollbar>
            ) : null}
          </Collapse>
          {noScrollbar ? (
            <Box
              sx={{
                width: "100%",
                display: "flex",
                alignItems: "center",
                justifyContent: "center",
                pt: 6,
              }}
            >
              <Component />
            </Box>
          ) : (
            <AppScrollbar
              sx={{
                mt: 2,
                p: 4,
                borderRadius: 3,
                maxHeight: maxHeight,
              }}
            >
              <Box
                sx={{
                  width: "100%",
                  display: "flex",
                  alignItems: "center",
                  justifyContent: "center",
                }}
              >
                <Component />
              </Box>
            </AppScrollbar>
          )}
        </CardContent>
      </Card>
    </AppAnimate>
  );
};

export default AppComponentCard;
