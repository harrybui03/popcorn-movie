import React from "react";
import Header from './Header'
import Footer from '../defaultPage/Footer'
import LogOut from './LogOut'
import ChangePW from './ChangePW'
import { useLocation } from 'react-router-dom';
import './styles.css'
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faExclamationCircle } from "@fortawesome/free-solid-svg-icons";

function PaymentFail(params) {
    const location = useLocation();
    const searchParams = new URLSearchParams(location.search);
    const token = searchParams.get('token');
    const payerID = searchParams.get('PayerID');

    return (
        <>
            <Header {...params} ></Header>
            <div style={{maxWidth: '40rem'}} className="container bg-light mx-auto my-3 p-5 rounded-5 border shadow">
                <h3>Thanh toán thất bại!</h3>
                <FontAwesomeIcon icon={faExclamationCircle} className="display-1 border border-danger rounded-circle bg-danger text-light m-3" />
                <p>Cảm ơn bạn đã đến với H3DC</p>
                <p>Vui lòng đặt lại đơn hàng khác.</p>
                <a href="/" style={{width: '10rem'}} className="btn m-3 btn-primary">Đặt lại</a>
                <a href="/" style={{width: '10rem'}} className="btn m-3 btn-primary">Về trang chủ</a>
            </div>
            <LogOut></LogOut>
            <ChangePW {...params}></ChangePW>
            <Footer></Footer>
        </>
    )
}

export default PaymentFail;