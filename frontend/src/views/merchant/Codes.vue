<template>
  <div class="codes-management">
    <!-- 搜索和操作栏 -->
    <el-card class="search-bar">
      <el-form :inline="true" :model="searchForm">
        <el-form-item label="防伪码">
          <el-input v-model="searchForm.code" placeholder="请输入防伪码" clearable />
        </el-form-item>
        <el-form-item label="批次号">
          <el-input v-model="searchForm.batch" placeholder="请输入批次号" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择状态" clearable>
            <el-option label="未使用" value="unused" />
            <el-option label="已使用" value="used" />
            <el-option label="已过期" value="expired" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">搜索</el-button>
          <el-button @click="handleReset">重置</el-button>
          <el-button type="success" @click="navigateToGenerate">生成防伪码</el-button>
          <el-button type="warning" @click="handleExport">导出数据</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <!-- 防伪码列表 -->
    <el-card>
      <template #header>
        <span>防伪码列表</span>
      </template>
      <el-table :data="codeList" style="width: 100%">
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="code" label="防伪码" min-width="200">
          <template #default="scope">
            <el-tooltip :content="scope.row.code" placement="top">
              <span class="code-text">{{ formatCode(scope.row.code) }}</span>
            </el-tooltip>
          </template>
        </el-table-column>
        <el-table-column prop="batch" label="批次号" width="120" />
        <el-table-column prop="product" label="产品名称" min-width="150" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="scope">
            <el-tag :type="getStatusType(scope.row.status)">
              {{ getStatusText(scope.row.status) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="createTime" label="生成时间" width="160" />
        <el-table-column prop="verifyTime" label="验证时间" width="160" />
        <el-table-column prop="verifyLocation" label="验证地点" min-width="120" />
        <el-table-column label="操作" width="120" fixed="right">
          <template #default="scope">
            <el-button size="small" @click="handleView(scope.row)">查看</el-button>
            <el-button 
              size="small" 
              type="danger" 
              :disabled="scope.row.status !== 'unused'"
              @click="handleInvalidate(scope.row)"
            >
              作废
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <!-- 分页 -->
      <div class="pagination">
        <el-pagination
          v-model:current-page="pagination.current"
          v-model:page-size="pagination.size"
          :total="pagination.total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
      </div>
    </el-card>

    <!-- 防伪码详情对话框 -->
    <el-dialog
      v-model="detailVisible"
      title="防伪码详情"
      width="600px"
    >
      <el-descriptions :column="2" border>
        <el-descriptions-item label="防伪码">{{ currentCode?.code }}</el-descriptions-item>
        <el-descriptions-item label="批次号">{{ currentCode?.batch }}</el-descriptions-item>
        <el-descriptions-item label="产品名称">{{ currentCode?.product }}</el-descriptions-item>
        <el-descriptions-item label="状态">
          <el-tag :type="getStatusType(currentCode?.status)">
            {{ getStatusText(currentCode?.status) }}
          </el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="生成时间">{{ currentCode?.createTime }}</el-descriptions-item>
        <el-descriptions-item label="验证时间">{{ currentCode?.verifyTime || '-' }}</el-descriptions-item>
        <el-descriptions-item label="验证地点">{{ currentCode?.verifyLocation || '-' }}</el-descriptions-item>
        <el-descriptions-item label="验证IP">{{ currentCode?.verifyIp || '-' }}</el-descriptions-item>
        <el-descriptions-item label="验证设备">{{ currentCode?.verifyDevice || '-' }}</el-descriptions-item>
      </el-descriptions>
      <template #footer>
        <el-button @click="detailVisible = false">关闭</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'

const router = useRouter()

const searchForm = reactive({
  code: '',
  batch: '',
  status: ''
})

const codeList = ref([])
const pagination = reactive({
  current: 1,
  size: 10,
  total: 0
})

const detailVisible = ref(false)
const currentCode = ref(null)

// 模拟数据
const mockCodes = Array.from({ length: 100 }, (_, index) => {
  const status = index % 3 === 0 ? 'used' : index % 5 === 0 ? 'expired' : 'unused'
  const isUsed = status === 'used'
  
  return {
    id: index + 1,
    code: `ABC${String(index + 1).padStart(6, '0')}XYZ${String(index + 1000).padStart(4, '0')}`,
    batch: `BATCH${Math.floor(index / 10) + 1}`,
    product: `产品${Math.floor(index / 20) + 1}`,
    status: status,
    createTime: `2024-01-${String(15 - index % 15).padStart(2, '0')} 10:00:00`,
    verifyTime: isUsed ? `2024-01-${String(15 - index % 15).padStart(2, '0')} 14:30:00` : null,
    verifyLocation: isUsed ? ['北京市', '上海市', '广州市', '深圳市'][index % 4] : null,
    verifyIp: isUsed ? `192.168.1.${index % 255}` : null,
    verifyDevice: isUsed ? ['Chrome', 'Safari', 'Firefox', 'Edge'][index % 4] : null
  }
})

// 方法
const handleSearch = () => {
  pagination.current = 1
  loadCodes()
}

const handleReset = () => {
  Object.assign(searchForm, {
    code: '',
    batch: '',
    status: ''
  })
  pagination.current = 1
  loadCodes()
}

const navigateToGenerate = () => {
  router.push('/merchant/codes/generate')
}

const handleExport = () => {
  ElMessage.success('导出功能开发中')
}

const handleView = (row) => {
  currentCode.value = row
  detailVisible.value = true
}

const handleInvalidate = async (row) => {
  try {
    await ElMessageBox.confirm(
      `确定要作废防伪码"${formatCode(row.code)}"吗？此操作不可恢复！`,
      '警告',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'error'
      }
    )
    
    // 模拟API调用
    row.status = 'expired'
    ElMessage.success('作废成功')
  } catch {
    // 用户取消操作
  }
}

const formatCode = (code) => {
  if (!code) return ''
  if (code.length <= 12) return code
  return code.substring(0, 6) + '...' + code.substring(code.length - 6)
}

const getStatusType = (status) => {
  const typeMap = {
    unused: 'success',
    used: 'primary',
    expired: 'danger'
  }
  return typeMap[status] || 'info'
}

const getStatusText = (status) => {
  const textMap = {
    unused: '未使用',
    used: '已使用',
    expired: '已过期'
  }
  return textMap[status] || '未知'
}

const handleSizeChange = (size) => {
  pagination.size = size
  pagination.current = 1
  loadCodes()
}

const handleCurrentChange = (current) => {
  pagination.current = current
  loadCodes()
}

const loadCodes = () => {
  // 模拟API调用
  const filtered = mockCodes.filter(code => {
    const codeMatch = !searchForm.code || code.code.includes(searchForm.code)
    const batchMatch = !searchForm.batch || code.batch.includes(searchForm.batch)
    const statusMatch = !searchForm.status || code.status === searchForm.status
    return codeMatch && batchMatch && statusMatch
  })
  
  const start = (pagination.current - 1) * pagination.size
  const end = start + pagination.size
  codeList.value = filtered.slice(start, end)
  pagination.total = filtered.length
}

onMounted(() => {
  loadCodes()
})
</script>

<style scoped>
.codes-management {
  padding: 0;
}

.search-bar {
  margin-bottom: 20px;
}

.code-text {
  font-family: 'Courier New', monospace;
  font-weight: bold;
  color: #409eff;
  cursor: pointer;
}

.pagination {
  margin-top: 20px;
  text-align: right;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .el-form--inline .el-form-item {
    margin-right: 0;
    margin-bottom: 10px;
    width: 100%;
  }
  
  .el-form--inline .el-form-item__content {
    width: 100%;
  }
  
  .code-text {
    font-size: 12px;
  }
}
</style>