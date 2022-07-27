import React, { useMemo } from "react";
import clsx from "clsx";
import VerticalCollapse from "../VerticalCollapse";
import VerticalItem from "../VerticalItem";
import FormattedMessage from "../../../../../utility/FormattedMessage";
import { useSidebarContext } from "../../../../../utility/AppContextProvider/SidebarContextProvider";
import VerticalNavGroupItem from "./VerticalNavGroupItem";
import {
  GetMessageId,
  RouterConfig,
} from "../../../../../types/models/RouterConfig";

interface VerticalNavGroupProps {
  item?: RouterConfig;
  level?: any;
}

const VerticalNavGroup: React.FC<VerticalNavGroupProps> = ({ item, level }) => {
  const { sidebarTextColor } = useSidebarContext();
  //   const hasPermission = useMemo(
  //     () => checkPermission(item!.permittedRole, user.role),
  //     [item, user.role]
  //   );

  //   if (!hasPermission) {
  //     return null;
  //   }
  return (
    <>
      <VerticalNavGroupItem
        level={level}
        sidebarTextColor={sidebarTextColor}
        component="div"
        className={clsx("nav-item nav-item-header")}
      >
        {<FormattedMessage id={GetMessageId(item)} />}
      </VerticalNavGroupItem>

      {item!.children && (
        <>
          {item!.children.map((item) => (
            <React.Fragment key={item.id}>
              {item.type === "group" && (
                <NavVerticalGroup item={item} level={level} />
              )}

              {item.type === "collapse" && (
                <VerticalCollapse item={item} level={level} />
              )}

              {item.type === "item" && (
                <VerticalItem item={item} level={level} />
              )}
            </React.Fragment>
          ))}
        </>
      )}
    </>
  );
};

const NavVerticalGroup = VerticalNavGroup;

export default NavVerticalGroup;
