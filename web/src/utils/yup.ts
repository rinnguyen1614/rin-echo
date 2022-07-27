import * as yup from "yup";

const rPhone =
  /^((\\+[1-9]{1,4}[ \\-]*)|(\\([0-9]{2,3}\\)[ \\-]*)|([0-9]{2,4})[ \\-]*)*?[0-9]{3,4}?[ \\-]*[0-9]{3,4}?$/;

const rPassword =
  /^(?=.*[A-Za-z])(?=.*\d)(?=.*[@$!%*#?&])[A-Za-z\d@$!%*#?&]{8,}$/;

yup.addMethod(
  yup.string,
  "password",
  function (
    message: string = "Password must contain minimum of eight characters, at least one letter, one number and one special character"
  ) {
    return this.matches(rPassword, {
      name: "password",
      message,
      excludeEmptyString: true,
    });
  }
);

yup.addMethod(
  yup.string,
  "phone",
  function (message: string = "Phone number is not valid") {
    return this.matches(rPhone, {
      name: "phone",
      message,
      excludeEmptyString: true,
    });
  }
);

export default yup;
