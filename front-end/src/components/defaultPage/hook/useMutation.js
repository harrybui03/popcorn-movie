import { useMutation, useQueryClient } from '@tanstack/react-query'
import { fetchGraphQL } from '../../../services/graphql-services'
import useGraphql from "../graphql";
import useAuth from '../../../hooks/useAuth';
import { useNavigate } from 'react-router-dom';
import { toast } from "react-toastify";

function useLogin({setIsLoginSuccess,setLoginErrorMessage }) {
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
            auth.setToken(JSON.stringify(token));
            navigate("/")
        }
    },
    onError: (error) => {
      toast.error('Fail to Login')
    }
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

function useSignUp({setIsSignupSuccess ,setSignupErrorMessage}) {
  const {signup} = useGraphql()
  const {mutate , isSuccess } = useMutation({
    mutationFn: (input) => 
      fetchGraphQL(signup.query , {
        input: input,
      }),
      onSuccess: (data) =>{
        // Show popup
        console.log(data)
      },
      onError: (error) => {
        console.log(error)
      }
  })

  function onSignup({email , displayName , password , confirmPassword}) {
    const passwordRegex = /^(?=.*[A-Z])(?=.*[a-z])(?=.*\d)(?=.*[!@#$%^&*()_+{}[\]:;<>,.?~\\\/-]).{8,}$/;
        if (email === '' || password === '' || confirmPassword === '' || displayName === '') {
            setIsSignupSuccess(false);
            setSignupErrorMessage("Vui lòng điền đủ thông tin.");
        } else if (password !== confirmPassword) {
            setIsSignupSuccess(false);
            setSignupErrorMessage("Vui lòng xác nhận lại mật khẩu.");
        } else if (!passwordRegex.test(password)) {
            setIsSignupSuccess(false);
            setSignupErrorMessage("Mật khẩu ít nhất 8 kí tự (phải bao gồm chữ hoa, chữ thường, chữ số và kí tự đặt biệt)");
        }  else {
         try {
            mutate({email , password: password , displayName , confirmPassword})
         } catch (error) {
             console.error('Error during sign-in:', error);
         }
     }
  }

  return {
    isSuccess,
    onSignup
  }
}

export  {useLogin , useSignUp}