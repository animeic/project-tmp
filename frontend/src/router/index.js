import { createRouter,createWebHistory } from "vue-router"

// import $store from '@/store'

const routes = [
  {
      path: '/:catchAll(.*)',
      name: '404',
      component: ()=>import('../views/404.vue')
  },
  // {
  //   path: '/',
  //   name: 'home',
  //   component: ()=>import('../views/Home.vue')
  // },
  {
    path: '/login',
    name: 'login',
    component: ()=>import('../views/Login.vue')
  },
    // 暂时禁用，仅供自己使用
  // {
  //   path: '/register',
  //   name: 'register',
  //   component: ()=>import('../views/Register.vue')
  // },
  {
    path: '/repass',
    name: 'repass',
    component: ()=>import('../views/Repass.vue')
  },

  // admin新增文章
  {
    path: '/admin/article/post',
    name: 'article-post',
    component: ()=>import('../views/admin/ArticlePost.vue')
  },
  // web 文章详情d
  {
    path: '/article/:id',
    name: 'article-detail',
    component: ()=>import('../views/article/detail.vue')
  },
  // web 文章列表
  // {
  //   path: '/article',
  //   name: 'article-list',
  //   component: ()=>import('../views/article/list.vue')
  // },
  {
    path: '/',
    name: 'home',
    component: ()=>import('../views/article/list.vue')
  },
  // user profile
  {
    path: '/user/profile/:id',
    name: 'user-profile',
    component: ()=>import('../views/Profile.vue')
  },
  // edit detail
  {
    path: '/admin/article/detail/:id',
    name: 'admin-article-detail',
    component: ()=>import('../views/admin/ArticleEdit.vue')
  },
  {
    path: '/permission/list',
    name: 'permission-list',
    component: ()=>import('../views/rbac/PermissionList.vue')
  },
  {
    path: '/permission/edit',
    name: 'permission-edit',
    component: ()=>import('../views/rbac/PermissionEdit.vue')
  },
  {
    path: '/permission/add',
    name: 'permission-add',
    component: ()=>import('../views/rbac/PermissionAdd.vue')
  },
  {
    path: '/comment',
    name: 'comment',
    component: ()=>import('../components/Comment.vue')
  },

]

const router = createRouter({
    history: createWebHistory("/"),
    routes
})
// router.beforeEach((to,from,next) => {
  // let token = localStorage.getItem("token")
  // // 有token时访问
  // if(token){
  //   // 不应该访问登录页面和注册页面
  //   if(to.path === "/login"){
  //     // router.push("/")
  //     // $store.dispatch("getToken")
  //
  //
  //     // if($store.state.tokenInfo){
  //     //   router.push("/")
  //     // }
  //     // 应该验证token是否
  //     // 如果登录了跳转到首页
  //   }
  //   next()
  // }else{
  //   // 主页和登录界面不需要token
  //   if(to.path === "/" || to.path === "/login"){
  //     next()
  //   }else{
  //     // 其他页面如果没有token则跳转到登录页面
  //     next('/login')
  //   }
  // }
// })

export default router