/*
 * @Description:
 * @version: 1.0.0
 * @Author: William
 * @Date: 2023-08-02 09:33:32
 * @LastEditors: William
 * @LastEditTime: 2023-08-02 11:33:12
 */
import React, { memo } from 'react';
import { Button, Popup, Space } from 'tdesign-react';
import {
  SettingIcon,
} from 'tdesign-icons-react';
import { useAppDispatch } from 'modules/store';
import { toggleSetting } from 'modules/global';
import Style from './HeaderIcon.module.less';

export default memo(() => {
  const dispatch = useAppDispatch();

  return (
    <Space align='center'>
      <Popup content='页面设置' placement='bottom' showArrow destroyOnClose>
        <Button
          className={Style.menuIcon}
          shape='square'
          size='large'
          variant='text'
          onClick={() => dispatch(toggleSetting())}
          icon={<SettingIcon />}
        />
      </Popup>
    </Space>
  );
});
