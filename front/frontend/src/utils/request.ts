import axios from 'axios';

const env = import.meta.env.MODE || 'development';

const SUCCESS_CODE = 1;
const TIMEOUT = 5000;

declare module "axios" {
  interface AxiosResponse<T = any> {
    errorinfo: string;
    code: number;
    data: T;
  }
  export function create(config?: AxiosRequestConfig): AxiosInstance;
}


export const instance = axios.create({
  baseURL: '',
  timeout: TIMEOUT,
  withCredentials: true,
});

instance.interceptors.response.use(
  // eslint-disable-next-line consistent-return
  (response) => {
    if (response.status === 200) {
      const { data } = response;
      if (data.code === SUCCESS_CODE) {
        return data;
      }
      return Promise.reject(data);
    }
    return Promise.reject(response?.data);
  },
  (e) => Promise.reject(e),
);

export default instance;
