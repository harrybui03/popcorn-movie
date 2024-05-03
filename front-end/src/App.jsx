import { useState, useEffect } from 'react'
import './App.css'
import 'bootstrap/dist/css/bootstrap.css'
import { jwtDecode } from "jwt-decode"
import { BrowserRouter as Router, Route, Routes, BrowserRouter } from 'react-router-dom';
import PrivateRoute from './components/PrivateRoute'
import useAuth from './hooks/useAuth'
import AdminPage from './page/AdminPage'
import CustomerPage from './page/CustomerPage'
import DefaultPage from './page/DefaultPage'
import TicketManagerPage from './page/TicketManagerPage'
import ReceptionistPage from './page/ReceptionistPage'
import OrderTicket from './components/customerPage/OrderTicket'
import PaymentSuccess from './components/customerPage/PaymentSuccess';
import PaymentFail from './components/customerPage/PaymentFail';
import HistoryBooking from './components/customerPage/HistoryBooking';
import VerifyAccount from './components/defaultPage/VerifyAccount';

function App() {
  const getRoles = (accessToken) => {
    if (accessToken) {
      let decodedJwt = jwtDecode(accessToken);
      return decodedJwt.roles;
    }
    return [];
  }
  const { accessToken, setAccessToken } = useAuth();
  const [roles, setRoles] = useState(getRoles(accessToken));
  useEffect(() => {
    // Update roles when accessToken changes
    setRoles(getRoles(accessToken));
  }, [accessToken]);

  const getUserName = (accessToken) => {
    if (accessToken) {
      let decodedJwt = jwtDecode(accessToken);
      return decodedJwt.username;
    }
    return [];
  }
  const getEmail = (accessToken) => {
    if (accessToken) {
      let decodedJwt = jwtDecode(accessToken);
      return decodedJwt.sub;
    }
    return [];
  }
  const username = getUserName(accessToken);
  const email = getEmail(accessToken);
  return (
    <>
      <BrowserRouter>
        <Routes>
          <Route
            path='/' element={
              roles.find(role => role === "ROLE_CUSTOMER") ? (
                <PrivateRoute>
                  <CustomerPage username={username} email={email} />
                </PrivateRoute>
              ) : roles.find(role => role === "ROLE_STAFF") ? (
                <PrivateRoute>
                  <ReceptionistPage username={username} email={email} />
                </PrivateRoute>
              ) : roles.find(role => role === "ROLE_ADMIN") ? (
                <PrivateRoute>
                  <AdminPage username={username} email={email} />
                </PrivateRoute>
              ) : roles.find(role => role === "ROLE_TICKET_MANAGER") ? (
                <PrivateRoute>
                  <TicketManagerPage username={username} email={email} />
                </PrivateRoute>
              ) : (
                <DefaultPage />
              )
            }
          />
          <Route path='/movies/:id' element={
            roles.find(role => role === "ROLE_CUSTOMER") ? (
              <PrivateRoute>
                <OrderTicket username={username} email={email} />
              </PrivateRoute>
            ) : (
              <DefaultPage />
            )
          } />
          <Route path='/payments/success' element={
            roles.find(role => role === "ROLE_CUSTOMER") ? (
              <PrivateRoute>
                <PaymentSuccess username={username} email={email} />
              </PrivateRoute>
            ) : (
              <DefaultPage />
            )
          } />
          <Route path='/payments/cancel' element={
            roles.find(role => role === "ROLE_CUSTOMER") ? (
              <PrivateRoute>
                <PaymentFail username={username} email={email} />
              </PrivateRoute>
            ) : (
              <DefaultPage />
            )
          } />
          <Route path='/history-booking' element={
            roles.find(role => role === "ROLE_CUSTOMER") ? (
              <PrivateRoute>
                <HistoryBooking username={username} email={email} />
              </PrivateRoute>
            ) : (
              <DefaultPage />
            )
          } />
          <Route path='/verify-account' element={
            <VerifyAccount />
          } />
          <Route
            path='/manage/:currentChosen' element={
              roles.find(role => role === "ROLE_TICKET_MANAGER") ? (
                <PrivateRoute>
                  <TicketManagerPage username={username} email={email} />
                </PrivateRoute>
              ) : (
                <DefaultPage />
              )
            }
          />
          <Route
            path='/admin/:currentChosen' element={
              roles.find(role => role === "ROLE_ADMIN") ? (
                <PrivateRoute>
                  <AdminPage username={username} email={email} />
                </PrivateRoute>
              ) : (
                <DefaultPage />
              )
            }
          />
        </Routes>
      </BrowserRouter>

    </>
  )
}

export default App
