import { createRouter, createWebHistory } from "vue-router";

import Login from "../views/Login.vue"
import Register from "../views/Register.vue"
import Dashboard from "../views/Dashboard.vue"
import Conversations from "../views/Conversation.vue"
import ConversationDetail from "../views/ConversationDetail.vue"
import Tickets from "../views/Tickets.vue"
import TicketDetail from "../views/TicketDetail.vue"
import Channel from "../views/Channel.vue"
import MainLayout from "../layouts/MainLayout.vue"


const routes = [
    {
        path: "/login",
        name: "Login",
        component: Login,
    },
    {
        path: "/register",
        name: "Register",
        component: Register,
    },
    {
        path: "/channel",
        name: "Channel",
        component: Channel,
    },
    {
        path: "/",
        name: "MainLayout",
        component: MainLayout,
        meta: { requiresAuth: true },
        redirect: "/dashboard",
        children: [
            {
                path: "dashboard",
                name: "Dashboard",
                component: Dashboard,
                meta: { roles: ['agent', 'admin'] }
            },
            {
                path: "conversations",
                name: "Conversations",
                component: Conversations,
                meta: { roles: ['agent', 'admin'] }
            },
            {
                path: "conversation/:id",
                name: "ConversationDetail",
                component: ConversationDetail,
                meta: { roles: ['agent', 'admin'] }
            },
            {
                path: "tickets",
                name: "Tickets",
                component: Tickets,
                meta: { roles: ['agent', 'admin'] }
            },
            {
                path: "ticket/:id",
                name: "TicketDetail",
                component: TicketDetail,
                meta: { roles: ['agent', 'admin'] }
            }
        ]
    },
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

router.beforeEach((to, from, next) => {
    const token = localStorage.getItem("token")
    let user = null
    try {
        user = JSON.parse(localStorage.getItem("user"))
    } catch (e) { }

    if (to.meta.requiresAuth && !token) {
        next({ name: "Login" })
        return
    }

    if (to.meta.roles && to.meta.roles.length > 0) {
        if (!user || !user.role || !to.meta.roles.includes(user.role)) {
            console.warn("User doesn't have the required role")

            next({ name: "Dashboard" })
            return
        }
    }

    next()
})

export default router