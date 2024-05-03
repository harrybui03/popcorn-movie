import React, { useEffect, useState, useRef } from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faXmark } from "@fortawesome/free-solid-svg-icons";
import { faFacebookSquare, faGoogle } from "@fortawesome/free-brands-svg-icons";
import useAuth from '../../hooks/useAuth';

function LoginSignup() {
    const [isChecked, setIsChecked] = useState(true);
    const [isLoginSuccess, setIsLoginSuccess] = useState(true);
    const [loginErrorMessage, setLoginErrorMessage] = useState("");
    const [isSignupSuccess, setIsSignupSuccess] = useState(true);
    const [signupErrorMessage, setSignupErrorMessage] = useState("");
    const handleCheckboxChange = () => {
        setIsSignupSuccess(true);
        setSignupErrorMessage("");
        setIsLoginSuccess(true);
        setLoginErrorMessage("");
        setIsChecked(!isChecked);
    };
    const auth = useAuth();
    const [email, setEmail] = useState('');
    const [pwd, setPwd] = useState('');
    const handleSignIn = async () => {
        const passwordRegex = /^(?=.*[A-Z])(?=.*[a-z])(?=.*\d)(?=.*[!@#$%^&*()_+{}[\]:;<>,.?~\\\/-]).{8,}$/;
        if (email === '' || pwd === '') {
            setIsLoginSuccess(false);
            setLoginErrorMessage("Vui lòng điền đủ thông tin.");
        } else if (!passwordRegex.test(pwd)) {
            setIsLoginSuccess(false);
            setLoginErrorMessage("Mật khẩu ít nhất 8 kí tự (phải bao gồm chữ hoa, chữ thường, chữ số và kí tự đặt biệt)");
        } else {
            try {
                const response = await fetch('http://localhost:8080/auth/login', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({ email, password: pwd }),
                });
                const data = await response.json();
                if (response.status === 200) {
                    setIsLoginSuccess(true);
                    setLoginErrorMessage("");
                    auth.setAccessToken(data.accessToken);
                    window.location.href = '/';
                } else {
                    console.log(data);
                    setIsLoginSuccess(false);
                    setLoginErrorMessage("Email hoặc mật khẩu không hợp lệ, vui lòng kiểm tra lại.");
                }

            } catch (error) {
                console.error('Error during sign-in:', error);
            }
        }

    }

    const [emailSU, setEmailSU] = useState('');
    const [pwdSU, setPwdSU] = useState('');
    const [username, setUsername] = useState('');
    const [confirmPwd, setConfirmPwd] = useState('');
    const handleSignUp = async () => {
        const passwordRegex = /^(?=.*[A-Z])(?=.*[a-z])(?=.*\d)(?=.*[!@#$%^&*()_+{}[\]:;<>,.?~\\\/-]).{8,}$/;
        if (emailSU === '' || pwdSU === '' || confirmPwd === '' || username === '') {
            setIsSignupSuccess(false);
            setSignupErrorMessage("Vui lòng điền đủ thông tin.");
        } else if (pwdSU !== confirmPwd) {
            setIsSignupSuccess(false);
            setSignupErrorMessage("Vui lòng xác nhận lại mật khẩu.");
        } else if (!passwordRegex.test(pwdSU)) {
            console.log(pwdSU, passwordRegex.test(pwdSU));
            setIsSignupSuccess(false);
            setSignupErrorMessage("Mật khẩu ít nhất 8 kí tự (phải bao gồm chữ hoa, chữ thường, chữ số và kí tự đặt biệt)");
        } else {
            try {
                const response = await fetch('http://localhost:8080/auth/signup', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({ email: emailSU, password: pwdSU, displayName: username }),
                });
                const data = await response.json();
                console.log(data);
                if (response.status === 200) {
                    if (data.message === "Email is already in use") {
                        setIsSignupSuccess(false);
                        setSignupErrorMessage("Email đã được đăng ký.");
                    } else {
                        setIsSignupSuccess(true);
                        setSignupErrorMessage("");
                        alert("Đăng ký thành công, vui lòng kiểm tra email để kích hoạt tài khoản");
                    }
                } else {
                    setIsSignupSuccess(false);
                    setSignupErrorMessage("Đăng ký thất bại, vui lòng thử lại.");
                }

            } catch (error) {
                console.error('Error during sign-in:', error);
            }
        }

    }


    return (
        <div className="login-modal">
            <div className="row">
                <div className="col-md-6 mx-auto" >
                    <div className="login-box login-modal-container" style={{ backgroundColor: "#f0f0f0" }}>
                        <button className="login-close-btn">
                            <FontAwesomeIcon icon={faXmark} />
                        </button>
                        <div className="login-snip py-3">
                            <input id="tab-1" type="radio" name="tab" className="sign-in" checked={isChecked}
                                onChange={handleCheckboxChange} /><label htmlFor="tab-1" className="tab">Đăng nhập</label>
                            <input id="tab-2" type="radio" name="tab" className="sign-up" checked={!isChecked}
                                onChange={handleCheckboxChange} /><label htmlFor="tab-2" className="tab">Đăng ký</label>
                            <div className="login-space">
                                <div className="login">
                                    <div className="group group-login-signup">
                                        <label htmlFor="user-login" className="label">Email</label>
                                        <input onChange={(e) => setEmail(e.target.value)} id="user-login" type="text" className="input py-2" placeholder="Nhập email" />
                                    </div>
                                    <div className="group group-login-signup">
                                        <label htmlFor="pass-login" className="label">Mật khẩu</label>
                                        <input onChange={(e) => setPwd(e.target.value)} id="pass-login" type="password" className="input py-2" data-type="password" placeholder="Nhập mật khẩu" />
                                    </div>
                                    <div className="group group-login-signup d-flex justify-content-start">
                                        <input id="check" type="checkbox" className="check" />
                                        <label htmlFor="check"><span className="icon"></span> Lưu đăng nhập</label>
                                    </div>
                                    <div className="group group-login-signup">
                                        <button onClick={handleSignIn} className="button py-2">Đăng nhập</button>
                                    </div>
                                    {!isLoginSuccess && (
                                        <div style={{ color: 'red', marginTop: '10px' }}>
                                            {loginErrorMessage}
                                        </div>
                                    )}
                                    <div className="foot d-flex justify-content-between">
                                        <label htmlFor="tab-2">Chưa có tài khoản?</label>
                                        <a href="#">Quên mật khẩu?</a>
                                    </div>
                                    <hr className="text-white"></hr>
                                    <div>
                                        <h5 className="m-4 text-light">Hoặc đăng nhập với</h5>
                                        <div className="d-flex justify-content-center">
                                            <h3 className="px-1 mx-3 rounded bg-primary text-white">
                                                <FontAwesomeIcon icon={faFacebookSquare} />
                                            </h3>
                                            <h3 className="px-1 mx-3 rounded bg-danger text-white">
                                                <FontAwesomeIcon icon={faGoogle} />
                                            </h3>
                                        </div>
                                    </div>
                                </div>
                                <div className="sign-up-form">
                                    <div className="group group-login-signup">
                                        <label htmlFor="user" className="label">Tên đăng nhập</label>
                                        <input id="user" onChange={(e) => setUsername(e.target.value)} type="text" className="input py-2" placeholder="Nhập tên đăng nhập" />
                                    </div>
                                    <div className="group group-login-signup">
                                        <label htmlFor="pass" className="label">Mật khẩu</label>
                                        <input id="pass" onChange={(e) => setPwdSU(e.target.value)} type="password" className="input py-2" data-type="password" placeholder="Nhập mật khẩu" />
                                    </div>
                                    <div className="group group-login-signup">
                                        <label htmlFor="confirmpass" className="label">Xác nhận mật khẩu</label>
                                        <input id="confirmpass" onChange={(e) => setConfirmPwd(e.target.value)} type="password" className="input py-2" data-type="password" placeholder="Nhập lại mật khẩu" />
                                    </div>
                                    <div className="group group-login-signup">
                                        <label htmlFor="email" className="label">Email</label>
                                        <input id="email" onChange={(e) => setEmailSU(e.target.value)} type="text" className="input py-2" placeholder="Nhập email" />
                                    </div>
                                    <div className="group group-login-signup my-4">
                                        <button onClick={handleSignUp} className="button py-2">Đăng ký</button>
                                    </div>
                                    {!isSignupSuccess && (
                                        <div style={{ color: 'red', marginTop: '10px' }}>
                                            {signupErrorMessage}
                                        </div>
                                    )}
                                    <div className="foot d-flex justify-content-end">
                                        <label htmlFor="tab-1">Đã có tài khoản?</label>
                                    </div>
                                    <hr className="text-white"></hr>
                                    <div>
                                        <h5 className="m-4 text-light">Hoặc đăng ký với</h5>
                                        <div className="d-flex justify-content-center">
                                            <h3 className="px-1 mx-3 rounded bg-primary text-white">
                                                <FontAwesomeIcon icon={faFacebookSquare} />
                                            </h3>
                                            <h3 className="px-1 mx-3 rounded bg-danger text-white">
                                                <FontAwesomeIcon icon={faGoogle} />
                                            </h3>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>)
}

export default LoginSignup