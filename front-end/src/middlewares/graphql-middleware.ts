import { RequestMiddleware, ResponseMiddleware } from 'graphql-request'

function getAccessToken() {
  const localStorageValue = localStorage.getItem('token');
  if (!localStorageValue) {
    window.location.href = '/'
    return 
  }

  const auth = JSON.parse(JSON.parse(localStorageValue))

  const accessToken = auth.accessToken
  return accessToken

}

export const requestMiddleware: RequestMiddleware = async (request) => {
  const token = await getAccessToken()
  const headers = {
    "Authorization":token
  }
  return {
    ...request,
    headers: { ...request.headers, ...headers },
  }
}

export const responseMiddleware: ResponseMiddleware = (response) => {
  if (!(response instanceof Error) && response.errors) {
    console.error(
      `Request error:
        status ${String(response.status)}
        details: ${response.errors.map((_) => _.message).join(`, `)}`
    )
  }
}
