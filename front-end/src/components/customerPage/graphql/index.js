import { buildQuery } from "../../../services/graphql-services";
function useGraphql() {
  const queryKey = "Theaters";
  const getAllTheaters = buildQuery({
    operation: "Theaters",
    params:{input:"ListTheatersInput!"},
    options: {
      type: "query",
    },
    node: `
    data{
        id
        address
        name
        phoneNumber
      }
      pagination{
        total
      }
        `,
  });

  const getAllShowTimes = buildQuery({
    operation: "ShowTimes",
    params:{input:"ListShowTimeInput!"},
    options: {
      type: "query",
    },
    node: `
    data{
        id
        startAt
        endAt
        movie {
          title
        }
        room{
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
  });

  const getAllTickets = buildQuery({
    operation: "Tickets",
    params:{input:"ListTicketInput!"},
    options: {
      type: "query",
    },
    node: `
    data{
      ID
      IsBooked
      Price
      Seat{
        ID
        SeatNumber
        Category
      }
    }
    pagination{
      total
    }
        `,
  }); 

  const getAllTransactions = buildQuery({
    operation: "Transactions",
    params:{input:"ListTransactionInput!"},
    options: {
      type: "query",
    },
    node: `
    data{
      id
      total 
      createdAt
    }
    pagination{
      total
    }
        `,
  }); 

  const getAllSeats = buildQuery({
    operation:"Seats",
    params:{input: "ListSeatInput!"},
    options: {
      type: "query",
    },
    node:`
    data{
      ID
      SeatNumber
      Category
    }
    pagination{
      total
    }
    `
  })

  const createTransaction = buildQuery({
    operation:"CreateTransaction",
    params:{input:"CreateTransactionInput!"},
    options:{type: "mutation"},
    node:`
    OrderCode
    Status
    CheckoutUrl
    QRCode
    `
  })

  const changePassword = buildQuery({
    operation:"ChangePassword",
    params:{input:"ChangePasswordInput!"},
    options:{type:"mutation"},
    node:`output`,
  })
  



  return {getAllTheaters , getAllShowTimes,getAllTickets , getAllTransactions , getAllSeats, createTransaction,changePassword , queryKey}

}
export default useGraphql;