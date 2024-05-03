import { RequestMiddleware, ResponseMiddleware } from 'graphql-request'

function getAccessToken() {
  return 'eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImtpZCI6InEtMjNmYWxldlpoaEQzaG05Q1Fia1A1TVF5VSJ9.eyJhdWQiOiI0NWI0NzBjMi01ZjI3LTRlN2MtYWUzNi0wMzlhODhmMjJhNWQiLCJpc3MiOiJodHRwczovL2xvZ2luLm1pY3Jvc29mdG9ubGluZS5jb20vNDIzM2I2NWYtNDA0ZC00ZWRjLWJhNDYtMDI2NWI2NWVlNWQ5L3YyLjAiLCJpYXQiOjE3MTI5MTQ0MzEsIm5iZiI6MTcxMjkxNDQzMSwiZXhwIjoxNzEyOTE4NjQ5LCJhaW8iOiJBV1FBbS84V0FBQUExQlVTZlY1Nm5aRTdjQklTL3VEVTdHQld6SmIwN0s4bTFISjFHMFZKeXUyRExEV2FrZUhieHJBNFRrWjRQYjBJeGpaK3Y5Y05ZdExDWGhrOEhrME00K0pSVHpEaXRUZU1GV1NIdEpDbHBYVkRCNzdlWVUxaklwalBQK3MxM3VINiIsImF6cCI6IjQ1YjQ3MGMyLTVmMjctNGU3Yy1hZTM2LTAzOWE4OGYyMmE1ZCIsImF6cGFjciI6IjEiLCJuYW1lIjoiQW5kcmV3IFBoYW0gKFRFQ0hWSUZZLklUUykiLCJvaWQiOiI0YjViMGVjMS0yMjA3LTRlZWQtOTBjYy05NTJmNDk0ZWRjYjMiLCJwcmVmZXJyZWRfdXNlcm5hbWUiOiJhbmRyZXcucGhhbUB0ZWNodmlmeS5jb20udm4iLCJyaCI6IjAuQVZNQVg3WXpRazFBM0U2NlJnSmx0bDdsMmNKd3RFVW5YM3hPcmpZRG1vanlLbDNGQUdVLiIsInNjcCI6IkhSTS5SZWFkIiwic3ViIjoiUWFxaVc5NE00bHMyX2VLdEZuMDVzb0VmV3FNSFJUbGZEaHIxelNVRkRQVSIsInRpZCI6IjQyMzNiNjVmLTQwNGQtNGVkYy1iYTQ2LTAyNjViNjVlZTVkOSIsInV0aSI6IjNKNjVBbGxmM0UyV0p6RDlzWlYxQVEiLCJ2ZXIiOiIyLjAifQ.b8QYlaWB9Fnccf3sr5LjYFM0c3PnQy3kciZ5tsTLPRWXi39fG5VQLYYxAh1aVwlh2TW8AJBzXsjg_qgsmlJJqSEq3Llful9hFVZETlnR09zUHFrBrQ5M-pp3DahA7WFx2TKH6ZmdkQXresRmB583VQkU5CLJcubgXVyKu79U4ap3As1blrQWsVVSFmxyiHngzq8L7SrRJ2j0XuO8m_EDj6Mgd2GdHgHyny72q7k_zxM1R0xz4Z2TWuCfkjL1XillHtI5fmHjcE04S-kGOBOAcFhLACFeuFp7WNeytwwIyLtHZ-1AgdyHQjWfacbX4nYbUbirJtcIJQh-zqO6_SKobQ'
}

export const requestMiddleware: RequestMiddleware = async (request) => {
  const token = await getAccessToken()
  const headers = {
    Authorization: `Bearer ${token}`, // Set your token in the Authorization header
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
