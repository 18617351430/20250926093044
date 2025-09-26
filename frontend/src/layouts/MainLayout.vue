<template>
  <div class="main-layout">
    <!-- 顶部导航栏 -->
    <el-header class="header">
      <div class="header-left">
        <h1 class="logo">防伪码管理系统</h1>
      </div>
      <div class="header-right">
        <el-dropdown @command="handleCommand">
          <span class="user-info">
            <el-avatar :size="32" :src="userAvatar" />
            <span class="user-name">{{ userInfo.name }}</span>
            <el-icon><arrow-down /></el-icon>
          </span>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item command="profile">个人中心</el-dropdown-item>
              <el-dropdown-item command="logout" divided>退出登录</el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </div>
    </el-header>

    <!-- 主要内容区域 -->
    <div class="main-content">
      <!-- 侧边栏 -->
      <el-aside class="sidebar" :width="isCollapse ? '64px' : '200px'">
        <div class="sidebar-toggle" @click="toggleSidebar">
          <el-icon>
            <expand v-if="isCollapse" />
            <fold v-else />
          </el-icon>
        </div>
        
        <el-menu
          :default-active="activeMenu"
          :collapse="isCollapse"
          :collapse-transition="false"
          router
          class="sidebar-menu"
        >
          <!-- 平台管理员菜单 -->
          <template v-if="authStore.isPlatform">
            <el-menu-item index="/platform/dashboard">
              <el-icon><odometer /></el-icon>
              <template #title>控制台</template>
            </el-menu-item>
            
            <el-sub-menu index="platform-management">
              <template #title>
                <el-icon><setting /></el-icon>
                <span>平台管理</span>
              </template>
              <el-menu-item index="/platform/merchants">商户管理</el-menu-item>
              <el-menu-item index="/platform/rules">规则配置</el-menu-item>
              <el-menu-item index="/platform/statistics">数据统计</el-menu-item>
            </el-sub-menu>
          </template>

          <!-- 商户用户菜单 -->
          <template v-if="authStore.isMerchant">
            <el-menu-item index="/merchant/dashboard">
              <el-icon><odometer /></el-icon>
              <template #title>控制台</template>
            </el-menu-item>
            
            <el-sub-menu index="code-management">
              <template #title>
                <el-icon><document /></el-icon>
                <span>防伪码管理</span>
              </template>
              <el-menu-item index="/merchant/codes">防伪码列表</el-menu-item>
              <el-menu-item index="/merchant/generate">生成防伪码</el-menu-item>
              <el-menu-item index="/merchant/batches">批次管理</el-menu-item>
            </el-sub-menu>
            
            <el-menu-item index="/merchant/verify">
              <el-icon><search /></el-icon>
              <template #title>防伪验证</template>
            </el-menu-item>
            
            <el-menu-item index="/merchant/reports">
              <el-icon><data-analysis /></el-icon>
              <template #title>验证报告</template>
            </el-menu-item>
          </template>
        </el-menu>
      </el-aside>

      <!-- 页面内容 -->
      <el-main class="content">
        <router-view />
      </el-main>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { ElMessage, ElMessageBox } from 'element-plus'
import {
  ArrowDown,
  Expand,
  Fold,
  Odometer,
  Setting,
  Document,
  Search,
  DataAnalysis
} from '@element-plus/icons-vue'

const router = useRouter()
const route = useRoute()
const authStore = useAuthStore()

const isCollapse = ref(false)
const activeMenu = computed(() => route.path)

const userInfo = computed(() => authStore.currentUser)
const userAvatar = computed(() => {
  // 根据用户类型返回不同的头像
  if (authStore.isPlatform) {
    return 'https://api.dicebear.com/7.x/avataaars/svg?seed=admin'
  } else {
    return 'https://api.dicebear.com/7.x/avataaars/svg?seed=merchant'
  }
})

const toggleSidebar = () => {
  isCollapse.value = !isCollapse.value
}

const handleCommand = async (command) => {
  switch (command) {
    case 'profile':
      ElMessage.info('个人中心功能开发中')
      break
    case 'logout':
      await handleLogout()
      break
  }
}

const handleLogout = async () => {
  try {
    await ElMessageBox.confirm('确定要退出登录吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    
    authStore.logout()
    ElMessage.success('退出登录成功')
    router.push('/login')
  } catch (error) {
    // 用户取消退出
  }
}

// 初始化检查认证状态
onMounted(() => {
  if (!authStore.isAuthenticated) {
    router.push('/login')
  }
})
</script>

<style scoped>
.main-layout {
  height: 100vh;
  display: flex;
  flex-direction: column;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 20px;
  background: #fff;
  border-bottom: 1px solid #e6e6e6;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.logo {
  margin: 0;
  font-size: 20px;
  font-weight: 600;
  color: #409eff;
}

.user-info {
  display: flex;
  align-items: center;
  cursor: pointer;
  padding: 8px 12px;
  border-radius: 6px;
  transition: background-color 0.3s;
}

.user-info:hover {
  background-color: #f5f7fa;
}

.user-name {
  margin: 0 8px;
  font-weight: 500;
}

.main-content {
  flex: 1;
  display: flex;
  overflow: hidden;
}

.sidebar {
  background: #fff;
  border-right: 1px solid #e6e6e6;
  display: flex;
  flex-direction: column;
  transition: width 0.3s;
}

.sidebar-toggle {
  height: 48px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  border-bottom: 1px solid #e6e6e6;
  transition: background-color 0.3s;
}

.sidebar-toggle:hover {
  background-color: #f5f7fa;
}

.sidebar-menu {
  flex: 1;
  border: none;
}

.content {
  padding: 20px;
  background: #f5f7fa;
  overflow-y: auto;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .sidebar {
    position: absolute;
    top: 60px;
    left: 0;
    z-index: 1000;
    height: calc(100vh - 60px);
  }
  
  .content {
    padding: 15px;
  }
  
  .user-name {
    display: none;
  }
}
</style>