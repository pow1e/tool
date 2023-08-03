/*
 * @Description:
 * @version: 1.0.0
 * @Author: William
 * @Date: 2023-08-02 16:17:35
 * @LastEditors: William
 * @LastEditTime: 2023-08-03 11:09:40
 */
import request from 'utils/request';

export interface CardInfo {
  description: string;
  index: number;
  // isSetup: boolean;
  name: string;
  type: number;
  directive: string;
}

interface IResult {
  list: CardInfo[];
}


export const getProductList = async () => {
  const result = await request.get<IResult>('api/v1/shortcuts');
  let list = result?.data || [];
  return {
    list,
  };
};
