/*
 * @Description:
 * @version: 1.0.0
 * @Author: William
 * @Date: 2023-08-02 16:17:35
 * @LastEditors: William
import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit';
import { RootState } from '../store';
import { getProductList, CardInfo } from 'services/command';
import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit';
import { RootState } from '../store';
import { getCommandList, CardInfo } from 'services/command';

const namespace = 'card';

interface IInitialState {
  pageLoading: boolean;
  loading: boolean;
  productList: CardInfo[];
  CommandList: CardInfo[];
}

const initialState: IInitialState = {
  pageLoading: true,
  loading: true,
  productList: [],
  CommandList: [],
};

export const getList = createAsyncThunk(
  `/${namespace}`,
  async () => {
    const result = await getProductList();
    const result = await getCommandList();
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
        state.productList = action.payload?.list || [];
        state.CommandList = action.payload?.list || [];
      })
      .addCase(getList.rejected, (state) => {
        state.loading = false;
      });
  },
});

export const { switchPageLoading } = listCardSlice.actions;

export const selectListCard = (state: RootState) => state.card;

export default listCardSlice.reducer;
