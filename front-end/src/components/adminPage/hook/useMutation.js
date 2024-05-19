import { useMutation, useQueryClient } from "@tanstack/react-query";
import useGraphql from "../graphql";
import { useState } from "react";
import { fetchGraphQLWithToken } from "../../../services/graphql-services";

function useCreateUser({onSuccess , onError}){
    
    const [messageInsertUser, setMessageInsertUser] = useState({ isShow: false, text: '', success: false });

    const queryClient = useQueryClient()
    const { createUser , queryKey } = useGraphql();
    const { mutate , isPending } = useMutation({
      mutationKey:[queryKey],
      mutationFn: (input) =>
        fetchGraphQLWithToken(createUser.query, {
          input: input,
        }),
      onSuccess: (data) => {
        onSuccess()
        queryClient.invalidateQueries({queryKey:[queryKey]})
      },
      onError: () => {
        onError()
      }
    })

    function handleCreateUser(data){
        const passwordRegex = /^(?=.*[A-Z])(?=.*[a-z])(?=.*\d)(?=.*[!@#$%^&*()_+{}[\]:;<>,.?~\\\/-]).{8,}$/;
        if (data.email === '' || data.password === '' || data.confirmPassword === '' || data.displayName === '') {
            setMessageInsertUser({ isShow: true, text: "Vui lòng điền đủ thông tin.", success: false });
        } else if (data.password !== data.confirmPassword) {
            setMessageInsertUser({ isShow: true, text: "Vui lòng xác nhận lại mật khẩu.", success: false });
        } else if (!passwordRegex.test(data.password)) {
            setIsLoginSuccess(false);
            setLoginErrorMessage("Mật khẩu ít nhất 8 kí tự (phải bao gồm chữ hoa, chữ thường, chữ số và kí tự đặt biệt)");
        } else {
            try {
                const {confirmPassword,...other} = data
                mutate(other)
            } catch (error) {
                console.error('Error Insert:', error);
            }
        }
    }

    return {
        handleCreateUser,
        messageInsertUser,
        isPending
    }
}

function useLockUser() {
    const queryClient = useQueryClient()
    const { updateUser , queryKey } = useGraphql();
    const { mutate  } = useMutation({
      mutationKey:[queryKey],
      mutationFn: (input) =>
        fetchGraphQLWithToken(updateUser.query, {
          input: input,
        }),
      onSuccess: (data) => {
        queryClient.invalidateQueries({queryKey:[queryKey]})
      },
    })

    function handleLockUser(data) {
        try {
            mutate(data)
        } catch (error) {
            console.error('Error Update:', error);
        }
    }

    return {
        handleLockUser
    }
}

export {useCreateUser ,useLockUser }