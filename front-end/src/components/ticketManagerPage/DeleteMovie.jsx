import React, { useEffect, useState, useRef } from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faXmark } from "@fortawesome/free-solid-svg-icons";
import useAuth from "../../hooks/useAuth";


function DeleteMovie({ deleteId, setNotification }) {
    const auth = useAuth();

    const handleDeleteMovie = async () => {
        try {
            const response = await fetch(`http://localhost:8080/movies/${deleteId}`, {
                method: 'Delete',
                headers: {
                    'Authorization': `Bearer ${auth.accessToken}`,
                }
            });
            const result = await response.json();
            if (result.message === 'Delete success fully') {
                setNotification({ title: 'Thông báo', body: 'Đã xóa phim thành công!', footer: 'OK', status: 'success' });
                $('#deleteMovie').modal('hide');
                $('#notification').modal('show');
            } else {
                setNotification({ title: 'Thông báo', body: 'Phim đã có lịch chiếu, chỉ có thể xóa phim chưa có lịch chiếu!', footer: 'OK', status: 'danger' });
                $('#deleteMovie').modal('hide');
                $('#notification').modal('show');
            }
        } catch (error) {
            console.error('Error:', error);
            throw error;
        }
    }

    useEffect(() => {
        return () => {
            $('.close-deleteMovie').on('click', function () {
                $('#deleteMovie').modal('hide');
            });
        };
    }, []);
    return (
        <div className="modal fade" id="deleteMovie" tabIndex="-1" role="dialog" aria-labelledby="exampleModalLabel"
            aria-hidden="true">
            <div className="modal-dialog" role="document">
                <div className="modal-content">
                    <div className="modal-header">
                        <h5 className="modal-title" id="exampleModalLabel">Xác nhận xóa</h5>
                        <button type="button" className="btn btn-secondary close-deleteMovie" data-dismiss="modal" aria-label="Close">
                            <FontAwesomeIcon icon={faXmark} />
                        </button>
                    </div>
                    <div className="modal-body">
                        Bạn có chắc muốn xóa phim này?
                    </div>
                    <div className="modal-footer">
                        <button type="button" className="btn btn-secondary close-deleteMovie" data-dismiss="modal">Hủy</button>
                        <button type="button" onClick={handleDeleteMovie} className="btn btn-primary">Xóa</button>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default DeleteMovie