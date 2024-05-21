import { useQuery } from "@tanstack/react-query";
import useGraphql from "../graphql";
import { fetchGraphQL, fetchGraphQLWithToken } from "../../../services/graphql-services";

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

  function convertDate(dateString) {
    const isDateString = dateString.includes('/');
    let dateObject;
    
    if (isDateString) {
        // If the input is in the format "day/month/year"
        const [day, month, year] = dateString.split('/');
        dateObject = new Date(`${year}-${month}-${day}`);
    } else {
        // If the input is in ISO format like "2024-05-26T07:00:00.000Z"
        dateObject = new Date(dateString);
    }

    // Set hours, minutes, and seconds to zero
    dateObject.setUTCHours(0);
    dateObject.setUTCMinutes(0);
    dateObject.setUTCSeconds(0);
    dateObject.setUTCMilliseconds(0);

    // Convert back to string in the desired format
    const convertedTimestamp = dateObject.toISOString();

    return convertedTimestamp;
}


  function useGetAllShowTimes(theaterId , movieId , dateChosen) {
    const { getAllShowTimes, queryKey } = useGraphql();
    const { isLoading, error, data, refetch } = useQuery({
      gcTime: 0,
      queryKey: [queryKey , theaterId , movieId, dateChosen],
      queryFn: async () => fetchGraphQL(getAllShowTimes.query, {
        input:{
            filter:{
                movieId:movieId,
                theaterId:theaterId,
                date:convertDate(dateChosen)
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
      queryFn: async () => fetchGraphQLWithToken(getAllTransactions.query, {
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