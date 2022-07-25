import { ToolbarProps } from "@mui/material";
import PropTypes from "prop-types";
import {
  Exporter,
  Identifier,
  sanitizeListRestProps,
  useListContext,
} from "ra-core";
import * as React from "react";
import { cloneElement, isValidElement, useMemo } from "react";
import {
  CreateButton,
  ExportButton,
  TopToolbar,
  useCreatePath,
} from "react-admin";
import { matchPath, useLocation } from "react-router";

export const TreeActions = (props: TreeActionsProps) => {
  const { className, exporter, addRootButton, resource, ...rest } = props;
  const { total } = useListContext(props);
  const location = useLocation();
  const createPath = useCreatePath();

  const match = useMemo(
    () =>
      matchPath(createPath({ resource, type: "create" }), location.pathname),
    [resource, location, createPath]
  );

  return useMemo(
    () => (
      <TopToolbar className={className} {...sanitizeListRestProps(rest)}>
        {!match &&
          !!addRootButton &&
          (isValidElement(addRootButton) ? (
            cloneElement(addRootButton, {
              resource: valueOrDefault(addRootButton.props.resource, resource),
            })
          ) : (
            <CreateButton
              label={"rin.tree.actions.add_root"}
              resource={resource}
            />
          ))}
        {exporter !== false && (
          <ExportButton disabled={total === 0} resource={resource} />
        )}
      </TopToolbar>
    ),
    /* eslint-disable react-hooks/exhaustive-deps */
    [resource, total, className, exporter, addRootButton, match]
  );
};

TreeActions.propTypes = {
  className: PropTypes.string,
  exporter: PropTypes.oneOfType([PropTypes.func, PropTypes.bool]),
  addRootButton: PropTypes.oneOfType([PropTypes.node, PropTypes.bool]),
  resource: PropTypes.string,
  onUnselectItems: PropTypes.func.isRequired,
  selectedIds: PropTypes.arrayOf(PropTypes.any),
  total: PropTypes.number,
};

TreeActions.defaultProps = {
  selectedIds: [],
  onUnselectItems: () => null,
};

export interface TreeActionsProps extends ToolbarProps {
  className?: string;
  resource?: string;
  exporter?: Exporter | boolean;
  addRootButton?: React.ReactElement | boolean;
  selectedIds?: Identifier[];
  onUnselectItems?: () => void;
  total?: number;
}

const valueOrDefault = (value, defaultValue) =>
  typeof value === "undefined" ? defaultValue : value;
