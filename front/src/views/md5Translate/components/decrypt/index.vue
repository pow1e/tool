<!--
 * @Description: 
 * @version: 1.0.0
 * @Author: William
 * @Date: 2023-07-24 11:11:32
 * @LastEditors: William
 * @LastEditTime: 2023-07-25 11:31:40
-->
<template>
<el-form>
    <el-form-item label="密文">
        <el-input v-model="input" placeholder="请输入密文" size="100px"></el-input>
    </el-form-item>
    <el-form-item>
        <el-button type="primary" @click="encrypt">加密</el-button>
    </el-form-item>
    <el-form-item label="加密结果">
        <el-input v-model="output" placeholder="这里是解密后的内容" readonly></el-input>
    </el-form-item>
</el-form>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
const input = ref('')
const output = ref('Dingzhen')
const encrypt = () => {
    console.log(input.value)
}
</script>

<style lang="scss" scoped>

</style>