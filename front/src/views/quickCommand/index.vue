<!--
 * @Description:
 * @version: 1.0.0
 * @Author: William
 * @Date: 2023-07-12 15:27:12
 * @LastEditors: William
 * @LastEditTime: 2023-07-21 08:58:56
-->
<template>
  <el-button type="primary" size="large" class="top-button" @click="dv = true">
    新增配置
  </el-button>
  <div class="main-area">
    <el-scrollbar height="425px">
      <el-empty description="暂无数据，请新增配置" v-if="!data" />
      <Card v-for="item in data" :key="item" :title="item.name" :desc="item.description" :command="item.directives"
        :permit="item.type.toString()" :id="item.id" :handle-update-data="handleUpdateData" v-bind:get-data="getList"/>
    </el-scrollbar>
  </div>
  <el-dialog v-model="dv" title="新增配置" width="50%" draggable align-center destroy-on-close>
    <el-form label-width="80px">
      <el-form-item label="服务名称">
        <el-input v-model="serviceTitle"></el-input>
      </el-form-item>
      <el-form-item label="服务描述">
        <el-input v-model="serviceDesc"></el-input>
      </el-form-item>
      <el-form-item label="启动命令">
        <el-input v-model="serviceComm"></el-input>
      </el-form-item>
      <el-form-item label="权限">
        <el-radio-group v-model="radio">
          <el-radio label="0">管理员</el-radio>
          <el-radio label="1">普通用户</el-radio>
        </el-radio-group>
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dv = false">取 消</el-button>
        <el-button type="primary" @click="submit">
          确 认
        </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script lang="ts" setup>
import { onMounted, ref } from 'vue'
import Card from './components/Card/index.vue'
import { ElMessage } from 'element-plus'
import { addQuickCommand, getQuickCommandList } from './api/quickCommand'

const dv = ref(false)
const serviceTitle = ref(null)
const serviceDesc = ref(null)
const serviceComm = ref(null)
const radio = ref('0')
const submit = () => {
  if (serviceTitle.value == null || serviceDesc.value == null || serviceComm.value == null) {
    ElMessage.closeAll();
    ElMessage({
      message: '请填写完整',
      type: 'error'
    })
  }
  else if (serviceTitle.value == '' || serviceDesc.value == '' || serviceComm.value == '') {
    ElMessage.closeAll();
    ElMessage({
      message: '请填写完整',
      type: 'error'
    })
  }
  else {
    const addData = {
      "name": serviceTitle.value,
      "description": serviceDesc.value,
      "directives": serviceComm.value,
      "type": radio.value
    }
    addQuickCommand(addData).then(result => {
      if (result.code == 1) {
        ElMessage.closeAll();
        ElMessage({
          message: '添加成功',
          type: 'success'
        })
        getList()
      }
      else {
        ElMessage.closeAll();
        ElMessage({
          message: '添加失败',
          type: 'error'
        })
      }
    }).catch(error => {
      ElMessage.closeAll();
      ElMessage({
        message: `${error}`,
        type: 'error'
      })
    })
    serviceTitle.value = null
    serviceDesc.value = null
    serviceComm.value = null
    dv.value = false
  }
}
// 对接接口
const data = ref({})
const getList = async () => {
  const res = await getQuickCommandList()
  data.value = res.data
}
onMounted(() => {
  getList()
})
const handleUpdateData = (new_data: { title: string; desc: string; command: string; permit: string; id: string; }) => {
  for (let i = 0; i < data.value.length; i++) {
    if (data.value[i].id == new_data.id) {
      data.value[i].name = new_data.title
      data.value[i].description = new_data.desc
      data.value[i].directives = new_data.command
      data.value[i].type = new_data.permit
    }
  }
}

</script>

<style lang="scss" scoped>
.top-button {
  margin-bottom: 5%;
}

.main-area {
  width: 100%;

  .el-scrollbar__view {
    width: 100%;
  }

  .scrollbar-demo-item {
    width: 100%;
    height: 100px;
    margin: 10px;
    border-radius: 4px;
    background: var(--el-color-primary-light-9);
    color: var(--el-color-primary);
  }
}
</style>
