import moment from "moment";

export const getFormattedDateTime = (
  value = 0,
  unit = "days",
  format = "YYYY-MM-DD"
) => {
  if (value === 0) {
    return moment().format(format);
  } else {
    // @ts-ignore
    return moment().add(value, unit).format(format);
  }
};

export const timeFromNow = (date: string) => {
  const timestamp = moment(date).format("X");
  const newDate = moment.unix(Number(timestamp));
  return moment(newDate).fromNow();
};
