import api from "../api/axios";

export const conversationService = {
    async getConversations(params = {}) {
        const response = await api.get("/conversation/", { params })
        return response.data
    },

    async getConversation(id) {
        const response = await api.get(`/conversation/${id}`)
        return response.data
    },

    async reply(id, message) {
        const response = await api.post(`/conversation/${id}/agent-reply`, { message })
        return response.data
    },

    async escalate(id, title, description, priority) {
        const response = await api.post(`/conversation/${id}/escalate`, { title, description, priority })
        return response.data
    }
}
