import React, { ReactNode } from "react";
import { styled } from "@mui/material/styles";
import SimpleBarReact from "simplebar-react";

interface AppsContentContainerProps {
  children: ReactNode;
  isDetailView?: boolean;
  fullView?: boolean;

  [x: string]: any;
}

const AppsContentContainer: React.FC<AppsContentContainerProps> = styled(
  SimpleBarReact
)((props: AppsContentContainerProps) => {
  return {
    width: "100%",
    paddingTop: 8,
    paddingBottom: 8,
    display: "flex",
    flexDirection: "column",
    height: `calc(100% - ${props.isDetailView ? 60 : 129}px)`,
    [props.theme.breakpoints.up("sm")]: {
      height: `calc(100% - ${props.fullView ? 0 : 60}px)`,
    },
    "& .simplebar-content": {
      height: "100%",
    },
  };
});

interface AppsContentProps {
  children: ReactNode;
  isDetailView?: boolean;
  fullView?: boolean;

  [x: string]: any;
}

const AppsContent: React.FC<AppsContentProps> = (props) => {
  return (
    <AppsContentContainer {...props}>{props.children}</AppsContentContainer>
  );
};

export default AppsContent;
