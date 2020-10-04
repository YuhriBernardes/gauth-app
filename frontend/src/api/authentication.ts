import api, { ApiError, ApiResponse } from "./base"

interface AuthenticationRequest {
  userName: string
  password: string
}

interface AuthenticationResponse {
  token: string
}

export default async function (
  request: AuthenticationRequest
): Promise<ApiResponse<AuthenticationResponse>> {
  const response = await api.post<AuthenticationResponse>("/auth", request)

  const { data, status, headers } = response
  return { data, status, headers }
}
