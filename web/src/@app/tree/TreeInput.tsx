import ChevronRightIcon from "@mui/icons-material/ChevronRight";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import { TreeView as MuiTreeView, TreeViewPropsBase } from "@mui/lab";
import {
  FormControl,
  FormControlProps,
  FormGroup,
  FormHelperText,
  FormLabel,
  styled,
} from "@mui/material";
import { differenceWith, get, unionWith } from "lodash";
import PropTypes from "prop-types";
import {
  ChoicesProps,
  FieldTitle,
  useChoicesContext,
  useInput,
  Validator,
  warning,
} from "ra-core";
import {
  CommonInputProps,
  InputHelperText,
  Labeled,
  LinearProgress,
  sanitizeInputRestProps,
} from "ra-ui-materialui";
import { FC, useCallback, useEffect, useMemo, useRef } from "react";
import { useTree } from "./useTree";
import { TreeInputItem } from "./TreeInputItem";

const TreeInput: FC<TreeInputPros> = (props) => {
  const {
    choices: choicesProp,
    className,
    format,
    helperText,
    label,
    isLoading: isLoadingProp,
    isFetching: isFetchingProp,
    margin = "dense",
    onBlur,
    onChange,
    optionText = "name",
    optionValue = "id",
    optionParent = "parent_id",
    parse,
    resource: resourceProp,
    source: sourceProp,
    translateChoice = true,
    validate,
    defaultExpanded: customDefaultExpanded,
    fullWidth,
    onSelect,
    ...rest
  } = props;

  const { allChoices, isLoading, resource, source } = useChoicesContext({
    choices: choicesProp,
    isFetching: isFetchingProp,
    isLoading: isLoadingProp,
    resource: resourceProp,
    source: sourceProp,
  });

  warning(
    source === undefined,
    `If you're not wrapping the TreeInput inside a ReferenceArrayInput, you must provide the source prop`
  );

  warning(
    allChoices === undefined,
    `If you're not wrapping the TreeInput inside a ReferenceArrayInput, you must provide the choices prop`
  );

  const {
    field: { onChange: formOnChange, onBlur: formOnBlur, value },
    fieldState: { error, invalid, isTouched },
    formState: { isSubmitted },
    isRequired,
  } = useInput({
    format,
    parse,
    resource,
    source,
    validate,
    onChange,
    onBlur,
    ...rest,
  });

  const [tree, { set }] = useTree(resource, optionValue, optionParent);

  const defaultExpanded = useMemo(() => {
    return customDefaultExpanded ?? Object.keys(tree.data);
  }, [customDefaultExpanded, tree]);

  useEffect(() => {
    set(allChoices);
  }, [set, allChoices]);

  const handleCheck = useCallback(
    (event, isChecked) => {
      const comparator = (value, other) => value == other;
      let newValue: any;

      if (
        Object.keys(tree.data).every(
          (k) => typeof get(tree.data[k], optionValue) === "number"
        )
      ) {
        try {
          // try to convert string value to number, e.g. '123'
          newValue = JSON.parse(event.target.value);
        } catch (e) {
          // impossible to convert value, e.g. 'abc'
          newValue = event.target.value;
        }
      } else {
        newValue = event.target.value;
      }

      const node = tree.data[newValue];
      let newValues = [newValue, ...node.all_children_ids];
      if (isChecked) {
        newValues = unionWith(value, newValues, comparator);
        for (var k of node.all_parent_ids) {
          const parent = tree.data[k];
          if (
            // eslint-disable-next-line no-loop-func
            parent.all_children_ids.every((v: any) =>
              newValues.find((n) => n == v)
            )
          ) {
            newValues.push(k);
          }
        }

        formOnChange(newValues);
      } else {
        if (get(node, optionParent)) {
          newValues = [...newValues, ...node.all_parent_ids];
        }

        formOnChange(differenceWith(value, newValues, comparator)); // eslint-disable-line eqeqeq
      }
      formOnBlur(); // Ensure field is flagged as touched
    },
    [formOnBlur, optionParent, formOnChange, value, tree, optionValue]
  );

  const handleSelect = useCallback(
    (event, node) => {
      onSelect && onSelect(event, node);
      event.stopPropagation();
    },
    [onSelect]
  );

  if (isLoading && allChoices?.length === 0) {
    return (
      <Labeled
        label={label}
        source={source}
        resource={resource}
        className={className}
        isRequired={isRequired}
        margin={margin}
      >
        <LinearProgress />
      </Labeled>
    );
  }

  return (
    <StyledFormControl
      margin={margin}
      error={(isTouched || isSubmitted) && invalid}
      fullWidth={fullWidth}
      {...sanitizeRestProps(rest)}
    >
      <MuiTreeView
        //defaultSelected
        defaultExpanded={defaultExpanded}
        {...sanitizeInputRestProps(rest)}
      >
        <FormLabel component="legend" className={TreeInputClasses.label}>
          <FieldTitle
            label={label}
            source={source}
            resource={resource}
            isRequired={isRequired}
          />
        </FormLabel>
        <FormGroup row={false}>
          {tree.rootIds.map((id: any) => (
            <TreeInputItem
              tree={tree}
              key={id}
              choice={tree.data[id]}
              id={id}
              onChange={handleCheck}
              onClick={(event) => handleSelect(event, tree.data[id])}
              optionText={optionText}
              optionValue={optionValue}
              translateChoice={translateChoice}
              value={value}
              {...sanitizeRestProps(rest)}
            />
          ))}
        </FormGroup>
      </MuiTreeView>

      <FormHelperText>
        <InputHelperText
          touched={isTouched || isSubmitted}
          error={error?.message}
          helperText={helperText}
        />
      </FormHelperText>
    </StyledFormControl>
  );
};

export type TreeInputPros = Omit<CommonInputProps, "source" | "resource"> &
  ChoicesProps &
  Omit<TreeViewPropsBase, "defaultValue" | "onBlur" | "onChange"> &
  FormControlProps & {
    resource: string;
    source?: string;
    validate?: Validator | Validator[];
    optionParent?: string;
    onSelect?: (event: React.ChangeEvent<{}>, node: any) => void;
  };

const sanitizeRestProps = ({
  defaultExpandIcon,
  defaultCollapseIcon,
  defaultEndIcon,
  ...rest
}: any) => sanitizeInputRestProps(rest);

const PREFIX = "CustomTreeInput";

export const TreeInputClasses = {
  label: `${PREFIX}-label`,
};

const StyledFormControl = styled(FormControl, {
  name: PREFIX,
  overridesResolver: (props, styles) => styles.root,
})(({ theme }) => ({
  [`& .${TreeInputClasses.label}`]: {
    transform: "translate(0, 8px) scale(0.75)",
    transformOrigin: `top ${theme.direction === "ltr" ? "left" : "right"}`,
  },
}));

TreeInput.propTypes = {
  choices: PropTypes.arrayOf(PropTypes.any),
  classes: PropTypes.object,
  className: PropTypes.string,
  label: PropTypes.string,
  optionText: PropTypes.oneOfType([
    PropTypes.string,
    PropTypes.func,
    PropTypes.element,
  ]).isRequired,
  optionValue: PropTypes.string.isRequired,
  optionParent: PropTypes.string.isRequired,
  resource: PropTypes.string,
  source: PropTypes.string,
  translateChoice: PropTypes.bool,
  defaultCollapseIcon: PropTypes.node,
  defaultEndIcon: PropTypes.node,
  defaultExpanded: PropTypes.array,
  defaultExpandIcon: PropTypes.node,
};

TreeInput.defaultProps = {
  defaultCollapseIcon: <ExpandMoreIcon />,
  defaultExpandIcon: <ChevronRightIcon />,
  optionValue: "id",
  optionParent: "parent_id",
};

export default TreeInput;
