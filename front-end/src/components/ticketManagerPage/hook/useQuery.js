import { useQuery } from "@tanstack/react-query";
import { fetchGraphQLWithToken } from "../../../services/graphql-services";
import useGraphql from "../graphql";

function useGetAllRooms(theaterChosen){
    const {getAllRooms} = useGraphql()
    const {isLoading,error,data,refetch} = useQuery({
        gcTime:0,
        queryKey:[getAllRooms.queryKey , theaterChosen],
        queryFn: async () => fetchGraphQLWithToken(getAllRooms.query,{
            input:{
                filter:{
                    theaterID:theaterChosen
                },
                pagination:{
                    page:1,
                    limit:100
                }
            }
        })
    })

    const roomsData = data?.[getAllRooms.operation]?.data
    const pagination = data?.[getAllRooms.operation]?.pagination
    return {
        isLoading,
        error,
        data: roomsData,
        pagination,
        refetch,
    }
}

function useGetAllShowtimes(theaterID,roomID , date) {
    const {getAllShowtimes , queryKey} = useGraphql()
    const {isLoading,error,data,refetch} = useQuery({
        gcTime:0,
        queryKey:[queryKey , theaterID, roomID , date],
        queryFn: async () => fetchGraphQLWithToken(getAllShowtimes.query,{
            input:{
                filter:{
                    theaterId:theaterID,
                    date:date,
                    roomId:roomID
                },
                pagination:{
                    page:1,
                    limit:100
                }
            }
        })
    })

    const showtimesData = data?.[getAllShowtimes.operation]?.data
    const pagination = data?.[getAllShowtimes.operation]?.pagination
    return {
        isLoading,
        error,
        data: showtimesData,
        pagination,
        refetch,
    }
}

function useGetRevenue(year) {
    const {getRevenue , queryKey} = useGraphql()
    const {isLoading,error,data,refetch} = useQuery({
        gcTime:0,
        queryKey:[queryKey,  year],
        queryFn: async () => fetchGraphQLWithToken(getRevenue.query,{
            input:{
                year:year,
            }
        })
    })

    return {
        isLoading,
        error,
        data:data?.[getRevenue.operation],
        refetch,
    }
}

export  {useGetAllRooms , useGetAllShowtimes, useGetRevenue};