import React, { useEffect, useState } from "react";
import Header from './Header'
import Footer from '../defaultPage/Footer'
import LogOut from './LogOut'
import ChangePW from './ChangePW'
import LoadingSpinner from "../defaultPage/Loading";
import useAuth from "../../hooks/useAuth";
import { useGetAllTransactions } from "./hook/useQuery";


function HistoryBooking(params) {
    const auth = useAuth();
    console.log(auth)
    const historyData = useGetAllTransactions(auth.id)
    const history = historyData?.data??[]
    console.log(history)
    const convertTimeStringToDate = (dateTimeString) => {
        const date = new Date(dateTimeString);

        // Extract date components
        const year = date.getUTCFullYear();
        const month = String(date.getUTCMonth() + 1).padStart(2, '0');
        const day = String(date.getUTCDate()).padStart(2, '0');

        // Extract time components
        const hours = String(date.getUTCHours()).padStart(2, '0');
        const minutes = String(date.getUTCMinutes()).padStart(2, '0');
        const seconds = String(date.getUTCSeconds()).padStart(2, '0');

        // Format the date and time
        const formattedDate = `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;

        return formattedDate;
    }

    return (
        <>
            <Header {...params} ></Header>
            <div className="container my-5 mx-auto">
                {history === null ? (
                    <div className='d-flex justify-content-center my-5'>
                        <LoadingSpinner />
                    </div>
                ) : history.length > 0 ? (
                    <table className="table table-hover table-bordered shadow" style={{ borderRadius: '10px', overflow: 'hidden' }}>
                        <thead className="table-primary">
                            <tr style={{ cursor: 'context-menu' }}>
                                <th scope="col">
                                    <div className='d-flex justify-content-center align-items-center'>
                                        <div className='mx-2'>Mã giao dịch</div>
                                        {/* <FontAwesomeIcon icon={sortBy.by !== 'id' ? faUnsorted : sortBy.isDown ? faSortDown : faSortUp} /> */}
                                    </div>
                                </th>
                                <th scope="col">
                                    <div className='d-flex justify-content-center align-items-center'>
                                        <div className='mx-2'>Thời gian</div>
                                        {/* <FontAwesomeIcon icon={sortBy.by !== 'title' ? faUnsorted : sortBy.isDown ? faSortDown : faSortUp} /> */}
                                    </div>
                                </th>
                                <th scope="col">
                                    <div className='d-flex justify-content-center align-items-center'>
                                        <div className='mx-2'>Tổng tiền</div>
                                        {/* <FontAwesomeIcon icon={sortBy.by !== 'openingDay' ? faUnsorted : sortBy.isDown ? faSortDown : faSortUp} /> */}
                                    </div>
                                </th>
                            </tr>
                        </thead>
                        <tbody>
                            {history.map((item) => (
                                <tr key={item.id}>
                                    <th className="align-middle" scope="row">{item.id}</th>
                                    <td className="align-middle">{convertTimeStringToDate(item.createdAt)}</td>
                                    <td className="align-middle">{item.total} <u>đ</u></td>
                                </tr>
                            ))}
                        </tbody>
                    </table>

                ) : (
                    <table className="table table-hover table-bordered shadow" style={{ borderRadius: '10px', overflow: 'hidden' }}>
                        <thead className="table-primary">
                            <tr style={{ cursor: 'context-menu' }}>
                                <th scope="col">
                                    <div className='d-flex justify-content-center align-items-center'>
                                        <div className='mx-2'>Mã giao dịch</div>
                                        {/* <FontAwesomeIcon icon={sortBy.by !== 'id' ? faUnsorted : sortBy.isDown ? faSortDown : faSortUp} /> */}
                                    </div>
                                </th>
                                <th scope="col">
                                    <div className='d-flex justify-content-center align-items-center'>
                                        <div className='mx-2'>Thời gian</div>
                                        {/* <FontAwesomeIcon icon={sortBy.by !== 'title' ? faUnsorted : sortBy.isDown ? faSortDown : faSortUp} /> */}
                                    </div>
                                </th>
                                <th scope="col">
                                    <div className='d-flex justify-content-center align-items-center'>
                                        <div className='mx-2'>Tổng tiền</div>
                                        {/* <FontAwesomeIcon icon={sortBy.by !== 'openingDay' ? faUnsorted : sortBy.isDown ? faSortDown : faSortUp} /> */}
                                    </div>
                                </th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr>
                                <th colSpan={3} className="align-middle" scope="row">Chưa có giao dịch nào</th>
                            </tr>
                        </tbody>
                    </table>
                )}
            </div>
            <LogOut></LogOut>
            <ChangePW {...params}></ChangePW>
            <Footer></Footer>
        </>
    );
}

export default HistoryBooking;