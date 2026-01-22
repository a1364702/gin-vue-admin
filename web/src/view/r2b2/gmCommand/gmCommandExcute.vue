<template>
    <div class="gva-form-box" style="max-width: autosize; margin: 0 auto; padding: 20px;">
        <el-form ref="formRef" :model="formData" :rules="rules" label-position="right" label-width="100px" @submit.prevent>
            <!-- 服务器树选择 + 执行按钮 -->
            <el-form-item label="服务器:" prop="executeServer">
                <div style="display: flex; gap: 12px; align-items: flex-start;">
                    <el-tree-select :model-value="nowExecuteServer" :data="serverTreeData" placeholder="请选择服务器（可多选）"
                        style="flex: 1;" multiple filterable clearable check-strictly node-key="id"
                        :props="{ label: 'label', children: 'children', value: 'id' }"
                        @update:model-value="handleServerSelect" @visible-change="handleVisibleChange" />
                    <el-button type="primary" @click="handleSubmit">执行</el-button>
                </div>
            </el-form-item>
            <!-- 命令内容 -->
            <el-form-item label="命令内容:" prop="command">
                <el-input v-model="formData.command" type="textarea" placeholder="输入 help 查看命令帮助"
                    :autosize="{ minRows: 4, maxRows: 10 }" clearable />
                <el-input v-model="commandResult" type="textarea" placeholder="这里显示命令的结果"
                    :autosize="{ minRows: 4, maxRows: 10 }" clearable readonly />
            </el-form-item>
        </el-form>
    </div>
</template>
  
<script setup>
import { ref, reactive, watch } from 'vue'
import { ElMessage } from 'element-plus'

import {
    getServerInfoList
} from '@/api/r2b2/serverInfo'

import {
    excuteCommand
} from '@/api/r2b2/gmCommand'

// 表单数据
const formData = reactive({
    command: '',
    executor: '',
    result: '',
    executeTime: new Date(),
    executeServer: new Array(),
})

const nowExecuteServer = ref([])

const rules = reactive({
    nowExecuteServer: [{ required: true, message: '请选择至少一个服务器', trigger: 'change' }],
    command: [{ required: true, message: '请输入命令内容', trigger: 'blur' }]
})

const formRef = ref()
var nodeMap = ref()

// 命令执行结果（响应式）
const commandResult = ref('') // 初始为空字符串

// 模拟服务器树（注意：只有叶子节点才是真实服务器）
const serverTreeData = ref([
])


// 工具：判断是否为叶子节点
const isLeaf = (node) => !node.children || node.children.length === 0

// 工具：递归获取某节点下的所有叶子 ID
const getLeafIds = (nodes) => {
    const leaves = []
    const traverse = (list) => {
        for (const node of list) {
            if (isLeaf(node)) {
                leaves.push(node.id)
            } else {
                traverse(node.children)
            }
        }
    }
    traverse(nodes)
    return leaves
}

// 下拉面板展开/收起时触发
const handleVisibleChange = async (visible) => {
    if (visible) {
        // 只有在面板展开时才重新获取服务器数据
        await getServerInfoFunc()
    }
}

// 构建一个从 id 到 node 的映射（用于快速查找）
const buildNodeMap = (data) => {
    const map = {}
    const walk = (nodes) => {
        for (const node of nodes) {
            map[node.id] = node
            if (node.children) walk(node.children)
        }
    }
    walk(data)
    return map
}


const getServerInfoFunc = async () => {
    const res = await getServerInfoList()
    if (res.code === 0) {
        const serverList = res.data.list

        const result = []
        const groupMap = new Map()

        serverList.forEach(server => {
            const { serverGroup, serverName, serverId } = server

            if (serverGroup && serverGroup.trim()) {
                // 有分组的服务器：加入对应区域
                if (!groupMap.has(serverGroup)) {
                    groupMap.set(serverGroup, {
                        id: serverGroup, // 生成唯一 ID
                        label: serverGroup,
                        children: []
                    })
                }

                const groupNode = groupMap.get(serverGroup)
                groupNode.children.push({
                    id: `${serverId}`,
                    label: serverName
                })
            } else {
                // 没有分组的服务器：直接作为叶子节点
                result.push({
                    id: `${serverId}`,
                    label: serverName
                })
            }
        })
        // 将所有分组节点添加到结果中
        groupMap.forEach(node => {
            result.push(node)
        })
        serverTreeData.value = result
    }
    nodeMap = buildNodeMap(serverTreeData.value)
}
getServerInfoFunc()


// 处理服务器选择变化
const handleServerSelect = (selectedIds) => {
    const newLeaves = new Set()

    for (const id of selectedIds) {
        const node = nodeMap[id]
        if (!node) continue

        if (isLeaf(node)) {
            // 是叶子，直接加入
            newLeaves.add(id)
        } else {
            // 是父节点，展开所有叶子
            const leafIds = getLeafIds([node])
            leafIds.forEach(leafId => newLeaves.add(leafId))
        }
    }

    // 更新表单值（只保留叶子）
    nowExecuteServer.value = Array.from(newLeaves)
}

// 提交
const handleSubmit = async () => {
    try {
        await formRef.value.validate()
        commandResult.value = ''
        const executeServer = nowExecuteServer.value.map(id => {
            const num = parseInt(id, 10)
            if (isNaN(num)) {
                throw new Error(`无效服务器ID: ${id}`)
            }
            return num
        })
        formData.executeServer = executeServer
        const result = await excuteCommand(formData)
        if (result.code === 0) {
            commandResult.value = typeof result.data === 'string'
                ? result.data
                : JSON.stringify(result.data, null, 2)
        } else {
            ElMessage.error(result.msg || '命令执行失败')
            commandResult.value = result.msg || '执行出错'
        }

    } catch (error) {
        if (nowExecuteServer.value.length === 0) {
            ElMessage.error("请选择服务器")
            return
        }
        if (formData.command.length === 0) {
            ElMessage.error("请输入命令")
            return
        }
    }
}
</script>