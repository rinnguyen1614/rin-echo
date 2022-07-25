export const createPathPattern = (basename, resource, type) => {
  switch (type) {
    case "list":
      return removeDoubleSlashes(`${basename}/${resource}`);
    case "create":
      return removeDoubleSlashes(`${basename}/${resource}/create`);
    case "edit":
      return removeDoubleSlashes(`${basename}/${resource}/:id`);
    case "show":
      return removeDoubleSlashes(`${basename}/${resource}/:id/show`);
    default:
      return type;
  }
};
export const removeDoubleSlashes = (path: string) => path.replace("//", "/");
