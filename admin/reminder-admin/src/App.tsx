import {Admin, Resource,} from "react-admin";
import {dataProvider} from "./dataProvider";
import authProvider from "./authProvider";

import Reminders from "./components/reminder/index"
import Users from "./components/user/index"
import {Login} from "./LoginPage";
import ReminderIcon from "@material-ui/icons/NotificationImportant";
import PersonIcon from "@material-ui/icons/Person";
import {RoleAdmin} from "./components/user/role";
import {CustomLayout} from "./Layout";

const randomBackground = () => {
    const timeStamp = Date.now()
    return timeStamp % 2 == 0 ? "/admin/lightfall.jpeg" : "/admin/deep_blue.jpeg"
}

const CustomLoginPage = () => <Login backgroundImage={randomBackground()}/>;

export const App = () => (
    <Admin layout={CustomLayout}
           loginPage={CustomLoginPage}
           dataProvider={dataProvider}
           authProvider={authProvider}
           darkTheme={{palette: {mode: 'dark'}}}
    >
        {permissions => (
            <>
                <Resource
                    name="reminders"
                    icon={ReminderIcon}
                    {...Reminders}
                />
                {permissions == RoleAdmin ? <Resource
                    name="users"
                    icon={PersonIcon}
                    {...Users}
                /> : null}

            </>
        )}

    </Admin>
);
