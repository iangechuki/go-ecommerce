import api from "./index"
export const AuthService = {
    login: async (email: string, password: string) => {
      const response = await api.post("/auth/login", { email, password });
      console.log(response)
    },
}