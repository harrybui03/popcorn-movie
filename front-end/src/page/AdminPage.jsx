import React, { useEffect, useState } from "react";
import { useParams } from 'react-router-dom';
import Header from "../components/ticketManagerPage/Header";
import Home from '../components/adminPage/Home'
import SideBar from '../components/adminPage/Sidebar'
import ManageUser from '../components/adminPage/ManageUser'
import ManageComment from '../components/adminPage/ManageComment'
import Footer from '../components/defaultPage/Footer'
import LogOut from '../components/customerPage/LogOut'
import ChangePW from '../components/customerPage/ChangePW'
import '../components/adminPage/styles.css'

function AdminDashBoard(params) {
  const { currentChosen } = useParams();

  const [isCollapse, setIsCollapse] = useState(false);
  const updateData = () => {
    setIsCollapse(!isCollapse);
  }

  return (
    <>
      <div className="container-fluid">
        <div className="row">
          <SideBar
            chosen={currentChosen}
            isCollapse={isCollapse}></SideBar>
          <div className='col p-0 d-flex flex-column justify-content-between' style={{ flexShrink: '0' }}>
            <Header updateData={updateData} isCollapse={isCollapse} params={params} ></Header>
            {currentChosen === 'home' ? (
              <Home></Home>
            ) : currentChosen === 'manage-users' ? (
              <ManageUser/>
            ) : currentChosen === 'manage-comments' ? (
              <ManageComment />
            ) : (
              <Home></Home>
            )}
            <Footer></Footer>
            <LogOut></LogOut>
            <ChangePW {...params}></ChangePW>
          </div>
        </div>
      </div>
    </>
  )
}

export default AdminDashBoard