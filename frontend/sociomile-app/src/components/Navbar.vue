<template>
    <div>
        <router-link to="/dashboard" style="padding: 10px;">Dashboard</router-link>
        <router-link to="/conversations" style="padding: 10px;">Conversations</router-link>
        <router-link to="/tickets" style="padding: 10px;">Tickets</router-link>
        <router-link to="/channel" style="padding: 10px;">Webhook Channel</router-link>

        <button @click="logout" style="padding: 10px;">Logout</button> 
    </div>
</template>

<script setup>
import { useRouter } from "vue-router"

const router = useRouter()

const logout = () => {
    localStorage.removeItem("token")
    localStorage.removeItem("user")
    router.push({ name: "Login" })
}
</script>