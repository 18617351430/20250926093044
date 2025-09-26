import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export const useAuthStore = defineStore('auth', () => {
  // 状态
  const token = ref(localStorage.getItem('auth_token') || '')
  const userInfo = ref(JSON.parse(localStorage.getItem('user_info') || '{}'))
  const userType = ref(localStorage.getItem('user_type') || '')

  // 计算属性
  const isAuthenticated = computed(() => !!token.value)
  const isPlatform = computed(() => userType.value === 'platform')
  const isMerchant = computed(() => userType.value === 'merchant')
  const currentUser = computed(() => userInfo.value)

  // 方法
  const login = (loginData) => {
    return new Promise((resolve, reject) => {
      // 模拟API调用
      setTimeout(() => {
        if (loginData.username === 'admin' && loginData.password === 'admin123') {
          // 平台管理员登录
          const userData = {
            id: 1,
            username: 'admin',
            name: '平台管理员',
            type: 'platform',
            permissions: ['*']
          }
          
          setAuthData('platform_token_' + Date.now(), userData, 'platform')
          resolve(userData)
        } else if (loginData.username === 'merchant' && loginData.password === 'merchant123') {
          // 商户登录
          const userData = {
            id: 1001,
            username: 'merchant',
            name: '测试商户',
            type: 'merchant',
            merchantId: 1001,
            permissions: ['codes:read', 'codes:write', 'verifies:read']
          }
          
          setAuthData('merchant_token_' + Date.now(), userData, 'merchant')
          resolve(userData)
        } else {
          reject(new Error('用户名或密码错误'))
        }
      }, 1000)
    })
  }

  const logout = () => {
    token.value = ''
    userInfo.value = {}
    userType.value = ''
    
    localStorage.removeItem('auth_token')
    localStorage.removeItem('user_info')
    localStorage.removeItem('user_type')
  }

  const setAuthData = (newToken, newUserInfo, newUserType) => {
    token.value = newToken
    userInfo.value = newUserInfo
    userType.value = newUserType
    
    localStorage.setItem('auth_token', newToken)
    localStorage.setItem('user_info', JSON.stringify(newUserInfo))
    localStorage.setItem('user_type', newUserType)
  }

  const checkPermission = (permission) => {
    if (!isAuthenticated.value) return false
    if (userInfo.value.permissions?.includes('*')) return true
    return userInfo.value.permissions?.includes(permission)
  }

  // 初始化时检查token有效性
  const initAuth = () => {
    if (token.value && userInfo.value.id) {
      // 这里可以添加token有效性检查逻辑
      return true
    }
    return false
  }

  return {
    // 状态
    token,
    userInfo,
    userType,
    
    // 计算属性
    isAuthenticated,
    isPlatform,
    isMerchant,
    currentUser,
    
    // 方法
    login,
    logout,
    setAuthData,
    checkPermission,
    initAuth
  }
})