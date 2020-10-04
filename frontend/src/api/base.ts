import axios from "axios"

export interface ApiError {
  errMsg: string
}

export enum ResponseStatus {
  UNAUTHORIZED = 401,
  BAD_REQUEST = 400,
}

export interface ApiResponse<T> {
  data: T
  status: number
  headers: any
}

export default axios.create({
  baseURL: "http://localhost:3001/api",
})
