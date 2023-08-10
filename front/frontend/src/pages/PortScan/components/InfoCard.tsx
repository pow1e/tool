/*
 * @Description:
 * @version: 1.0.0
 * @Author: William
 * @Date: 2023-08-02 16:17:35
 * @LastEditors: William
 * @LastEditTime: 2023-08-09 11:31:22
 */
import React from 'react';
import { Card, Textarea } from 'tdesign-react';
// import { InfoCard } from 'services/portscan';
import Style from './InfoCard.module.less';

const InfoCards = ({ info }: { info: any }) => {
  const status = info?.Info.Status
  const statusCode = status.match(/\b(\d{3})\b/)[0];
  return (
    <div>
      <Card
        className={Style.panel}
        bordered={false}
        size={'small'}
        title={
          <a href={info?.Info.Url} className={Style.name}>{info?.Info.Url}</a>
        }
        actions={<div className={Style.actions}>{info?.port}</div>}
      >
        <div className={Style.description}>端口开放情况：{info?.isOpen ? '开放' : '关闭'}</div>
        <div className={Style.description}>状态码：{statusCode}</div>
        <div className={Style.description}>网站标题：{info?.Info.Title}</div>
        <Textarea
          readonly
          className={Style.detail}
          value={JSON.stringify(info?.Info,null,2)}
          autosize={true}
        />
      </Card>
    </div>
  );
};

export default React  .memo(InfoCards);
