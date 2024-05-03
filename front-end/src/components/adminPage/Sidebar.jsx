import React, { useEffect, useState } from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faHouse, faUsers, faComment } from "@fortawesome/free-solid-svg-icons";
import './styles.css'

function SideBar({ chosen, isCollapse }) {
    const handleChosenData = (select) => {
        window.location.href = `/admin/${select}`;
    }
    return (
        <>
            <nav id="sidebar" className={`py-4 col-2 ${!isCollapse ? '' : 'collapse'}`}>
                <div onClick={() => handleChosenData('home')} className='fs-5 fw-bold my-5' style={{ cursor: 'context-menu' }}>
                    <FontAwesomeIcon icon={faHouse} /> Trang chủ quản lý
                </div>
                
                <div onClick={() => handleChosenData('manage-users')} id='manage-users' className={`items p-3 ${chosen === 'manage-users' ? 'activeItems' : ''}`}><FontAwesomeIcon className='mx-2' icon={faUsers} /> Quản lý tài khoản</div>
                <div onClick={() => handleChosenData('manage-comments')} id='manage-comments' className={`items p-3 ${chosen === 'manage-comments' ? 'activeItems' : ''}`}><FontAwesomeIcon className='mx-2' icon={faComment} /> Quản lý bình luận</div>
            </nav>
            <nav id="sidebar-background" className={`col-2 ${!isCollapse ? '' : 'collapse'}`} style={{flexShrink: '0'}}>
            </nav>
        </>
    )
}

export default SideBar