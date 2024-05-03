import React, { useEffect, useState } from "react";
import Header from './Header'
import Footer from '../defaultPage/Footer'
import LogOut from './LogOut'
import ChangePW from './ChangePW'
import LoadingSpinner from "../defaultPage/Loading";
import useAuth from "../../hooks/useAuth";


function HistoryBooking(params) {
    const auth = useAuth();
    const [history, setHistory] = useState(null);

    const getHistory = async () => {
        try {
            const response = await fetch(`http://localhost:8080/transactions/history`, {
                method: 'GET',
                headers: {
                    'Authorization': `Bearer ${auth.accessToken}`,
                },
            });

            if (!response.ok) {
                throw new Error(`HTTP error! Status: ${response.status}`);
            }
            const data = await response.json();
            data.sort((a, b) => {
                const date1 = new Date(a.createdAt);
                const date2 = new Date(b.createdAt);
                return date2 - date1;
            })

            setHistory(data);
        } catch (error) {
            console.error('Error:', error);
            throw error;
        }
    }

    useEffect(() => {
        getHistory();
    }, []);

    const convertTimeStringToDate = (dateTimeString) => {
        return dateTimeString.replace('T', ' ');
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