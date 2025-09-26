<template>
  <div class="generate-codes">
    <el-card>
      <template #header>
        <span>生成防伪码</span>
      </template>

      <el-form ref="generateForm" :model="generateForm" :rules="generateRules" label-width="120px">
        <el-form-item label="产品名称" prop="product">
          <el-input v-model="generateForm.product" placeholder="请输入产品名称" style="width: 300px" />
        </el-form-item>
        
        <el-form-item label="生成数量" prop="quantity">
          <el-input-number 
            v-model="generateForm.quantity" 
            :min="1" 
            :max="10000" 
            placeholder="请输入生成数量"
          />
          <span class="form-tip">单次最多生成10000个防伪码</span>
        </el-form-item>
        
        <el-form-item label="批次号" prop="batch">
          <el-input v-model="generateForm.batch" placeholder="请输入批次号" style="width: 300px" />
          <span class="form-tip">如不填写将自动生成</span>
        </el-form-item>
        
        <el-form-item label="有效期" prop="expiryDate">
          <el-date-picker
            v-model="generateForm.expiryDate"
            type="date"
            placeholder="选择有效期"
            :disabled-date="disabledDate"
            style="width: 300px"
          />
          <span class="form-tip">如不选择将永久有效</span>
        </el-form-item>
        
        <el-form-item label="防伪码格式" prop="format">
          <el-radio-group v-model="generateForm.format">
            <el-radio label="standard">标准格式 (字母+数字)</el-radio>
            <el-radio label="custom">自定义格式</el-radio>
          </el-radio-group>
        </el-form-item>
        
        <el-form-item v-if="generateForm.format === 'custom'" label="自定义格式" prop="customFormat">
          <el-input v-model="generateForm.customFormat" placeholder="例如：ABC{6}XYZ{4}" style="width: 300px" />
          <div class="format-examples">
            <p>格式说明：</p>
            <ul>
              <li><code>A</code> - 大写字母</li>
              <li><code>a</code> - 小写字母</li>
              <li><code>0</code> - 数字</li>
              <li><code>{n}</code> - 重复n次</li>
            </ul>
            <p>示例：<code>ABC{6}XYZ{4}</code> 生成类似 ABC123456XYZ7890</p>
          </div>
        </el-form-item>
        
        <el-form-item>
          <el-button type="primary" @click="handleGenerate" :loading="generating">
            {{ generating ? '生成中...' : '立即生成' }}
          </el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <!-- 生成结果 -->
    <el-card v-if="generationResult" class="result-card">
      <template #header>
        <span>生成结果</span>
        <el-button type="success" @click="handleDownload" style="float: right">
          <el-icon><Download /></el-icon>
          下载防伪码列表
        </el-button>
      </template>
      
      <el-alert
        title="生成成功"
        :description="`成功生成 ${generationResult.total} 个防伪码，批次号：${generationResult.batch}`"
        type="success"
        show-icon
        :closable="false"
      />
      
      <div class="preview-codes">
        <h4>预览（前10个）：</h4>
        <el-table :data="generationResult.preview" style="width: 100%">
          <el-table-column prop="index" label="序号" width="60" />
          <el-table-column prop="code" label="防伪码" />
        </el-table>
      </div>
    </el-card>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { ElMessage } from 'element-plus'
import { Download } from '@element-plus/icons-vue'

const generateForm = reactive({
  product: '',
  quantity: 100,
  batch: '',
  expiryDate: null,
  format: 'standard',
  customFormat: ''
})

const generateRules = {
  product: [{ required: true, message: '请输入产品名称', trigger: 'blur' }],
  quantity: [
    { required: true, message: '请输入生成数量', trigger: 'blur' },
    { type: 'number', min: 1, max: 10000, message: '数量必须在1-10000之间', trigger: 'blur' }
  ]
}

const generating = ref(false)
const generationResult = ref(null)

// 方法
const disabledDate = (time) => {
  return time.getTime() < Date.now() - 8.64e7 // 禁用今天之前的日期
}

const handleGenerate = () => {
  // 表单验证
  // 这里应该调用表单验证方法
  
  generating.value = true
  
  // 模拟API调用
  setTimeout(() => {
    const batch = generateForm.batch || `BATCH${Date.now()}`
    const total = generateForm.quantity
    const preview = Array.from({ length: Math.min(10, total) }, (_, index) => ({
      index: index + 1,
      code: generateCode()
    }))
    
    generationResult.value = {
      batch,
      total,
      preview
    }
    
    generating.value = false
    ElMessage.success(`成功生成 ${total} 个防伪码`)
  }, 2000)
}

const handleReset = () => {
  Object.assign(generateForm, {
    product: '',
    quantity: 100,
    batch: '',
    expiryDate: null,
    format: 'standard',
    customFormat: ''
  })
  generationResult.value = null
}

const handleDownload = () => {
  ElMessage.success('下载功能开发中')
}

const generateCode = () => {
  if (generateForm.format === 'custom' && generateForm.customFormat) {
    return generateCustomCode()
  }
  
  // 标准格式：3位字母 + 6位数字 + 3位字母 + 4位数字
  const letters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'
  const numbers = '0123456789'
  
  let code = ''
  for (let i = 0; i < 3; i++) {
    code += letters.charAt(Math.floor(Math.random() * letters.length))
  }
  for (let i = 0; i < 6; i++) {
    code += numbers.charAt(Math.floor(Math.random() * numbers.length))
  }
  for (let i = 0; i < 3; i++) {
    code += letters.charAt(Math.floor(Math.random() * letters.length))
  }
  for (let i = 0; i < 4; i++) {
    code += numbers.charAt(Math.floor(Math.random() * numbers.length))
  }
  
  return code
}

const generateCustomCode = () => {
  // 简单的自定义格式生成器
  // 实际项目中应该实现更复杂的格式解析
  const format = generateForm.customFormat
  let code = ''
  
  for (let i = 0; i < format.length; i++) {
    const char = format[i]
    
    if (char === '{') {
      // 解析重复次数
      let repeatStr = ''
      i++
      while (format[i] !== '}' && i < format.length) {
        repeatStr += format[i]
        i++
      }
      const repeat = parseInt(repeatStr) || 1
      
      // 获取前一个字符作为模板
      const templateChar = code[code.length - 1]
      if (templateChar) {
        for (let j = 0; j < repeat; j++) {
          code += generateChar(templateChar)
        }
      }
    } else {
      code += generateChar(char)
    }
  }
  
  return code
}

const generateChar = (template) => {
  const lettersUpper = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'
  const lettersLower = 'abcdefghijklmnopqrstuvwxyz'
  const numbers = '0123456789'
  
  switch (template) {
    case 'A':
      return lettersUpper.charAt(Math.floor(Math.random() * lettersUpper.length))
    case 'a':
      return lettersLower.charAt(Math.floor(Math.random() * lettersLower.length))
    case '0':
      return numbers.charAt(Math.floor(Math.random() * numbers.length))
    default:
      return template
  }
}
</script>

<style scoped>
.generate-codes {
  padding: 0;
}

.form-tip {
  margin-left: 10px;
  color: #909399;
  font-size: 12px;
}

.format-examples {
  margin-top: 10px;
  padding: 10px;
  background-color: #f5f7fa;
  border-radius: 4px;
  font-size: 12px;
}

.format-examples p {
  margin: 5px 0;
}

.format-examples ul {
  margin: 5px 0;
  padding-left: 20px;
}

.format-examples code {
  background-color: #e8f4fd;
  padding: 2px 4px;
  border-radius: 3px;
  font-family: 'Courier New', monospace;
}

.result-card {
  margin-top: 20px;
}

.preview-codes {
  margin-top: 20px;
}

.preview-codes h4 {
  margin-bottom: 10px;
  color: #303133;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .el-form-item {
    margin-bottom: 20px;
  }
  
  .el-form-item__content {
    width: 100%;
  }
  
  .el-input, .el-date-picker {
    width: 100% !important;
  }
}
</style>