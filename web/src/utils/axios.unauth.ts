import axios, { AxiosInstance } from "axios";
const baseURL: string =
  process.env.NEXT_PUBLIC_BACKEND_URL || "http://localhost:8190/api";

const axiosInstance: AxiosInstance = axios.create({
  baseURL,
  timeout: 10000,
  headers: {
    "Content-Type": "application/json",
  },
});

export default axiosInstance;
