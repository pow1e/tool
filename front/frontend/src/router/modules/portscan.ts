/*
 * @Description:
 * @version: 1.0.0
 * @Author: William
 * @Date: 2023-08-08 10:38:43
 * @LastEditors: William
 * @LastEditTime: 2023-08-08 10:49:57
 */
import { lazy } from 'react';
import { WifiIcon } from 'tdesign-icons-react';
import { IRouter } from '../index';

const result: IRouter[] = [

  {
    path: '/portscan',
    Component: lazy(() => import('pages/PortScan')),
    meta: {
      title: '端口扫描',
      Icon: WifiIcon,
      single: false
    },
  }
];
export default result;
