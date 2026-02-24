import api from "../api/axios";

export const channelService = {
    async webhook(tenant_id, content) {
        const response = await api.post("/channel/webhook", { tenant_id, content })
        return response.data
    }
}
