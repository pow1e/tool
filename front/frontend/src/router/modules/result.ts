/*
 * @Description:
 * @version: 1.0.0
 * @Author: William
 * @Date: 2023-08-02 09:33:33
 * @LastEditors: William
 * @LastEditTime: 2023-08-02 15:45:40
 */
import { lazy } from 'react';
import { CheckCircleIcon } from 'tdesign-icons-react';
import { IRouter } from '../index';

const result: IRouter[] = [
  {
    path: '/md5',
    meta: {
      title: 'MD5加密/解密',
      Icon: CheckCircleIcon,
    },
    Component: lazy(() => import('pages/Md5')),
  },
];

export default result;
