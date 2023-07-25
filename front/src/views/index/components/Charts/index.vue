<!--
 * @Description: 
 * @version: 1.0.0
 * @Author: William
 * @Date: 2023-07-13 16:34:30
 * @LastEditors: William
 * @LastEditTime: 2023-07-13 17:31:34
-->
<template>
  <div :id="id"></div>
</template>

<script lang="ts" setup>
  // 引入 echarts 核心模块，核心模块提供了 echarts 使用必须要的接口。
  import * as echarts from 'echarts';
  import { onMounted } from 'vue';
  onMounted(() => {
    const chart = echarts.init(document.getElementById('chart'));
    chart.setOption({
      title: {
        text: '内存使用率',
        left: 'center',
      },
      series: [
        {
          
          type: 'pie',
          data: [
            {
              value: 335,
              name: 'MySQL'
            },
            {
              value: 234,
              name: 'Redis'
            },
            {
              value: 1548,
              name: '剩余内存'
            }
          ],
          radius: ['40%', '70%']
        }
      ]
    })
    const widAndHei = () => {
      props.id.style.width = '600px';
      props.id.style.height = '400px';
    }
  })
  const props = defineProps({
    id: {
      type: String,
      default: 'chart'
    }
  })
  
  
</script>

<style lang="scss" scoped>
  #chart {
    width: 600px;
    height: 400px;
  }
</style>