import React, { useEffect, useState, useRef } from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faXmark } from "@fortawesome/free-solid-svg-icons";

function RequireLogin() {
    const handleRequireLogin = () => {
        $('#requireLogin').modal('hide');
        $('.login-modal').addClass('display');
    }
    useEffect(() => {
        return () => {
            $('.close-requireLogin').on('click', function () {
                $('#requireLogin').modal('hide');
            });
        }
    }, []);
    return (
        <div className="modal fade" id="requireLogin" tabIndex="-1" role="dialog" aria-labelledby="exampleModalLabel"
            aria-hidden="true">
            <div className="modal-dialog" role="document">
                <div className="modal-content">
                    <div className="modal-header">
                        <h5 className="modal-title" id="exampleModalLabel">Thông báo</h5>
                        <button type="button" className="btn btn-secondary close-requireLogin" data-dismiss="modal" aria-label="Close">
                            <FontAwesomeIcon icon={faXmark} />
                        </button>
                    </div>
                    <div className="modal-body">
                        Vui lòng đăng nhập trước khi thực hiện thao tác này!
                    </div>
                    <div className="modal-footer">
                        <button type="button" className="btn btn-secondary close-requireLogin" data-dismiss="modal">Hủy</button>
                        <button type="button" onClick={handleRequireLogin} className="btn btn-primary accept-requireLogin">Đăng nhập</button>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default RequireLogin