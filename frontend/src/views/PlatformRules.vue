<template>
  <div class="platform-rules">
    <div class="page-header">
      <h1>规则配置</h1>
      <p>管理系统防伪码生成和验证规则</p>
    </div>

    <!-- 规则配置表单 -->
    <el-card>
      <template #header>
        <span>防伪码规则配置</span>
      </template>
      
      <el-form :model="ruleForm" label-width="120px">
        <el-form-item label="防伪码长度">
          <el-input-number
            v-model="ruleForm.codeLength"
            :min="8"
            :max="32"
            controls-position="right"
          />
          <span class="form-tip">建议长度：12-20位</span>
        </el-form-item>
        
        <el-form-item label="字符类型">
          <el-checkbox-group v-model="ruleForm.charTypes">
            <el-checkbox label="numbers">数字</el-checkbox>
            <el-checkbox label="uppercase">大写字母</el-checkbox>
            <el-checkbox label="lowercase">小写字母</el-checkbox>
          </el-checkbox-group>
        </el-form-item>
        
        <el-form-item label="有效期设置">
          <el-radio-group v-model="ruleForm.expiryType">
            <el-radio label="never">永不过期</el-radio>
            <el-radio label="custom">自定义</el-radio>
          </el-radio-group>
          <el-input-number
            v-if="ruleForm.expiryType === 'custom'"
            v-model="ruleForm.expiryDays"
            :min="1"
            :max="3650"
            controls-position="right"
            style="margin-left: 10px;"
          />
          <span v-if="ruleForm.expiryType === 'custom'" class="form-tip">天</span>
        </el-form-item>
        
        <el-form-item label="验证限制">
          <el-input-number
            v-model="ruleForm.maxVerifyTimes"
            :min="1"
            :max="100"
            controls-position="right"
          />
          <span class="form-tip">次（0表示不限制）</span>
        </el-form-item>
        
        <el-form-item label="批量生成限制">
          <el-input-number
            v-model="ruleForm.batchLimit"
            :min="100"
            :max="100000"
            controls-position="right"
          />
          <span class="form-tip">个/批次</span>
        </el-form-item>
        
        <el-form-item>
          <el-button type="primary" @click="handleSave">保存配置</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { ElMessage } from 'element-plus'

const ruleForm = reactive({
  codeLength: 16,
  charTypes: ['numbers', 'uppercase', 'lowercase'],
  expiryType: 'never',
  expiryDays: 365,
  maxVerifyTimes: 10,
  batchLimit: 10000
})

const handleSave = () => {
  ElMessage.success('规则配置保存成功')
  // 这里应该调用API保存配置
}

const handleReset = () => {
  Object.assign(ruleForm, {
    codeLength: 16,
    charTypes: ['numbers', 'uppercase', 'lowercase'],
    expiryType: 'never',
    expiryDays: 365,
    maxVerifyTimes: 10,
    batchLimit: 10000
  })
  ElMessage.info('配置已重置')
}
</script>

<style scoped>
.platform-rules {
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

/* 响应式设计 */
@media (max-width: 768px) {
  .platform-rules {
    padding: 15px;
  }
  
  .el-form-item {
    margin-bottom: 20px;
  }
  
  .el-form-item__label {
    width: 100px !important;
  }
}
</style>