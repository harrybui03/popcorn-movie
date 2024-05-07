import { useQuery } from "@tanstack/react-query";
import { fetchGraphQL } from "../../../services/graphql-services";
import useGraphql from "../graphql";

function useDemoQuery() {
  const { getAllJobTitles, queryKey } = useGraphql();
  const { isLoading, error, data, refetch } = useQuery({
    gcTime: 0,
    queryKey: [queryKey],
    queryFn: async () => fetchGraphQL(getAllJobTitles.query, {
      input:{
        filter:{
          status:"UPCOMING"
        },
            pagination:{
              page:1,
              limit:10
            }
      }
    }),
  });
  return {
    isLoading,
    error,
    data,
    refetch,
  };
}

export default useDemoQuery;
