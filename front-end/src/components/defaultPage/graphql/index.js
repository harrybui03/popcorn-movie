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

  const signup = buildQuery({
    operation:"Signup",
    options: {
      type:"mutation",
    },
    node:`
    id
    displayName
    email
    `,
    params:{
      input:"RegisterInput!"
    }
  })

  const forgotPassword = buildQuery({
    operation:"ForgotPassword",
    options:{
      type:"mutation",
    },
    node:`
      output
    `,
    params:{
      input:"String!"
    }
  })

  const resetPassword = buildQuery({
    operation:"ResetPassword",
    options:{
      type:"mutation",
    },
    node:
    `output`,
    params:{
      input:"ResetPasswordInput!"
    }
  })

  return { getAllMovies, login,signup, queryKey , getMovieByID ,forgotPassword , resetPassword};
}
export default useGraphql;
