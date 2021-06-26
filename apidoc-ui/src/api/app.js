import { sendRequest } from "@/utils/request";

export const url = {
  config: "/apidoc/config",
  apiData: "/apidoc/apiData",
  verifyAuth: "/apidoc/verifyAuth",
  createCrud: "/apidoc/createCrud",
  mdDetail: "/apidoc/mdDetail"
};

export const getConfig = param => {
  return sendRequest(url.config, param, "get");
};
export const getApiData = param => {
  return sendRequest(url.apiData, param, "get");
};
export const verifyAuth = param => {
  return sendRequest(url.verifyAuth, param, "post");
};
export const createCrud = param => {
  return sendRequest(url.createCrud, param, "post");
};
export const getMdDetail = param => {
  return sendRequest(url.mdDetail, param, "get");
};
