import mock from "../MockConfig";
import userList from "../../db/userList";

mock.onGet("/api/user/list").reply(200, userList);
