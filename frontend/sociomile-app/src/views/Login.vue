<template>
    <div>
        <h2>Login</h2>
        <input v-model="email" placeholder="Email" />
        <input type="password" v-model="password" placeholder="Password" />
        <button @click="submit">Login</button>
    </div>
</template>

<script setup>
import { ref } from "vue"
import { loginService } from "../services/auth.service"
import { useRouter } from "vue-router"

const email = ref("")
const password = ref("")
const router = useRouter()

const submit = async () => {
    try {
        const response = await loginService.login(email.value, password.value)
        localStorage.setItem("token", response.data.access_token)
        router.push({ name: "Dashboard" })
    } catch (error) {
        console.error(error)
    }
}
</script>