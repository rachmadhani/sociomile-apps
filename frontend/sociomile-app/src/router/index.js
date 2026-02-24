import { createRouter, createWebHistory } from "vue-router";

import Login from "../views/Login.vue"
import Register from "../views/Register.vue"
import Dashboard from "../views/Dashboard.vue"
import Conversations from "../views/Conversation.vue"


const routes = [
    {
        path: "/",
        name: "Login",
        component: Login,
    },
    {
        path: "/register",
        name: "Register",
        component: Register,
    },
    {
        path: "/dashboard",
        name: "Dashboard",
        component: Dashboard,
        meta: { requiresAuth: true }
    },
    {
        path: "/conversations",
        name: "Conversations",
        component: Conversations,
        meta: { requiresAuth: true }
    },
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

router.beforeEach((to, from, next) => {
    const token = localStorage.getItem("token")
    if (to.meta.requiresAuth && !token) {
        next({ name: "Login" })
        return
    } else {
        next()
    }
})

export default router