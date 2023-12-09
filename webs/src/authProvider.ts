import {AuthProvider} from "react-admin";
import {API_HOST} from "./config";

export const authProvider: AuthProvider = {
    login: ({username, password}) => {
        const request = new Request(`${API_HOST}/auth/authenticate`, {
            method: 'POST',
            body: JSON.stringify({username, password}),
            headers: new Headers({'Content-Type': 'application/json'}),
        });
        return fetch(request)
            .then(response => {
                if (response.status < 200 || response.status >= 300) {
                    throw new Error(response.statusText);
                }
                return response.json();
            })
            .then(data => {
                localStorage.setItem('user', JSON.stringify(data.user));
                localStorage.setItem('token', data.token);
            })
            .catch(() => {
                throw new Error('Network Error!')
            });
    },
    logout: () => {
        localStorage.removeItem("user");
        localStorage.removeItem("token");
        return Promise.resolve();
    },
    checkError: (error) => {
        const status = error.status;
        if (status === 401 || status === 403) {
            localStorage.removeItem('user');
            localStorage.removeItem('token');
            return Promise.reject();
        }
    },
    checkAuth: () =>
        localStorage.getItem("user") ? Promise.resolve() : Promise.reject(),
    getPermissions: () => {
        const user = JSON.parse(localStorage.getItem("user"))
        return user ? Promise.resolve(user.role) : Promise.resolve()
    },
    getIdentity: () => {
        const persistedUser = localStorage.getItem("user");
        const user = persistedUser ? JSON.parse(persistedUser) : {};
        return Promise.resolve({
            id: user.id,
            fullName: user.username,
        })
    },
};

export function getToken() {
    return localStorage.getItem("token");
}

export default authProvider;
