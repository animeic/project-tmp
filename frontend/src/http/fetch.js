import { createFetch } from '@vueuse/core'
const useMyFetch = createFetch({
    baseUrl: '', // 基礎路由
    options: {
        immediate: false, // 是否在使用 useMyFetch 時自動運行 (推薦手動運行)
        timeout: 30000, // 請求過期時間
// 在請求前修改配置，如：注入 token 值
        async beforeFetch({ options }) {
            const myToken = await getMyToken()
            options.headers.Authorization = `Bearer ${myToken}`
            return { options }
        },
// 在請求後處理數據，如：攔截錯誤、處理過期
        afterFetch({ data, response }) {
// code...
            return { data, response }
        },
// 請求錯誤
        onFetchError({ data, error }) {
            console.error(error)
            return { data, error }
        },
    },
    fetchOptions: {
        mode: 'cors',
        credentials: 'include', // 請求時攜帶 cookie 值
    },
})