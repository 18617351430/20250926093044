<template>
  <div class="merchant-statistics">
    <div class="page-header">
      <h1>数据统计</h1>
      <p>商户数据统计和分析</p>
    </div>

    <!-- 统计概览 -->
    <el-row :gutter="20" class="stats-overview">
      <el-col :xs="12" :sm="6" v-for="stat in overviewStats" :key="stat.title">
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
              <div class="stat-trend" :class="stat.trend">
                <el-icon v-if="stat.trend === 'up'"><top /></el-icon>
                <el-icon v-else><bottom /></el-icon>
                {{ stat.change }}
              </div>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <!-- 图表区域 -->
    <el-row :gutter="20" class="charts-section">
      <el-col :xs="24" :lg="12">
        <el-card shadow="hover">
          <template #header>
            <span>验证趋势</span>
            <el-radio-group v-model="trendPeriod" size="small">
              <el-radio-button label="week">本周</el-radio-button>
              <el-radio-button label="month">本月</el-radio-button>
              <el-radio-button label="year">本年</el-radio-button>
            </el-radio-group>
          </template>
          <div class="chart-container">
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
        <el-card shadow="hover">
          <template #header>
            <span>批次分布</span>
          </template>
          <div class="chart-container">
            <div class="mock-chart">
              <div class="chart-placeholder">
                <el-icon :size="48"><pie-chart /></el-icon>
                <p>批次分布图表</p>
              </div>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <!-- 详细统计表格 -->
    <el-card shadow="hover">
      <template #header>
        <span>批次统计详情</span>
      </template>
      <el-table :data="batchStats" stripe>
        <el-table-column prop="batchName" label="批次名称" />
        <el-table-column prop="batchNo" label="批次号" width="120" />
        <el-table-column prop="codeCount" label="防伪码数量" align="center">
          <template #default="{ row }">
            <el-tag>{{ row.codeCount }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="todayVerify" label="今日验证" align="center">
          <template #default="{ row }">
            <el-tag type="success">{{ row.todayVerify }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="monthVerify" label="本月验证" align="center">
          <template #default="{ row }">
            <el-tag type="warning">{{ row.monthVerify }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="totalVerify" label="总验证数" align="center">
          <template #default="{ row }">
            <el-tag type="info">{{ row.totalVerify }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="validRate" label="有效验证率" align="center">
          <template #default="{ row }">
            <el-progress 
              :percentage="row.validRate" 
              :show-text="false" 
              style="width: 80px;"
            />
            <span style="margin-left: 8px;">{{ row.validRate }}%</span>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import {
  Document,
  Check,
  Close,
  DataLine,
  PieChart,
  Top,
  Bottom
} from '@element-plus/icons-vue'

const trendPeriod = ref('month')

const overviewStats = ref([
  {
    title: '防伪码总数',
    value: '105,000',
    icon: Document,
    color: '#409eff',
    trend: 'up',
    change: '+5.4%'
  },
  {
    title: '今日验证',
    value: '156',
    icon: Check,
    color: '#67c23a',
    trend: 'up',
    change: '+12.3%'
  },
  {
    title: '本月验证',
    value: '4,230',
    icon: Check,
    color: '#e6a23c',
    trend: 'up',
    change: '+8.7%'
  },
  {
    title: '无效验证',
    value: '23',
    icon: Close,
    color: '#f56c6c',
    trend: 'down',
    change: '-3.2%'
  }
])

const batchStats = ref([
  {
    batchName: '2024年第一季度防伪码',
    batchNo: 'BATCH2024001',
    codeCount: '50,000',
    todayVerify: '89',
    monthVerify: '2,560',
    totalVerify: '85,600',
    validRate: 94
  },
  {
    batchName: '2024年第二季度防伪码',
    batchNo: 'BATCH2024002',
    codeCount: '30,000',
    todayVerify: '45',
    monthVerify: '1,230',
    totalVerify: '28,900',
    validRate: 92
  },
  {
    batchName: '2023年12月防伪码',
    batchNo: 'BATCH2023121',
    codeCount: '10,000',
    todayVerify: '12',
    monthVerify: '890',
    totalVerify: '9,800',
    validRate: 88
  }
])
</script>

<style scoped>
.merchant-statistics {
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

.stats-overview {
  margin-bottom: 20px;
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
  margin-bottom: 4px;
}

.stat-trend {
  font-size: 12px;
  display: flex;
  align-items: center;
}

.stat-trend.up {
  color: #67c23a;
}

.stat-trend.down {
  color: #f56c6c;
}

.charts-section {
  margin-bottom: 20px;
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

/* 响应式设计 */
@media (max-width: 768px) {
  .merchant-statistics {
    padding: 15px;
  }
  
  .stat-value {
    font-size: 20px;
  }
}
</style>