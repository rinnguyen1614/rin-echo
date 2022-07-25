import { get } from "lodash";

export const flatDeepNode = (
  node: any,
  nodesMap: Record<string, any> = {},
  fieldId: string = "id",
  parentIdField: string = "parent_id"
): any => {
  let copiedNode = { ...node };
  populateNode(copiedNode, parentIdField);
  const key = get(copiedNode, fieldId);
  const parentKey = get(copiedNode, parentIdField);
  if (!nodesMap[key]) {
    nodesMap[key] = copiedNode;
  }

  if (parentKey) {
    const parent = nodesMap[parentKey];
    if (parent) {
      copiedNode.all_parent_ids = [
        ...copiedNode.all_parent_ids,
        ...parent.all_parent_ids,
      ];
    }
  }

  for (var child of copiedNode.children) {
    copiedNode.all_children_ids.push(get(child, fieldId));
    child = flatDeepNode(child, nodesMap, fieldId, parentIdField);
    if (child["all_children_ids"]) {
      copiedNode.all_children_ids = [
        ...copiedNode.all_children_ids,
        ...child.all_children_ids,
      ];
    }
  }

  return copiedNode;
};

export const addNode = (
  nodes: Record<string, any>,
  newNode: any,
  fieldId: string = "id",
  parentIdField: string = "parent_id"
) => {
  if (typeof nodes === "undefined") return;
  const key = -1;
  if (nodes[key]) {
    return;
  }
  populateNode(newNode, parentIdField);
  newNode[fieldId] = key;
  const parentKey = newNode[parentIdField];
  if (parentKey && nodes[parentKey]) {
    let parent = nodes[parentKey];
    parent.children.push(newNode);
    parent.all_parent_ids.forEach((id: any) => {
      nodes[id].all_children_ids.push(newNode.id);
    });
    newNode.all_parent_ids = [parentKey, ...parent.all_parent_ids];
  }

  nodes[key] = newNode;
};

export const removeNode = (
  nodes: Record<string, any>,
  id: any,
  fieldId: string = "id",
  parentIdField: string = "parent_id"
) => {
  if (typeof nodes === "undefined") return;
  const node = nodes[id];
  if (!node) {
    return;
  }

  const parentKey = node[parentIdField];
  if (parentKey && nodes[parentKey]) {
    let parent = nodes[parentKey];
    parent.children = parent.children.filter((child: any) => child.id != id);
  }
  node.all_parent_ids.forEach((pId: any) => {
    nodes[pId].all_children_ids = nodes[pId].all_children_ids.filter(
      (i: any) => i != id
    );
  });

  delete nodes[id];
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
