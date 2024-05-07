import { GraphQLClient } from 'graphql-request'
import { requestMiddleware, responseMiddleware } from '../middlewares/graphql-middleware'
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

export const graphQLClient = new GraphQLClient('http://localhost:8000/query', {
  requestMiddleware,
  responseMiddleware,
})

export const fetchGraphQL = async <T extends object>(
  query: any,
  variables?: any
): Promise<T> => {
  return await graphQLClient.request(query, variables)
}