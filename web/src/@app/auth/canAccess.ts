export const canAccess = ({ permissions, resource, action }) => {
  const permission = permissions?.find((per) => resource === per.name);
  return (
    permission &&
    (permission.actions.includes(action) ||
      permission.actions.includes(`${resource}:${action}`))
  );
};
