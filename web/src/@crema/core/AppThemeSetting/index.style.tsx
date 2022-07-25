import { styled } from "@mui/material/styles";
import ToggleButton from "@mui/material/ToggleButton";
import { Box } from "@mui/material";

export const CustomizerItemWrapper = styled(Box)(({ theme }) => ({
  "&:not(:last-of-type)": {
    borderBottom: [`1px solid ${theme.palette.divider}`],
    paddingBottom: 20,
    marginBottom: 20,
    [theme.breakpoints.up("xl")]: {
      paddingBottom: 30,
      marginBottom: 30,
    },
  },
}));
export const StyledToggleButton = styled(ToggleButton)(({ theme }) => ({
  height: 36,
  backgroundColor: theme.palette.background.paper,
  color: theme.palette.primary.main,
  borderColor: theme.palette.primary.main,
  "&:not(:first-of-type)": {
    borderLeft: `1px solid ${theme.palette.primary.main} !important`,
  },
  [theme.breakpoints.up("xl")]: {
    height: 44,
    minWidth: 96,
  },
  "&:hover,&:focus": {
    backgroundColor: theme.palette.background.paper,
    color: theme.palette.primary.main,
  },
  "&.active": {
    backgroundColor: theme.palette.primary.main,
    color: theme.palette.primary.contrastText,
    "&:hover,&:focus": {
      backgroundColor: theme.palette.primary.main,
      color: theme.palette.primary.contrastText,
    },
  },
}));
