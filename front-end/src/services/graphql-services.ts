import { GraphQLClient } from 'graphql-request'
import { getAccessToken, requestMiddleware, responseMiddleware } from '../middlewares/graphql-middleware'
import * as dotenv from 'dotenv';

dotenv.config();
interface IbuildQuery {
  operation: string
  node: string
  params?: Record<string, string>
  options: {
    type: 'query' | 'mutation'
  }
}

export interface IbuildQueryReturn {
  operation: string
  query: string
}

export const buildQuery = (props: IbuildQuery): IbuildQueryReturn => {
  const { operation, node, params, options } = props
  let paramsQuery = ''
  let operationQuery = ''
  if (params && Object.keys(params).length > 0) {
    Object.keys(params).forEach((key) => {
      paramsQuery += `$${key}: ${params[key]}\n`
      operationQuery += `${key}: $${key}\n`
    })
  }
  const query = `
      ${options.type} ${operation}(
        ${paramsQuery}
      ) {
        ${operation}(
          ${operationQuery}
        ) {
          ${node}
        }
      }
    `
  return {
    query,
    operation,
  }
}

const reactURL = process.env.REACT_APP_GRAPHQL_URL ?? 'http://localhost:8000/query'
export const graphQLClient = new GraphQLClient(reactURL)

export const fetchGraphQL = async <T extends object>(
  query: any,
  variables?: any
): Promise<T> => {
  return await graphQLClient.request(query, variables)
}

export const graphQLClientWithToken = new GraphQLClient(reactURL, {
  requestMiddleware,
  responseMiddleware,
})

export const fetchGraphQLWithToken = async <T extends object>(
  query: any,
  variables?: any
): Promise<T> => {
  return await graphQLClientWithToken.request(query, variables)
}

export const fetchUploadFile = async (
  formData:FormData
):Promise<Response> => {
  return await fetch(reactURL, {
    method: 'POST',
    body: formData,
    headers: {
      "Authorization": getAccessToken(),
      // 'Content-Type':'multipart/form-data; boundary=<calculated when request is sent>',
      //  'Accept': '*/*',
    }
  })
}