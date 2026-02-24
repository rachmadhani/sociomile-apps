<template>
    <div>
        <a href="/dashboard" style="padding: 10px;">Dashboard</a>
        <a href="/conversation" style="padding: 10px;">Conversation</a>
        <a href="/ticket" style="padding: 10px;">Ticket</a>

        <button @click="logout" style="padding: 10px;">Logout</button> 
    </div>
</template>

<script setup>
import { useRouter } from "vue-router"

const router = useRouter()

const logout = () => {
    localStorage.removeItem("token")
    router.push({ name: "Login" })
}
</script>