// in src/MyLayout.js
import { Layout } from 'react-admin';

import { CustomAppBar } from './AppBar';

export const CustomLayout = (props) => (
    <Layout {...props} appBar={CustomAppBar} />
);