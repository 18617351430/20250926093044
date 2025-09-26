import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useAuthStore = defineStore('auth', () => {
  const token = ref(localStorage.getItem('auth_token') || '')
  const userInfo = ref(JSON.parse(localStorage.getItem('user_info') || 'null'))
  const userType = ref(localStorage.getItem('user_type') || '') // platform, merchant

  const isAuthenticated = ref(!!token.value)

  const login = (loginData) => {
    // 模拟登录成功
    const mockUser = {
      id: loginData.username === 'admin' ? 1 : 2,
      username: loginData.username,
      name: loginData.username === 'admin' ? '平台管理员' : '示例商户',
      type: loginData.username === 'admin' ? 'platform' : 'merchant'
    }

    const mockToken = 'mock_jwt_token_' + Date.now()

    token.value = mockToken
    userInfo.value = mockUser
    userType.value = mockUser.type
    isAuthenticated.value = true

    // 存储到本地存储
    localStorage.setItem('auth_token', mockToken)
    localStorage.setItem('user_info', JSON.stringify(mockUser))
    localStorage.setItem('user_type', mockUser.type)

    return Promise.resolve({
      success: true,
      message: '登录成功',
      data: {
        token: mockToken,
        user: mockUser
      }
    })
  }

  const logout = () => {
    token.value = ''
    userInfo.value = null
    userType.value = ''
    isAuthenticated.value = false

    // 清除本地存储
    localStorage.removeItem('auth_token')
    localStorage.removeItem('user_info')
    localStorage.removeItem('user_type')

    return Promise.resolve()
  }

  const isPlatform = () => {
    return userType.value === 'platform'
  }

  const isMerchant = () => {
    return userType.value === 'merchant'
  }

  return {
    token,
    userInfo,
    userType,
    isAuthenticated,
    login,
    logout,
    isPlatform,
    isMerchant
  }
})