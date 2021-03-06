import {createRouter, createWebHistory, RouteRecordRaw} from 'vue-router'
import Home from '../views/Home.vue'

const routes: Array<RouteRecordRaw> = [
    {
        path: '/',
        name: 'Home',
        component: Home
    },
    {
        path: '/projects',
        name: 'Projects',
        component: () => import(/* webpackChunkName: "projects" */ '../views/Projects.vue')
    },
    {
        path: '/projects/:id',
        name: 'Project',
        component: () => import(/* webpackChunkName: "projects" */ '../views/Project.vue'),
        children: [
            {
                path: '',
                component: () => import(/* webpackChunkName: "projects" */ '../views/project/Ideas.vue')
            },
            {
                path: 'settings',
                component: () => import(/* webpackChunkName: "projects" */ '../views/project/Settings.vue')
            }
        ]
    },
    {
        path: '/privacy',
        name: 'Privacy',
        component: () => import(/* webpackChunkName: "privacy" */ '../views/Privacy.vue')
    },
    {
        path: '/tos',
        name: 'ToS',
        component: () => import(/* webpackChunkName: "tos" */ '../views/ToS.vue')
    },
]

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
})

export default router
