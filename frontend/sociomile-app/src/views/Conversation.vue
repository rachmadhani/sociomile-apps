<template>
    <div>
        <h2>Conversations</h2>
        
        <div v-if="loading">Loading...</div>
        <div v-else-if="error">{{ error }}</div>
        
        <div v-else>
            <ul>
                <li v-for="conv in conversations" :key="conv.id" style="margin-bottom: 10px; border: 1px solid #ccc; padding: 10px;">
                    <div><strong>ID:</strong> {{ conv.id }}</div>
                    <div><strong>Status:</strong> {{ conv.status }}</div>
                    <div>
                        <router-link :to="`/conversation/${conv.id}`">View Details</router-link>
                    </div>
                </li>
            </ul>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { conversationService } from '../services/conversation.service'

const conversations = ref([])
const loading = ref(true)
const error = ref(null)

onMounted(async () => {
    try {
        const response = await conversationService.getConversations()
        conversations.value = response.data || []
    } catch (err) {
        error.value = "Failed to load conversations."
        console.error(err)
    } finally {
        loading.value = false
    }
})
</script>