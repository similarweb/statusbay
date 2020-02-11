import React from 'react';
import ReactDOM from 'react-dom';
import * as io from 'socket.io-client';
import App from './App';
import { AppSettingsProvider } from './context/AppSettingsContext';
import { SocketIOProvider } from './context/SocketIOContext';
import './i18n/index';

ReactDOM.render(<SocketIOProvider io={io}><AppSettingsProvider><App /></AppSettingsProvider></SocketIOProvider>, document.getElementById('root'));
