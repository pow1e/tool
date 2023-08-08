/*
 * @Description:
 * @version: 1.0.0
 * @Author: William
 * @Date: 2023-08-02 16:17:35
 * @LastEditors: William
 * @LastEditTime: 2023-08-04 16:18:33
 */
import React, { useState } from 'react';
import { Card, Dropdown, Button, MessagePlugin, Form, Dialog, Radio, Input, Popconfirm } from 'tdesign-react';
import { Icon, PlayCircleFilledIcon } from 'tdesign-icons-react';
import { CardInfo, changeCommand, delCommand, runCommand } from 'services/command';
import Style from './CommandCard.module.less';
import FormItem from 'tdesign-react/es/form/FormItem';

const CommandCard = ({ info, pageInit }: { info: CardInfo; pageInit: any }) => {
  const onStart = () => {
    runCommand(info?.id).result.then((res) => {
      const code = res.code.toString();
      const msg = res.msg.toString();
      if (code === '1') {
        MessagePlugin.success('启动成功');
      }
    }
    ).catch((err) => {
      const msg = err.msg.toString();
      MessagePlugin.error(`启动失败，返回信息：${msg}`);
    });
  }
  const [forms] = Form.useForm();
  const data = {
    id: info?.id,
    name: Form.useWatch('name', forms),
    description: Form.useWatch('description', forms),
    directives: Form.useWatch('directives', forms),
    type: Form.useWatch('type', forms)
  };
  const onSubmit = () => {
    changeCommand(data).result.then((res) => {
      const code = res.code.toString();
      const msg = res.msg.toString();
      if (code === '1') {
        MessagePlugin.success('修改成功');
      } else {
        MessagePlugin.error(msg);
      }
      setVisible(false);
      pageInit();
    }).catch((err) => {
      const msg = err.msg.toString();
      MessagePlugin.error(`启动失败，返回信息：${msg}`);
    });
  };
  const delDataSingle = {
    ids: [info?.id]
  };
  const onDel = () => {
    delCommand(delDataSingle).result.then((res) => {
      const code = res.code.toString();
      const msg = res.msg.toString();
      if (code === '1') {
        MessagePlugin.success('删除成功');
      } else {
        MessagePlugin.error(msg);
      }
      setDelVisible(false);
      pageInit();
    }).catch((err) => {
      const msg = err.msg.toString();
      MessagePlugin.error(`启动失败，返回信息：${msg}`);
    });
  }
  const [visible, setVisible] = React.useState(false);
  const [delVisible, setDelVisible] = React.useState(false);
  return (
    <div>
      <Card
        className={Style.panel}
        bordered={false}
        size={'small'}
        actions={
          <PlayCircleFilledIcon size="3em" onClick={onStart} />
        }
        title={
          <div className={Style.name}>{info?.name}</div>
        }
        footer={
          <div className={Style.footer}>
            <Dropdown
              trigger={'click'}
              options={[
                {
                  content: '管理',
                  value: 1,
                  onClick: () => {
                    setVisible(true);
                  }
                },
                {
                  content: '删除',
                  value: 2,
                  onClick: () => {
                    setDelVisible(true);
                  }
                },
              ]}
            >
              <Button theme='default' variant='text'>
                <Icon name='more' size='16' />
              </Button>
            </Dropdown>
          </div>
        }
      >
        <div className={Style.description}>{info?.description}</div>
      </Card>
      <Dialog
        header='管理快捷指令'
        closeBtn
        closeOnEscKeydown
        closeOnOverlayClick
        cancelBtn="取 消"
        confirmBtn="提 交"
        mode='modal'
        onClose={() => {
          setVisible(false);
        }}
        placement='center'
        preventScrollThrough
        showOverlay
        theme='default'
        visible={visible}
        onConfirm={onSubmit}
        destroyOnClose={true}
      >
        <Form labelAlign="top" form={forms}>
          <FormItem label="服务名称" name="name" initialData={info?.name}>
            <Input />
          </FormItem>
          <FormItem label="服务描述" name="description" initialData={info?.description}>
            <Input />
          </FormItem>
          <FormItem label="启动命令" name="directives" initialData={info?.directives}>
            <Input />
          </FormItem>
          <FormItem label="权限" name="type" initialData={info?.type}>
            <Radio.Group>
              <Radio value="0">普通用户</Radio>
              <Radio value="1">管理员</Radio>
            </Radio.Group>
          </FormItem>
        </Form>
      </Dialog>
      <Dialog
        header='删除快捷指令'
        closeBtn
        closeOnEscKeydown
        closeOnOverlayClick
        cancelBtn="取 消"
        confirmBtn="确 定"
        mode='modal'
        onClose={() => {
          setDelVisible(false);
        }}
        placement='center'
        preventScrollThrough
        showOverlay
        theme='warning'
        visible={delVisible}
        onConfirm={onDel}
        destroyOnClose={true}
      >
        <div>确定删除该快捷指令吗？</div>
      </Dialog>
    </div>
  );
};

export default React.memo(CommandCard);
