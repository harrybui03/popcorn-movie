import React, { useEffect, useState, useRef } from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faTimes } from "@fortawesome/free-solid-svg-icons";
import useAuth from '../../hooks/useAuth';
import './styles.css';

function ChangePW(params) {
    const auth = useAuth()
    const [isChangePWSuccess, setIsChangePWSuccess] = useState(true);
    const [ChangePWErrorMessage, setChangePWErrorMessage] = useState("");
    const [oldPW, setOldPW] = useState('');
    const [newPW, setNewPW] = useState('');
    const [confirmPW, setConfirmPW] = useState('');

    const handleChangePW = async () => {
        setIsChangePWSuccess(true);
        setChangePWErrorMessage("");
        const passwordRegex = /^(?=.*[A-Z])(?=.*[a-z])(?=.*\d)(?=.*[!@#$%^&*()_+{}[\]:;<>,.?~\\\/-]).{8,}$/;
        if (oldPW === '' || newPW === '' || confirmPW === '') {
            setIsChangePWSuccess(false);
            setChangePWErrorMessage("Vui lòng điền đủ thông tin.");
        } else if (newPW !== confirmPW) {
            setIsChangePWSuccess(false);
            setChangePWErrorMessage("Vui lòng xác nhận lại mật khẩu.");
        } else if (newPW === oldPW) {
            setIsChangePWSuccess(false);
            setChangePWErrorMessage("Mật khẩu mới phải khác mật khẩu cũ.");
        } else if (!passwordRegex.test(newPW)) {
            setIsChangePWSuccess(false);
            setChangePWErrorMessage("Mật khẩu ít nhất 8 kí tự (phải bao gồm chữ hoa, chữ thường, chữ số và kí tự đặt biệt).");
        } else {
            try {
                const response = await fetch('http://localhost:8080/auth/forgot-password', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({ email: params.email, oldPassword: oldPW, newPassword: newPW }),
                });
                if (response.status === 200) {
                    alert('Đổi mật khẩu thành công');
                    const response = await fetch('http://localhost:8080/auth/login', {
                        method: 'POST',
                        headers: {
                            'Content-Type': 'application/json',
                        },
                        body: JSON.stringify({ email: params.email, password: newPW }),
                    });
                    if (response.status === 200) {
                        const data = await response.json();
                        auth.setAccessToken(data.accessToken);
                        window.location.href = '/';
                    } else {
                        console.error('Error during change password');
                    }
                } else {
                    setIsChangePWSuccess(false);
                    setChangePWErrorMessage("Đổi mật khẩu thất bại, vui lòng kiểm tra mật khẩu cũ và thử lại.");
                }

            } catch (error) {
                console.error('Error during change password:', error);
            }
        }
    }

    useEffect(() => {
        return () => {
            $('.close-changePW').on('click', function () {
                $('#changePWModal').modal('hide');
            });
        }
    }, []);

    return (
        <div className="modal fade" id="changePWModal" tabIndex="-1" role="dialog" aria-labelledby="exampleModalLabel"
            aria-hidden="true">
            <div className="modal-dialog" role="document">
                <div className="modal-content">
                    <div className="modal-header">
                        <h5 className="modal-title" id="exampleModalLabel">Đổi mật khẩu</h5>
                        <button type="button" className="btn btn-secondary close-changePW" data-dismiss="modal" aria-label="Close">
                            <FontAwesomeIcon icon={faTimes} />
                        </button>
                    </div>
                    <div className="modal-body d-flex flex-column justify-content-between align-items-start">
                        <label htmlFor="oldPass" style={{ display: 'block' }} className="label">Mật khẩu cũ</label>
                        <input id="oldPass" style={{ width: '100%' }} onChange={(e) => setOldPW(e.target.value)} type="password" className="input py-2 text-dark bg-light" data-type="password" placeholder="Nhập mật khẩu cũ" />

                        <label htmlFor="newPass" style={{ display: 'block' }} className="label">Mật khẩu mới</label>
                        <input id="newPass" style={{ width: '100%' }} onChange={(e) => setNewPW(e.target.value)} type="password" className="input py-2 text-dark bg-light" data-type="password" placeholder="Nhập mật khẩu mới" />

                        <label htmlFor="confirmNewPass" style={{ display: 'block' }} className="label">Xác nhận mật khẩu</label>
                        <input id="confirmNewPass" style={{ width: '100%' }} onChange={(e) => setConfirmPW(e.target.value)} type="password" className="input py-2 text-dark bg-light" data-type="password" placeholder="Nhập lại mật khẩu mới" />

                        {!isChangePWSuccess && (
                            <div style={{ color: 'red', marginTop: '10px' }}>
                                {ChangePWErrorMessage}
                            </div>
                        )}
                    </div>
                    <div className="modal-footer">
                        <button type="button" className="btn btn-secondary close-changePW" data-dismiss="modal">Hủy</button>
                        <button type="button" onClick={handleChangePW} className="btn btn-primary">Xác nhận</button>
                    </div>
                </div>
            </div>
        </div>)
}

export default ChangePW