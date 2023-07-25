<!--
 * @Description: 
 * @version: 
 * @Author: William
 * @Date: 2023-07-12 15:27:12
 * @LastEditors: William
 * @LastEditTime: 2023-07-21 09:54:58
-->

<template>
  <div class="menu-container">
    <el-menu 
      :default-active="currentPath"
      class="el-sidebar"
      router 
    >
      <el-menu-item v-for="i in data" :key="i.index" :index.sync="i.index" :route="i.path" @click="activeIndex = i.index">
        <i :class="i.icon"></i>
        <span>{{ i.name }}</span>
      </el-menu-item>
      <div class="fake-block"></div>
      <el-divider />
      <div class="button-container">
        <el-switch v-model="isDark" :inactive-icon="Moon" :active-icon="Sunny" id="DarkSwitch" inline-prompt
          size="large"></el-switch>
      </div>
    </el-menu>
  </div>
</template>

<script lang="ts" setup>
import { computed, ref } from 'vue'
import { useDark, useToggle } from "@vueuse/core";
import { Sunny, Moon } from '@element-plus/icons-vue'
import { useRoute } from 'vue-router'
const isDark = useDark({
  storageKey: 'useDarkKEY', //状态存储的key
  valueDark: 'dark', //深色模式的值
  valueLight: 'light',  //浅色模式的值
});
const toggle = useToggle(isDark);
const data = [
  {
    "index": '/',
    "name": "主页",
    "icon": "el-icon-s-home",
    "path": "/",
  },
  {
    "index": '/qc',
    "name": "快速指令",
    "icon": "el-icon-s-order",
    "path": "/qc"
  },
  {
    "index": "/md5",
    "name": "MD5转换",
    "icon": "el-icon-s-data",
    "path": "/md5"
  }
]
const route = useRoute()
const currentPath = computed(() => route.path)
</script>

<style lang="scss" scoped>
.menu-container {
  display: flex;
  height: 550px;

  .menu-item-text {
    display: flex;
    align-items: center;
  }

  .button-container {
    display: flex;
    justify-content: center;
    margin-top: 20px;
    /* 调整按钮距离上方元素的间距 */
  }

  .el-sidebar {
    width: 200px;

    .el-menu {
      height: 100%;
    }
  }

  .fake-block {
    height: 200px;
  }
}
</style>