<template>
  <div class="platform-merchants">
    <div class="page-header">
      <h1>商户管理</h1>
      <p>管理所有注册商户信息</p>
    </div>

    <!-- 搜索和操作栏 -->
    <el-card class="search-card">
      <el-form :model="searchForm" inline>
        <el-form-item label="商户名称">
          <el-input
            v-model="searchForm.name"
            placeholder="请输入商户名称"
            clearable
          />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择状态" clearable>
            <el-option label="正常" value="active" />
            <el-option label="禁用" value="inactive" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">搜索</el-button>
          <el-button @click="handleReset">重置</el-button>
          <el-button type="success" @click="handleAdd">新增商户</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <!-- 商户列表 -->
    <el-card>
      <el-table :data="merchantList" stripe>
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="name" label="商户名称" min-width="120" />
        <el-table-column prop="contact" label="联系人" width="100" />
        <el-table-column prop="phone" label="联系电话" width="120" />
        <el-table-column prop="email" label="邮箱" min-width="150" />
        <el-table-column prop="codeCount" label="防伪码数量" width="100" align="center">
          <template #default="{ row }">
            <el-tag>{{ row.codeCount }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="verifyCount" label="验证次数" width="100" align="center">
          <template #default="{ row }">
            <el-tag type="success">{{ row.verifyCount }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="row.status === 'active' ? 'success' : 'danger'">
              {{ row.status === 'active' ? '正常' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="createTime" label="创建时间" width="160" />
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button size="small" @click="handleEdit(row)">编辑</el-button>
            <el-button
              size="small"
              :type="row.status === 'active' ? 'warning' : 'success'"
              @click="handleToggleStatus(row)"
            >
              {{ row.status === 'active' ? '禁用' : '启用' }}
            </el-button>
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
  name: '',
  status: ''
})

const merchantList = ref([])
const pagination = reactive({
  current: 1,
  size: 10,
  total: 0
})

// 模拟商户数据
const mockMerchants = [
  {
    id: 1001,
    name: '茅台集团',
    contact: '张经理',
    phone: '13800138001',
    email: 'zhang@maotai.com',
    codeCount: 50000,
    verifyCount: 12560,
    status: 'active',
    createTime: '2024-01-01 10:00:00'
  },
  {
    id: 1002,
    name: '五粮液股份',
    contact: '李总监',
    phone: '13800138002',
    email: 'li@wuliangye.com',
    codeCount: 30000,
    verifyCount: 8560,
    status: 'active',
    createTime: '2024-01-02 14:30:00'
  },
  {
    id: 1003,
    name: '洋河酒厂',
    contact: '王主管',
    phone: '13800138003',
    email: 'wang@yanghe.com',
    codeCount: 20000,
    verifyCount: 4230,
    status: 'inactive',
    createTime: '2024-01-03 09:15:00'
  }
]

const loadMerchants = () => {
  // 模拟API调用
  setTimeout(() => {
    merchantList.value = mockMerchants
    pagination.total = mockMerchants.length
  }, 500)
}

const handleSearch = () => {
  pagination.current = 1
  loadMerchants()
}

const handleReset = () => {
  Object.keys(searchForm).forEach(key => {
    searchForm[key] = ''
  })
  pagination.current = 1
  loadMerchants()
}

const handleAdd = () => {
  ElMessage.info('新增商户功能开发中')
}

const handleEdit = (row) => {
  ElMessage.info(`编辑商户: ${row.name}`)
}

const handleToggleStatus = async (row) => {
  const action = row.status === 'active' ? '禁用' : '启用'
  
  try {
    await ElMessageBox.confirm(
      `确定要${action}商户 "${row.name}" 吗？`,
      '提示',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }
    )
    
    row.status = row.status === 'active' ? 'inactive' : 'active'
    ElMessage.success(`${action}成功`)
  } catch (error) {
    // 用户取消操作
  }
}

const handleDelete = async (row) => {
  try {
    await ElMessageBox.confirm(
      `确定要删除商户 "${row.name}" 吗？此操作不可恢复。`,
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
  loadMerchants()
}

const handleCurrentChange = (current) => {
  pagination.current = current
  loadMerchants()
}

onMounted(() => {
  loadMerchants()
})
</script>

<style scoped>
.platform-merchants {
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

.pagination {
  margin-top: 20px;
  text-align: right;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .platform-merchants {
    padding: 15px;
  }
  
  .search-card .el-form--inline .el-form-item {
    margin-right: 0;
    margin-bottom: 10px;
  }
}
</style>