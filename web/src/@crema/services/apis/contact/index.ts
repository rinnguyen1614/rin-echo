import { AxiosRequestConfig } from "axios";
import { ContactObj, LabelObj } from "../../../types/models/apps/Contact";
import contactData from "../../db/apps/contact/contactList";
import folderList from "../../db/apps/contact/folderList";
import labelList from "../../db/apps/contact/labelList";
import mock from "../MockConfig";

let contactList = contactData;

mock.onGet("/api/contactApp/folders/list").reply(200, folderList);

mock.onGet("/api/contactApp/labels/list").reply(200, labelList);

mock
  .onGet("/api/contactApp/contact/List")
  .reply((config: AxiosRequestConfig) => {
    const { params } = config;
    let folderContactList: ContactObj[];
    if (params.type === "folder") {
      if (params.name === "starred") {
        folderContactList = contactList.filter((contact) => contact.isStarred);
      } else if (params.name === "frequent") {
        folderContactList = contactList.filter((contact) => contact.isFrequent);
      } else {
        folderContactList = contactList;
      }
    } else {
      const labelType = labelList.find(
        (label) => label.alias === params.name
      )!.id;
      folderContactList = contactList.filter(
        (contact) => contact.label === labelType
      );
    }
    const index = params.page * 15;
    const total = folderContactList.length;
    const list =
      folderContactList.length > 15
        ? folderContactList.slice(index, index + 15)
        : folderContactList;
    return [200, { list, total }];
  });

mock
  .onPut("/api/contactApp/update/starred")
  .reply((request: AxiosRequestConfig) => {
    const { contactIds, status } = JSON.parse(request.data);
    contactList = contactList.map((contact) => {
      if (contactIds.includes(contact.id)) {
        contact.isStarred = !!status;
        return contact;
      }
      return contact;
    });
    const updatedList = contactList.filter((contact) =>
      contactIds.includes(contact.id)
    );
    return [200, updatedList];
  });

mock
  .onPost("/api/contactApp/delete/contact")
  .reply((request: AxiosRequestConfig) => {
    const { contactIds, type, name, page } = JSON.parse(request.data);
    let folderContactList: ContactObj[];
    if (type === "folder") {
      if (name === "starred") {
        contactList = contactList.filter(
          (contact) => !contactIds.includes(contact.id)
        );
        folderContactList = contactList.filter((contact) => contact.isStarred);
      } else if (name === "frequent") {
        contactList = contactList.filter(
          (contact) => !contactIds.includes(contact.id)
        );
        folderContactList = contactList.filter((contact) => contact.isFrequent);
      } else {
        contactList = contactList.filter(
          (contact) => !contactIds.includes(contact.id)
        );
        folderContactList = contactList;
      }
    } else {
      const labelType = labelList.find(
        (label: LabelObj) => label.alias === name
      )!.id;
      contactList = contactList.filter(
        (contact) => !contactIds.includes(contact.id)
      );
      folderContactList = contactList.filter(
        (contact) => contact.label === labelType
      );
    }
    const index = page * 15;
    const total = folderContactList.length;
    const list =
      folderContactList.length > 15
        ? folderContactList.slice(index, index + 15)
        : folderContactList;
    return [200, { list, total }];
  });

mock
  .onPut("/api/contactApp/update/label")
  .reply((request: AxiosRequestConfig) => {
    const { contactIds, type } = JSON.parse(request.data);
    contactList = contactList.map((contact) => {
      if (contactIds.includes(contact.id)) {
        contact.label = type;
        return contact;
      }
      return contact;
    });
    const updatedContacts = contactList.filter((contact) =>
      contactIds.includes(contact.id)
    );
    return [200, updatedContacts];
  });

mock.onPut("/api/contactApp/contact/").reply((request: AxiosRequestConfig) => {
  const { contact } = JSON.parse(request.data);
  contactList = contactList.map((item) =>
    item.id === contact.id ? contact : item
  );
  return [200, contact];
});

mock.onPost("/api/contactApp/compose").reply((request: AxiosRequestConfig) => {
  const { contact } = JSON.parse(request.data);
  contactList = [contact, ...contactList];
  return [200, contact];
});

mock.onGet("/api/contactApp/contact/").reply((config: AxiosRequestConfig) => {
  const { params } = config;
  const response = contactList.find(
    (contact) => contact.id === parseInt(params.id)
  );
  return [200, response];
});
