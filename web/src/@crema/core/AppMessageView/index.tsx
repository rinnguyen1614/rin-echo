import React, { useEffect } from "react";
import { styled } from "@mui/material/styles";
import clsx from "clsx";
import CheckCircleIcon from "@mui/icons-material/CheckCircle";
import ErrorIcon from "@mui/icons-material/Error";
import InfoIcon from "@mui/icons-material/Info";
import CloseIcon from "@mui/icons-material/Close";
import IconButton from "@mui/material/IconButton";
import SnackbarContent from "@mui/material/SnackbarContent";
import WarningIcon from "@mui/icons-material/Warning";
import Snackbar from "@mui/material/Snackbar";
import { hideMessage } from "../../redux/actions";
import { useDispatch } from "react-redux";
import { Slide } from "@mui/material";
import { amber, green } from "@mui/material/colors";

const PREFIX = "AppMessageView";

const classes = {
  success: `${PREFIX}-success`,
  error: `${PREFIX}-error`,
  info: `${PREFIX}-info`,
  warning: `${PREFIX}-warning`,
  icon: `${PREFIX}-icon`,
  iconVariant: `${PREFIX}-iconVariant`,
  message: `${PREFIX}-message`,
};

const StyledSnackbar = styled(Snackbar)(({ theme }) => ({
  [`& .${classes.success}`]: {
    backgroundColor: green[600],
  },

  [`& .${classes.error}`]: {
    backgroundColor: theme.palette.error.main,
  },

  [`& .${classes.info}`]: {
    backgroundColor: theme.palette.primary.light,
  },

  [`& .${classes.warning}`]: {
    backgroundColor: amber[700],
  },

  [`& .${classes.icon}`]: {
    fontSize: 20,
  },

  [`& .${classes.iconVariant}`]: {
    opacity: 0.9,
    marginRight: theme.spacing(1),
  },

  [`& .${classes.message}`]: {
    display: "flex",
    alignItems: "center",
  },
}));

const variantIcon = {
  success: CheckCircleIcon,
  warning: WarningIcon,
  error: ErrorIcon,
  info: InfoIcon,
};

function TransitionLeft(props: any) {
  return <Slide {...props} direction="left" />;
}

interface AppMessageViewProps {
  className?: string;
  message: string;
  variant: "success" | "error" | "warning" | "info";

  [x: string]: any;
}

const AppMessageView: React.FC<AppMessageViewProps> = (props) => {
  const [open, setOpen] = React.useState(false);
  const { className, message, variant, ...other } = props;
  const Icon = variantIcon[variant];
  const dispatch = useDispatch();
  const onClose = () => {
    setOpen(false);
    setTimeout(() => dispatch(hideMessage()), 500);
  };
  useEffect(() => {
    if (message) {
      setOpen(true);
    }
  }, [message]);

  return (
    <StyledSnackbar
      anchorOrigin={{
        vertical: "bottom",
        horizontal: "right",
      }}
      open={open}
      onClose={onClose}
      autoHideDuration={3500}
      TransitionComponent={TransitionLeft}
    >
      <SnackbarContent
        className={clsx(classes[variant], className)}
        aria-describedby="client-snackbar"
        message={
          <span id="client-snackbar" className={classes.message}>
            <Icon className={clsx(classes.icon, classes.iconVariant)} />
            {message}
          </span>
        }
        action={[
          <IconButton
            key="close"
            aria-label="close"
            color="inherit"
            onClick={onClose}
            size="large"
          >
            <CloseIcon className={classes.icon} />
          </IconButton>,
        ]}
        {...other}
      />
    </StyledSnackbar>
  );
};

export default AppMessageView;
