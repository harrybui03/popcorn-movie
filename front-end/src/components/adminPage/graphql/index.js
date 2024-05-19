import { buildQuery } from "../../../services/graphql-services"

function useGraphql(){
    const queryKey = "User"
    const getAllUsers = buildQuery({
        operation:"Users",
        params:{input:"ListUserInput!"},
        options:{
            type:"query",
        },
        node:`
        data{
            id
            displayName
            email
            isLocked
            role
          }
          pagination{
            total
          }
        `
    })

    const createUser = buildQuery({
      operation: "CreateUser",
      options:{
        type: "mutation"
      },
      node:`
        id
        displayName
        isLocked
        role
      `,
      params:{
        input:"CreateUserInput!",
      }
    })

    const updateUser = buildQuery({
      operation:"UpdateUser",
      options:{
        type:"mutation"
      },
      node:`
      id
      displayName
      isLocked
      role
      `,
      params:{
        input:"UpdateUserInput!"
      }
    })

    return {getAllUsers , queryKey , createUser , updateUser}
}

export default useGraphql