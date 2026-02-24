<template>
    <div>
        <h2>Tickets List</h2>
        
        <div v-if="loading">Loading...</div>
        <div v-else-if="error">{{ error }}</div>
        
        <div v-else>
            <ul>
                <li v-for="ticket in tickets" :key="ticket.id" style="margin-bottom: 10px; border: 1px solid #ccc; padding: 10px;">
                    <div><strong>ID:</strong> {{ ticket.id }}</div>
                    <div><strong>Subject:</strong> {{ ticket.subject }}</div>
                    <div><strong>Priority:</strong> {{ ticket.priority }}</div>
                    <div><strong>Status:</strong> {{ ticket.status }}</div>
                    <div>
                        <router-link :to="`/ticket/${ticket.id}`">Manage Ticket</router-link>
                    </div>
                </li>
            </ul>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { ticketService } from '../services/ticket.service'

const tickets = ref([])
const loading = ref(true)
const error = ref(null)

onMounted(async () => {
    try {
        const response = await ticketService.getTickets()
        tickets.value = response.data || []
    } catch (err) {
        error.value = "Failed to load tickets."
        console.error(err)
    } finally {
        loading.value = false
    }
})
</script>
