<template>
  <div class="merchant-codes">
    <div class="page-header">
      <h1>防伪码列表</h1>
      <p>管理您的防伪码信息</p>
    </div>

    <!-- 搜索和操作栏 -->
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="防伪码">
          <el-input
            v-model="searchForm.code"
            placeholder="请输入防伪码"
            clearable
          />
        </el-form-item>
        <el-form-item label="批次号">
          <el-input
            v-model="searchForm.batchNo"
            placeholder="请输入批次号"
            clearable
          />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择状态" clearable>
            <el-option label="未验证" value="unverified" />
            <el-option label="已验证" value="verified" />
            <el-option label="已过期" value="expired" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">搜索</el-button>
          <el-button @click="handleReset">重置</el-button>
          <el-button type="success" @click="handleExport">导出数据</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <!-- 防伪码列表 -->
    <el-card>
      <el-table :data="codeList" stripe>
        <el-table-column prop="code" label="防伪码" min-width="180">
          <template #default="{ row }">
            <el-tooltip :content="row.code" placement="top">
              <span class="code-text">{{ formatCode(row.code) }}</span>
            </el-tooltip>
          </template>
        </el-table-column>
        <el-table-column prop="batchNo" label="批次号" width="120" />
        <el-table-column prop="createTime" label="生成时间" width="160" />
        <el-table-column prop="expiryTime" label="过期时间" width="160">
          <template #default="{ row }">
            <span :class="{ 'expired': isExpired(row.expiryTime) }">
              {{ row.expiryTime || '永不过期' }}
            </span>
          </template>
        </el-table-column>
        <el-table-column prop="verifyTimes" label="验证次数" width="100" align="center">
          <template #default="{ row }">
            <el-tag :type="row.verifyTimes > 0 ? 'success' : 'info'">
              {{ row.verifyTimes }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="lastVerifyTime" label="最后验证" width="160">
          <template #default="{ row }">
            {{ row.lastVerifyTime || '未验证' }}
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row)">
              {{ getStatusText(row) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="120" fixed="right">
          <template #default="{ row }">
            <el-button size="small" @click="handleView(row)">查看</el-button>
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
  code: '',
  batchNo: '',
  status: ''
})

const codeList = ref([])
const pagination = reactive({
  current: 1,
  size: 10,
  total: 0
})

// 模拟防伪码数据
const mockCodes = [
  {
    id: 1,
    code: 'ABC123456XYZ7890DEF',
    batchNo: 'BATCH2024001',
    createTime: '2024-01-15 10:00:00',
    expiryTime: '2025-01-15 10:00:00',
    verifyTimes: 3,
    lastVerifyTime: '2024-01-20 14:30:00',
    status: 'verified'
  },
  {
    id: 2,
    code: 'DEF987654UVW3210GHI',
    batchNo: 'BATCH2024001',
    createTime: '2024-01-15 10:01:00',
    expiryTime: '2025-01-15 10:01:00',
    verifyTimes: 0,
    lastVerifyTime: null,
    status: 'unverified'
  },
  {
    id: 3,
    code: 'JKL456789MNO1234PQR',
    batchNo: 'BATCH2024002',
    createTime: '2024-01-16 09:30:00',
    expiryTime: '2024-02-16 09:30:00',
    verifyTimes: 1,
    lastVerifyTime: '2024-01-18 16:45:00',
    status: 'verified'
  }
]

const formatCode = (code) => {
  if (code.length > 12) {
    return code.substring(0, 6) + '...' + code.substring(code.length - 6)
  }
  return code
}

const isExpired = (expiryTime) => {
  if (!expiryTime) return false
  return new Date(expiryTime) < new Date()
}

const getStatusType = (row) => {
  if (isExpired(row.expiryTime)) return 'danger'
  if (row.verifyTimes > 0) return 'success'
  return 'info'
}

const getStatusText = (row) => {
  if (isExpired(row.expiryTime)) return '已过期'
  if (row.verifyTimes > 0) return '已验证'
  return '未验证'
}

const loadCodes = () => {
  // 模拟API调用
  setTimeout(() => {
    codeList.value = mockCodes
    pagination.total = mockCodes.length
  }, 500)
}

const handleSearch = () => {
  pagination.current = 1
  loadCodes()
}

const handleReset = () => {
  Object.keys(searchForm).forEach(key => {
    searchForm[key] = ''
  })
  pagination.current = 1
  loadCodes()
}

const handleExport = () => {
  ElMessage.info('导出功能开发中')
}

const handleView = (row) => {
  ElMessage.info(`查看防伪码: ${row.code}`)
}

const handleDelete = async (row) => {
  try {
    await ElMessageBox.confirm(
      `确定要删除防伪码 "${row.code}" 吗？此操作不可恢复。`,
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
  loadCodes()
}

const handleCurrentChange = (current) => {
  pagination.current = current
  loadCodes()
}

onMounted(() => {
  loadCodes()
})
</script>

<style scoped>
.merchant-codes {
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

.code-text {
  font-family: 'Courier New', monospace;
  font-weight: 600;
  color: #409eff;
  cursor: pointer;
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
  .merchant-codes {
    padding: 15px;
  }
  
  .search-card .el-form--inline .el-form-item {
    margin-right: 0;
    margin-bottom: 10px;
  }
}
</style>