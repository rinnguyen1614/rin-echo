import { AxiosRequestConfig } from "axios";
import { ConnectionObj } from "../../../types/models/apps/Chat";
import chatList from "../../db/apps/chat/chatList";
import connectionList from "../../db/apps/chat/connectionList";
import mock from "../MockConfig";

let connectionData = connectionList;
let chatData = chatList;

mock.onGet("/api/chatApp/connections").reply(200, connectionData);

mock
  .onGet("/api/chatApp/connection/messages")
  .reply((request: AxiosRequestConfig) => {
    const { id } = request.params;
    const response = chatData.find((chat) => chat.channelId === parseInt(id));
    if (response) {
      return [200, response];
    }
    return [200, null];
  });

mock.onPost("/api/chatApp/message").reply((request: AxiosRequestConfig) => {
  const { channelId, message } = JSON.parse(request.data);
  const id = (Math.random() * 10000).toFixed();
  const data = { ...message, id };
  let user: ConnectionObj = connectionData.find(
    (connection) => connection.channelId === channelId
  )!;
  user = { ...user, lastMessage: data };
  connectionData = connectionData.map((item: ConnectionObj) =>
    item.channelId === user!.channelId ? user : item
  );
  let userMessages = chatData.find((chat) => chat.channelId === channelId);
  if (userMessages) {
    userMessages.messageData = userMessages.messageData.concat(data);
  } else {
    userMessages = {
      channelId,
      messageData: [data],
    };
    chatData = chatData.concat(userMessages);
  }
  return [200, { user, userMessages }];
});

mock.onPut("/api/chatApp/message").reply((request: AxiosRequestConfig) => {
  const { channelId, message } = JSON.parse(request.data);
  let user = connectionData.find(
    (connection) => connection.channelId === channelId
  )!;
  if (user!.lastMessage!.id === message.id) {
    user = { ...user, lastMessage: message };
    connectionData = connectionData.map((item) =>
      item.channelId === user.channelId ? user : item
    );
  }
  const userMessages = chatData.find((chat) => chat.channelId === channelId);
  if (userMessages)
    userMessages.messageData = userMessages.messageData.map((item) =>
      item.id === message.id ? message : item
    );
  return [200, { user, userMessages }];
});

mock
  .onPost("/api/chatApp/delete/message")
  .reply((request: AxiosRequestConfig) => {
    const { channelId, messageId } = JSON.parse(request.data);
    const userMessages = chatData.find((chat) => chat.channelId === channelId)!;
    let user = connectionData.find(
      (connection) => connection.channelId === channelId
    )!;
    if (userMessages) {
      userMessages.messageData = userMessages.messageData.filter(
        (item) => item.id !== messageId
      );
      if (user!.lastMessage!.id === messageId) {
        const lastMessage =
          userMessages.messageData[userMessages.messageData.length - 1];
        user = {
          ...user,
          lastMessage: {
            id: lastMessage.id!,
            message: lastMessage.message!,
            type: "",
            time: lastMessage.time!,
          },
        };
        connectionData = connectionData.map((item) =>
          item.id === user!.id ? user : item
        )!;
      }
    }
    return [200, { user, userMessages }];
  });

mock
  .onPost("/api/chatApp/delete/user/messages")
  .reply((request: AxiosRequestConfig) => {
    const { channelId } = JSON.parse(request.data);
    chatData = chatData.filter((chat) => chat.channelId !== channelId);
    const user = connectionData.find(
      (connection) => connection.channelId === channelId
    )!;
    // let user.lastMessage;
    connectionData = connectionData.map((item) =>
      item.id === user.id ? user : item
    );
    return [200, user];
  });
