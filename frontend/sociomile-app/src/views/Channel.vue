<template>
    <div>
        <h2>Incoming Webhook Settings</h2>
        
        <div style="background: #f9f9f9; padding: 15px; border: 1px solid #ddd;">
            <div style="margin-bottom: 10px;">
                <label>Select Agent ID: </label>
                <input v-model="tenantId" placeholder="Agent UUID (Acts as Tenant)" style="width: 100%; max-width: 400px; margin-top: 5px;" />
            </div>

            <div style="margin-bottom: 10px;">
                <label>Customer Message Content: </label>
                <textarea v-model="content" rows="4" style="width: 100%; max-width: 400px; margin-top: 5px;" placeholder="I need help with my account..."></textarea>
            </div>

            <button @click="triggerWebhook" :disabled="submitting">Trigger Webhook</button>
            <span v-if="successMsg" style="color: green; margin-left: 10px;">{{ successMsg }}</span>
            <span v-if="errorMsg" style="color: red; margin-left: 10px;">{{ errorMsg }}</span>
        </div>
    </div>
</template>

<script setup>
import { ref } from 'vue'
import { channelService } from '../services/channel.service'

const tenantId = ref("")
const content = ref("")
const submitting = ref(false)
const successMsg = ref("")
const errorMsg = ref("")

const triggerWebhook = async () => {
    if (!tenantId.value.trim() || !content.value.trim()) {
        errorMsg.value = "Both Agent ID and Message Content are required."
        return
    }

    submitting.value = true
    successMsg.value = ""
    errorMsg.value = ""

    try {
        await channelService.webhook(tenantId.value, content.value)
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
