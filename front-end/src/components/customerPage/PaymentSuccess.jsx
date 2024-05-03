import React, { useEffect, useState } from "react";
import Header from './Header'
import Footer from '../defaultPage/Footer'
import LogOut from './LogOut'
import ChangePW from './ChangePW'
import { useLocation } from 'react-router-dom';
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faCheckCircle } from "@fortawesome/free-solid-svg-icons";
import useAuth from "../../hooks/useAuth";

function PaymentSuccess(params) {
    const auth = useAuth();
    const location = useLocation();
    const searchParams = new URLSearchParams(location.search);
    const token = searchParams.get('token');
    const payerID = searchParams.get('PayerID');
    const product = localStorage.getItem('listItem');
    const checkoutCapture = async (data) => {
        try {
            const response = await fetch(`http://localhost:8080/checkout/capture?token=${token}`, {
                method: 'POST',
                headers: {
                    'Authorization': `Bearer ${auth.accessToken}`,
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(data)
            });

            if (!response.ok) {
                throw new Error(`HTTP error! Status: ${response.status}`);
            }
            return response;
        } catch (error) {
            console.error('Error:', error);
            throw error;
        }
    }

    useEffect(() => {
        const js = JSON.parse(product);
        localStorage.removeItem('listItem');
        checkoutCapture(js);
    }, []);

    return (
        <>
            <Header {...params} ></Header>
            <div style={{ maxWidth: '40rem' }} className="container bg-light mx-auto my-3 p-5 rounded-5 border shadow">
                <h3>Thanh toán thành công!</h3>
                <FontAwesomeIcon icon={faCheckCircle} className="display-1 border border-success rounded-circle bg-success text-light m-3" />
                <p>Cảm ơn bạn đã đến với H3DC</p>
                <p>Đơn hàng của bạn đã được xác nhận.</p>
                <div className="border shadow border-black rounded-5 mx-auto p-3" style={{ maxWidth: '25rem', backgroundColor: '#d0f9ef' }}>
                    <p>Mã đơn hàng: {token}</p>
                    <p>ID Người thanh toán: {payerID}</p>
                </div>
                <a href="/history-booking" style={{ width: '10rem' }} className="btn m-3 btn-primary">Xem lịch sử đặt vé</a>
                <a href="/" style={{ width: '10rem' }} className="btn m-3 btn-primary">Về trang chủ</a>
            </div>
            <LogOut></LogOut>
            <ChangePW {...params}></ChangePW>
            <Footer></Footer>
        </>
    )
}

export default PaymentSuccess;