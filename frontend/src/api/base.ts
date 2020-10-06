import axios from "axios"
import qs from "qs"

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

const GAUTH_SERVER_HOST = process.env.REACT_APP_SERVER_HOST || "localhost"
const GAUTH_SERVER_PORT = process.env.REACT_APP_SERVER_PORT || "3001"

const getServerPort = () => {
  if (GAUTH_SERVER_PORT === "80") return ""
  return `:${GAUTH_SERVER_PORT}`
}

export default axios.create({
  baseURL: `http://${GAUTH_SERVER_HOST}${getServerPort()}/api`,
  transformRequest: (data: any, headers: any): any => {
    if (headers["Content-Type"] === "application/x-www-form-urlencoded") {
      return qs.stringify(data)
    }
    return data
  },
})
