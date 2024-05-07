import { useQuery } from "@tanstack/react-query";
import useGraphql from "../graphql";
import { fetchGraphQL } from "../../../services/graphql-services";

function useGetAllThearters() {
    const { getAllTheaters, queryKey } = useGraphql();
    const { isLoading, error, data, refetch } = useQuery({
      gcTime: 0,
      queryKey: [queryKey],
      queryFn: async () => fetchGraphQL(getAllTheaters.query, {
        input:{
              pagination:{
                page:1,
                limit:10
              }
        }
      }),
    });
    const theatersData = data?.[getAllTheaters.operation]?.data
    const pagination = data?.[getAllTheaters.operation]?.pagination
    return {
      isLoading,
      error,
      data: theatersData,
      pagination,
      refetch,
    };
  }

  function useGetAllShowTimes(theaterId , movieId) {
    const { getAllShowTimes, queryKey } = useGraphql();
    const { isLoading, error, data, refetch } = useQuery({
      gcTime: 0,
      queryKey: [queryKey , theaterId , movieId],
      queryFn: async () => fetchGraphQL(getAllShowTimes.query, {
        input:{
            filter:{
                movieId:movieId
            } , 
              pagination:{
                page:1,
                limit:10
              }
        }
      }),
    });
    const showTimesData = data?.[getAllShowTimes.operation]?.data
    const pagination = data?.[getAllShowTimes.operation]?.pagination
    return {
      isLoading,
      error,
      data: showTimesData,
      pagination,
      refetch,
    };
  }

  
  export {useGetAllThearters , useGetAllShowTimes};