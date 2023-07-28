<template>
  <div class="encrypt">
    <el-form label-width="120px">
      <el-form-item label="原文">
        <el-input v-model="input" placeholder="请输入原文" class="input"></el-input>
        <el-button type="primary" @click="encrypt" class="button">加密</el-button>
      </el-form-item>
      <el-form-item label="加密结果" />
      <el-form-item label="16位大写">
        <el-input v-model="sixteenUpper" placeholder="这里是加密后的内容" readonly></el-input>
      </el-form-item>
      <el-form-item label="16位小写">
        <el-input v-model="sixteenLower" placeholder="这里是加密后的内容" readonly></el-input>
      </el-form-item>
      <el-form-item label="32位大写">
        <el-input v-model="thirtyTwoUpper" placeholder="这里是加密后的内容" readonly></el-input>
      </el-form-item>
      <el-form-item label="32位小写">
        <el-input v-model="thirtyTwoLower" placeholder="这里是加密后的内容" readonly></el-input>
      </el-form-item>
    </el-form>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { getMd5Encrypt } from '../../api/md5Translate';
import { ElMessage } from 'element-plus';
const input = ref('')
const sixteenUpper = ref('')
const sixteenLower = ref('')
const thirtyTwoUpper = ref('')
const thirtyTwoLower = ref('')
const encrypt = () => {
  const data = {
    "plaintext": input.value
  }
  getMd5Encrypt(data).then(res => {
    ElMessage.closeAll()
    if (res.code === 1) {
      ElMessage.success(res.msg)
      sixteenLower.value = res.data['16lowercase']
      sixteenUpper.value = res.data['16uppercase']
      thirtyTwoUpper.value = res.data['32uppercase']
      thirtyTwoLower.value = res.data['32lowercase']
    }
    else if (res.data.code === 0)
      ElMessage.error(res.data.msg)
  })
}
</script>

<style lang="scss" scoped>
.encrypt {
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