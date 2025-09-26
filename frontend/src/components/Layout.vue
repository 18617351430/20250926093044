<template>
  <el-container class="layout-container">
    <!-- 侧边栏 -->
    <el-aside :width="isCollapse ? '64px' : '200px'" class="sidebar">
      <div class="logo">
        <h2 v-if="!isCollapse">防伪码系统</h2>
        <h2 v-else>防伪</h2>
      </div>
      <el-menu
        :default-active="activeMenu"
        :collapse="isCollapse"
        :collapse-transition="false"
        router
        class="sidebar-menu"
      >
        <!-- 平台管理员菜单 -->
        <template v-if="userType === 'platform'">
          <el-menu-item index="/platform/dashboard">
            <el-icon><House /></el-icon>
            <span>控制台</span>
          </el-menu-item>
          <el-menu-item index="/platform/merchants">
            <el-icon><User /></el-icon>
            <span>商户管理</span>
          </el-menu-item>
          <el-menu-item index="/platform/codes">
            <el-icon><Document /></el-icon>
            <span>防伪码管理</span>
          </el-menu-item>
          <el-menu-item index="/platform/rules">
            <el-icon><Setting /></el-icon>
            <span>规则设置</span>
          </el-menu-item>
          <el-menu-item index="/platform/statistics">
            <el-icon><DataAnalysis /></el-icon>
            <span>统计分析</span>
          </el-menu-item>
        </template>

        <!-- 商户菜单 -->
        <template v-else-if="userType === 'merchant'">
          <el-menu-item index="/merchant/dashboard">
            <el-icon><House /></el-icon>
            <span>控制台</span>
          </el-menu-item>
          <el-menu-item index="/merchant/codes">
            <el-icon><Document /></el-icon>
            <span>防伪码管理</span>
          </el-menu-item>
          <el-menu-item index="/merchant/codes/generate">
            <el-icon><Plus /></el-icon>
            <span>生成防伪码</span>
          </el-menu-item>
          <el-menu-item index="/merchant/verifies">
            <el-icon><Search /></el-icon>
            <span>验证记录</span>
          </el-menu-item>
          <el-menu-item index="/merchant/profile">
            <el-icon><User /></el-icon>
            <span>商户信息</span>
          </el-menu-item>
        </template>
      </el-menu>
    </el-aside>

    <!-- 主内容区 -->
    <el-container>
      <!-- 顶部导航 -->
      <el-header class="header">
        <div class="header-left">
          <el-button
            :icon="isCollapse ? 'Expand' : 'Fold'"
            @click="toggleSidebar"
            circle
            size="small"
          />
          <el-breadcrumb separator="/" class="breadcrumb">
            <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
            <el-breadcrumb-item>{{ currentRouteName }}</el-breadcrumb-item>
          </el-breadcrumb>
        </div>
        <div class="header-right">
          <el-dropdown @command="handleCommand">
            <span class="user-info">
              <el-avatar :size="32" :src="userInfo.avatar" />
              <span class="username">{{ userInfo.username }}</span>
              <el-icon><ArrowDown /></el-icon>
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item command="profile">个人信息</el-dropdown-item>
                <el-dropdown-item command="logout" divided>退出登录</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </div>
      </el-header>

      <!-- 主要内容 -->
      <el-main class="main-content">
        <router-view />
      </el-main>
    </el-container>
  </el-container>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth'
import {
  House,
  User,
  Document,
  Setting,
  DataAnalysis,
  Plus,
  Search,
  Expand,
  Fold,
  ArrowDown
} from '@element-plus/icons-vue'

const route = useRoute()
const router = useRouter()
const authStore = useAuthStore()

const isCollapse = ref(false)
const userInfo = ref({
  username: '',
  avatar: ''
})

// 计算属性
const activeMenu = computed(() => route.path)
const userType = computed(() => authStore.userType)
const currentRouteName = computed(() => {
  const routeMap = {
    '/platform/dashboard': '控制台',
    '/platform/merchants': '商户管理',
    '/platform/codes': '防伪码管理',
    '/platform/rules': '规则设置',
    '/platform/statistics': '统计分析',
    '/merchant/dashboard': '控制台',
    '/merchant/codes': '防伪码管理',
    '/merchant/codes/generate': '生成防伪码',
    '/merchant/verifies': '验证记录',
    '/merchant/profile': '商户信息'
  }
  return routeMap[route.path] || '未知页面'
})

// 方法
const toggleSidebar = () => {
  isCollapse.value = !isCollapse.value
}

const handleCommand = (command) => {
  if (command === 'logout') {
    authStore.logout()
    router.push('/login')
  } else if (command === 'profile') {
    // 跳转到个人信息页面
    if (userType.value === 'platform') {
      // 平台管理员个人信息
    } else {
      router.push('/merchant/profile')
    }
  }
}

// 生命周期
onMounted(() => {
  // 初始化用户信息
  userInfo.value = {
    username: authStore.userInfo?.username || '用户',
    avatar: authStore.userInfo?.avatar || ''
  }
})
</script>

<style scoped>
.layout-container {
  height: 100vh;
}

.sidebar {
  background-color: #001529;
  transition: width 0.3s;
}

.logo {
  height: 60px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  border-bottom: 1px solid #333;
}

.logo h2 {
  margin: 0;
  font-size: 18px;
  font-weight: 600;
}

.sidebar-menu {
  border: none;
  background-color: #001529;
}

.sidebar-menu .el-menu-item {
  color: #bfbfbf;
}

.sidebar-menu .el-menu-item:hover {
  background-color: #1890ff;
  color: white;
}

.sidebar-menu .el-menu-item.is-active {
  background-color: #1890ff;
  color: white;
}

.header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 20px;
  background-color: white;
  border-bottom: 1px solid #f0f0f0;
  box-shadow: 0 1px 4px rgba(0, 21, 41, 0.08);
}

.header-left {
  display: flex;
  align-items: center;
  gap: 16px;
}

.breadcrumb {
  font-size: 14px;
}

.header-right {
  display: flex;
  align-items: center;
}

.user-info {
  display: flex;
  align-items: center;
  gap: 8px;
  cursor: pointer;
  padding: 8px 12px;
  border-radius: 4px;
  transition: background-color 0.3s;
}

.user-info:hover {
  background-color: #f5f5f5;
}

.username {
  font-size: 14px;
  color: #333;
}

.main-content {
  background-color: #f0f2f5;
  padding: 20px;
  overflow-y: auto;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .sidebar {
    position: fixed;
    z-index: 1000;
    height: 100vh;
  }
  
  .header {
    padding: 0 10px;
  }
  
  .username {
    display: none;
  }
}
</style>