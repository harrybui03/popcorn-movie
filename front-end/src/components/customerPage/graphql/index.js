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

  return {getAllTheaters , getAllShowTimes , queryKey}

}
export default useGraphql;