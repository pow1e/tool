/*
 * @Description:
 * @version: 1.0.0
 * @Author: William
 * @Date: 2023-08-08 10:47:01
 * @LastEditors: William
 * @LastEditTime: 2023-08-09 16:11:35
 */
import React from 'react';
import { Card, Input, Button, Space, Row, Col, RangeInput, Loading } from 'tdesign-react';
import InfoCards from './components/InfoCard';
import { success, error } from 'components/Notification';
import { getInfoList } from 'services/portscan';
import Style from './index.module.less';

const CardList = () => {
  const [loading, setLoading] = React.useState(false); // 添加 loading 状态
  const [ipaddr, setIpaddr] = React.useState<string>('');
  const [result, setResult] = React.useState<any[]>([]); // 修改为 useState 来存储结果数据
  const [timeout, setTimeout] = React.useState<string>('');
  const [ports, setPorts] = React.useState<string[]>(['0', '65535']);

  const portsChange = (e: any) => {
    setPorts(e);
  };
  const timeoutChange = (e: any) => {
    setTimeout(e);
  };
  const ipaddrChange = (e: any) => {
    setIpaddr(e);
  };
  const getInfo = () => {

    if (ipaddr === '') {
      error('IP地址不能为空');
      return;
    } else if (!/^((2(5[0-5]|[0-4]\d))|[0-1]?\d{1,2})(\.((2(5[0-5]|[0-4]\d))|[0-1]?\d{1,2})){3}$/.test(ipaddr)) {
      error('IP地址格式不正确');
      return;
    } else if (Number(ports[0]) > Number(ports[1])) {
      error('起始端口号不能大于结束端口号');
      return;
    } else if (Number(ports[1]) === 0) {
      error('结束端口号不能为空');
      return;
    } else if (Number(ports[0]) < 0 || Number(ports[1]) > 65535 || Number(ports[1]) < 0) {
      error('端口号范围为0-65535');
      return;
    }
    else if (timeout !== '' && !/^[0-9]*$/.test(timeout)) {
      error('超时时间必须为数字');
      return;
    }
    else {
      setLoading(true);
      const getInfoPromise = new Promise((resolve, reject) => {
        const data = {
          ip: ipaddr,
          timeout: Number(timeout),
          ports: ports,
        }
        getInfoList(data)
          .then((res) => {
            success('扫描成功');
            resolve(res);
          })
          .catch((err) => {
            error(err.msg.toString());
            reject(err);
          })
          .finally(() => {
            setLoading(false); // 请求结束后，设置 loading 状态为 false
          });
      });
      getInfoPromise
        .then((res: any) => {
          setResult(res.list);
        })
        .catch((err: any) => {
          error(err.msg.toString())
        }).finally(() => {
        });
    }
  };
  return (
    <div>
      <Space direction="vertical" style={{ width: "100%" }}>
        <Card bordered={false}>
          <Row gutter={20}>
            <Col flex={7}>
              <Input
                placeholder="请输入IP地址"
                style={{ width: "100%" }}
                align="center"
                value={ipaddr}
                onChange={ipaddrChange}
              />
            </Col>
            <Col flex={2} span={4}>
              <RangeInput
                placeholder={['请输入起始端口', '请输入结束端口']}
                defaultValue={['0', '65535']}
                value={ports}
                onChange={portsChange}
                inputProps={{ align: 'center' }}
              />
            </Col>
            <Col flex={2} span={4}>
              <Input
                placeholder="请输入超时时间 (可选)"
                suffix="s"
                align="center"
                value={timeout}
                onChange={timeoutChange}
              />
            </Col>
          </Row>
          <Button style={{ width: "100%", margin: '20px 0' }} onClick={getInfo}>
            扫 描
          </Button>
        </Card>

        {loading ? ( // 加载状态显示
          <div className={Style.loading}>
            <Loading text="端口扫描中，请勿离开该页面..." loading />
          </div>
        ) : (
          result.map((info: any, index: any) => (
            <InfoCards info={info} key={index} />
          )))}
      </Space>
    </div>
  );
};

export default React.memo(CardList);
