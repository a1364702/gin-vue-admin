
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="命令内容:" prop="command">
    <el-input v-model="formData.command" :clearable="true" placeholder="请输入命令内容" />
</el-form-item>
        <el-form-item label="执行人:" prop="executor">
    <el-select v-model="formData.executor" placeholder="请选择执行人" filterable style="width:100%" :clearable="false">
        <el-option v-for="(item,key) in dataSource.executor" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="执行结果:" prop="result">
    <RichEdit v-model="formData.result"/>
</el-form-item>
        <el-form-item label="执行时间:" prop="executeTime">
    <el-date-picker v-model="formData.executeTime" type="date" style="width:100%" placeholder="选择日期" :clearable="false" />
</el-form-item>
        <el-form-item label="服务器:" prop="executeServer">
    <el-select v-model="formData.executeServer" placeholder="请选择服务器" filterable style="width:100%" :clearable="true">
        <el-option v-for="(item,key) in dataSource.executeServer" :key="key" :label="item.label" :value="item.value" />
    </el-select>
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
    getGmCommandDataSource,
  createGmCommand,
  updateGmCommand,
  findGmCommand
} from '@/api/r2b2/gmCommand'

defineOptions({
    name: 'GmCommandForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const formData = ref({
            command: '',
            executor: '',
            result: '',
            executeTime: new Date(),
            executeServer: undefined,
        })
// 验证规则
const rule = reactive({
               command : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               executor : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               executeTime : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               executeServer : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getGmCommandDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findGmCommand({ ID: route.query.id })
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
               res = await createGmCommand(formData.value)
               break
             case 'update':
               res = await updateGmCommand(formData.value)
               break
             default:
               res = await createGmCommand(formData.value)
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
