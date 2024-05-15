import { useQuery } from "@tanstack/react-query";
import { fetchGraphQL } from "../../../services/graphql-services";
import useGraphql from "../graphql";

function useGetAllMovies(status) {
  const { getAllMovies, queryKey } = useGraphql();
  const { isLoading, error, data, refetch } = useQuery({
    gcTime: 0,
    queryKey: [queryKey , status],
    queryFn: async () => fetchGraphQL(getAllMovies.query, {
      input:{
        filter:{
          status:status
        },
            pagination:{
              page:1,
              limit:100
            }
      }
    }),
  });
  const moviesData = data?.[getAllMovies.operation]?.data
  const pagination = data?.[getAllMovies.operation]?.pagination
  return {
    isLoading,
    error,
    data: moviesData,
    pagination,
    refetch,
  };
}

function useGetMovieByID(id) {
  const {getMovieByID , queryKey} = useGraphql()
  const {isLoading,error, data, refetch} = useQuery({
    gcTime: 0,
    queryKey: [queryKey , id],
    queryFn: async () => fetchGraphQL(getMovieByID.query , {
      input: id
    })
  })


  return {
    isLoading,
    error,
    data,
    refetch,
  }
}

export {useGetAllMovies , useGetMovieByID};
