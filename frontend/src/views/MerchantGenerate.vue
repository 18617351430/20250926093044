<template>
  <div class="merchant-generate">
    <div class="page-header">
      <h1>生成防伪码</h1>
      <p>批量生成新的防伪码</p>
    </div>

    <!-- 生成表单 -->
    <el-card>
      <template #header>
        <span>防伪码生成配置</span>
      </template>
      
      <el-form :model="generateForm" label-width="120px">
        <el-form-item label="生成数量">
          <el-input-number
            v-model="generateForm.quantity"
            :min="1"
            :max="10000"
            controls-position="right"
          />
          <span class="form-tip">最大支持10000个/批次</span>
        </el-form-item>
        
        <el-form-item label="批次名称">
          <el-input
            v-model="generateForm.batchName"
            placeholder="请输入批次名称"
            maxlength="50"
            show-word-limit
          />
        </el-form-item>
        
        <el-form-item label="有效期设置">
          <el-radio-group v-model="generateForm.expiryType">
            <el-radio label="never">永不过期</el-radio>
            <el-radio label="custom">自定义</el-radio>
          </el-radio-group>
          <el-date-picker
            v-if="generateForm.expiryType === 'custom'"
            v-model="generateForm.expiryDate"
            type="date"
            placeholder="选择过期日期"
            value-format="YYYY-MM-DD"
            style="margin-left: 10px; width: 200px;"
          />
        </el-form-item>
        
        <el-form-item label="防伪码格式">
          <el-radio-group v-model="generateForm.format">
            <el-radio label="mixed">数字+字母混合</el-radio>
            <el-radio label="numbers">纯数字</el-radio>
            <el-radio label="letters">纯字母</el-radio>
          </el-radio-group>
        </el-form-item>
        
        <el-form-item label="码长度">
          <el-input-number
            v-model="generateForm.length"
            :min="8"
            :max="32"
            controls-position="right"
          />
        </el-form-item>
        
        <el-form-item>
          <el-button type="primary" @click="handleGenerate" :loading="generating">
            {{ generating ? '生成中...' : '开始生成' }}
          </el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <!-- 生成结果 -->
    <el-card v-if="generationResult" class="result-card">
      <template #header>
        <span>生成结果</span>
        <el-button type="success" size="small" @click="handleDownload">下载文件</el-button>
      </template>
      
      <div class="result-info">
        <el-alert
          title="生成成功"
          type="success"
          :description="`成功生成 ${generationResult.successCount} 个防伪码，批次号：${generationResult.batchNo}`"
          show-icon
          :closable="false"
        />
        
        <div class="code-preview">
          <h4>防伪码预览（前10个）：</h4>
          <div class="code-list">
            <div
              v-for="(code, index) in generationResult.previewCodes"
              :key="index"
              class="code-item"
            >
              {{ code }}
            </div>
          </div>
        </div>
      </div>
    </el-card>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { ElMessage } from 'element-plus'

const generateForm = reactive({
  quantity: 100,
  batchName: '',
  expiryType: 'never',
  expiryDate: '',
  format: 'mixed',
  length: 16
})

const generating = ref(false)
const generationResult = ref(null)

const generateRandomCode = (length, format) => {
  const numbers = '0123456789'
  const letters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'
  const mixed = numbers + letters
  
  let charset = mixed
  if (format === 'numbers') charset = numbers
  if (format === 'letters') charset = letters
  
  let result = ''
  for (let i = 0; i < length; i++) {
    result += charset.charAt(Math.floor(Math.random() * charset.length))
  }
  return result
}

const handleGenerate = async () => {
  if (generateForm.quantity <= 0) {
    ElMessage.error('请填写正确的生成数量')
    return
  }
  
  if (!generateForm.batchName.trim()) {
    ElMessage.error('请填写批次名称')
    return
  }
  
  generating.value = true
  
  // 模拟生成过程
  setTimeout(() => {
    const codes = []
    for (let i = 0; i < generateForm.quantity; i++) {
      codes.push(generateRandomCode(generateForm.length, generateForm.format))
    }
    
    generationResult.value = {
      successCount: generateForm.quantity,
      batchNo: `BATCH${Date.now()}`,
      previewCodes: codes.slice(0, 10)
    }
    
    generating.value = false
    ElMessage.success('防伪码生成成功')
  }, 2000)
}

const handleReset = () => {
  Object.assign(generateForm, {
    quantity: 100,
    batchName: '',
    expiryType: 'never',
    expiryDate: '',
    format: 'mixed',
    length: 16
  })
  generationResult.value = null
}

const handleDownload = () => {
  ElMessage.info('下载功能开发中')
}
</script>

<style scoped>
.merchant-generate {
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

.form-tip {
  margin-left: 10px;
  color: #999;
  font-size: 12px;
}

.result-card {
  margin-top: 20px;
}

.result-info {
  padding: 10px 0;
}

.code-preview {
  margin-top: 20px;
}

.code-preview h4 {
  margin: 0 0 10px 0;
  font-size: 14px;
  color: #666;
}

.code-list {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  gap: 10px;
  margin-top: 10px;
}

.code-item {
  padding: 8px 12px;
  background: #f5f7fa;
  border-radius: 4px;
  font-family: 'Courier New', monospace;
  font-size: 12px;
  text-align: center;
  border: 1px solid #e4e7ed;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .merchant-generate {
    padding: 15px;
  }
  
  .code-list {
    grid-template-columns: 1fr;
  }
  
  .el-form-item {
    margin-bottom: 20px;
  }
  
  .el-form-item__label {
    width: 100px !important;
  }
}
</style>