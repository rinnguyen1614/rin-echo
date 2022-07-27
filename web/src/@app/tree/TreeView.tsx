import {
  ChevronRight as ChevronRightIcon,
  ExpandMore as ExpandMoreIcon,
} from "@mui/icons-material";
import {
  TreeItem as MuiTreeItem,
  TreeView as MuiTreeView,
  TreeViewPropsBase,
} from "@mui/lab";
import { Box, Card, Paper, styled } from "@mui/material";
import { union } from "lodash";
import PropTypes from "prop-types";
import { parse } from "query-string";
import React, {
  cloneElement,
  isValidElement,
  MouseEvent,
  ReactElement,
  useCallback,
  useEffect,
  useMemo,
  useRef,
} from "react";
import {
  Empty,
  FunctionToElement,
  ListToolbar,
  ListViewProps,
  RecordContextProvider,
  ResourceContextProvider,
  sanitizeListRestProps,
  Title,
  useCreatePath,
  useListContext,
  useRecordSelection,
  useTranslate,
} from "react-admin";
import { matchPath, useLocation, useNavigate } from "react-router";
import { Record } from "../types";
import { useRecordExpand } from "../useRecordExpand";
import { useTree } from "./useTree";
import { TreeActions as DefaultActions } from "./TreeActions";
import { TreeAside } from "./TreeAside";

const defaultActions = <DefaultActions />;
const defaultEmpty = <Empty />;
const DefaultComponent = Card;

export const TreeView = <RecordType extends Record = any>(
  props: TreeViewProps<RecordType>
) => {
  const {
    actions = defaultActions,
    className,
    onSelect,
    component: Content = DefaultComponent,
    create,
    edit,
    show,
    empty = defaultEmpty,
    emptyWhileLoading,
    title,
    fieldId,
    parentFieldId,
    resource,
    resourceTree = resource,
    addRootButton,
    nodeText,
    nodeActions,
    defaultExpanded,
    autoExpandAll,
    ...rest
  } = props;

  const navigate = useNavigate();
  const createPath = useCreatePath();
  const location = useLocation();
  const translate = useTranslate();

  const matchCreate = useMemo(
    () => matchPath(`/${resource}/create`, location.pathname),
    [location, resource]
  );

  const parentIdParam = useMemo(() => {
    if (!!matchCreate) {
      const searchParams = parse(location.search);
      return Number(searchParams.parent_id || 0);
    }
    return 0;
  }, [matchCreate, location.search]);

  const matchEdit = useMemo(
    () => !matchCreate && matchPath(`/${resource}/:id`, location.pathname),
    [matchCreate, location, resource]
  );

  const { defaultTitle, data, isLoading, total } =
    useListContext<RecordType>(props);

  const [tree, { set, addNode, removeNode }] = useTree(
    resourceTree,
    fieldId,
    parentFieldId
  );

  const [expanded, { expand, clear: clearExpand }] = useRecordExpand(resource);

  const [selected, { select }] = useRecordSelection(resource);

  const lastExpanded = useRef(null);

  const expandAllOfNode = useCallback(
    (id, include?: boolean) => {
      if (tree?.data) {
        const newExpand = tree.data[id]?.all_parent_ids || [];
        if (include) {
          newExpand.push(id);
        }

        newExpand.length &&
          expand(union(newExpand.map(String), lastExpanded.current));
      }
    },

    [tree.data, expand]
  );

  const unsaveNewNode = useCallback(
    () => tree.data[-1] && removeNode(-1),
    [removeNode, tree.data]
  );

  const isFirstRender = useRef(true);
  useEffect(() => {
    if (isFirstRender.current) {
      isFirstRender.current = false;
      lastExpanded.current = null;
      clearExpand();
      return;
    }
    lastExpanded.current = expanded;
  }, [clearExpand, expanded]);

  useEffect(() => {
    set(data);
  }, [data, set]);

  useEffect(() => {
    if (autoExpandAll) {
      expand(
        Object.keys(tree.data).filter(
          (id) => tree.data[id].all_children_ids.length
        )
      );
    } else if (defaultExpanded?.length) {
      expand(defaultExpanded);
    }
  }, [defaultExpanded, expand, autoExpandAll, tree.data]);

  // for creation
  // if location is creation and it contains create element, default addNode and expand all parent of node.
  useEffect(() => {
    if (!!matchCreate) {
      addNode({
        id: -1,
        name: translate("tree.new_node"),
        parent_id: parentIdParam,
      });
      expandAllOfNode(parentIdParam, true);
      select([-1]);
    }
  }, [matchCreate, expandAllOfNode, select, addNode, parentIdParam, translate]);

  // for edition
  // if location is edition and it contains edit element, expand all parent of node and focus node.
  useEffect(() => {
    if (!!matchEdit) {
      expandAllOfNode(matchEdit.params.id);
      select([matchEdit.params.id]);
    }
  }, [matchEdit, expandAllOfNode, select]);

  const handleNodeToggle = useCallback(
    (event, nodeIds: any[]) => {
      expand(nodeIds);
    },
    [expand]
  );

  const handleNodeSelect = useCallback(
    (event, nodeIds: any[]) => {
      unsaveNewNode();
      select(nodeIds);
    },
    [select, unsaveNewNode]
  );

  const handleNodeFocus = useCallback((event, value) => {}, []);

  const handleSelect = useCallback(
    (event, node: any) => {
      unsaveNewNode();
      onSelect && onSelect(node, event);
      if (node.id !== -1) {
        const type = edit ? "edit" : show ? "show" : null;
        type && navigate(createPath({ type: type, resource, id: node.id }));
      }
      event.stopPropagation();
    },
    [onSelect, navigate, createPath, unsaveNewNode, edit, show, resource]
  );

  if (!data && isLoading && emptyWhileLoading) {
    return null;
  }

  const renderTreeItem = (node, rowIndex) => {
    if (!node) return null;
    const id = node[fieldId];
    return (
      <RecordContextProvider key={id} value={node}>
        <MuiTreeItem
          key={id}
          nodeId={id + ""}
          label={
            <Box onClick={(event) => handleSelect(event, node)}>
              {isValidElement(nodeText) ? nodeText : nodeText(node, id)}
            </Box>
          }
        >
          {node.children && Array.isArray(node.children)
            ? node.children?.map((child: any, rowIndex: number) =>
                renderTreeItem(child, rowIndex)
              )
            : null}
        </MuiTreeItem>
      </RecordContextProvider>
    );
  };

  const renderTree = () =>
    tree &&
    tree.data && (
      <div className={TreeViewClasses.main}>
        <Card>
          {actions && (
            <ListToolbar
              actions={cloneElement(actions, {
                addRootButton: valueOrDefault(
                  actions.props.addRootButton,
                  addRootButton
                ),
                resource: valueOrDefault(actions.props.resource, resource),
              })}
            />
          )}
          <Content className={TreeViewClasses.content}>
            <MuiTreeView
              {...sanitizeListRestProps(rest)}
              expanded={expanded.map(String)}
              selected={selected.map(String)}
              onNodeToggle={handleNodeToggle}
              onNodeSelect={handleNodeSelect}
              onNodeFocus={handleNodeFocus}
              defaultExpanded={defaultExpanded}
            >
              {tree.rootIds?.map((key, rowIndex) =>
                renderTreeItem(tree.data[key], rowIndex)
              )}
            </MuiTreeView>
          </Content>
        </Card>
      </div>
    );

  const renderEmpty = () =>
    empty !== false &&
    cloneElement(empty, { hasCreate: !matchCreate, resource: resource });

  const shouldRenderEmptyPage = !isLoading && total === 0 && empty !== false;

  return (
    <ResourceContextProvider value={resource}>
      <Root className={className}>
        {shouldRenderEmptyPage ? renderEmpty() : renderTree()}
        <TreeAside
          edit={edit}
          show={show}
          create={create}
          resource={resource}
          onClose={unsaveNewNode}
        />
      </Root>
    </ResourceContextProvider>
  );
};

export interface TreeViewProps<RecordType extends Record = any>
  extends Omit<
      ListViewProps,
      "classes" | "pagination" | "sx" | "children" | "hasCreate" | "aside"
    >,
    Omit<
      TreeViewPropsBase,
      | "classes"
      | "title"
      | "children"
      | "onSelect"
      | "disableSelection"
      | "disabledItemsFocusable"
      | "defaultValue"
      | "defaultChecked"
    > {
  onSelect?: (node: any, event: MouseEvent<HTMLButtonElement>) => void;
  create?: ReactElement;
  edit?: ReactElement;
  show?: ReactElement;
  isLoading?: boolean;
  fieldId?: string;
  parentFieldId?: string;
  resourceTree?: string;
  addRootButton?: React.ReactElement | boolean;
  nodeText: ReactElement | FunctionToElement<RecordType>;
  nodeActions?: ReactElement | boolean;
  autoExpandAll?: boolean;
}

TreeView.propTypes = {
  className: PropTypes.string,
  defaultCollapseIcon: PropTypes.node,
  defaultEndIcon: PropTypes.node,
  defaultExpanded: PropTypes.array,
  defaultExpandIcon: PropTypes.node,
  rowStyle: PropTypes.func,
  onSelect: PropTypes.func,
  fieldId: PropTypes.string,
  parentFieldId: PropTypes.string,
  addRootButton: PropTypes.oneOfType([PropTypes.node, PropTypes.bool]),
  nodeActions: PropTypes.oneOfType([PropTypes.node, PropTypes.bool]),
  nodeText: PropTypes.oneOfType([PropTypes.func, PropTypes.element]),
  autoExpandAll: PropTypes.bool,
};

TreeView.defaultProps = {
  defaultCollapseIcon: <ExpandMoreIcon />,
  defaultExpandIcon: <ChevronRightIcon />,
  fieldId: "id",
  parentFieldId: "parent_id",
};

const PREFIX = "CustomTreeView";

export const TreeViewClasses = {
  main: `${PREFIX}-main`,
  content: `${PREFIX}-content`,
  actions: `${PREFIX}-actions`,
  noResults: `${PREFIX}-noResults`,
};

const Root = styled("div", {
  name: PREFIX,
  overridesResolver: (props, styles) => styles.root,
})(({ theme }) => ({
  display: "flex",
  [`& .${TreeViewClasses.main}`]: {
    flex: "1 1 auto",
    display: "flex",
    flexDirection: "column",
    [theme.breakpoints.up("sm")]: {
      paddingRight: 20,
    },
  },
  [`& .${TreeViewClasses.content}`]: {
    position: "relative",
    [theme.breakpoints.down("md")]: {
      boxShadow: "none",
    },
    overflow: "inherit",
    padding: 15,
  },

  [`& .${TreeViewClasses.actions}`]: {
    zIndex: 2,
    display: "flex",
    justifyContent: "flex-end",
    flexWrap: "wrap",
  },
  [`& .${TreeViewClasses.noResults}`]: { padding: 20 },
  [`& .MuiTreeItem-content`]: {
    width: "auto",
  },
}));

const valueOrDefault = (value, defaultValue) =>
  typeof value === "undefined" ? defaultValue : value;
