import { buildQuery } from "../../../services/graphql-services";
function useGraphql() {
  const queryKey = "Movie";
  const getAllMovies = buildQuery({
    operation: "Movies",
    params:{input:"ListMovieInput!"},
    options: {
      type: "query",
    },
    node: `
        data
        {
          id
          title
          genre
          status
          language
          director
          cast
          poster
          rated
          duration
          trailer
          openingDay
          story
        }
        pagination{
          total
        }
        `,
  });

  const getMovieByID = buildQuery({
    operation: "GetMovieByID",
    params:{input:"ID!"},
    options: {
      type: "query",
    },
    node:`
    id
    title
    genre
    status
    language
    director
    cast
    poster
    rated
    duration
    trailer
    openingDay
    story
    `
  })


  const login = buildQuery({
    operation: "Login",
    options: {
      type: "mutation",
    },
    node: `
      accessToken
      refreshToken
        `,
    params: {
      input: "LoginInput!",
    },
  });

  return { getAllMovies, login, queryKey , getMovieByID };
}
export default useGraphql;
