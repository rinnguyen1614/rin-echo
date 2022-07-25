import { Chip } from "@mui/material";
import { makeStyles } from "@mui/styles";
import classnames from "classnames";
import { FC, memo, ReactElement } from "react";
import {
  ChipFieldProps,
  Datagrid,
  DateField,
  ListProps,
  NullableBooleanInput,
  TextField,
  TextInput,
  useRecordContext,
  useTranslate,
} from "react-admin";
import ListBase from "@app/auth/ListBase";
import ListActions from "@app/ListActions";
import AppAnimate from "@crema/core/AppAnimate";
import AppsContainer from "@crema/core/AppsContainer";
import AppsContent from "@crema/core/AppsContainer/AppsContent";

const useStyles = makeStyles(
  {
    chip: { margin: 4, cursor: "inherit" },
  },
  { name: "RaChipField" }
);

const RoleOptions: FC<ChipFieldProps> = memo<ChipFieldProps>((props) => {
  const { className } = props;
  const record = useRecordContext(props);
  const classes = useStyles(props);
  const translate = useTranslate();
  return record ? (
    <>
      {record.is_default ? (
        <Chip
          key={record.id + "_default"}
          className={classnames(classes.chip, className)}
          label={translate("resources.roles.fields.is_default")}
        />
      ) : null}
      {record.is_static ? (
        <Chip
          key={record.id + "_static"}
          color="primary"
          className={classnames(classes.chip, className)}
          label={translate("resources.roles.fields.is_static")}
        />
      ) : null}
    </>
  ) : null;
});

const roleFilters = [
  <TextInput
    label="Search"
    source="q=name:like,slug:like"
    alwaysOn
    variant="outlined"
  />,
  <NullableBooleanInput
    label="Default"
    source="is_default"
    variant="outlined"
  />,
];

const RoleList = (props: ListProps): ReactElement => {
  const translate = useTranslate();

  return (
    <AppAnimate animation="transition.slideUpIn" delay={200}>
      <AppsContainer title="Roles" fullView>
        <AppsContent>
          <ListBase
            {...props}
            actions={<ListActions />}
            filters={roleFilters}
            filter={{ select: "id,slug,name,is_default,is_static, created_at" }}
          >
            <Datagrid optimized rowClick="edit">
              <TextField source="slug" />
              <TextField source="name" />
              <RoleOptions />
              <DateField
                source="created_at"
                locales="us-US"
                showTime={true}
                label={translate("rin.model.created_at")}
              />
            </Datagrid>
          </ListBase>
        </AppsContent>
      </AppsContainer>
    </AppAnimate>
  );
};

export default RoleList;
