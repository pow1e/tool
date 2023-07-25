<!--
 * @Description:
 * @version:
 * @Author: William
 * @Date: 2023-07-12 16:52:01
 * @LastEditors: William
 * @LastEditTime: 2023-07-21 09:13:39
-->
<template>
  <el-card class="box-card">
    <template #header>
      <div class="card-header">
        <span>服务：{{ title }}</span>
        <el-button icon="CaretRight" circle @click="submitButton" class="card-button-run" />
        <el-icon class="card-button-right-middle" @click="dialogVisible = true" :size="25">
          <Edit />
        </el-icon>
        <el-icon class="card-button-right-border" @click="delDialogVisible = true" :size="25">
          <Delete />
        </el-icon>
      </div>
    </template>
    <div class="text item">
      <span>服务描述：{{ desc }}</span>
    </div>
  </el-card>
  <el-dialog v-model="dialogVisible" title="修改配置" width="50%" draggable destroy-on-close align-center>
    <el-form label-width="80px">
      <el-form-item label="id">
        <el-input v-model="id" disabled></el-input>
      </el-form-item>
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
        <el-radio-group v-model="permit">
          <el-radio label="0">普通用户</el-radio>
          <el-radio label="1">管理员</el-radio>
        </el-radio-group>
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="putData">
          确 认
        </el-button>
      </span>
    </template>
  </el-dialog>
  <el-dialog v-model="delDialogVisible" title="删除快捷指令" width="30%" draggable destroy-on-close align-center>
    <span style="font-size: 16px;">确认<span style="color: red;">删除</span>该快捷指令吗？</span>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="delDialogVisible = false">取 消</el-button>
        <el-button type="danger" @click="delData">
          确 认
        </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script lang="ts" setup>
import { inject, ref, emit } from 'vue'
import { ElMessage } from 'element-plus'
import { deleteQuickCommand, runQuickCommand, updateQuickCommand } from '../../api/quickCommand';

const props = defineProps({
  title: {
    type: String,
    default: '标题'
  },
  desc: {
    type: String,
    default: '描述'
  },
  command: {
    type: String,
    default: null
  },
  permit: {
    type: String,
    default: '0'
  },
  id: {
    type: String,
    default: null
  },
  handleUpdateData: {
    type: Function,
    default: () => { }
  },
  getData: {
    type: Function,
    default: () => { }
  }
})
const cardSwitch = ref(false)
const dialogVisible = ref(false)
const delDialogVisible = ref(false)
// 设置值
const serviceTitle = ref(props.title)
const serviceDesc = ref(props.desc)
const serviceComm = ref(props.command)
const permit = ref(props.permit)
const id = ref(props.id)
const submitButton = () => {
  ElMessage.closeAll()
  runQuickCommand(props.id).then(res => {
    if (res.code == 1) {
      ElMessage({
        message: '执行成功',
        type: 'success'
      })
    } else {
      ElMessage({
        message: '执行失败',
        type: 'error'
      })
    }
  })
}
const putData = () => {
  const errorMessage =
    serviceTitle.value == null || serviceDesc.value == null || serviceComm.value == null ? '请填写完整' :
      serviceTitle.value == '' || serviceDesc.value == '' || serviceComm.value == '' ? '请填写完整' :
        serviceTitle.value == props.title && serviceDesc.value == props.desc && serviceComm.value == props.command && permit.value == props.permit ? '请修改后提交' :
          '';

  if (errorMessage) {
    ElMessage({
      message: errorMessage,
      type: 'error'
    })
  } else {
    updateQuickCommand({
      id: props.id,
      name: serviceTitle.value,
      description: serviceDesc.value,
      directives: serviceComm.value,
      type: permit.value
    }).then(res => {
      if (res.code == 1) {
        ElMessage({
          message: '修改成功',
          type: 'success'
        })
        const updatedData = {
          id: id.value,
          title: serviceTitle.value,
          desc: serviceDesc.value,
          command: serviceComm.value,
          permit: permit.value
        }
        props.handleUpdateData(updatedData)
      } else {
        ElMessage({
          message: '修改失败',
          type: 'error'
        })
      }
    })
    dialogVisible.value = false
  }
}
const delData = () => {
  ElMessage.closeAll()
  deleteQuickCommand(props.id).then(res => {
    if (res.code == 1) {
      delDialogVisible.value = false
      ElMessage({
        message: '删除成功',
        type: 'success'
      })
      props.getData()
    } else {
      delDialogVisible.value = false
      ElMessage({
        message: '删除失败',
        type: 'error'
      })
    }
  })
}

</script>

<style lang="scss" scoped>
.box-card {
  margin-bottom: 20px;

  .card-header {
    .card-button-right-middle {
      width: 1.25em;
      height: 1.25em;
      float: right;
      margin-right: 10px;
    }

    .card-button-right-middle:hover {
      cursor: pointer;
      color: var(--el-color-primary);
    }

    .card-button-right-border {
      width: 1.25em;
      height: 1.25em;
      float: right;
      margin-right: 10px;
      margin-left: 5%;
    }

    .card-button-right-border:hover {
      cursor: pointer;
      color: var(--el-color-primary);
    }

    .card-button-run {
      justify-content: normal;
      float: right;
    }
  }
}
</style>
