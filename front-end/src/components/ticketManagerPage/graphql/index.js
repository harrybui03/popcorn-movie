import { buildQuery } from "../../../services/graphql-services";

function useGraphql() {
    const queryKey = "ShowTime";
    const getAllRooms = buildQuery({
        operation:"Rooms",
        params:{input:"ListRoomInput!"},
        options:{
            type:"query",
        },
        node:`
        data{
            id
            roomNumber
            theater{
              id
              address
              name
              phoneNumber
            }
            seat{
              ID
              SeatNumber
              Category
            }
          }
          pagination{
            total
          }
        `
    })

    const getAllShowtimes = buildQuery({
        operation:"ShowTimes",
        params:{input:"ListShowTimeInput!"},
        options:{
            type:"query",
        },
        node:`
        data{
            id
            startAt
            endAt
            movie {
              title
            }
            room{
              id
              roomNumber
              theater{
                address
                name
              }
            }
          }
          pagination{
            total
          }
        `
    })

    const generateTickets = buildQuery({
        operation:"GenerateTicket",
        options:{
            type:"mutation"
        },
        node:`
        data{
            id
            startAt
            endAt
            movie {
              title
            }
            room{
              id
              roomNumber
              theater{
                address
                name
              }
            }
          }
          pagination{
            total
          }
        `,
        params:{
            input:"GenerateTicketInput!"
        }
    })

    const createShowTime = buildQuery({
        operation:"CreateShowTime",
        options:{
            type:"mutation"
        },
        node: `
        id
        startAt
        endAt
        `,
        params:{
            input:"CreateShowTimeInput!"
        }
    })

    const updateShowTime = buildQuery({
        operation:"UpdateShowTime",
        options:{
            type:"mutation"
        },
        node: `
        id
        startAt
        endAt
        `,
        params:{
            input:"UpdateShowTimeInput!"
        }
    })

    const deleteShowTime = buildQuery({
        operation:"DeleteShowTime",
        options:{
            type:"mutation"
        },
        params:{
            input:"ID!"
        }
    })

    const deleteMovie = buildQuery({
        operation:"DeleteMovie",
        options:{
            type:"mutation"
        },
        params:{
            input:"ID!"
        }
    })

    const getRevenue = buildQuery({
        operation:"GetRevenue",
        options:{
            type:"query"
        },
        params:{
            input:"RevenueInput!"
        },
        node:`
        total
        arr{
          total
          month
        }
        `,
       
    })

    

    

    return {queryKey , getAllRooms,getAllShowtimes , generateTickets , createShowTime , updateShowTime,deleteShowTime,getRevenue , deleteMovie}
}

export default useGraphql;