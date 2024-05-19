import { useQuery } from "@tanstack/react-query";
import useGraphql from "../graphql";
import { fetchGraphQLWithToken } from "../../../services/graphql-services";

function useGetAllUsers() {
    const {getAllUsers , queryKey} =  useGraphql()
    const {isLoading, error, data, refetch } = useQuery({
        gcTime: 0,
        queryKey: [queryKey ],
        queryFn: async () => fetchGraphQLWithToken(getAllUsers.query, {
        input:{
            pagination:{
              page:1,
              limit:100
            }
      }
    } ),
    })

    const usersData = data?.[getAllUsers.operation]?.data
    const pagination = data?.[getAllUsers.operation]?.pagination
  return {
    isLoading,
    error,
    data: usersData,
    pagination,
    refetch,
  };
}

export {useGetAllUsers}