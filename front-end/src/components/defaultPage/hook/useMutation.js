import { useMutation, useQueryClient } from '@tanstack/react-query'
import { fetchGraphQL } from '../../../services/graphql-services'
import useGraphql from "../graphql";
import useAuth from '../../../hooks/useAuth';
import { useNavigate } from 'react-router-dom';

function useLogin({setIsLoginSuccess,setLoginErrorMessage}) {
  const queryClient = useQueryClient()
  const navigate = useNavigate()
  const { login  } = useGraphql();
  const auth = useAuth();
  const { mutate , isSuccess } = useMutation({
    mutationFn: (input) =>
      fetchGraphQL(login.query, {
        input: input,
      }),
    onSuccess: (data) => {
        const accessToken = data?.[login.operation]?.accessToken
        const refreshToken = data?.[login.operation]?.refreshToken
        if (accessToken && refreshToken) {
            const token = {accessToken ,refreshToken }
            console.log(token)
            auth.setToken(JSON.stringify(token));
            navigate("/")
        }
    },
  })

  function onLogin({email , pwd}) {
     const passwordRegex = /^(?=.*[A-Z])(?=.*[a-z])(?=.*\d)(?=.*[!@#$%^&*()_+{}[\]:;<>,.?~\\\/-]).{8,}$/;
     if (email === '' || pwd === '') {
         setIsLoginSuccess(false);
         setLoginErrorMessage("Vui lòng điền đủ thông tin.");
     } else if (!passwordRegex.test(pwd)) {
         setIsLoginSuccess(false);
         setLoginErrorMessage("Mật khẩu ít nhất 8 kí tự (phải bao gồm chữ hoa, chữ thường, chữ số và kí tự đặt biệt)");
     } else {
         try {
            mutate({email , password: pwd})
         } catch (error) {
             console.error('Error during sign-in:', error);
         }
     }
  }
  return {
    isSuccess,
    onLogin,
  };
}

export default useLogin