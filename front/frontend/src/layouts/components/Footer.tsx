/*
 * @Description:
 * @version: 1.0.0
 * @Author: William
 * @Date: 2023-08-03 16:30:49
 * @LastEditors: William
 * @LastEditTime: 2023-08-04 17:28:38
 */
import React from 'react';
import { Layout, Row } from 'tdesign-react';
import Style from './Footer.module.less';
const { Footer: TFooter } = Layout;

const Footer = () => {
  return (
    <TFooter>
      {/* <Row justify='center'>Copyright © 2022-{new Date().getFullYear()} Tencent. All Rights Reserved</Row> */}
      <Row justify='center' className={Style.footer}>Further Development By wuqianer & willi404, Copyright © 2022-{new Date().getFullYear()} Tencent. All Rights Reserved</Row>
    </TFooter>
  );
};

export default React.memo(Footer);
