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
        toast.success('Đăng nhập thành công')
        if (accessToken && refreshToken) {
            const token = {accessToken ,refreshToken }
            auth.setToken(JSON.stringify(token));
            navigate("/")
        }
    },
    onError: (error) => {
      toast.error('Sai Email hoặc mật khẩu')
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
        toast.success('Đăng kí thành công')
      },
      onError: (error) => {
        toast.error('Email đã tồn tại trong hệ thống')
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

function useForgotPassword() {
  const {forgotPassword} = useGraphql()
  const {mutate , isSuccess} = useMutation({
    mutationFn: (input) =>
    fetchGraphQL(forgotPassword.query, {
        input:input
    }),
    onSuccess: (data) =>{
        toast.success('Một thư email khôi phục mật khẩu đã được gửi cho địa chỉ email tài khoản của bạn')
    },
    onError:(error) => {
        console.log('Error')

        toast.error('Email không tồn tại')
    }
  })

  function onForgot(email) {
    if (email === '') {
      toast.error('Vui lòng điền đủ thông tin')
    }
    else {
      try {
        mutate(email)
     } catch (error) {
         console.error('Error during forgotpassword:', error);
     }
    }
    
  }

  return {
    onForgot,
    isSuccess
  }
}

function useResetPassword() {
  const {resetPassword} = useGraphql()
  const navigate = useNavigate()
  const {mutate , isSuccess} = useMutation({
    mutationFn: (input) => 
      fetchGraphQL(resetPassword.query , {
        input:input
      }),
      onSuccess: (data) => {
        toast.success('Reset mật khẩu thành công ')
        navigate("/")
      },
      onError:(error) => {
        toast.error('Reset mật khẩu thất bại')
      }
  })

  function onReset({newPassword , confirmNewPassword , token}){
    try {
      console.log({newPassword , confirmNewPassword , token})
      mutate({newPassword ,confirmNewPassword , token})
   } catch (error) {
       console.error('Error during sign-in:', error);
   }
  }

  return {
    isSuccess,
    onReset
  }
}

export  {useLogin , useSignUp , useForgotPassword ,useResetPassword}