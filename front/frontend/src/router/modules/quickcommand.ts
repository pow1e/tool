/*
 * @Description:
 * @version: 1.0.0
 * @Author: William
 * @Date: 2023-08-03 16:30:49
 * @LastEditors: William
 * @LastEditTime: 2023-08-08 10:45:24
 */
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
