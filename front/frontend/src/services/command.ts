/*
 * @Description:
 * @version: 1.0.0
 * @Author: William
 * @Date: 2023-08-02 16:17:35
 * @LastEditors: William
 * @LastEditTime: 2023-08-04 16:12:48
 */
import request from 'utils/request';

export interface CardInfo {
  id: string;
  description: string;
  index: number;
  // isSetup: boolean;
  name: string;
  type: number;
  directives: string;
}

interface IResult {
  list: CardInfo[];
}


export const getCommandList = async () => {
  const result = await request.get<IResult>('api/v1/shortcuts');
  let list = result?.data || [];
  return {
    list,
  };
};

export const addCommand = (form: any) => {
  const result = request({
    url: 'api/v1/shortcuts',
    method: 'post',
    data: form
  })
  return {
    result
  }
}

export const changeCommand = (form: any) => {
  const result = request({
    url: 'api/v1/shortcuts',
    method: 'put',
    data: form
  })
  return {
    result
  }
}

export const delCommand = (form: any) => {
  const result = request({
    url: 'api/v1/shortcuts/',
    method: 'delete',
    data: form
  })
  return {
    result
  }
}

export const runCommand = (id: string) => {
  const result = request({
    url: 'api/v1/shortcuts/run',
    method: 'get',
    params: {
      id
    }
  })
  return {
    result
  }
}

