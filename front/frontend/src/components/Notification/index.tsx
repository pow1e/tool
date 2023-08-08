/*
 * @Description:
 * @version: 1.0.0
 * @Author: William
 * @Date: 2023-08-07 10:00:30
 * @LastEditors: William
 * @LastEditTime: 2023-08-07 14:13:56
 */
import { NotificationPlugin } from 'tdesign-react';

export function success(message: string) {
  const success = NotificationPlugin.success({
    title: '成功',
    content: message,
    duration: 3000,
    closeBtn: true,
    offset: [-20, 20],
  });
  return success;
};

export function error(message: string) {
  const error = NotificationPlugin.error({
    title: '失败',
    content: `执行失败：${message}`,
    duration: 3000,
    closeBtn: true,
    offset: [-20, 20],
  });
  return error;
};
