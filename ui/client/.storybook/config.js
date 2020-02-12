import { configure } from '@storybook/react';
import { addDecorator } from '@storybook/react';
import { withKnobs } from '@storybook/addon-knobs';
import {muiTheme} from 'storybook-addon-material-ui';
import Theme from "../src/Themes/default";
import '../src/i18n/index';

addDecorator(withKnobs);
addDecorator(muiTheme(Theme('light')));
// addDecorator(storyFn => <div style={{ textAlign: 'center' }}>{storyFn()}</div>);

// automatically import all files ending in *.stories.js
configure(require.context('../src', true, /\.stories\.js$/), module);
