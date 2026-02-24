import Api from "../api/axios"

export const loginService = {
    async login(email, password) {
        const response = await Api.post("/auth/login", { email, password })
        return response.data
    },
    async register(email, password) {
        const response = await Api.post("/auth/register", { email, password })
        return response.data
    },
}