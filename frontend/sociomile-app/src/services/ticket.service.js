import api from "../api/axios";

export const ticketService = {
    async getTickets(params = {}) {
        const response = await api.get("/ticket/", { params })
        return response.data
    },

    async updateStatus(id, tenant_id, status) {
        const response = await api.post(`/ticket/${id}/update-status`, { tenant_id, status })
        return response.data
    }
}
