<!--
 * @Description: 
 * @version: 
 * @Author: William
 * @Date: 2023-07-10 15:57:03
 * @LastEditors: William
 * @LastEditTime: 2023-07-26 14:14:36
-->
<template>
  <el-container>
    <el-aside>
      <sidebar />
    </el-aside>
    <el-main>
      <router-view v-slot="{ Component }">
        <transition name="el-fade-in-linear">
          <component :is="Component" />
        </transition>
      </router-view>
    </el-main>
  </el-container>
</template>

<script lang="ts" setup>
import sidebar from './components/sidebar/index.vue'

</script>

<style lang="scss" scoped>
.el-container {
  max-height: 550px;

  .el-aside {
    width: auto;
  }

}
</style>