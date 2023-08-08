/*
 * @Description:
 * @version: 1.0.0
 * @Author: William
 * @Date: 2023-08-08 10:47:01
 * @LastEditors: William
 * @LastEditTime: 2023-08-08 11:27:55
 */
import React from 'react';
import { Card, Input, Textarea, Button } from 'tdesign-react';

const cardList = () => {

  return (
    <div>
      <Card
        bordered={false}
      >
        <Input placeholder="请输入IP地址" size="large" style={{ width: "100%" }} align="center" />
        <Button
          size="large"
          style={{ width: "100%", margin: '20px 0 20px 0' }}
        >
          扫 描
        </Button>
        <Textarea
          readonly
          style={{ width: "100%" }}
          autosize={{ minRows: 15, maxRows: 20 }}
        />
      </Card>
    </div>
  )
}

export default React.memo(cardList);
