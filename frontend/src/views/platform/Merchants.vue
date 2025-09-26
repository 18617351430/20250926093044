<template>
  <div class="merchants-management">
    <!-- 搜索和操作栏 -->
    <el-card class="search-bar">
      <el-form :inline="true" :model="searchForm">
        <el-form-item label="商户名称">
          <el-input v-model="searchForm.name" placeholder="请输入商户名称" clearable />
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
      <template #header>
        <span>商户列表</span>
      </template>
      <el-table :data="merchantList" style="width: 100%">
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="name" label="商户名称" min-width="120" />
        <el-table-column prop="contact" label="联系人" width="100" />
        <el-table-column prop="phone" label="联系电话" width="120" />
        <el-table-column prop="email" label="邮箱" min-width="150" />
        <el-table-column prop="totalCodes" label="防伪码总数" width="100" />
        <el-table-column prop="status" label="状态" width="80">
          <template #default="scope">
            <el-tag :type="scope.row.status === 'active' ? 'success' : 'danger'">
              {{ scope.row.status === 'active' ? '正常' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="createTime" label="创建时间" width="160" />
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="scope">
            <el-button size="small" @click="handleEdit(scope.row)">编辑</el-button>
            <el-button 
              size="small" 
              :type="scope.row.status === 'active' ? 'warning' : 'success'"
              @click="handleToggleStatus(scope.row)"
            >
              {{ scope.row.status === 'active' ? '禁用' : '启用' }}
            </el-button>
            <el-button size="small" type="danger" @click="handleDelete(scope.row)">删除</el-button>
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

    <!-- 新增/编辑商户对话框 -->
    <el-dialog
      v-model="dialogVisible"
      :title="dialogTitle"
      width="500px"
      :before-close="handleDialogClose"
    >
      <el-form ref="merchantForm" :model="merchantForm" :rules="merchantRules" label-width="80px">
        <el-form-item label="商户名称" prop="name">
          <el-input v-model="merchantForm.name" placeholder="请输入商户名称" />
        </el-form-item>
        <el-form-item label="联系人" prop="contact">
          <el-input v-model="merchantForm.contact" placeholder="请输入联系人" />
        </el-form-item>
        <el-form-item label="联系电话" prop="phone">
          <el-input v-model="merchantForm.phone" placeholder="请输入联系电话" />
        </el-form-item>
        <el-form-item label="邮箱" prop="email">
          <el-input v-model="merchantForm.email" placeholder="请输入邮箱" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="merchantForm.status">
            <el-radio label="active">正常</el-radio>
            <el-radio label="inactive">禁用</el-radio>
          </el-radio-group>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="handleDialogClose">取消</el-button>
          <el-button type="primary" @click="handleSubmit">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted } from 'vue'
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

const dialogVisible = ref(false)
const isEdit = ref(false)
const merchantForm = reactive({
  id: '',
  name: '',
  contact: '',
  phone: '',
  email: '',
  status: 'active'
})

const merchantRules = {
  name: [{ required: true, message: '请输入商户名称', trigger: 'blur' }],
  contact: [{ required: true, message: '请输入联系人', trigger: 'blur' }],
  phone: [{ required: true, message: '请输入联系电话', trigger: 'blur' }],
  email: [
    { required: true, message: '请输入邮箱', trigger: 'blur' },
    { type: 'email', message: '请输入正确的邮箱地址', trigger: 'blur' }
  ]
}

// 计算属性
const dialogTitle = computed(() => {
  return isEdit.value ? '编辑商户' : '新增商户'
})

// 模拟数据
const mockMerchants = Array.from({ length: 50 }, (_, index) => ({
  id: index + 1,
  name: `商户${index + 1}`,
  contact: `联系人${index + 1}`,
  phone: `138${String(index).padStart(8, '0')}`,
  email: `merchant${index + 1}@example.com`,
  totalCodes: Math.floor(Math.random() * 10000),
  status: index % 5 === 0 ? 'inactive' : 'active',
  createTime: `2024-01-${String(15 - index % 15).padStart(2, '0')} 10:00:00`
}))

// 方法
const handleSearch = () => {
  pagination.current = 1
  loadMerchants()
}

const handleReset = () => {
  Object.assign(searchForm, {
    name: '',
    status: ''
  })
  pagination.current = 1
  loadMerchants()
}

const handleAdd = () => {
  isEdit.value = false
  Object.assign(merchantForm, {
    id: '',
    name: '',
    contact: '',
    phone: '',
    email: '',
    status: 'active'
  })
  dialogVisible.value = true
}

const handleEdit = (row) => {
  isEdit.value = true
  Object.assign(merchantForm, { ...row })
  dialogVisible.value = true
}

const handleToggleStatus = async (row) => {
  const newStatus = row.status === 'active' ? 'inactive' : 'active'
  const action = newStatus === 'active' ? '启用' : '禁用'
  
  try {
    await ElMessageBox.confirm(
      `确定要${action}商户"${row.name}"吗？`,
      '提示',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }
    )
    
    // 模拟API调用
    row.status = newStatus
    ElMessage.success(`${action}成功`)
  } catch {
    // 用户取消操作
  }
}

const handleDelete = async (row) => {
  try {
    await ElMessageBox.confirm(
      `确定要删除商户"${row.name}"吗？此操作不可恢复！`,
      '警告',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'error'
      }
    )
    
    // 模拟API调用
    merchantList.value = merchantList.value.filter(item => item.id !== row.id)
    ElMessage.success('删除成功')
  } catch {
    // 用户取消操作
  }
}

const handleDialogClose = () => {
  dialogVisible.value = false
}

const handleSubmit = () => {
  // 表单验证
  // 这里应该调用表单验证方法
  
  if (isEdit.value) {
    // 编辑商户
    const index = merchantList.value.findIndex(item => item.id === merchantForm.id)
    if (index !== -1) {
      merchantList.value[index] = { ...merchantForm }
    }
    ElMessage.success('编辑成功')
  } else {
    // 新增商户
    const newMerchant = {
      ...merchantForm,
      id: merchantList.value.length + 1,
      totalCodes: 0,
      createTime: new Date().toLocaleString()
    }
    merchantList.value.unshift(newMerchant)
    ElMessage.success('新增成功')
  }
  
  dialogVisible.value = false
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

const loadMerchants = () => {
  // 模拟API调用
  const filtered = mockMerchants.filter(merchant => {
    const nameMatch = !searchForm.name || merchant.name.includes(searchForm.name)
    const statusMatch = !searchForm.status || merchant.status === searchForm.status
    return nameMatch && statusMatch
  })
  
  const start = (pagination.current - 1) * pagination.size
  const end = start + pagination.size
  merchantList.value = filtered.slice(start, end)
  pagination.total = filtered.length
}

onMounted(() => {
  loadMerchants()
})
</script>

<style scoped>
.merchants-management {
  padding: 0;
}

.search-bar {
  margin-bottom: 20px;
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
}
</style>