/*
 * @Description:
 * @version: 1.0.0
 * @Author: William
 * @Date: 2023-08-08 16:54:57
 * @LastEditors: William
 * @LastEditTime: 2023-08-08 16:59:54
 */
import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit';
import { RootState } from '../store';
import { InfoCard, getInfoList } from 'services/portscan';

const namespace = 'portscan';

interface IInitialState {
  pageLoading: boolean;
  loading: boolean;
  InfoList: InfoCard[];
}

const initialState: IInitialState = {
  pageLoading: true,
  loading: true,
  InfoList: [],
};

export const getList = createAsyncThunk(
  `/${namespace}`,
  async (data: any) => {
    const result = await getInfoList(data);
    return {
      list: result?.list
    };
  },
);


const listCardSlice = createSlice({
  name: namespace,
  initialState,
  reducers: {
    switchPageLoading: (state, action: PayloadAction<boolean>) => {
      state.pageLoading = action.payload;
    },
  },
  extraReducers: (builder) => {
    builder
      .addCase(getList.pending, (state) => {
        state.loading = true;
      })
      .addCase(getList.fulfilled, (state, action) => {
        state.loading = false;
        state.pageLoading = false;
        state.InfoList = action.payload?.list || [];
      })
      .addCase(getList.rejected, (state) => {
        state.loading = false;
      });
  },
});

export const { switchPageLoading } = listCardSlice.actions;

export const selectInfoCard = (state: RootState) => state.card;

export default listCardSlice.reducer;
