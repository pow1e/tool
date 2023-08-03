/*
 * @Description:
 * @version: 1.0.0
 * @Author: William
 * @Date: 2023-08-02 09:33:32
 * @LastEditors: William
 * @LastEditTime: 2023-08-02 11:35:54
 */
import { useMemo } from 'react';
import { useAppSelector } from 'modules/store';
import { selectGlobal } from 'modules/global';
import { getChartColor } from 'utils/color';
import { CHART_COLORS } from 'configs/color';
import lodashSet from 'lodash/set';
import lodashMap from 'lodash/map';
import { ETheme } from '../types';

export type TChartColorKey = keyof typeof CHART_COLORS[ETheme.light];
/**
 * 根据当前主题色返回动态的图表颜色列表
 * @param options 图表的固定配置
 * @param configs 需要动态变换颜色的字段
 * @returns string[]
 */
export default function useDynamicChart(
  options: Record<string, any>,
  configs?: Partial<Record<TChartColorKey, Array<string>>>,
) {
  const { theme, color } = useAppSelector(selectGlobal);
  return useMemo(() => {
    const dynamicColor = getChartColor(theme, color);
    const newOptions = {
      ...options,
    };
    // 设置动态的图表颜色
    lodashSet(newOptions, 'color', dynamicColor.colorList);
    if (configs) {
      lodashMap(configs, (config : any, configKey: TChartColorKey) => {
        config?.map((val : any) => lodashSet(newOptions, val, dynamicColor[configKey]));
      });
    }
    return newOptions;
  }, [theme, color, options]);
}
