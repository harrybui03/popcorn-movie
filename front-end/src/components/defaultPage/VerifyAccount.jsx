import React, { useEffect, useState } from "react";

import { useLocation } from 'react-router-dom';
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faCheckCircle } from "@fortawesome/free-solid-svg-icons";

function VerifyAccount() {
    const location = useLocation();
    const searchParams = new URLSearchParams(location.search);
    const token = searchParams.get('token');

    const verifyAccount = async () => {
        try {
            const response = await fetch(`http://localhost:8080/auth/verify-account?token=${token}`);
            const data = await response.json();
            console.log(data);
        } catch (error) {
            console.error('Error:', error);
            throw error;
        }
    }

    useEffect(() => {
        verifyAccount();
    }, []);

    return (
        <>
            <div style={{ maxWidth: '40rem', height: '95vh'}} className="container bg-light mx-auto my-3 p-5 rounded-5 border shadow">
                <h3>Tài khoản của bạn đã được kích hoạt!</h3>
                <FontAwesomeIcon icon={faCheckCircle} className="display-1 border border-success rounded-circle bg-success text-light m-3" />
                <p>Cảm ơn bạn đã đến với H3DC</p>
                <p>Vui lòng đăng nhập lại để đặt vé.</p>
                <a href="/" style={{ width: '10rem' }} className="btn m-3 btn-primary">Về trang chủ</a>
            </div>
        </>
    )
}

export default VerifyAccount;