import { useMutation } from "@tanstack/react-query";
import useGraphql from "../graphql";
import { fetchGraphQLWithToken } from "../../../services/graphql-services";
import { toast } from "react-toastify";


function useCreateTransaction() {
    const { createTransaction } = useGraphql();
    const { mutate } = useMutation({
        mutationFn: (input) =>
            fetchGraphQLWithToken(createTransaction.query, {
                input: input
            }),
        onSuccess: (data) => {
            console.log(data?.CreateTransaction?.CheckoutUrl)
            console.log('Redirecting to:', data?.CreateTransaction?.CheckoutUrl);
            window.location.href = data?.CreateTransaction?.CheckoutUrl; // Redirect to the CheckoutUrl
        },
        onError: (error) => {
            toast.error('Fail to payment');
        }
    });

    function onCreateTransaction({ ticketIDs, foods }) {
        try {
            mutate({ ticketIDs, foods });
        } catch (error) {
            console.log('Error during payment');
        }
    }

    return { onCreateTransaction };
}


function useChangePassword(setIsChangePWSuccess , setChangePWErrorMessage) {
    const {changePassword} = useGraphql()
    const {mutate} = useMutation({
        mutationFn:(input) =>
            fetchGraphQLWithToken(changePassword.query , {
                input:input
            }),
        onSuccess: (data) => {
            toast.success('Đổi mật khẩu thành công')
        },
        onError: (data) => {
            toast.error('Đổi mật khẩu thất bại, vui lòng kiểm tra mật khẩu cũ và thử lại.')
        }
    })

    function onChangePassword({oldPassword , newPassword , confirmNewPassword}) {
        setIsChangePWSuccess(true);
        setChangePWErrorMessage("");
        const passwordRegex = /^(?=.*[A-Z])(?=.*[a-z])(?=.*\d)(?=.*[!@#$%^&*()_+{}[\]:;<>,.?~\\\/-]).{8,}$/;
        if (oldPassword === '' || newPassword === '' || confirmNewPassword === '') {
            setIsChangePWSuccess(false);
            setChangePWErrorMessage("Vui lòng điền đủ thông tin.");
        } else if (newPassword !== confirmNewPassword) {
            setIsChangePWSuccess(false);
            setChangePWErrorMessage("Vui lòng xác nhận lại mật khẩu.");
        } else if (newPassword === oldPassword) {
            setIsChangePWSuccess(false);
            setChangePWErrorMessage("Mật khẩu mới phải khác mật khẩu cũ.");
        } else if (!passwordRegex.test(newPassword)) {
            setIsChangePWSuccess(false);
            setChangePWErrorMessage("Mật khẩu ít nhất 8 kí tự (phải bao gồm chữ hoa, chữ thường, chữ số và kí tự đặt biệt).");
        } else {
            try {
                mutate({oldPassword , newPassword , confirmNewPassword})
            } catch (error) {
                console.error('Error during change password:', error);
            }
        }
    }

    return {onChangePassword}
}

export {useCreateTransaction , useChangePassword}