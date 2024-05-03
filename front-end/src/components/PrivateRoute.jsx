import { useState } from "react";
import useAuth from "../hooks/useAuth";
import useLocalStorage from "../hooks/useLocalStorage";
import LoadingSpinner from "./defaultPage/Loading"
import { redirect } from "react-router-dom";


const PrivateRoute = ({ children }) => {
    const auth = useAuth()
    const [accessToken, setAccessToken] = useLocalStorage('accessToken', '')
    const [isLoading, setIsLoading] = useState(true);
    const [isValid, setIsValid] = useState(null);

    async function fetchData() {
        try {
            const response = await fetch(`http://localhost:8080/auth/validate?token=${auth.accessToken}`, {
                method: 'GET',
                headers: {
                    'Authorization': `Bearer ${auth.accessToken}`
                }
            });
            if (!response.ok) {
                setIsValid(null);
                setIsLoading(true);
                auth.setAccessToken('');
                window.location.href = '/';
            } else {
                const data = await response.json();
                setIsValid(data);
                setIsLoading(false);
            }

        } catch (error) {
            console.error('Error:', error);
            // Xử lý lỗi nếu cần thiết
            setIsValid(false);
            setIsLoading(true);
        }
    }

    if (auth.accessToken) {
        fetchData();
    } else {
        auth.setAccessToken('');
        window.location.href = '/';
    }
    if (isLoading) {
        return <div className='d-flex justify-content-center align-items-center my-5' style={{ height: '100vh' }}>
            <LoadingSpinner />
        </div>;
    } else if (isValid) {
        return children
    } else {
        auth.setAccessToken('');
        window.location.href = '/';
    }

}

export default PrivateRoute