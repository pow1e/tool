/*
 * @Description:
 * @version: 1.0.0
 * @Author: William
 * @Date: 2023-08-03 16:30:49
 * @LastEditors: William
 * @LastEditTime: 2023-08-04 09:50:39
 */

import React from 'react';
import ReactDOM from 'react-dom/client';
import { Provider } from 'react-redux';
import { HashRouter } from 'react-router-dom';
import store from 'modules/store';
import App from 'layouts/index';

import 'tdesign-react/es/style/index.css';

import './styles/index.less';

const env = import.meta.env.MODE || 'development';
const baseRouterName = env === 'site' ? '/starter/react/' : '';

// eslint-disable-next-line @typescript-eslint/no-non-null-assertion
const root = document.getElementById('app')!;

const renderApp = () => {
  ReactDOM.createRoot(root).render(
    <Provider store={store}>
      <HashRouter basename={baseRouterName}>
        <App />
      </HashRouter>
    </Provider>,
  );
};

renderApp();
