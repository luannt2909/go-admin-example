import {fetchUtils} from "react-admin"
import simpleRestProvider from "ra-data-simple-rest";
import {getToken} from "./authProvider"
import {API_HOST} from "./config";

const httpClient = (url, options = {}) => {
    const token = getToken();
    const user = { token: `Bearer ${token}`, authenticated: !!token };

    return fetchUtils.fetchJson(url, {...options, user});
};
export const dataProvider = simpleRestProvider(API_HOST, httpClient);
