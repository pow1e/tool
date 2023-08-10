/*
 * @Description:
 * @version: 1.0.0
 * @Author: William
 * @Date: 2023-08-03 16:30:49
 * @LastEditors: William
 * @LastEditTime: 2023-08-10 08:47:18
 */
import { error, success } from 'components/Notification';
import { useState } from 'react';
import { getdecrypt, getencrypt } from 'services/crypto';
import { Radio, Space, Card, Input, Button, Form, Select } from 'tdesign-react';
const { FormItem } = Form;

export default function RadioExample() {
  const [move, selectMove] = useState('encrypt');
  const [encryptValue, setEncryptValue] = useState('');
  const [value, setValue] = useState('0');
  const handleChange = (value: any) => {
    setValue(value);
  };
  const [keyValue, setKeyValue] = useState('');
  const keyInputChange = (e: any) => {
    setKeyValue(e);
  };
  const addRules = {
    name: [
      {
        required: true,
        message: '请输入服务名称',
        type: 'error',
      },
    ],
    description: [
      {
        required: true,
        message: '请输入服务描述',
        type: 'error',
      },
    ],
    directives: [
      {
        required: true,
        message: '请输入启动命令',
        type: 'error',
      },
    ],
    type: [
      {
        required: true,
        message: '请选择权限',
        type: 'error',
      },
    ],
  };
  const encryptInputChange = (e: any) => {
    setEncryptValue(e);
  };
  const encryptData = {
    "plaintext": encryptValue,
    "key": keyValue
  }
  const [form] = Form.useForm();
  const onEncrypt = () => {
    if (encryptValue === '') {
      error('加密内容不能为空');
      return;
    }
    getencrypt(encryptData).result.then((res) => {
      const code: string = res.code.toString();
      const result: any = res.data;
      const msg: string = res.msg.toString();
      if (code === '1') {
        success(msg);
        form.setFieldsValue({
          '16lowercase': result.md5['16lowercase'],
          '16uppercase': result.md5['16uppercase'],
          '32lowercase': result.md5['32lowercase'],
          '32uppercase': result.md5['32uppercase'],
          'base64': result.base64,
          'aes': result.aes,
        })
      }
    }
    ).catch((err) => {
      const msg: string = err.msg.toString();
      error(msg);
    })
  };
  const [inputValue, setInputValue] = useState('');
  const [decryptValue, setDecryptValue] = useState('');

  const data = {
    ciphertext: inputValue,
  }
  const onDecrypt = (e: any) => {
      getdecrypt(data).result.then((res) => {
        const code = res.code.toString();
        const msg = res.msg.toString();
        if (code === '1') {
          setDecryptValue(res.data);
          success(msg);
        }
      }).catch((err) => {
        const msg = err.msg.toString();
        error(msg);
      })
  };
  const handleInputChange = (e: any) => {
    setInputValue(e);
  };
  const options =
    [
      {
        label: 'MD5',
        value: '0',
      },
      {
        label: 'Base64',
        value: '1',
      },
      {
        label: 'AES',
        value: '2',
      }
    ]

  return (
    <Card
      title={
        <Space direction="vertical">
          <Radio.Group variant="primary-filled" value={move} onChange={(value) => selectMove(value)}>
            <Radio.Button value="encrypt">加密</Radio.Button>
            <Radio.Button value="decrypt">解密</Radio.Button>
          </Radio.Group>
        </Space>
      }
      bordered={false}
      actions={move === 'decrypt' ? (
        <Select
          options={options}
          value={value}
          onChange={handleChange}
          autoWidth={true}
        >
        </Select>
      ) : null}
    >
      {move === 'encrypt' ? (
        <div>
          <Input
            placeholder="请输入需要加密的内容"
            align="center"
            value={encryptValue}
            onChange={encryptInputChange}
          />
          <Input
            placeholder="请输入自定义 Key (用于加密AES - 可选)"
            align="center"
            value={keyValue}
            onChange={keyInputChange}
            style={{ margin: '20px 0 0 0' }}
          />
          <Button
            style={{ width: "100%", margin: '20px 0 20px 0' }}
            onClick={onEncrypt}
          >加 密
          </Button>
          <Form
            labelAlign='left'
            form={form}
            rules={addRules}
          >
            <FormItem
              label="MD5 16位小写"
              name="16lowercase"
            >
              <Input placeholder="密文将在这里显示" readonly={true} align="center" />
            </FormItem>
            <FormItem
              label="MD5 16位大写"
              name="16uppercase"
            >
              <Input placeholder="密文将在这里显示" readonly={true} align="center" />
            </FormItem>
            <FormItem
              label="MD5 32位小写"
              name="32lowercase"
            >
              <Input placeholder="密文将在这里显示" readonly={true} align="center" />
            </FormItem>
            <FormItem
              label="MD5 32位大写"
              name="32uppercase"
            >
              <Input placeholder="密文将在这里显示" readonly={true} align="center" />
            </FormItem>
            <FormItem
              label="Base64"
              name="base64"
            >
              <Input placeholder="密文将在这里显示" readonly={true} align="center" />
            </FormItem>
            <FormItem
              label="AES"
              name="aes"
            >
              <Input placeholder="密文将在这里显示" readonly={true} align="center" />
            </FormItem>
          </Form>
        </div>
      ) : (
        <div>
          <Input
            placeholder="请输入内容"
            align="center"
            value={inputValue}
            onChange={handleInputChange}
          />
          <Button
            style={{ margin: '20px 0 20px 0', width: '100%' }}
            onClick={onDecrypt}
          >解 密
          </Button>
          <Input
            placeholder="原文将在这里显示"
            align="center"
            readonly={true}
            value={decryptValue}
          />
        </div>
      )}
    </Card>

  );
}
