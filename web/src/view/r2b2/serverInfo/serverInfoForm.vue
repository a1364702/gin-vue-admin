
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="ID:" prop="serverId">
    <el-input v-model.number="formData.serverId" :clearable="false" placeholder="请输入ID" />
</el-form-item>
        <el-form-item label="名称:" prop="serverName">
    <el-input v-model="formData.serverName" :clearable="false" placeholder="请输入名称" />
</el-form-item>
        <el-form-item label="组别:" prop="serverGroup">
    <el-input v-model="formData.serverGroup" :clearable="false" placeholder="请输入组别" />
</el-form-item>
        <el-form-item label="IP地址:" prop="serverIp">
    <el-input v-model="formData.serverIp" :clearable="false" placeholder="请输入IP地址" />
</el-form-item>
        <el-form-item label="端口:" prop="serverPort">
    <el-input v-model.number="formData.serverPort" :clearable="false" placeholder="请输入端口" />
</el-form-item>
        <el-form-item label="描述:" prop="serverDesc">
    <el-input v-model="formData.serverDesc" :clearable="false" placeholder="请输入描述" />
</el-form-item>
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createServerInfo,
  updateServerInfo,
  findServerInfo
} from '@/api/r2b2/serverInfo'

defineOptions({
    name: 'ServerInfoForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const formData = ref({
            serverId: 0,
            serverName: '',
            serverGroup: '',
            serverIp: '',
            serverPort: 0,
            serverDesc: '',
        })
// 验证规则
const rule = reactive({
               serverId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               serverName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               serverGroup : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               serverIp : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               serverPort : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findServerInfo({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      btnLoading.value = true
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return btnLoading.value = false
            let res
           switch (type.value) {
             case 'create':
               res = await createServerInfo(formData.value)
               break
             case 'update':
               res = await updateServerInfo(formData.value)
               break
             default:
               res = await createServerInfo(formData.value)
               break
           }
           btnLoading.value = false
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
