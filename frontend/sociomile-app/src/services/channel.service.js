import api from "../api/axios";

export const channelService = {
    async webhook(tenant_id, customer_external_id, message) {
        const response = await api.post("/channel/webhook", { tenant_id, customer_external_id, message })
        return response.data
    }
}
