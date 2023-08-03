/*
 * @Description:
 * @version: 1.0.0
 * @Author: William
 * @Date: 2023-08-02 16:17:35
 * @LastEditors: William
 * @LastEditTime: 2023-08-03 14:20:54
 */
import React from 'react';
import { Card, Dropdown, Button, MessagePlugin } from 'tdesign-react';
import { Icon, PlayCircleFilledIcon } from 'tdesign-icons-react';
import { CardInfo } from 'services/command';
import Style from './CommandCard.module.less';

const CommandCard = ({ info }: { info: CardInfo }) => {
  const onStart = () => {

   MessagePlugin.success("执行成功")
  }
  return (
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
              },
              {
                content: '删除',
                value: 2,
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
  );
};

export default React.memo(CommandCard);
