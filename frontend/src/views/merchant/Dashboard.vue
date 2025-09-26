<template>
  <div class="merchant-dashboard">
    <!-- 商户统计卡片 -->
    <el-row :gutter="20" class="stats-cards">
      <el-col :xs="12" :sm="6" :md="6" :lg="6" :xl="6">
        <el-card class="stat-card">
          <div class="stat-content">
            <div class="stat-icon total-codes">
              <el-icon><Document /></el-icon>
            </div>
            <div class="stat-info">
              <div class="stat-value">{{ stats.totalCodes }}</div>
              <div class="stat-label">总防伪码数</div>
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :xs="12" :sm="6" :md="6" :lg="6" :xl="6">
        <el-card class="stat-card">
          <div class="stat-content">
            <div class="stat-icon used-codes">
              <el-icon><Check /></el-icon>
            </div>
            <div class="stat-info">
              <div class="stat-value">{{ stats.usedCodes }}</div>
              <div class="stat-label">已使用码数</div>
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :xs="12" :sm="6" :md="6" :lg="6" :xl="6">
        <el-card class="stat-card">
          <div class="stat-content">
            <div class="stat-icon today-verifies">
              <el-icon><Search /></el-icon>
            </div>
            <div class="stat-info">
              <div class="stat-value">{{ stats.todayVerifies }}</div>
              <div class="stat-label">今日验证数</div>
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :xs="12" :sm="6" :md="6" :lg="6" :xl="6">
        <el-card class="stat-card">
          <div class="stat-content">
            <div class="stat-icon valid-rate">
              <el-icon><SuccessFilled /></el-icon>
            </div>
            <div class="stat-info">
              <div class="stat-value">{{ stats.validRate }}%</div>
              <div class="stat-label">有效验证率</div>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <!-- 快速操作 -->
    <el-row :gutter="20" class="quick-actions">
      <el-col :xs="12" :sm="8" :md="6" :lg="4" :xl="4">
        <el-card class="action-card" shadow="hover" @click="navigateToGenerate">
          <div class="action-content">
            <div class="action-icon generate">
              <el-icon><Plus /></el-icon>
            </div>
            <div class="action-text">生成防伪码</div>
          </div>
        </el-card>
      </el-col>
      <el-col :xs="12" :sm="8" :md="6" :lg="4" :xl="4">
        <el-card class="action-card" shadow="hover" @click="navigateToCodes">
          <div class="action-content">
            <div class="action-icon manage">
              <el-icon><Document /></el-icon>
            </div>
            <div class="action-text">管理防伪码</div>
          </div>
        </el-card>
      </el-col>
      <el-col :xs="12" :sm="8" :md="6" :lg="4" :xl="4">
        <el-card class="action-card" shadow="hover" @click="navigateToVerifies">
          <div class="action-content">
            <div class="action-icon history">
              <el-icon><Search /></el-icon>
            </div>
            <div class="action-text">验证记录</div>
          </div>
        </el-card>
      </el-col>
      <el-col :xs="12" :sm="8" :md="6" :lg="4" :xl="4">
        <el-card class="action-card" shadow="hover" @click="navigateToProfile">
          <div class="action-content">
            <div class="action-icon profile">
              <el-icon><User /></el-icon>
            </div>
            <div class="action-text">商户信息</div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <!-- 验证趋势 -->
    <el-card class="verification-trend">
      <template #header>
        <span>验证趋势（最近7天）</span>
      </template>
      <div class="chart-container">
        <!-- 这里可以放置验证趋势图表 -->
        <div class="chart-placeholder">
          <el-icon><TrendCharts /></el-icon>
          <p>验证趋势图表</p>
        </div>
      </div>
    </el-card>

    <!-- 最近验证记录 -->
    <el-card class="recent-verifications">
      <template #header>
        <span>最近验证记录</span>
        <el-button type="primary" text @click="navigateToVerifies">查看全部</el-button>
      </template>
      <el-table :data="recentVerifications" style="width: 100%">
        <el-table-column prop="time" label="验证时间" width="180" />
        <el-table-column prop="code" label="防伪码" width="200" />
        <el-table-column prop="product" label="产品名称" />
        <el-table-column prop="result" label="验证结果" width="100">
          <template #default="scope">
            <el-tag :type="scope.row.result === '有效' ? 'success' : 'danger'">
              {{ scope.row.result }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="location" label="验证地点" />
      </el-table>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import {
  Document,
  Check,
  Search,
  SuccessFilled,
  Plus,
  User,
  TrendCharts
} from '@element-plus/icons-vue'

const router = useRouter()

const stats = ref({
  totalCodes: 0,
  usedCodes: 0,
  todayVerifies: 0,
  validRate: 0
})

const recentVerifications = ref([])

// 模拟数据
const mockData = {
  stats: {
    totalCodes: 12500,
    usedCodes: 8432,
    todayVerifies: 156,
    validRate: 97.8
  },
  recentVerifications: [
    {
      time: '2024-01-15 14:25:30',
      code: 'ABC123456789',
      product: '高端白酒A系列',
      result: '有效',
      location: '北京市朝阳区'
    },
    {
      time: '2024-01-15 13:45:12',
      code: 'DEF987654321',
      product: '化妆品B套装',
      result: '有效',
      location: '上海市浦东新区'
    },
    {
      time: '2024-01-15 12:30:45',
      code: 'GHI555666777',
      product: '电子产品C型号',
      result: '无效',
      location: '广州市天河区'
    },
    {
      time: '2024-01-15 11:15:20',
      code: 'JKL111222333',
      product: '服装D品牌',
      result: '有效',
      location: '深圳市南山区'
    }
  ]
}

// 导航方法
const navigateToGenerate = () => {
  router.push('/merchant/codes/generate')
}

const navigateToCodes = () => {
  router.push('/merchant/codes')
}

const navigateToVerifies = () => {
  router.push('/merchant/verifies')
}

const navigateToProfile = () => {
  router.push('/merchant/profile')
}

onMounted(() => {
  // 模拟API调用
  setTimeout(() => {
    stats.value = mockData.stats
    recentVerifications.value = mockData.recentVerifications
  }, 500)
})
</script>

<style scoped>
.merchant-dashboard {
  padding: 0;
}

.stats-cards {
  margin-bottom: 20px;
}

.stat-card {
  border-radius: 8px;
  border: none;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
}

.stat-content {
  display: flex;
  align-items: center;
  gap: 16px;
}

.stat-icon {
  width: 60px;
  height: 60px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 24px;
  color: white;
}

.stat-icon.total-codes {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

.stat-icon.used-codes {
  background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%);
}

.stat-icon.today-verifies {
  background: linear-gradient(135deg, #4facfe 0%, #00f2fe 100%);
}

.stat-icon.valid-rate {
  background: linear-gradient(135deg, #43e97b 0%, #38f9d7 100%);
}

.stat-info {
  flex: 1;
}

.stat-value {
  font-size: 24px;
  font-weight: bold;
  color: #303133;
  margin-bottom: 4px;
}

.stat-label {
  font-size: 14px;
  color: #909399;
}

.quick-actions {
  margin-bottom: 20px;
}

.action-card {
  border-radius: 8px;
  border: none;
  cursor: pointer;
  transition: transform 0.3s;
}

.action-card:hover {
  transform: translateY(-4px);
}

.action-content {
  text-align: center;
  padding: 20px 0;
}

.action-icon {
  width: 60px;
  height: 60px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0 auto 12px;
  font-size: 24px;
  color: white;
}

.action-icon.generate {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

.action-icon.manage {
  background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%);
}

.action-icon.history {
  background: linear-gradient(135deg, #4facfe 0%, #00f2fe 100%);
}

.action-icon.profile {
  background: linear-gradient(135deg, #43e97b 0%, #38f9d7 100%);
}

.action-text {
  font-size: 14px;
  font-weight: 500;
  color: #303133;
}

.verification-trend {
  border-radius: 8px;
  border: none;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
  margin-bottom: 20px;
}

.chart-container {
  height: 300px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.chart-placeholder {
  text-align: center;
  color: #909399;
}

.chart-placeholder .el-icon {
  font-size: 48px;
  margin-bottom: 16px;
}

.recent-verifications {
  border-radius: 8px;
  border: none;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
}

/* 响应式设计 */
@media (max-width: 768px) {
  .stats-cards {
    margin-bottom: 10px;
  }
  
  .stat-content {
    flex-direction: column;
    text-align: center;
    gap: 8px;
  }
  
  .stat-icon {
    width: 50px;
    height: 50px;
    font-size: 20px;
  }
  
  .stat-value {
    font-size: 20px;
  }
  
  .quick-actions {
    margin-bottom: 10px;
  }
  
  .action-content {
    padding: 15px 0;
  }
  
  .action-icon {
    width: 50px;
    height: 50px;
    font-size: 20px;
  }
  
  .chart-container {
    height: 200px;
  }
}
</style>