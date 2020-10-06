import api, { ApiResponse } from "./base"

interface AuthenticationRequest {
  login: string
  password: string
}

interface AuthenticationResponse {
  token: string
}

export default async function (
  request: AuthenticationRequest
): Promise<ApiResponse<AuthenticationResponse>> {
  const response = await api.post<AuthenticationResponse>("/tokens", request, {
    headers: {
      "Content-Type": "application/x-www-form-urlencoded",
    },
  })

  const { data, status, headers } = response
  return { data, status, headers }
}
