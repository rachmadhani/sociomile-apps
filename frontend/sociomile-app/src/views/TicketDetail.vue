<template>
    <div>
        <h2>Ticket Detail</h2>
        
        <div v-if="loading">Loading ticket...</div>
        <div v-else-if="error">{{ error }}</div>
        
        <div v-else>
            <div style="background: #f9f9f9; padding: 15px; border: 1px solid #ddd; margin-bottom: 20px;">
                <p><strong>ID:</strong> {{ ticket.id }}</p>
                <p><strong>TenantID:</strong> {{ ticket.tenant_id }}</p>
                <p><strong>Priority:</strong> {{ ticket.priority }}</p>
                <p><strong>Status:</strong> <span style="font-weight: bold;">{{ ticket.status }}</span></p>
            </div>

            <!-- Admin Update Status Form -->
            <div v-if="user?.role === 'admin'" style="border-top: 1px solid #ccc; padding-top: 20px;">
                <h4>Update Status (Admin Only)</h4>
                <div style="margin-bottom: 10px;">
                    <label>New Status: </label>
                    <select v-model="newStatus">
                        <option value="open">Open</option>
                        <option value="in_progress">In Progress</option>
                        <option value="resolved">Resolved</option>
                        <option value="closed">Closed</option>
                    </select>
                </div>
                <button @click="updateStatus" :disabled="submittingUpdate">Update Ticket</button>
            </div>
            <div v-else style="margin-top: 20px; color: #666; font-style: italic;">
                Only admins can update ticket statuses.
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { ticketService } from '../services/ticket.service'

const route = useRoute()
const ticketId = route.params.id

const ticket = ref(null)
const loading = ref(true)
const error = ref(null)

const newStatus = ref("in_progress")
const submittingUpdate = ref(false)

const user = JSON.parse(localStorage.getItem('user') || '{}')

const loadTicket = async () => {
    loading.value = true
    error.value = null
    try {
        const response = await ticketService.getTickets() 
        
        const found = response.data.find(t => t.id === ticketId)
        if (found) {
            ticket.value = found
            newStatus.value = found.status
        } else {
            error.value = "Ticket not found."
        }
    } catch (err) {
        error.value = "Failed to load ticket details."
        console.error(err)
    } finally {
        loading.value = false
    }
}

const updateStatus = async () => {
    submittingUpdate.value = true
    try {
        await ticketService.updateStatus(ticketId, ticket.value.tenant_id, newStatus.value)
        alert("Status updated successfully!")
        await loadTicket() 
    } catch (err) {
        alert("Failed to update status. Are you an Admin?")
        console.error(err)
    } finally {
        submittingUpdate.value = false
    }
}

onMounted(() => {
    loadTicket()
})
</script>
