<template>
  <div class="not-found">
    <div class="not-found-content">
      <div class="not-found-icon">
        <el-icon :size="120"><warning /></el-icon>
      </div>
      <h1>404</h1>
      <p>抱歉，您访问的页面不存在</p>
      <div class="not-found-actions">
        <el-button type="primary" @click="goHome">返回首页</el-button>
        <el-button @click="goBack">返回上页</el-button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'
import { Warning } from '@element-plus/icons-vue'

const router = useRouter()

const goHome = () => {
  router.push('/')
}

const goBack = () => {
  router.back()
}
</script>

<style scoped>
.not-found {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  padding: 20px;
}

.not-found-content {
  text-align: center;
  background: white;
  padding: 60px 40px;
  border-radius: 10px;
  box-shadow: 0 15px 35px rgba(0, 0, 0, 0.1);
  max-width: 500px;
  width: 100%;
}

.not-found-icon {
  margin-bottom: 30px;
  color: #e6a23c;
}

.not-found h1 {
  font-size: 72px;
  font-weight: 700;
  margin: 0 0 20px 0;
  color: #333;
}

.not-found p {
  font-size: 18px;
  color: #666;
  margin-bottom: 30px;
}

.not-found-actions {
  display: flex;
  gap: 15px;
  justify-content: center;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .not-found-content {
    padding: 40px 20px;
  }
  
  .not-found h1 {
    font-size: 48px;
  }
  
  .not-found-actions {
    flex-direction: column;
  }
}
</style>