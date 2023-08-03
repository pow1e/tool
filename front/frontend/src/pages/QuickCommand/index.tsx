/*
 * @Description:
 * @version: 1.0.0
 * @Author: William
 * @Date: 2023-08-02 16:17:35
 * @LastEditors: William
 * @LastEditTime: 2023-08-03 14:21:24
 */
import React, { useEffect } from 'react';
import { Row, Col, Button, Input, Loading, MessagePlugin, Form, Radio } from 'tdesign-react';
// import { SearchIcon } from 'tdesign-icons-react';
import { useAppDispatch, useAppSelector } from 'modules/store';
import { selectListCard, getList } from 'modules/list/card';
import CommandCard from './components/CommandCard';
import Style from './index.module.less';
import Dialog from 'tdesign-react/es/dialog';
import FormItem from 'tdesign-react/es/form/FormItem';

const CardList = () => {
  const dispatch = useAppDispatch();
  const pageState = useAppSelector(selectListCard);

  const pageInit = async () => {
    await dispatch(
      getList(),
    );
  };

  useEffect(() => {
    pageInit();
  }, []);

  const [visible, setVisible] = React.useState(false);
  const handleClick = () => {
    setVisible(true);
  };
  const onSubmit = () => {
    setVisible(false);
    MessagePlugin.success('提交成功');
  };
  return (
    <div>
      <div className={Style.toolBar}>
        <Button
          onClick={() => {
            handleClick();
          }}
        >
          新建快捷指令
        </Button>
        {/* <Input className={Style.search} suffixIcon={<SearchIcon />} placeholder='请输入你需要搜索的内容' /> */}
      </div>
      {pageState.pageLoading ? (
        <div className={Style.loading}>
          <Loading text='加载数据中...' loading size='large' />
        </div>
      ) : (
        <>
          <div className={Style.cardList}>
            <Row
              gutter={[16, 24]}
              style={{ margin: '0 0' }}
            >
                {pageState.productList.map((info, index) => (
                    <Col key={index} span={6} lg={4}>
                        <CommandCard info={info} />
                    </Col>
                ))}
            </Row>
          </div>
          <Dialog
            header='新建快捷指令'
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
          >
            <Form
              labelAlign="top"
            >
              <FormItem label="服务名称" name="name">
                <Input />
              </FormItem>
              <FormItem label="服务描述" name="description">
                <Input />
              </FormItem>
              <FormItem label="启动命令" name="directives">
                <Input />
              </FormItem>
              <FormItem label="权限" name="type">
                <Radio.Group defaultValue="0">
                  <Radio
                    value="0"
                  >
                    管理员
                  </Radio>
                  <Radio
                    value="1"
                  >
                    普通用户
                  </Radio>
                </Radio.Group>

              </FormItem>
            </Form>
          </Dialog>
        </>
      )}
    </div>
  );
};

export default React.memo(CardList);
