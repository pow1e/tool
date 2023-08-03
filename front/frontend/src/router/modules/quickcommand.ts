import { lazy } from 'react';
import { ViewModuleIcon } from 'tdesign-icons-react';
import { IRouter } from '../index';

const result: IRouter[] = [

  {
    path: '/card',
    Component: lazy(() => import('pages/QuickCommand')),
    meta: {
      title: '快速指令',
      Icon: ViewModuleIcon,
      single: false
    },
  }
];
export default result;
