import { configure } from '@storybook/react';
import { addDecorator } from '@storybook/react';
import { withKnobs } from '@storybook/addon-knobs';
import {muiTheme} from 'storybook-addon-material-ui';
import Theme from "../src/Themes/default";
import '../src/i18n/index';

addDecorator(withKnobs);
// addDecorator(muiTheme(Theme('light')));
configure(require.context('../src', true, /\.stories\.js$/), module);
