import { useContext } from 'react'
import AuthContext from '../components/AuthProvider'

const useAuth = () => {
    const auth = useContext(AuthContext);
    if(auth === undefined){
        throw new Error("useAuth must be within AuthProvider");
    }
    return auth;
}

export default useAuth;