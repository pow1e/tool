/*
 * @Description:
 * @version: 1.0.0
 * @Author: William
 * @Date: 2023-08-02 09:33:33
 * @LastEditors: William
 * @LastEditTime: 2023-08-08 10:46:36
 */
import { lazy } from 'react';
import { LockOnIcon } from 'tdesign-icons-react';
import { IRouter } from '../index';

const result: IRouter[] = [
  {
    path: '/crypto',
    meta: {
      title: '加密/解密',
      Icon: LockOnIcon,
    },
    Component: lazy(() => import('pages/Crypto')),
  },
];

export default result;
