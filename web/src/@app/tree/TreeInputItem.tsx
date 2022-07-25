import { Checkbox, FormControlLabel, styled } from "@mui/material";
import { TreeItem as MuiTreeItem } from "@mui/lab";
import { useChoices } from "ra-core";

export const TreeInputItem = (props: any) => {
  const {
    id,
    tree,
    choice,
    className,
    onChange,
    optionText,
    optionValue,
    options,
    translateChoice,
    value,
    onClick,
    ...rest
  } = props;
  const { getChoiceText, getChoiceValue } = useChoices({
    optionText,
    optionValue,
    translateChoice,
  });

  const renderTree = (choice: any) => {
    const choiceName = getChoiceText(choice);
    const choiceValue = getChoiceValue(choice);

    const checked = value
      ? value.find((v) => v == choiceValue) !== undefined
      : false;

    const indeterminate =
      !checked &&
      choice.all_children_ids &&
      value &&
      value.some(
        (v) => choice.all_children_ids.find((c) => c == v) !== undefined
      );

    return (
      <StyledMuiTreeItem
        key={choiceValue}
        nodeId={choiceValue + ""}
        className={className}
        onClick={onClick}
        label={
          <FormControlLabel
            htmlFor={`${id}_${choiceValue}`}
            key={choiceValue}
            onChange={onChange}
            control={
              <Checkbox
                id={`${id}_${choiceValue}`}
                color="primary"
                value={String(choiceValue)}
                checked={checked}
                indeterminate={indeterminate}
                onClick={(e) => e.stopPropagation()}
                {...options}
                {...rest}
              />
            }
            label={choiceName}
          />
        }
      >
        {Array.isArray(choice.children)
          ? choice.children?.map((child: any) =>
              renderTree(tree.data[getChoiceValue(child)])
            )
          : null}
      </StyledMuiTreeItem>
    );
  };

  return <>{renderTree(choice)}</>;
};

const PREFIX = "CustomTreeInputItem";

export const TreeInputItemClasses = {};

const StyledMuiTreeItem = styled(MuiTreeItem, {
  name: PREFIX,
  overridesResolver: (props, styles) => styles.root,
})({
  display: "block",
});
