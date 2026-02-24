<template>
    <div>
        <h2>Channel</h2>
        
        <div style="background: #f9f9f9; padding: 15px; border: 1px solid #ddd;">
            <div style="margin-bottom: 10px;">
                <label>Select Agent: </label>
                <select v-model="tenantId" style="width: 100%; max-width: 400px; margin-top: 5px;">
                    <option value="">Select Agent</option>
                    <option v-for="agent in agents" :key="agent.id" :value="agent.TenantID">{{ agent.Name }} ({{ agent.Email }})</option>
                </select>
            </div>

            <div style="margin-bottom: 10px;">
                <label>Whatsapp Number: </label>
                <input v-model="customerExternalId" type="text" style="width: 100%; max-width: 400px; margin-top: 5px;" placeholder="Whatsapp Number">
            </div>

            <div style="margin-bottom: 10px;">
                <label>Customer Message Content: </label>
                <textarea v-model="content" rows="4" style="width: 100%; max-width: 400px; margin-top: 5px;" placeholder="I need help with my account..."></textarea>
            </div>

            <button @click="triggerWebhook" :disabled="submitting">Submit</button>
            <span v-if="successMsg" style="color: green; margin-left: 10px;">{{ successMsg }}</span>
            <span v-if="errorMsg" style="color: red; margin-left: 10px;">{{ errorMsg }}</span>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { channelService } from '../services/channel.service'
import { loginService } from '../services/auth.service'


const tenantId = ref("")
const content = ref("")
const customerExternalId = ref("")
const submitting = ref(false)
const successMsg = ref("")
const errorMsg = ref("")
const agents = ref([])

onMounted(async () => {
    const response = await loginService.getListAgent()
    agents.value = response.data
})

const triggerWebhook = async () => {
    if (!tenantId.value.trim() || !content.value.trim()) {
        errorMsg.value = "Both Agent ID and Message Content are required."
        return
    }

    if (!customerExternalId.value.trim()) {
        errorMsg.value = "Customer External ID is required."
        return
    }

    submitting.value = true
    successMsg.value = ""
    errorMsg.value = ""

    try {
        await channelService.webhook(tenantId.value, customerExternalId.value, content.value)
        successMsg.value = "Webhook triggered successfully! A new conversation was created."
        content.value = ""
    } catch (err) {
        errorMsg.value = "Failed to trigger webhook."
        console.error(err)
    } finally {
        submitting.value = false
    }
}
</script>
