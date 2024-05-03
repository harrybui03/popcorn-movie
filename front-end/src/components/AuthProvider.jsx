import React, { createContext, useContext } from 'react'
import useLocalStorage from '../hooks/useLocalStorage';

const AuthContext = createContext({})
export const AuthProvider = ({children}) => {
    const [accessToken, setAccessToken] = useLocalStorage('accessToken', '')
  return (
    <>
        <AuthContext.Provider value={{accessToken, setAccessToken}}>
            {children}
        </AuthContext.Provider>
    </>
  )
}
export default AuthContext;
