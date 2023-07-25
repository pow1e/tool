<!--
 * @Description: 
 * @version: 1.0.0
 * @Author: William
 * @Date: 2023-07-12 16:47:25
 * @LastEditors: William
 * @LastEditTime: 2023-07-25 11:12:05
-->
<template>
  <el-tabs v-model="activeName" class="md5Tabs" @tab-click="handleClick">
      <el-tab-pane label="MD5解密" name="first">
        <decrypt />
      </el-tab-pane>
      <el-tab-pane label="MD5加密" name="second">
        <encrypt />
      </el-tab-pane>
    </el-tabs>
</template>

<script lang="ts" setup>
import { ref } from 'vue' 
import type { TabsPaneContext } from 'element-plus'
import decrypt from './components/decrypt/index.vue'
import encrypt from './components/encrypt/index.vue'
const activeName = ref('first')

const handleClick = (tab: TabsPaneContext, event: Event) => {
  console.log(tab, event)
}
</script>

<style lang="scss" scoped>

</style>