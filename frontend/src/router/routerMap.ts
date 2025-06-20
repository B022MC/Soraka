const constantRouterMap = [
  {
    path: '/',
    component: () => import('@/layouts/AppSider.vue'),
    redirect: { name: 'home' },
    children: [
      {
        path: '/home',
        name: 'home',
        component: () => import('@/views/home/index.vue'),
        meta: {
          title: '启动页',
          icon: 'icon-Home'
        }
      },
      {
        path: '/life',
        name: 'life',
        component: () => import('@/views/life/index.vue'),
        meta: {
          title: '生涯',
          icon: 'icon-Person'
        }
      },
      {
        path: '/stats',
        name: 'stats',
        component: () => import('@/views/stats/index.vue'),
        meta: {
          title: '战绩查询',
          icon: 'icon-Search'
        }
      },
      {
        path: '/match',
        name: 'match',
        component: () => import('@/views/match/index.vue'),
        meta: {
          title: '对局信息',
          icon: 'icon-Game'
        }
      },
      {
        path: '/more',
        name: 'more',
        component: () => import('@/views/more/index.vue'),
        meta: {
          title: '其他功能',
          icon: 'icon-Wrench'
        }
      },

      // 设置页面，走底部按钮跳转，不显示在侧边栏
      {
        path: '/setting',
        name: 'setting',
        component: () => import('@/views/setting/index.vue'),
        meta: {
          title: '设置',
          icon: 'icon-xitong',
          hideInMenu: true
        }
      }
    ]
  }
]
export default constantRouterMap;