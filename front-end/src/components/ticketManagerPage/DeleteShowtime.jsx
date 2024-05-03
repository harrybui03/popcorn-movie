import React, { useEffect, useState, useRef } from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faXmark } from "@fortawesome/free-solid-svg-icons";
import useAuth from "../../hooks/useAuth";


function DeleteShowtime({ deleteId, setNotification }) {
    const auth = useAuth();

    async function deleteShowtime() {
        try {
            const response = await fetch(`http://localhost:8080/showtimes/${deleteId}`, {
                method: 'DELETE',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${auth.accessToken}`,
                },
            });

            if (!response.ok) {
                throw new Error(`HTTP error! Status: ${response.status}`);
            }
            const data = await response.json();
            if (data.message === 'Some tickets of movies are booked. Can not delete this showtime') {
                setNotification({ title: 'Thông báo', body: 'Đã có người đặt vé, không thể xóa!', footer: 'OK', status: 'danger' });
                $('#deleteShowtime').modal('hide');
                $('#notification').modal('show');
            } else {
                setNotification({ title: 'Thông báo', body: 'Xóa lịch chiếu thành công', footer: 'OK', status: 'success' });
                $('#deleteShowtime').modal('hide');
                $('#notification').modal('show');
            }
        } catch (error) {
            console.error('Error:', error);
            throw error;
        }
    }

    const handleDeleteShowtime = async () => {
        try {
            const response = await fetch(`http://localhost:8080/tickets?showtime=${deleteId}`, {
                method: 'DELETE',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${auth.accessToken}`,
                },
            });

            const data = await response.json();
            if (data.message === 'There are some booked ticket') {
                setNotification({ title: 'Thông báo', body: 'Đã có người đặt vé, không thể xóa!', footer: 'OK', status: 'danger' });
                $('#deleteShowtime').modal('hide');
                $('#notification').modal('show');
            } else {
                deleteShowtime();
            }
        } catch (error) {
            console.error('Error:', error);
            throw error;
        }
    }

    useEffect(() => {
        return () => {
            $('.close-deleteShowtime').on('click', function () {
                $('#deleteShowtime').modal('hide');
            });
        };
    }, []);
    return (
        <div className="modal fade" id="deleteShowtime" tabIndex="-1" role="dialog" aria-labelledby="exampleModalLabel"
            aria-hidden="true">
            <div className="modal-dialog" role="document">
                <div className="modal-content">
                    <div className="modal-header">
                        <h5 className="modal-title" id="exampleModalLabel">Xác nhận xóa</h5>
                        <button type="button" className="btn btn-secondary close-deleteShowtime" data-dismiss="modal" aria-label="Close">
                            <FontAwesomeIcon icon={faXmark} />
                        </button>
                    </div>
                    <div className="modal-body">
                        Bạn có chắc muốn xóa lịch chiếu này?
                    </div>
                    <div className="modal-footer">
                        <button type="button" className="btn btn-secondary close-deleteShowtime" data-dismiss="modal">Hủy</button>
                        <button type="button" onClick={handleDeleteShowtime} className="btn btn-primary">Xóa</button>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default DeleteShowtime