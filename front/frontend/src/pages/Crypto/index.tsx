/*
 * @Description:
 * @version: 1.0.0
 * @Author: William
 * @Date: 2023-08-03 16:30:49
 * @LastEditors: William
 * @LastEditTime: 2023-08-07 15:12:22
 */
import { error, success } from 'components/Notification';
import { useState } from 'react';
import { getdecrypt, getencrypt } from 'services/md5';
import { Radio, Space, Card, Input, Button, Form, Select } from 'tdesign-react';
const { FormItem } = Form;

export default function RadioExample() {
  const [move, selectMove] = useState('encrypt');
  const [encryptValue, setEncryptValue] = useState('');
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
    "plaintext": encryptValue
  }
  const [form] = Form.useForm();
  const onEncrypt = () => {
    getencrypt(encryptData).result.then((res) => {
      const code: string = res.code.toString();
      const result: any = res.data;
      const msg: string = res.msg.toString();
      if (code === '1') {
        success(msg);
        form.setFieldsValue({
          '16lowercase': result['16lowercase'],
          '16uppercase': result['16uppercase'],
          '32lowercase': result['32lowercase'],
          '32uppercase': result['32uppercase']
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
    if (e.vali)
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
  const [value, setValue] = useState('0');
  const handleChange = (value: any) => {
    setValue(value);
  };
  return (
    <Card
      title={
        <Space direction="vertical">
          <Radio.Group variant="primary-filled" size="large" value={move} onChange={(value) => selectMove(value)}>
            <Radio.Button value="encrypt">加密</Radio.Button>
            <Radio.Button value="decrypt">解密</Radio.Button>
          </Radio.Group>
        </Space>
      }
      bordered={false}
      actions={move === 'decrypt' ? (
        <Select
          size="large"
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
            placeholder="请输入内容"
            align="center"
            size="large"
            value={encryptValue}
            onChange={encryptInputChange}
          />
          <Button
            size="large"
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
              label="16位小写"
              name="16lowercase"
            >
              <Input size="large" placeholder="密文将在这里显示" readonly={true} align="center" />
            </FormItem>
            <FormItem
              label="16位大写"
              name="16uppercase"
            >
              <Input size="large" placeholder="密文将在这里显示" readonly={true} align="center" />
            </FormItem>
            <FormItem
              label="32位小写"
              name="32lowercase"
            >
              <Input size="large" placeholder="密文将在这里显示" readonly={true} align="center" />
            </FormItem>
            <FormItem
              label="32位大写"
              name="32uppercase"
            >
              <Input size="large" placeholder="密文将在这里显示" readonly={true} align="center" />
            </FormItem>
          </Form>
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
