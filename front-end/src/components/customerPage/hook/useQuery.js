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
                movieId:movieId,
                theaterId:theaterId,
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

  function useGetAllTickets(showTimeId ) {
    const { getAllTickets, queryKey } = useGraphql();
    const { isLoading, error, data, refetch } = useQuery({
      gcTime: 0,
      enabled:!!showTimeId,
      queryKey: [queryKey , showTimeId],
      queryFn: async () => fetchGraphQL(getAllTickets.query, {
        input:{
            filter:{
              showTimeID:showTimeId
            } , 
              pagination:{
                page:1,
                limit:100
              }
        }
      }),
    });
    const ticketsData = data?.[getAllTickets.operation]?.data
    const pagination = data?.[getAllTickets.operation]?.pagination
    return {
      isLoading,
      error,
      data: ticketsData,
      pagination,
      refetch,
    };
  }

  function useGetAllTransactions(userID) {
    const { getAllTransactions, queryKey } = useGraphql();
    const { isLoading, error, data, refetch } = useQuery({
      gcTime: 0,
      enabled:!!userID,
      queryKey: [queryKey , userID],
      queryFn: async () => fetchGraphQL(getAllTransactions.query, {
        input:{
            filter:{
              userID:userID
            } , 
              pagination:{
                page:0,
                limit:10
              }
        }
      }),
    });
    const transactionData = data?.[getAllTransactions.operation]?.data
    const pagination = data?.[getAllTransactions.operation]?.pagination
    return {
      isLoading,
      error,
      data: transactionData,
      pagination,
      refetch,
    };
  }

  function useGetAllSeats(roomID) {
    const {getAllSeats , queryKey} = useGraphql()
    const { isLoading, error, data, refetch } = useQuery({
      gcTime: 0,
      enabled:!!roomID,
      queryKey: [queryKey , roomID],
      queryFn: async () => fetchGraphQL(getAllSeats.query, {
        input:{
            filter:{
              roomID:roomID
            } , 
              pagination:{
                page:0,
                limit:10
              }
        }
      }),
    }); 

    const seatData = data?.[getAllSeats.operation]?.data
    const pagination = data?.[getAllSeats.operation]?.pagination
    return {
      isLoading,
      error,
      data: seatData,
      pagination,
      refetch,
    };
  }

  
  
  export {useGetAllThearters , useGetAllShowTimes , useGetAllTickets,useGetAllSeats , useGetAllTransactions};