/*
 * @Description:
 * @version: 1.0.0
 * @Author: William
 * @Date: 2023-08-03 16:30:49
 * @LastEditors: William
 * @LastEditTime: 2023-08-04 14:53:09
 */
import React, { useState } from 'react';
import { Radio, Space, Card, Input,Button } from 'tdesign-react';

export default function RadioExample() {
  const [move, selectMove] = useState('encrypt');
  return (
    <Card
      title={
        <Space direction="vertical">
          <Radio.Group variant="default-filled" size="large" value={move} onChange={(value) => selectMove(value)}>
            <Radio.Button value="encrypt">加密</Radio.Button>
            <Radio.Button value="descrypt">解密</Radio.Button>
          </Radio.Group>
        </Space>
      }
      bordered={false}
    >
      {/* <Space direction="horizontal" style={{ width: '100%'}} align="center" size={14}> */}
        <Input
          placeholder="请输入内容"
          align="center"
          size="large"
        />
        <Button
          size="large"
        style={{ float: 'right',margin: '20px 0 0 0'}}
        >加 密
        </Button>
      {/* </Space> */}
    </Card>
  );
}
