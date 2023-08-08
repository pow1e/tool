/*
 * @Description:
 * @version: 1.0.0
 * @Author: William
 * @Date: 2023-08-03 16:30:49
 * @LastEditors: William
 * @LastEditTime: 2023-08-04 17:24:29
 */
import React, { useState } from 'react';
import { getdecrypt } from 'services/md5';
import { Radio, Space, Card, Input, Button, MessagePlugin } from 'tdesign-react';

export default function RadioExample() {
  const [move, selectMove] = useState('encrypt');
  const onEncrypt = () => {
    MessagePlugin.success('加密成功');
  };
  const [inputValue, setInputValue] = useState('');
  const [decryptValue, setDecryptValue] = useState('');
  const data = {
    ciphertext: inputValue,
  }
  const onDecrypt = () => {
    getdecrypt(data).result.then((res) => {
      const code = res.code.toString();
      const msg = res.msg.toString();
      if (code === '1') {
        setDecryptValue(res.data);
        MessagePlugin.success('解密成功');
      }
    }).catch((err) => {
      const msg = err.msg.toString();
      MessagePlugin.error(`解密失败，返回信息：${msg}`);
    })
  };
  const handleInputChange = (e: any) => {
    setInputValue(e);
  };
  return (
    <Card
      title={
        <Space direction="vertical">
          <Radio.Group variant="primary-filled" size="large" value={move} onChange={(value) => selectMove(value)}>
            <Radio.Button value="encrypt">加密</Radio.Button>
            <Radio.Button value="descrypt">解密</Radio.Button>
          </Radio.Group>
        </Space>
      }
      bordered={false}
    >
      {move === 'encrypt' ? (
        <div>
          <Input
            placeholder="请输入内容"
            align="center"
            size="large"
          />
          <Button
            size="large"
            style={{ float: 'right', margin: '20px 0 0 0' }}
          >加 密
          </Button>
        </div>
      ) : (
        <div>
          <Input
            placeholder="请输入内容"
            align="center"
            size="large"
            value={inputValue}
            onChange={handleInputChange}
          />
          <Button
            size="large"
            style={{ margin: '20px 0 20px 0', width: '100%' }}
            onClick={onDecrypt}
          >解 密
          </Button>
          <Input
            placeholder="原文将在这里显示"
            align="center"
            size="large"
            readonly={true}
            value={decryptValue}
          />
        </div>
      )}
    </Card>
  );
}
