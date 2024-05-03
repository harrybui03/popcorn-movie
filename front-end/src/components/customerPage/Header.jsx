import React, { useEffect, useState, useRef } from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faUser } from "@fortawesome/free-solid-svg-icons";
import SearchButton from '../defaultPage/SearchButton'

function HomepageHeader(params) {
    useEffect(() => {
        return () => {
            $('.btn-logout').on('click', function () {
                $('#logoutModal').modal('show');
            });
            $('.close-logout').on('click', function () {
                $('#logoutModal').modal('hide');
            });
            $('.btn-changePW').on('click', function () {
                $('#changePWModal').modal('show');
            });
            $('.close-changePW').on('click', function () {
                $('#changePWModal').modal('hide');
            });
        }
    }, []);
    return (
        <nav className="navbar navbar-expand-lg navbar-light" style={{ backgroundColor: '#2b1a1a' }}>
            <div className="container-fluid d-flex justify-content-center align-items-center">
                <a className="navbar-brand" href="/"><img src="/Logo.png" style={{ height: '3rem' }} alt='logo'></img></a>
                <button className="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNavDropdown" aria-controls="navbarNavDropdown" aria-expanded="false" aria-label="Toggle navigation">
                    <span className="navbar-toggler-icon"></span>
                </button>
                <div className="collapse navbar-collapse" id="navbarNavDropdown">
                    <ul className="navbar-nav nav-bar">
                        <li className="nav-item">
                            <a className="nav-link text-light active" aria-current="page" href="/">Trang chủ</a>
                        </li>
                        <li className="nav-item">
                            <a className="nav-link text-light" href="/history-booking">Lịch sử đặt vé</a>
                        </li>
                    </ul>
                </div>
                <div>
                    <SearchButton />
                </div>
                <div className="collapse navbar-collapse d-flex justify-content-end mx-4" id="navbarNavDropdown">
                    <ul className="navbar-nav nav-bar">
                        <li className="nav-item dropdown ps-4">
                            <a className="nav-link dropdown-toggle text-light" href="#" id="navbarDropdownMenuLink" role="button" data-bs-toggle="dropdown" aria-expanded="false">
                                <FontAwesomeIcon icon={faUser} /> {params.username}
                            </a>
                            <ul className="dropdown-menu" aria-labelledby="navbarDropdownMenuLink">
                                <li><a className="dropdown-item btn-logout" href="#" >Đăng xuất</a></li>
                                <li><a className="dropdown-item btn-changePW" href="#">Đổi mật khẩu</a></li>
                            </ul>
                        </li>
                    </ul>
                </div>
            </div>
        </nav>
    )
}

export default HomepageHeader