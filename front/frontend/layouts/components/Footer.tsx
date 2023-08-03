import React from 'react';
import { Layout, Row } from 'tdesign-react';
import Style from './Footer.module.less';
const { Footer: TFooter } = Layout;

const Footer = () => {
  return (
    <TFooter>
      {/* <Row justify='center'>Copyright © 2022-{new Date().getFullYear()} Tencent. All Rights Reserved</Row> */}
      <Row justify='center' className={Style.footer}>Further Development By William, Copyright © 2022-{new Date().getFullYear()} Tencent. All Rights Reserved</Row>
    </TFooter>
  );
};

export default React.memo(Footer);
