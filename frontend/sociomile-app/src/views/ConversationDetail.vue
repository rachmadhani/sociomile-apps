<template>
    <div>
        <h2>Conversation Detail</h2>
        
        <div v-if="loading">Loading conversation...</div>
        <div v-else-if="error">{{ error }}</div>
        
        <div v-else>
            <div style="background: #f9f9f9; padding: 15px; border: 1px solid #ddd; margin-bottom: 20px;">
                <p><strong>ID:</strong> {{ conversation.id }}</p>
                <p><strong>Status:</strong> {{ conversation.status }}</p>
                <p><strong>Agent ID:</strong> {{ conversation.assigned_agent_id || 'Unassigned' }}</p>
            </div>

            <h3>Messages</h3>
            <div style="max-height: 400px; overflow-y: auto; border: 1px solid #ccc; padding: 10px; margin-bottom: 20px;">
                <div v-for="msg in conversation.messages" :key="msg.ID" 
                     :style="{ 
                         padding: '10px', 
                         margin: '10px 0', 
                         borderRadius: '5px',
                         background: msg.sender_type === 'agent' ? '#e3f2fd' : '#f5f5f5',
                         textAlign: msg.sender_type === 'agent' ? 'right' : 'left'
                     }">
                    <p style="margin: 0;"><strong>{{ msg.sender_type }}:</strong> {{ msg.content }}</p>
                    <small style="color: #666;">{{ new Date(msg.CreatedAt).toLocaleString() }}</small>
                </div>
                <div v-if="!conversation.messages || conversation.messages.length === 0">No messages yet.</div>
            </div>

            <!-- Agent Reply Form -->
            <div v-if="user?.role === 'agent' || user?.role === 'admin'" style="margin-bottom: 20px;">
                <h4>Reply</h4>
                <textarea v-model="replyContent" rows="3" style="width: 100%; margin-bottom: 10px;" placeholder="Type your message here..."></textarea>
                <button @click="submitReply" :disabled="submittingReply">Send Reply</button>
            </div>

            <!-- Escalate Form -->
            <div v-if="user?.role === 'agent' && conversation.status !== 'escalated'" style="border-top: 1px solid #ccc; padding-top: 20px;">
                <h4>Escalate to Ticket</h4>
                <div style="margin-bottom: 10px;">
                    <label>Category: </label>
                    <input v-model="escalateCategory" placeholder="e.g. Technical Support" />
                </div>
                <div style="margin-bottom: 10px;">
                    <label>Priority: </label>
                    <select v-model="escalatePriority">
                        <option value="low">Low</option>
                        <option value="medium">Medium</option>
                        <option value="high">High</option>
                        <option value="urgent">Urgent</option>
                    </select>
                </div>
                <button @click="escalateTicket" :disabled="submittingEscalate">Escalate to Admin</button>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { conversationService } from '../services/conversation.service'

const route = useRoute()
const conversationId = route.params.id

const conversation = ref(null)
const loading = ref(true)
const error = ref(null)

const replyContent = ref("")
const submittingReply = ref(false)

const escalateCategory = ref("")
const escalatePriority = ref("medium")
const submittingEscalate = ref(false)


const user = JSON.parse(localStorage.getItem('user') || '{}')

const loadConversation = async () => {
    loading.value = true
    error.value = null
    try {
        const response = await conversationService.getConversation(conversationId)
        conversation.value = response.data
    } catch (err) {
        error.value = "Failed to load conversation details."
        console.error(err)
    } finally {
        loading.value = false
    }
}

const submitReply = async () => {
    if (!replyContent.value.trim()) return
    submittingReply.value = true
    try {
        await conversationService.reply(conversationId, replyContent.value)
        replyContent.value = ""
        await loadConversation() 
    } catch (err) {
        alert("Failed to send reply")
        console.error(err)
    } finally {
        submittingReply.value = false
    }
}

const escalateTicket = async () => {
    if (!escalateCategory.value.trim()) {
        alert("Please enter a category")
        return
    }
    submittingEscalate.value = true
    try {
        await conversationService.escalate(conversationId, escalateCategory.value, escalatePriority.value)
        alert("Escalated successfully!")
        await loadConversation()
    } catch (err) {
        alert("Failed to escalate ticket")
        console.error(err)
    } finally {
        submittingEscalate.value = false
    }
}

onMounted(() => {
    loadConversation()
})
</script>
