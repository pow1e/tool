import request from 'utils/request';

// export interface InfoCard {
//   url: string,
//   port: number,
//   info: [],
//   code: number,
//   title: string
// }

export interface InfoCard {
  port: number,
  isOpen: boolean,
  Info: any
}
interface IResult {
  list: InfoCard[];
}


export const getInfoList = async (data: any) => {
  const result = await request.post<IResult>('api/v1/scan/port', data)
  let list = result?.data || []
  return {
    list
  }
};
