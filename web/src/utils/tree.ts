import { get } from "lodash";

export const flatDeepNode = (
  node: any,
  nodesMap: Record<string, any> = {},
  fieldId: string = "id",
  parentIdField: string = "parent_id"
): any => {
  //   let copiedNode = { ...node };

  populateNode(node, parentIdField);

  const key = get(node, fieldId);
  const parentKey = get(node, parentIdField);

  if (!nodesMap[key]) {
    nodesMap[key] = node;
  }

  if (parentKey) {
    const parent = nodesMap[parentKey];
    if (parent) {
      node.all_parent_ids = [...node.all_parent_ids, ...parent.all_parent_ids];
    }
  }

  for (var child of node.children) {
    node.all_children_ids.push(get(child, fieldId));
    child = flatDeepNode(child, nodesMap, fieldId, parentIdField);

    if (child["all_children_ids"]) {
      node.all_children_ids = [
        ...node.all_children_ids,
        ...child.all_children_ids,
      ];
    }
  }

  return node;
};

export const addNode = (
  nodes: Record<string, any>,
  newNode: any,
  fieldId: string = "id",
  parentIdField: string = "parent_id"
) => {
  const key = -1;

  if (typeof nodes === "undefined") return;

  if (nodes[key]) return;

  populateNode(newNode, parentIdField);

  newNode[fieldId] = key;
  const parentKey = newNode[parentIdField];
  if (parentKey && nodes[parentKey]) {
    let parent = nodes[parentKey];
    parent.children = [...parent.children, newNode];
    parent.all_parent_ids.forEach((id: any) => {
      nodes[id].all_children_ids.push(newNode.id);
    });

    newNode.all_parent_ids = [parentKey, ...parent.all_parent_ids];
  }

  nodes[key] = newNode;

  return nodes;
};

export const removeNode = (
  nodes: Record<string, any>,
  id: any,
  fieldId: string = "id",
  parentIdField: string = "parent_id"
) => {
  if (typeof nodes === "undefined") return;

  const node = nodes[id];
  if (!node) return;

  const parentKey = node[parentIdField];

  if (parentKey && nodes[parentKey]) {
    let parent = nodes[parentKey];
    parent.children = parent.children.filter(
      (child: any) => child[fieldId] != id
    );
  }

  node.all_parent_ids.forEach((pId: any) => {
    nodes[pId].all_children_ids = nodes[pId].all_children_ids.filter(
      (i: any) => i != id
    );
  });

  delete nodes[id];

  return nodes;
};

const populateNode = (node: any, parentIdField: string) => {
  if (!node["all_parent_ids"]) {
    const parentKey = node[parentIdField];
    node["all_parent_ids"] = parentKey ? [parentKey] : [];
  }
  if (!node["all_children_ids"]) {
    node["all_children_ids"] = [];
  }
  if (!node["children"]) {
    node["children"] = [];
  }
};
