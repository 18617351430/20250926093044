<template>
  <div class="verify-page">
    <div class="verify-container">
      <!-- 验证区域 -->
      <el-card class="verify-card" shadow="hover">
        <template #header>
          <div class="verify-header">
            <h2>防伪码验证</h2>
            <p>请输入防伪码进行验证</p>
          </div>
        </template>

        <!-- 验证表单 -->
        <el-form :model="verifyForm" :rules="verifyRules" ref="verifyFormRef">
          <el-form-item prop="code">
            <el-input
              v-model="verifyForm.code"
              placeholder="请输入16位防伪码"
              size="large"
              clearable
              @keyup.enter="handleVerify"
            >
              <template #prefix>
                <el-icon><search /></el-icon>
              </template>
            </el-input>
          </el-form-item>
          
          <el-form-item>
            <el-button 
              type="primary" 
              size="large" 
              @click="handleVerify"
              :loading="verifying"
              style="width: 100%"
            >
              {{ verifying ? '验证中...' : '立即验证' }}
            </el-button>
          </el-form-item>
        </el-form>

        <!-- 验证结果 -->
        <div v-if="verifyResult" class="verify-result">
          <div class="result-card" :class="resultClass">
            <div class="result-icon">
              <el-icon :size="48">
                <component :is="resultIcon" />
              </el-icon>
            </div>
            <div class="result-content">
              <h3>{{ resultTitle }}</h3>
              <p>{{ resultMessage }}</p>
              <div v-if="verifyResult.details" class="result-details">
                <el-descriptions :column="1" border>
                  <el-descriptions-item label="批次号">
                    {{ verifyResult.details.batchNo }}
                  </el-descriptions-item>
                  <el-descriptions-item label="商户名称">
                    {{ verifyResult.details.merchantName }}
                  </el-descriptions-item>
                  <el-descriptions-item label="验证时间">
                    {{ verifyResult.details.verifyTime }}
                  </el-descriptions-item>
                  <el-descriptions-item label="验证次数">
                    {{ verifyResult.details.verifyCount }}
                  </el-descriptions-item>
                </el-descriptions>
              </div>
            </div>
          </div>
        </div>
      </el-card>

      <!-- 批量验证 -->
      <el-card class="batch-verify-card" shadow="hover">
        <template #header>
          <div class="batch-header">
            <h3>批量验证</h3>
            <el-button type="text" @click="showBatchDialog = true">
              上传文件批量验证
            </el-button>
          </div>
        </template>

        <el-upload
          class="batch-upload"
          drag
          action="#"
          :auto-upload="false"
          :on-change="handleBatchUpload"
          :show-file-list="false"
        >
          <el-icon class="el-icon--upload"><upload-filled /></el-icon>
          <div class="el-upload__text">
            将文件拖到此处，或<em>点击上传</em>
          </div>
          <template #tip>
            <div class="el-upload__tip">
              支持 .txt、.csv 格式文件，每行一个防伪码
            </div>
          </template>
        </el-upload>
      </el-card>
    </div>

    <!-- 批量验证对话框 -->
    <el-dialog
      v-model="showBatchDialog"
      title="批量验证结果"
      width="600px"
    >
      <div v-if="batchResults.length > 0">
        <el-table :data="batchResults" stripe height="300">
          <el-table-column prop="code" label="防伪码" width="120">
            <template #default="{ row }">
              <span class="code-text">{{ row.code }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="status" label="状态" width="80">
            <template #default="{ row }">
              <el-tag :type="row.valid ? 'success' : 'danger'">
                {{ row.valid ? '有效' : '无效' }}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="message" label="结果信息" />
        </el-table>
        
        <div class="batch-summary">
          <el-statistic
            title="验证统计"
            :value="batchResults.filter(r => r.valid).length"
            :value-style="{ color: '#3f8600' }"
          >
            <template #suffix>
              / {{ batchResults.length }}
            </template>
          </el-statistic>
        </div>
      </div>
      <div v-else class="empty-batch">
        <el-empty description="暂无验证结果" />
      </div>
      
      <template #footer>
        <el-button @click="showBatchDialog = false">关闭</el-button>
        <el-button type="primary" @click="exportBatchResults">导出结果</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, computed } from 'vue'
import { ElMessage } from 'element-plus'
import { Search, UploadFilled, CircleCheck, CircleClose, Warning } from '@element-plus/icons-vue'

const verifyFormRef = ref()
const verifying = ref(false)
const showBatchDialog = ref(false)

const verifyForm = reactive({
  code: ''
})

const verifyResult = ref(null)

const verifyRules = {
  code: [
    { required: true, message: '请输入防伪码', trigger: 'blur' },
    { min: 16, max: 16, message: '防伪码必须为16位', trigger: 'blur' }
  ]
}

const batchResults = ref([])

const resultClass = computed(() => {
  if (!verifyResult.value) return ''
  return verifyResult.value.valid ? 'valid' : 'invalid'
})

const resultIcon = computed(() => {
  if (!verifyResult.value) return Warning
  return verifyResult.value.valid ? CircleCheck : CircleClose
})

const resultTitle = computed(() => {
  if (!verifyResult.value) return ''
  return verifyResult.value.valid ? '验证成功' : '验证失败'
})

const resultMessage = computed(() => {
  if (!verifyResult.value) return ''
  return verifyResult.value.message
})

const handleVerify = async () => {
  try {
    await verifyFormRef.value.validate()
    verifying.value = true

    // 模拟API调用
    setTimeout(() => {
      const mockResults = [
        {
          valid: true,
          message: '该防伪码为有效码，请放心使用',
          details: {
            batchNo: 'BATCH2024001',
            merchantName: '示例商户有限公司',
            verifyTime: '2024-09-26 10:15:30',
            verifyCount: 1
          }
        },
        {
          valid: false,
          message: '该防伪码不存在或已失效'
        },
        {
          valid: false,
          message: '该防伪码已被多次验证，可能存在风险'
        }
      ]

      const randomResult = mockResults[Math.floor(Math.random() * mockResults.length)]
      verifyResult.value = randomResult
      verifying.value = false

      ElMessage.success('验证完成')
    }, 1000)
  } catch (error) {
    verifying.value = false
  }
}

const handleBatchUpload = (file) => {
  // 模拟批量验证
  const mockBatchResults = [
    { code: 'ABCD1234EFGH5678', valid: true, message: '验证成功' },
    { code: 'EFGH5678IJKL9012', valid: false, message: '防伪码不存在' },
    { code: 'IJKL9012MNOP3456', valid: true, message: '验证成功' },
    { code: 'MNOP3456QRST7890', valid: false, message: '已超过验证次数' }
  ]

  batchResults.value = mockBatchResults
  showBatchDialog.value = true
  ElMessage.success('批量验证完成')
}

const exportBatchResults = () => {
  ElMessage.info('导出功能开发中')
}
</script>

<style scoped>
.verify-page {
  min-height: 100vh;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  padding: 40px 20px;
  display: flex;
  justify-content: center;
  align-items: flex-start;
}

.verify-container {
  max-width: 800px;
  width: 100%;
}

.verify-card {
  margin-bottom: 20px;
  border-radius: 12px;
}

.verify-header {
  text-align: center;
}

.verify-header h2 {
  margin: 0 0 8px 0;
  color: #333;
  font-size: 24px;
}

.verify-header p {
  margin: 0;
  color: #666;
  font-size: 14px;
}

.verify-result {
  margin-top: 20px;
}

.result-card {
  padding: 20px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  gap: 16px;
}

.result-card.valid {
  background: #f0f9ff;
  border: 1px solid #91d5ff;
}

.result-card.invalid {
  background: #fff2f0;
  border: 1px solid #ffccc7;
}

.result-icon {
  flex-shrink: 0;
}

.result-content h3 {
  margin: 0 0 8px 0;
  font-size: 18px;
}

.result-content p {
  margin: 0 0 12px 0;
  color: #666;
}

.result-details {
  margin-top: 12px;
}

.batch-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.batch-upload {
  width: 100%;
}

.batch-summary {
  margin-top: 16px;
  text-align: center;
}

.code-text {
  font-family: 'Courier New', monospace;
  font-size: 12px;
}

.empty-batch {
  padding: 40px 0;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .verify-page {
    padding: 20px 15px;
  }
  
  .result-card {
    flex-direction: column;
    text-align: center;
  }
}
</style>