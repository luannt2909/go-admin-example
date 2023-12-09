import {Admin, Resource,} from "react-admin";
import {dataProvider} from "./dataProvider";
import authProvider from "./authProvider";

import Users from "./components/user"
import PersonIcon from "@material-ui/icons/Person";
import {CustomLayout} from "./Layout";
import {Login} from "./LoginPage";

export const App = () => (
    <Admin
        layout={CustomLayout}
        loginPage={Login}
        dataProvider={dataProvider}
        authProvider={authProvider}
        darkTheme={{palette: {mode: 'dark'}}}
    >
        <Resource
            name="users"
            icon={PersonIcon}
            {...Users}
        />
    </Admin>
);
