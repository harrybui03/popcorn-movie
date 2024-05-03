import React, { useEffect, useState, useRef } from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faXmark } from "@fortawesome/free-solid-svg-icons";
import useAuth from '../../hooks/useAuth';

function LoginSignup() {
    const auth = useAuth()

    const handleLogout = () => {
        auth.setAccessToken('');
        window.location.href = '/';
    }

    useEffect(() => {
        return () => {
            $('.close-logout').on('click', function () {
                $('#logoutModal').modal('hide');
            });
        }
    }, []);

    return (
        <div className="modal fade" id="logoutModal" tabIndex="-1" role="dialog" aria-labelledby="exampleModalLabel"
            aria-hidden="true">
            <div className="modal-dialog" role="document">
                <div className="modal-content">
                    <div className="modal-header">
                        <h5 className="modal-title" id="exampleModalLabel">Xác nhận</h5>
                        <button type="button" className="btn btn-secondary close-logout" data-dismiss="modal" aria-label="Close">
                            <FontAwesomeIcon icon={faXmark} />
                        </button>
                    </div>
                    <div className="modal-body">
                        Bạn có chắc muốn đăng xuất?
                    </div>
                    <div className="modal-footer">
                        <button type="button" className="btn btn-secondary close-logout" data-dismiss="modal">Hủy</button>
                        <button type="button" onClick={handleLogout} className="btn btn-primary">Đăng xuất</button>
                    </div>
                </div>
            </div>
        </div>)
}

export default LoginSignup