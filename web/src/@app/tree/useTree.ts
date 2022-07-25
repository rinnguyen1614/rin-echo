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
  const [nodes, setNodes] = useStore(storeKey, defaultTree);
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
      return nodes.data[id];
    },
    [nodes]
  );

  const addNode = useCallback(
    (newNode: Record) => {
      if (typeof newNode === "undefined") return;
      setNodes((nodes) => {
        addNodeAndRoot(nodes, newNode, fieldId, parentIdField);
        return { ...nodes };
      });
    },
    [setNodes, fieldId, parentIdField]
  );

  const removeNode = useCallback(
    (id: Identifier) => {
      if (typeof id === "undefined") return;
      setNodes((nodes) => {
        const iRoot = nodes.rootIds.findIndex((el) => el == id); // eslint-disable-line eqeqeq
        if (iRoot !== -1) {
          nodes.rootIds = [
            ...nodes.rootIds.slice(0, iRoot),
            ...nodes.rootIds.slice(iRoot + 1),
          ];
        }
        util.removeNode(nodes.data, id, fieldId, parentIdField);

        return { ...nodes };
      });
    },
    [setNodes, fieldId, parentIdField]
  );

  const clear = useCallback(() => reset(), [reset]);

  return [
    nodes,
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

const addNodeAndRoot = (
  nodes: Tree,
  newNode: any,
  fieldId: string,
  parentIdField: string
) => {
  util.addNode(nodes.data, newNode, fieldId, parentIdField);
  const isRoot = !newNode[parentIdField];
  const id = newNode[fieldId];
  if (isRoot && !nodes.rootIds.includes(id)) {
    nodes.rootIds.push(id);
  }
  return nodes;
};
