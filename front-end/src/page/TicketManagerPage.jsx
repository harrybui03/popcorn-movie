import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import Header from '../components/ticketManagerPage/Header'
import Footer from '../components/defaultPage/Footer'
import LogOut from '../components/customerPage/LogOut'
import ChangePW from '../components/customerPage/ChangePW'
import SideBar from '../components/ticketManagerPage/SideBar';
import Home from '../components/ticketManagerPage/Home';
import Movie from '../components/ticketManagerPage/Movie';
import Showtime from '../components/ticketManagerPage/Showtime';
import Profit from '../components/ticketManagerPage/Profit';
import LoadingSpinner from '../components/defaultPage/Loading'
import DeleteMovie from '../components/ticketManagerPage/DeleteMovie';
import DeleteShowtime from '../components/ticketManagerPage/DeleteShowtime';
import Notification from '../components/ticketManagerPage/Notification';
import '../components/ticketManagerPage/styles.css';

function TicketManagerDashBoard(params) {
  const { currentChosen } = useParams();

  const [isCollapse, setIsCollapse] = useState(false);
  const updateData = () => {
    setIsCollapse(!isCollapse);
  }

  const [deleteId, setDeleteId] = useState(null);
  const [deleteIdShowtime, setDeleteIdShowtime] = useState(null);

  const [notification, setNotifacation] = useState({ title: '', body: '', footer: '', status: '' });

  return (
    <>

      <div className="container-fluid">
        <div className="row">
          <SideBar
            chosen={currentChosen}
            isCollapse={isCollapse}></SideBar>
          <div className='col p-0 d-flex flex-column justify-content-between'  style={{flexShrink: '0'}}>
            <Header updateData={updateData} isCollapse={isCollapse} params={params} ></Header>
            <Notification notification={notification}></Notification>
            <DeleteMovie setNotification={setNotifacation} deleteId={deleteId}></DeleteMovie>
            <DeleteShowtime setNotification={setNotifacation} deleteId={deleteIdShowtime}></DeleteShowtime>
            {currentChosen === 'home' ? (
              <Home></Home>
            ) : currentChosen === 'movie' ? (
              <Movie setDeleteId={setDeleteId}></Movie>
            ) : currentChosen === 'showtime' ? (
              <Showtime setDeleteIdShowtime={setDeleteIdShowtime}></Showtime>
            ) : currentChosen === 'profit' ? (
              <Profit></Profit>
            ) : (
              <Home></Home>
            )}

            <LogOut></LogOut>
            <ChangePW {...params}></ChangePW>
            <Footer></Footer>
          </div>
        </div>
      </div>

    </>
  )
}

export default TicketManagerDashBoard