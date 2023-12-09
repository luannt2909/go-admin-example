import {fetchUtils} from "react-admin"
import simpleRestProvider from "ra-data-simple-rest";
import {retrieveToken} from "./authProvider"

const httpClient = (url, options = {}) => {
    const token = retrieveToken();
    const user = { token: `Bearer ${token}`, authenticated: !!token };

    return fetchUtils.fetchJson(url, {...options, user});
};
export const dataProvider = simpleRestProvider(import.meta.env.VITE_SIMPLE_REST_URL, httpClient);
// export const dataProvider = simpleRestProvider(
//   import.meta.env.VITE_SIMPLE_REST_URL
// );
