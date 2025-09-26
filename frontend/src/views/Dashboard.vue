<template>
  <div class="dashboard">
    <!-- 页面标题 -->
    <div class="page-header">
      <h1>控制台</h1>
      <p>欢迎回来，{{ userInfo.name }}</p>
    </div>

    <!-- 统计卡片 -->
    <el-row :gutter="20" class="stats-cards">
      <el-col :xs="24" :sm="12" :md="6" v-for="stat in stats" :key="stat.title">
        <el-card class="stat-card" shadow="hover">
          <div class="stat-content">
            <div class="stat-icon" :style="{ backgroundColor: stat.color }">
              <el-icon :size="24">
                <component :is="stat.icon" />
              </el-icon>
            </div>
            <div class="stat-info">
              <div class="stat-value">{{ stat.value }}</div>
              <div class="stat-title">{{ stat.title }}</div>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <!-- 图表区域 -->
    <el-row :gutter="20" class="charts-section">
      <el-col :xs="24" :lg="12">
        <el-card class="chart-card" shadow="hover">
          <template #header>
            <span>验证趋势</span>
          </template>
          <div class="chart-container">
            <!-- 这里可以放置验证趋势图表 -->
            <div class="mock-chart">
              <div class="chart-placeholder">
                <el-icon :size="48"><data-line /></el-icon>
                <p>验证趋势图表</p>
              </div>
            </div>
          </div>
        </el-card>
      </el-col>
      
      <el-col :xs="24" :lg="12">
        <el-card class="chart-card" shadow="hover">
          <template #header>
            <span>验证分布</span>
          </template>
          <div class="chart-container">
            <!-- 这里可以放置验证分布图表 -->
            <div class="mock-chart">
              <div class="chart-placeholder">
                <el-icon :size="48"><pie-chart /></el-icon>
                <p>验证分布图表</p>
              </div>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <!-- 最近活动 -->
    <el-card class="recent-activities" shadow="hover">
      <template #header>
        <span>最近活动</span>
      </template>
      <el-timeline>
        <el-timeline-item
          v-for="activity in recentActivities"
          :key="activity.id"
          :timestamp="activity.time"
          :type="activity.type"
        >
          {{ activity.content }}
        </el-timeline-item>
      </el-timeline>
    </el-card>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useAuthStore } from '@/stores/auth'
import {
  Document,
  Check,
  Close,
  TrendCharts,
  DataLine,
  PieChart
} from '@element-plus/icons-vue'

const authStore = useAuthStore()

const userInfo = computed(() => authStore.currentUser)

// 统计数据
const stats = ref([])

// 最近活动
const recentActivities = ref([
  {
    id: 1,
    time: '2024-01-15 14:30:00',
    content: '用户验证了防伪码 ABC123456XYZ7890',
    type: 'success'
  },
  {
    id: 2,
    time: '2024-01-15 13:45:00',
    content: '生成了新的防伪码批次 BATCH2024001',
    type: 'primary'
  },
  {
    id: 3,
    time: '2024-01-15 12:20:00',
    content: '防伪码 DEF987654UVW3210 验证失败',
    type: 'warning'
  },
  {
    id: 4,
    time: '2024-01-15 11:10:00',
    content: '系统数据备份完成',
    type: 'info'
  }
])

// 根据用户类型设置不同的统计数据
const setStatsByUserType = () => {
  if (authStore.isPlatform) {
    stats.value = [
      {
        title: '总商户数',
        value: '156',
        icon: Document,
        color: '#409eff'
      },
      {
        title: '今日验证数',
        value: '2,845',
        icon: Check,
        color: '#67c23a'
      },
      {
        title: '有效验证',
        value: '2,680',
        icon: TrendCharts,
        color: '#e6a23c'
      },
      {
        title: '无效验证',
        value: '165',
        icon: Close,
        color: '#f56c6c'
      }
    ]
  } else {
    stats.value = [
      {
        title: '防伪码总数',
        value: '50,000',
        icon: Document,
        color: '#409eff'
      },
      {
        title: '今日验证数',
        value: '156',
        icon: Check,
        color: '#67c23a'
      },
      {
        title: '有效验证',
        value: '142',
        icon: TrendCharts,
        color: '#e6a23c'
      },
      {
        title: '无效验证',
        value: '14',
        icon: Close,
        color: '#f56c6c'
      }
    ]
  }
}

onMounted(() => {
  setStatsByUserType()
})
</script>

<style scoped>
.dashboard {
  padding: 20px;
}

.page-header {
  margin-bottom: 30px;
}

.page-header h1 {
  margin: 0 0 10px 0;
  font-size: 28px;
  font-weight: 600;
  color: #333;
}

.page-header p {
  margin: 0;
  color: #666;
  font-size: 14px;
}

.stats-cards {
  margin-bottom: 30px;
}

.stat-card {
  border-radius: 8px;
}

.stat-content {
  display: flex;
  align-items: center;
}

.stat-icon {
  width: 48px;
  height: 48px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  margin-right: 16px;
}

.stat-info {
  flex: 1;
}

.stat-value {
  font-size: 24px;
  font-weight: 600;
  color: #333;
  margin-bottom: 4px;
}

.stat-title {
  font-size: 14px;
  color: #666;
}

.charts-section {
  margin-bottom: 30px;
}

.chart-card {
  border-radius: 8px;
}

.chart-container {
  height: 300px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.mock-chart {
  text-align: center;
  color: #999;
}

.chart-placeholder p {
  margin-top: 10px;
  font-size: 14px;
}

.recent-activities {
  border-radius: 8px;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .dashboard {
    padding: 15px;
  }
  
  .page-header h1 {
    font-size: 24px;
  }
  
  .stat-value {
    font-size: 20px;
  }
}
</style>