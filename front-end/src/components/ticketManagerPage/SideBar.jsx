import React, { useState, useEffect } from 'react';
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faHouse, faFilm, faTicket, faChartLine } from "@fortawesome/free-solid-svg-icons";
import './styles.css'

function SideBar({ chosen, isCollapse }) {
    const handleChosenData = (select) => {
        window.location.href = `/manage/${select}`;
    }
    return (
        <>
            <nav id="sidebar" className={`py-4 col-2 ${!isCollapse ? '' : 'collapse'}`}>
                <div onClick={() => handleChosenData('home')} className='fs-5 fw-bold my-5' style={{ cursor: 'context-menu' }}>
                    <FontAwesomeIcon icon={faHouse} /> Trang chủ quản lý
                </div>
                
                <div onClick={() => handleChosenData('movie')} id='movie' className={`items p-3 ${chosen === 'movie' ? 'activeItems' : ''}`}><FontAwesomeIcon className='mx-2' icon={faFilm} /> Quản lý phim</div>
                <div onClick={() => handleChosenData('showtime')} id='showtime' className={`items p-3 ${chosen === 'showtime' ? 'activeItems' : ''}`}><FontAwesomeIcon className='mx-2' icon={faTicket} /> Xếp lịch chiếu</div>
                <div onClick={() => handleChosenData('profit')} id='profit' className={`items p-3 ${chosen === 'profit' ? 'activeItems' : ''}`}><FontAwesomeIcon className='mx-2' icon={faChartLine} /> Doanh thu</div>
            </nav>
            <nav id="sidebar-background" className={`col-2 ${!isCollapse ? '' : 'collapse'}`} style={{flexShrink: '0'}}>
            </nav>
        </>
    )
}

export default SideBar