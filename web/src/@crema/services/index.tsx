import mock from "./apis/MockConfig";
import "./apis";

mock.onAny().passThrough();
