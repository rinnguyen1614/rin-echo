import { useTranslate } from "react-admin";

const FormattedMessage = (props: any) => {
  const translate = useTranslate();
  const { id, options } = props;
  return <>{translate(id, options)}</>;
};

export default FormattedMessage;
