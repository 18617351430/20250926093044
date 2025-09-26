import axios from 'axios'
import { ElMessage } from 'element-plus'
import { useAuthStore } from '@/store/auth'

// 创建axios实例
const api = axios.create({
  baseURL: 'http://localhost:8080/api',
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json'
  }
})

// 请求拦截器
api.interceptors.request.use(
  (config) => {
    const authStore = useAuthStore()
    if (authStore.token) {
      config.headers.Authorization = `Bearer ${authStore.token}`
    }
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

// 响应拦截器
api.interceptors.response.use(
  (response) => {
    const { data } = response
    if (data.code === 0) {
      return data.data
    } else {
      ElMessage.error(data.message || '请求失败')
      return Promise.reject(new Error(data.message || '请求失败'))
    }
  },
  (error) => {
    if (error.response) {
      const { status, data } = error.response
      
      switch (status) {
        case 401:
          ElMessage.error('登录已过期，请重新登录')
          const authStore = useAuthStore()
          authStore.logout()
          window.location.href = '/login'
          break
        case 403:
          ElMessage.error('权限不足')
          break
        case 500:
          ElMessage.error('服务器内部错误')
          break
        default:
          ElMessage.error(data?.message || '请求失败')
      }
    } else if (error.request) {
      ElMessage.error('网络错误，请检查网络连接')
    } else {
      ElMessage.error('请求配置错误')
    }
    
    return Promise.reject(error)
  }
)

// API接口定义
export const authAPI = {
  // 平台登录
  platformLogin: (data) => api.post('/platform/auth/login', data),
  // 商户登录
  merchantLogin: (data) => api.post('/merchant/auth/login', data)
}

export const platformAPI = {
  // 商户管理
  getMerchants: (params) => api.get('/platform/merchants', { params }),
  createMerchant: (data) => api.post('/platform/merchants', data),
  updateMerchant: (id, data) => api.put(`/platform/merchants/${id}`, data),
  deleteMerchant: (id) => api.delete(`/platform/merchants/${id}`),
  
  // 规则管理
  getRules: () => api.get('/platform/rules'),
  updateRule: (data) => api.put('/platform/rules', data),
  
  // 统计
  getStatistics: () => api.get('/platform/statistics')
}

export const merchantAPI = {
  // 批次管理
  getBatches: (params) => api.get('/merchant/batches', { params }),
  createBatch: (data) => api.post('/merchant/batches', data),
  deleteBatch: (id) => api.delete(`/merchant/batches/${id}`),
  
  // 防伪码管理
  getCodes: (params) => api.get('/merchant/codes', { params }),
  generateCodes: (data) => api.post('/merchant/codes/generate', data),
  exportCodes: (batchId) => api.get(`/merchant/codes/export/${batchId}`),
  
  // 统计
  getStatistics: () => api.get('/merchant/statistics')
}

export const verifyAPI = {
  // 验证防伪码
  verifyCode: (code) => api.post('/verify', { code }),
  // 批量验证
  batchVerify: (codes) => api.post('/verify/batch', { codes })
}

export default api