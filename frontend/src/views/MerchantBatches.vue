<template>
  <div class="merchant-batches">
    <div class="page-header">
      <h1>批次管理</h1>
      <p>管理防伪码生成批次</p>
    </div>

    <!-- 搜索和操作栏 -->
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="批次号">
          <el-input
            v-model="searchForm.batchNo"
            placeholder="请输入批次号"
            clearable
          />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择状态" clearable>
            <el-option label="正常" value="active" />
            <el-option label="已过期" value="expired" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">搜索</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <!-- 批次列表 -->
    <el-card>
      <el-table :data="batchList" stripe>
        <el-table-column prop="batchNo" label="批次号" width="140" />
        <el-table-column prop="batchName" label="批次名称" min-width="120" />
        <el-table-column prop="codeCount" label="防伪码数量" width="100" align="center">
          <template #default="{ row }">
            <el-tag>{{ row.codeCount }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="verifiedCount" label="已验证数量" width="100" align="center">
          <template #default="{ row }">
            <el-tag type="success">{{ row.verifiedCount }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="createTime" label="创建时间" width="160" />
        <el-table-column prop="expiryTime" label="过期时间" width="160">
          <template #default="{ row }">
            <span :class="{ 'expired': isExpired(row.expiryTime) }">
              {{ row.expiryTime || '永不过期' }}
            </span>
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row)">
              {{ getStatusText(row) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button size="small" @click="handleView(row)">查看详情</el-button>
            <el-button size="small" type="success" @click="handleExport(row)">导出</el-button>
            <el-button size="small" type="danger" @click="handleDelete(row)">删除</el-button>
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
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'

const searchForm = reactive({
  batchNo: '',
  status: ''
})

const batchList = ref([])
const pagination = reactive({
  current: 1,
  size: 10,
  total: 0
})

// 模拟批次数据
const mockBatches = [
  {
    id: 1,
    batchNo: 'BATCH2024001',
    batchName: '2024年第一季度防伪码',
    codeCount: 50000,
    verifiedCount: 12560,
    createTime: '2024-01-15 10:00:00',
    expiryTime: '2025-01-15 10:00:00',
    status: 'active'
  },
  {
    id: 2,
    batchNo: 'BATCH2024002',
    batchName: '2024年第二季度防伪码',
    codeCount: 30000,
    verifiedCount: 8560,
    createTime: '2024-04-01 09:30:00',
    expiryTime: '2025-04-01 09:30:00',
    status: 'active'
  },
  {
    id: 3,
    batchNo: 'BATCH2023121',
    batchName: '2023年12月防伪码',
    codeCount: 10000,
    verifiedCount: 9800,
    createTime: '2023-12-01 14:00:00',
    expiryTime: '2024-12-01 14:00:00',
    status: 'expired'
  }
]

const isExpired = (expiryTime) => {
  if (!expiryTime) return false
  return new Date(expiryTime) < new Date()
}

const getStatusType = (row) => {
  if (isExpired(row.expiryTime)) return 'danger'
  return 'success'
}

const getStatusText = (row) => {
  if (isExpired(row.expiryTime)) return '已过期'
  return '正常'
}

const loadBatches = () => {
  // 模拟API调用
  setTimeout(() => {
    batchList.value = mockBatches
    pagination.total = mockBatches.length
  }, 500)
}

const handleSearch = () => {
  pagination.current = 1
  loadBatches()
}

const handleReset = () => {
  Object.keys(searchForm).forEach(key => {
    searchForm[key] = ''
  })
  pagination.current = 1
  loadBatches()
}

const handleView = (row) => {
  ElMessage.info(`查看批次: ${row.batchNo}`)
}

const handleExport = (row) => {
  ElMessage.info(`导出批次: ${row.batchNo}`)
}

const handleDelete = async (row) => {
  try {
    await ElMessageBox.confirm(
      `确定要删除批次 "${row.batchName}" 吗？此操作将删除该批次下的所有防伪码。`,
      '警告',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'error'
      }
    )
    
    ElMessage.success('删除成功')
  } catch (error) {
    // 用户取消删除
  }
}

const handleSizeChange = (size) => {
  pagination.size = size
  pagination.current = 1
  loadBatches()
}

const handleCurrentChange = (current) => {
  pagination.current = current
  loadBatches()
}

onMounted(() => {
  loadBatches()
})
</script>

<style scoped>
.merchant-batches {
  padding: 20px;
}

.page-header {
  margin-bottom: 20px;
}

.page-header h1 {
  margin: 0 0 10px 0;
  font-size: 24px;
  font-weight: 600;
  color: #333;
}

.page-header p {
  margin: 0;
  color: #666;
  font-size: 14px;
}

.search-card {
  margin-bottom: 20px;
}

.expired {
  color: #f56c6c;
  text-decoration: line-through;
}

.pagination {
  margin-top: 20px;
  text-align: right;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .merchant-batches {
    padding: 15px;
  }
  
  .search-card .el-form--inline .el-form-item {
    margin-right: 0;
    margin-bottom: 10px;
  }
}
</style>