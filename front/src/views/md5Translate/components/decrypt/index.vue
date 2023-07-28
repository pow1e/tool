<!--
 * @Description: 
 * @version: 1.0.0
 * @Author: William
 * @Date: 2023-07-24 11:11:32
 * @LastEditors: William
 * @LastEditTime: 2023-07-26 10:26:53
-->
<template>
  <div class="decrypt">
    <el-form label-width="120px">
      <el-form-item label="密文">
        <el-input v-model="input" placeholder="请输入密文" class="input"></el-input>
        <el-button type="primary" @click="encrypt" class="button">解密</el-button>
      </el-form-item>
      <el-form-item label="解密结果">
        <el-input v-model="output" placeholder="这里是解密后的内容" readonly></el-input>
      </el-form-item>
    </el-form>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { getMd5Decrypt } from '../../api/md5Translate';
import { ElMessage } from 'element-plus';
const input = ref('')
const output = ref('')
const encrypt = () => {
  const data = {
    "ciphertext": input.value
  }
  getMd5Decrypt(data).then(res => {
    ElMessage.closeAll()
    if (input.value === '') {
      ElMessage.error('请输入密文')
    }
    else {
      if (res.code === 1) {
        ElMessage.success(res.msg)
        output.value = res.data
      }

      else if (res.code === 0)
        ElMessage.error(res.msg)
    }
  })
    
}
</script>

<style lang="scss" scoped>
.decrypt {
  width: 100%;
  height: 100%;
    //   这里写居中样式
    margin-top: 10%;
    
  .input {
    width: calc(100% - 80px);
  }

  .button {
    margin-left: 10px;
  }
}
</style>