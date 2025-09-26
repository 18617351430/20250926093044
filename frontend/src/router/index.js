import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const routes = [
  {
    path: '/',
    redirect: '/login'
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/Login.vue'),
    meta: {
      title: '登录 - 防伪码管理系统',
      requiresAuth: false
    }
  },
  {
    path: '/verify',
    name: 'Verify',
    component: () => import('@/views/Verify.vue'),
    meta: {
      title: '防伪验证 - 防伪码管理系统',
      requiresAuth: false
    }
  },
  {
    path: '/platform',
    component: () => import('@/layouts/MainLayout.vue'),
    meta: {
      requiresAuth: true,
      requiresPlatform: true
    },
    children: [
      {
        path: 'dashboard',
        name: 'PlatformDashboard',
        component: () => import('@/views/Dashboard.vue'),
        meta: {
          title: '控制台 - 平台管理'
        }
      },
      {
        path: 'merchants',
        name: 'PlatformMerchants',
        component: () => import('@/views/PlatformMerchants.vue'),
        meta: {
          title: '商户管理 - 平台管理'
        }
      },
      {
        path: 'rules',
        name: 'PlatformRules',
        component: () => import('@/views/PlatformRules.vue'),
        meta: {
          title: '规则配置 - 平台管理'
        }
      },
      {
        path: 'statistics',
        name: 'PlatformStatistics',
        component: () => import('@/views/PlatformStatistics.vue'),
        meta: {
          title: '数据统计 - 平台管理'
        }
      }
    ]
  },
  {
    path: '/merchant',
    component: () => import('@/layouts/MainLayout.vue'),
    meta: {
      requiresAuth: true,
      requiresMerchant: true
    },
    children: [
      {
        path: 'dashboard',
        name: 'MerchantDashboard',
        component: () => import('@/views/Dashboard.vue'),
        meta: {
          title: '控制台 - 商户管理'
        }
      },
      {
        path: 'codes',
        name: 'MerchantCodes',
        component: () => import('@/views/MerchantCodes.vue'),
        meta: {
          title: '防伪码列表 - 商户管理'
        }
      },
      {
        path: 'generate',
        name: 'MerchantGenerate',
        component: () => import('@/views/MerchantGenerate.vue'),
        meta: {
          title: '生成防伪码 - 商户管理'
        }
      },
      {
        path: 'batches',
        name: 'MerchantBatches',
        component: () => import('@/views/MerchantBatches.vue'),
        meta: {
          title: '批次管理 - 商户管理'
        }
      },
      {
        path: 'verify',
        name: 'MerchantVerify',
        component: () => import('@/views/Verify.vue'),
        meta: {
          title: '防伪验证 - 商户管理'
        }
      },

    ]
  },
  {
    path: '/:pathMatch(.*)*',
    name: 'NotFound',
    component: () => import('@/views/NotFound.vue'),
    meta: {
      title: '页面不存在 - 防伪码管理系统'
    }
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// 路由守卫
router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()
  
  // 设置页面标题
  if (to.meta.title) {
    document.title = to.meta.title
  }

  // 检查是否需要认证
  if (to.meta.requiresAuth) {
    if (!authStore.isAuthenticated) {
      // 未登录，跳转到登录页
      next('/login')
      return
    }

    // 检查用户权限
    if (to.meta.requiresPlatform && !authStore.isPlatform) {
      // 非平台管理员访问平台页面
      next('/merchant/dashboard')
      return
    }

    if (to.meta.requiresMerchant && !authStore.isMerchant) {
      // 非商户用户访问商户页面
      next('/platform/dashboard')
      return
    }
  }

  // 如果已登录但访问登录页，跳转到对应首页
  if (to.path === '/login' && authStore.isAuthenticated) {
    if (authStore.isPlatform) {
      next('/platform/dashboard')
    } else {
      next('/merchant/dashboard')
    }
    return
  }

  next()
})

export default router