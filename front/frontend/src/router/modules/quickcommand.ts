import { lazy } from 'react';
import {LayersIcon } from 'tdesign-icons-react';
import { IRouter } from '../index';

const result: IRouter[] = [

  {
    path: '/card',
    Component: lazy(() => import('pages/QuickCommand')),
    meta: {
      title: '快捷指令',
      Icon: LayersIcon,
      single: false
    },
  }
];
export default result;
