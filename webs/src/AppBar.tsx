import {AppBar, Box, Button, IconButton, Toolbar, Tooltip} from '@mui/material';
import {RefreshIconButton, SidebarToggleButton, TitlePortal, ToggleThemeButton, UserMenu} from 'react-admin';
import GithubIcon from '@mui/icons-material/GitHub';

const ContributeButton = () => (
    <Tooltip title="Github">
        <IconButton color="inherit" href="https://github.com/luannt2909/go-admin-example">
            <GithubIcon/>
        </IconButton>
    </Tooltip>
);

const LogoButton = () => (
    <Tooltip title="Luciango">
        <Button color="inherit" size='large'
                href="https://luciango.com"
                sx={{fontFamily: 'Cursive', fontSize: 'large', fontWeight: 'bold'}}>
            ~ Luciango ~
        </Button>
    </Tooltip>
);

export const CustomAppBar = () => (
    <AppBar color="primary">
        <Toolbar variant="dense">
            <>
                <SidebarToggleButton/>
                <TitlePortal/>
                <Box component="span" flex={1}/>
                <Box component="span" flex={1}>
                    <LogoButton/>
                </Box>
                <Box component="span" flex={1}/>
                <ContributeButton/>
                <ToggleThemeButton/>
                <RefreshIconButton/>
                <UserMenu/>
            </>
        </Toolbar>
    </AppBar>
);