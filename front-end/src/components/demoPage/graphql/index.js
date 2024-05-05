import { buildQuery } from "../../../services/graphql-services";
function useGraphql() {
  const queryKey = "jobtitle";
  const getAllJobTitles = buildQuery({
    operation: "GetAllJobTitles",
    options: {
      type: "query",
    },
    node: `
          edges {
            node {
              id
              code
              name
            }
          }
          pagination {
            page
            perPage
            total
          }
        `,
    params: {
      pagination: "PaginationInput",
      filter: "JobTitleFilter",
      orderBy: "JobTitleOrder",
      freeWord: "JobTitleFreeWord",
    },
  });

  const createJobTitle = buildQuery({
    operation: "CreateJobTitle",
    options: {
      type: "mutation",
    },
    node: `
          data {
            id
          }
        `,
    params: {
      input: "NewJobTitleInput!",
    },
  });

  return { getAllJobTitles, createJobTitle, queryKey };
}
export default useGraphql;
