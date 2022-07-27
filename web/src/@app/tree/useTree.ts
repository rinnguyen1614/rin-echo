import { Identifier, useRemoveFromStore, useStore } from "ra-core";
import { useCallback } from "react";
import * as util from "../../utils/tree";
import { Record, RecordMap } from "../types";

export const useTree = (
  resource: string,
  fieldId: string = "id",
  parentIdField: string = "parent_id"
): [
  Tree,
  {
    set: (records: Record[]) => void;
    getNode: (id: Identifier) => any;
    addNode: (record: Record, parent_id?: Identifier) => void;
    removeNode: (id: Identifier) => void;
    clear: () => void;
  }
] => {
  const storeKey = `tree.${resource}`;
  const [tree, setNodes] = useStore(storeKey, defaultTree);
  const reset = useRemoveFromStore(storeKey);

  const set = useCallback(
    (records: Record[]) => {
      if (typeof records === "undefined") return;

      let data = {};
      let rootIds = [];
      records.forEach((node) => {
        util.flatDeepNode(node, data, fieldId, parentIdField);
        rootIds.push(node[fieldId]);
      });

      setNodes({
        data: data,
        rootIds: rootIds,
      });
    },
    [setNodes, fieldId, parentIdField]
  );

  const getNode = useCallback(
    (id: Identifier) => {
      if (typeof id === "undefined") return;
      return tree.data[id];
    },
    [tree]
  );

  const addNode = useCallback(
    (newNode: Record) => {
      if (typeof newNode === "undefined") return;
      setNodes((tree) => {
        let rootIds = [...tree.rootIds];
        const isRoot = !newNode[parentIdField];
        const id = newNode[fieldId];
        if (isRoot && !rootIds.includes(id)) {
          rootIds.push(id);
        }
        return {
          rootIds: rootIds,
          data: util.addNode(tree.data, newNode, fieldId, parentIdField),
        };
      });
    },
    [setNodes, fieldId, parentIdField]
  );

  const removeNode = useCallback(
    (id: Identifier) => {
      if (typeof id === "undefined") return;
      setNodes((tree) => {
        let rootIds = [...tree.rootIds];
        const iRoot = rootIds.findIndex((el) => el == id); // eslint-disable-line eqeqeq
        if (iRoot !== -1) {
          rootIds = [...rootIds.slice(0, iRoot), ...rootIds.slice(iRoot + 1)];
        }

        return {
          rootIds: rootIds,
          data: util.removeNode(tree.data, id, fieldId, parentIdField),
        };
      });
    },
    [setNodes, fieldId, parentIdField]
  );

  const clear = useCallback(() => reset(), [reset]);

  return [
    tree,
    {
      set,
      getNode,
      addNode,
      removeNode,
      clear,
    },
  ];
};

type Tree = {
  data: RecordMap;
  rootIds: Identifier[];
};

const defaultTree: Tree = {
  data: {},
  rootIds: [],
};
