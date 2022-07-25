import { Identifier, useRemoveFromStore, useStore } from "ra-core";
import { useMemo } from "react";

export const useRecordExpand = (
  resource: string
): [
  Identifier[],
  {
    expand: (ids: Identifier[]) => void;
    unexpand: (ids: Identifier[]) => void;
    toggle: (id: Identifier) => void;
    clear: () => void;
  }
] => {
  const storeKey = `${resource}.expandedIds`;
  const [ids, setIds] = useStore(storeKey, defaultExpand);
  const reset = useRemoveFromStore(storeKey);

  const expandModifiers = useMemo(
    () => ({
      expand: (idsToAdd: Identifier[]) => {
        if (!idsToAdd) return;
        setIds([...idsToAdd]);
      },
      unexpand(idsToRemove: Identifier[]) {
        if (!idsToRemove || idsToRemove.length === 0) return;
        setIds((ids) => {
          if (!Array.isArray(ids)) return [];
          return ids.filter((id) => !idsToRemove.includes(id));
        });
      },
      toggle: (id: Identifier) => {
        if (typeof id === "undefined") return;
        setIds((ids) => {
          if (!Array.isArray(ids)) return [...ids];
          const index = ids.indexOf(id);
          return index > -1
            ? [...ids.slice(0, index), ...ids.slice(index + 1)]
            : [...ids, id];
        });
      },
      clear: () => {
        reset();
      },
    }),
    [setIds, reset]
  );

  return [ids, expandModifiers];
};

const defaultExpand = [];
