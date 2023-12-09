import {AppBar, IconButton, Toolbar, Tooltip} from '@mui/material';
import {RefreshIconButton, SidebarToggleButton, TitlePortal, ToggleThemeButton, UserMenu} from 'react-admin';
import GithubIcon from '@mui/icons-material/GitHub';

const ContributeButton = () => (
    <Tooltip title="Github">
        <IconButton color="inherit" href="https://github.com/luannt2909/go-reminder-bot">
            <GithubIcon/>
        </IconButton>
    </Tooltip>
);
export const CustomAppBar = () => (
    <AppBar color="primary">
        <Toolbar variant="dense">
            <>
                <SidebarToggleButton/>
                <TitlePortal/>
                <ContributeButton/>
                <ToggleThemeButton/>
                <RefreshIconButton/>
                <UserMenu/>
            </>
        </Toolbar>
    </AppBar>
);