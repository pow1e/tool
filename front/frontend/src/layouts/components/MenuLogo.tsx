/*
 * @Description:
 * @version: 1.0.0
 * @Author: William
 * @Date: 2023-08-02 09:33:32
 * @LastEditors: William
 * @LastEditTime: 2023-08-02 12:46:19
 */
import React, { memo } from 'react';
import Style from './Menu.module.less';
// 保留 未来恢复
import FullLogo from 'assets/svg/assets-logo-full.svg?component';
import MiniLogo from 'assets/svg/assets-t-logo.svg?component';
import { useNavigate } from 'react-router-dom';

interface IProps {
  collapsed?: boolean;
}

export default memo((props: IProps) => {
  const navigate = useNavigate();

  const handleClick = () => {
    navigate('/');
  };

  return (
    <div className={Style.menuLogo} onClick={handleClick}>
      {/* {props.collapsed ? <h1 className={Style.logo}>L</h1> : <h1 className={Style.logo}>LOGO</h1>} */}
      {props.collapsed ? <MiniLogo /> : <FullLogo />}
    </div>
  );
});
