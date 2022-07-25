import jwtAxios from "../auth/jwt-auth";

const MockAdapter = require("axios-mock-adapter");

export default new MockAdapter(jwtAxios, { delayResponse: 100 });
