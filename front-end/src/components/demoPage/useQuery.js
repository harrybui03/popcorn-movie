import { useGraphql } from "./graphql/index";
import { useQuery } from "@tanstack/react-query";
import { fetchGraphQL } from "../../services/graphql-services";

function useDemoQuery() {
  const { getAllJobTitles, queryKey } = useGraphql();
  const { isLoading, error, data, refetch } = useQuery({
    gcTime: 0,
    queryKey: [queryKey],
    queryFn: async () => fetchGraphQL(getAllJobTitles.query, {}),
  });
  return {
    isLoading,
    error,
    data,
    refetch,
  };
}

export default useDemoQuery;
