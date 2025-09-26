<template>
  <div class="dashboard">
    <!-- 统计卡片 -->
    <el-row :gutter="20" class="stats-cards">
      <el-col :xs="12" :sm="6" :md="6" :lg="6" :xl="6">
        <el-card class="stat-card">
          <div class="stat-content">
            <div class="stat-icon total-merchants">
              <el-icon><User /></el-icon>
            </div>
            <div class="stat-info">
              <div class="stat-value">{{ stats.totalMerchants }}</div>
              <div class="stat-label">总商户数</div>
            </div>
          </div>
        </el-card>
      </el-col>
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

    <!-- 图表区域 -->
    <el-row :gutter="20" class="charts">
      <el-col :xs="24" :sm="24" :md="12" :lg="12" :xl="12">
        <el-card class="chart-card">
          <template #header>
            <span>验证趋势</span>
          </template>
          <div class="chart-container">
            <!-- 这里可以放置验证趋势图表 -->
            <div class="chart-placeholder">
              <el-icon><TrendCharts /></el-icon>
              <p>验证趋势图表</p>
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="24" :md="12" :lg="12" :xl="12">
        <el-card class="chart-card">
          <template #header>
            <span>商户分布</span>
          </template>
          <div class="chart-container">
            <!-- 这里可以放置商户分布图表 -->
            <div class="chart-placeholder">
              <el-icon><PieChart /></el-icon>
              <p>商户分布图表</p>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <!-- 最近活动 -->
    <el-card class="recent-activities">
      <template #header>
        <span>最近活动</span>
      </template>
      <el-table :data="recentActivities" style="width: 100%">
        <el-table-column prop="time" label="时间" width="180" />
        <el-table-column prop="merchant" label="商户" width="120" />
        <el-table-column prop="action" label="操作" />
        <el-table-column prop="details" label="详情" />
      </el-table>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import {
  User,
  Document,
  Search,
  SuccessFilled,
  TrendCharts,
  PieChart
} from '@element-plus/icons-vue'

const stats = ref({
  totalMerchants: 0,
  totalCodes: 0,
  todayVerifies: 0,
  validRate: 0
})

const recentActivities = ref([])

// 模拟数据
const mockData = {
  stats: {
    totalMerchants: 156,
    totalCodes: 28450,
    todayVerifies: 342,
    validRate: 98.5
  },
  recentActivities: [
    {
      time: '2024-01-15 14:30',
      merchant: '商户A',
      action: '生成防伪码',
      details: '生成1000个防伪码'
    },
    {
      time: '2024-01-15 13:15',
      merchant: '商户B',
      action: '验证防伪码',
      details: '验证成功：ABC123456'
    },
    {
      time: '2024-01-15 11:45',
      merchant: '商户C',
      action: '注册商户',
      details: '新商户注册完成'
    },
    {
      time: '2024-01-15 10:20',
      merchant: '商户A',
      action: '批量验证',
      details: '批量验证50个防伪码'
    }
  ]
}

onMounted(() => {
  // 模拟API调用
  setTimeout(() => {
    stats.value = mockData.stats
    recentActivities.value = mockData.recentActivities
  }, 500)
})
</script>

<style scoped>
.dashboard {
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

.stat-icon.total-merchants {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

.stat-icon.total-codes {
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

.charts {
  margin-bottom: 20px;
}

.chart-card {
  border-radius: 8px;
  border: none;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
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

.recent-activities {
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
  
  .chart-container {
    height: 200px;
  }
}
</style>