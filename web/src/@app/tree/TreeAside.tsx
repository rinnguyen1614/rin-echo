import { Fonts } from "@crema/shared/constants/AppEnums";
import FormattedMessage from "@crema/utility/FormattedMessage";
import {
  Box,
  Card,
  Drawer,
  Stack,
  styled,
  Theme,
  Typography,
  useMediaQuery,
} from "@mui/material";
import { useBasename, useResourceDefinitions } from "ra-core";
import { cloneElement, useCallback, useMemo } from "react";
import { matchPath, useLocation, useNavigate } from "react-router";
import { createPathPattern } from "../../utils/util";
import { AddChildButton } from "./AddChildButton";

export const TreeAside = (props) => {
  const { edit, show, create, resource, className, onClose, ...rest } = props;

  const navigate = useNavigate();
  const location = useLocation();
  const basename = useBasename();

  const matchCreate = useMemo(
    () =>
      matchPath(
        createPathPattern(basename, resource, "create"),
        location.pathname
      ),
    [location, resource, basename]
  );
  const matchEdit = useMemo(
    () =>
      !matchCreate &&
      matchPath(
        createPathPattern(basename, resource, "edit"),
        location.pathname
      ),
    [matchCreate, location, resource, basename]
  );

  const matchShow = useMemo(
    () =>
      matchPath(
        createPathPattern(basename, resource, "show"),
        location.pathname
      ),
    [location, resource, basename]
  );

  const handleClose = useCallback(() => {
    onClose && onClose();
    navigate(`/${resource}`);
  }, [navigate, resource, onClose]);

  const open = useMemo(
    () => !!matchEdit || !!matchCreate || !!matchShow,
    [matchEdit, matchCreate, matchShow]
  );

  const isSmall = useMediaQuery<Theme>((theme) => theme.breakpoints.down("sm"));

  const render = () => (
    <div className={TreeAsideClasses.content}>
      {!!matchEdit && (
        <Card>
          <Box sx={{ padding: 2, paddingBottom: 0 }}>
            <Stack
              direction="row"
              justifyContent="flex-end"
              alignItems="center"
              spacing={2}
              width="100%"
              minHeight={64}
            >
              {matchEdit.params.id && (
                <AddChildButton id={matchEdit.params.id} resource={resource} />
              )}
            </Stack>
          </Box>

          {cloneElement(edit, {
            id: valueOrDefault(edit.props.id, matchEdit.params.id),
            resource: valueOrDefault(edit.props.resource, resource),
          })}
        </Card>
      )}

      {!!matchShow && (
        <Card>
          {cloneElement(show, {
            id: valueOrDefault(show.props.id, matchShow.params.id),
            resource: valueOrDefault(show.props.resource, resource),
          })}
        </Card>
      )}

      {!!matchCreate && (
        <Card>
          <Box
            component="h4"
            sx={{
              m: 4,
              fontWeight: Fonts.SEMI_BOLD,
            }}
          >
            <FormattedMessage id="tree.create_node" />
          </Box>
          {cloneElement(create, {
            resource: valueOrDefault(create.props.resource, resource),
          })}
        </Card>
      )}
    </div>
  );

  return (
    <Root className={className}>
      <div className={TreeAsideClasses.main}>
        {!isSmall ? (
          render()
        ) : (
          <StyledDrawer
            open={open}
            onClose={handleClose}
            sx={{ zIndex: 100 }}
            anchor="bottom"
          >
            {render()}
          </StyledDrawer>
        )}
      </div>
    </Root>
  );
};

const PREFIX = "CustomTreeAside";

export const TreeAsideClasses = {
  main: `${PREFIX}-main`,
  content: `${PREFIX}-content`,
};

const Root = styled("div", {
  name: PREFIX,
  overridesResolver: (props, styles) => styles.root,
})(({ theme }) => ({
  [`& .${TreeAsideClasses.main}`]: {
    display: "flex",
    flexDirection: "column",
    pl: {
      lg: 8,
    },

    [theme.breakpoints.up("md")]: {
      width: 450,
    },
  },
  [`& .${TreeAsideClasses.content}`]: {},
}));

const StyledDrawer = styled(Drawer, {
  name: PREFIX,
  overridesResolver: (props, styles) => styles.root,
})(({ theme }) => ({
  [`& .${TreeAsideClasses.content}`]: {
    [theme.breakpoints.down("sm")]: {
      marginBottom: 70,
    },
  },
}));

const valueOrDefault = (value, defaultValue) =>
  typeof value === "undefined" ? defaultValue : value;
