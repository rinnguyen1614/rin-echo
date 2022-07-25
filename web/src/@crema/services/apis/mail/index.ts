import { AxiosRequestConfig } from "axios";
import { MailObj } from "../../../types/models/apps/Mail";
import connectionList from "../../db/apps/mail/connectionList";
import folderList from "../../db/apps/mail/folderList";
import labelList from "../../db/apps/mail/labelList";
import mailData from "../../db/apps/mail/mailList";
import mock from "../MockConfig";

let mailList = mailData;

mock.onGet("/api/mailApp/mail/").reply((config: AxiosRequestConfig) => {
  const { params } = config;
  const response = mailList.find((mail) => mail.id === parseInt(params.id));
  return [200, response];
});

mock.onGet("/api/mailApp/mails/list").reply(200, mailList);

mock.onGet("/api/mailApp/folders/list").reply(200, folderList);

mock.onGet("/api/mailApp/labels/list").reply(200, labelList);

mock.onGet("/api/mailApp/connection/list").reply(200, connectionList);

mock
  .onGet("/api/mailApp/folder/mail/List")
  .reply((config: AxiosRequestConfig) => {
    const { params } = config;
    let folderMailList: MailObj[] = [];
    if (params.type === "folder") {
      if (params.name === "starred") {
        folderMailList = mailList.filter((mail) => mail.isStarred);
      } else {
        const folderId = folderList.find(
          (folder) => folder.alias === params.name
        )!.id;
        folderMailList = mailList.filter(
          (mail) => mail.folderValue === folderId
        );
      }
    } else if (params.type === "label") {
      const labelType = labelList.find(
        (label) => label.alias === params.name
      )!.id;
      folderMailList = mailList.filter((mail) => mail.label.id === labelType);
    }
    const index = params.page * 15;
    const total = folderMailList.length;
    const list =
      folderMailList.length > 15
        ? folderMailList.slice(index, index + 15)
        : folderMailList;
    return [200, { list, total }];
  });

mock.onPut("/api/mailApp/update/mails").reply((request: AxiosRequestConfig) => {
  const mails = request.data;
  return [200, mails];
});

mock
  .onPut("/api/mailApp/update/starred")
  .reply((request: AxiosRequestConfig) => {
    const { mailIds, status } = JSON.parse(request.data);
    mailList = mailList.map((mail) => {
      if (mailIds.includes(mail.id)) {
        mail.isStarred = !!status;
        return mail;
      }
      return mail;
    });
    const updatedMails = mailList.filter((mail) => mailIds.includes(mail.id));
    return [200, updatedMails];
  });

mock
  .onPut("/api/mailApp/update/folder")
  .reply((request: AxiosRequestConfig) => {
    const { mailIds, type } = JSON.parse(request.data);
    mailList = mailList.map((mail) => {
      if (mailIds.includes(mail.id)) {
        mail.folderValue = type;
        return mail;
      }
      return mail;
    });
    return [200, mailIds];
  });

mock.onPut("/api/mailApp/update/label").reply((request: AxiosRequestConfig) => {
  const { mailIds, type } = JSON.parse(request.data);
  mailList = mailList.map((mail) => {
    if (mailIds.includes(mail.id)) {
      mail.label = type;
      return mail;
    }
    return mail;
  });
  const updatedMails = mailList.filter((mail) => mailIds.includes(mail.id));
  return [200, updatedMails];
});

mock.onPut("/api/mailApp/mail/").reply((request: AxiosRequestConfig) => {
  const { mail } = JSON.parse(request.data);
  mailList = mailList.map((item) => (item.id === mail.id ? mail : item));
  return [200, mail];
});

mock.onPut("/api/mailApp/update/read").reply((request: AxiosRequestConfig) => {
  const { mailIds, status } = JSON.parse(request.data);
  mailList = mailList.map((mail) => {
    if (mailIds.includes(mail.id)) {
      mail.isRead = status;
      return mail;
    }
    return mail;
  });
  const updatedMails = mailList.filter((mail) => mailIds.includes(mail.id));
  return [200, updatedMails];
});

mock.onPost("/api/mailApp/compose").reply((request: AxiosRequestConfig) => {
  const { mail } = JSON.parse(request.data);
  mailList = [mail, ...mailList];
  return [200, mail];
});
