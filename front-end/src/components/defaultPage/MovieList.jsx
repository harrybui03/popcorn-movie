import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faLessThan, faGreaterThan } from "@fortawesome/free-solid-svg-icons";
import React, { useEffect, useState, useRef } from "react";
import LoadingSpinner from './Loading'
import MovieCard from './MovieCard';
import { useNavigate } from 'react-router-dom';

function MovieList(username) {
    const navigate = useNavigate();
    const [isOnGoing, setIsOnGoing] = useState(true);
    const [loading, setLoading] = useState(true);
    const handleCheckboxChange = (check) => {
        setIsOnGoing(check);
    };
    const isLogin = username.username ? true : false;
    const [data, setData] = useState(null);
    const fetchData = async () => {
        try {
            setLoading(true);
            // Địa chỉ API và tham số
            const apiUrl = 'http://localhost:8080/movies';
            let params;
            if (isOnGoing) {
                params = { status: 'Ongoing' };
            } else {
                params = { status: 'Upcoming' };
            }

            // Chuyển đổi tham số thành chuỗi query
            const queryString = Object.keys(params)
                .map(key => `${encodeURIComponent(key)}=${encodeURIComponent(params[key])}`)
                .join('&');

            // Tạo URL với tham số
            const urlWithParams = `${apiUrl}?${queryString}`;
            // Gọi API bằng phương thức GET
            const response = await fetch(urlWithParams);

            // Kiểm tra trạng thái của response
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            else {
                // Chuyển đổi response thành JSON
                const result = await response.json();

                // Cập nhật state với dữ liệu từ API
                setData(result);
            }
        } catch (error) {
            console.error('Error fetching data:', error);
        } finally {
            setLoading(false); // Kết thúc loading sau khi nhận dữ liệu hoặc xảy ra lỗi
        }
    };

    useEffect(() => {
        fetchData();
    }, [isOnGoing]); // Dependency array để đảm bảo useEffect chỉ chạy một lần khi component được render

    const [detailMovie, setDetailMovie] = useState(false);
    const [movie, setMovie] = useState(null);
    const handleCardClick = (id) => {
        setLoading(true);
        setDetailMovie(true);
        setMovie(data.filter(e => e.id === id)[0]);
        setLoading(false);
    };
    useEffect(() => {
        // console.log(movie);
    }, [detailMovie]);

    const handleOrderClick = () => {
        if (isLogin) {
            navigate(`/movies/${movie.id}`);
        } else {
            $('#requireLogin').modal('show');
            console.log('456');
        }
    };

    return (
        <div className="p-5 movie-list">
            <div className="row">
                <div className="col-12">
                    <div className="btn-group movie-list-menu" role="group" aria-label="Basic radio toggle button group">
                        <span className="movie-list-border-solid"></span>
                        <h3 className="me-4">Danh mục</h3>
                        <input type="radio" className="btn-check movie-list-item" name="btnradio" id="btnradio1" autoComplete="off" checked={isOnGoing}
                            onChange={() => handleCheckboxChange(true)} onClick={() => setDetailMovie(false)} />
                        <label className="btn btn-outline-primary me-4 rounded-2" htmlFor="btnradio1">Phim đang chiếu</label>

                        <input type="radio" className="btn-check movie-list-item" name="btnradio" id="btnradio2" autoComplete="off" checked={!isOnGoing}
                            onChange={() => handleCheckboxChange(false)} onClick={() => setDetailMovie(false)} />
                        <label className="btn btn-outline-primary rounded-2" htmlFor="btnradio2">Phim sắp chiếu</label>

                    </div>
                </div>
            </div>

            {loading &&
                <div className='d-flex justify-content-center my-5'>
                    <LoadingSpinner />
                </div>}

            {(!detailMovie && data) ? (
                <div className="row mx-0 my-5" >
                    {data.map((item, index) => (
                        <MovieCard
                            isLogin={isLogin}
                            index={index}
                            movie={item}
                            key={item.id}
                            onClick={() => handleCardClick(item.id)}
                        />
                    ))}
                </div>
            ) : (detailMovie && movie) ? (
                <div className="row mx-0 my-5">
                    <div style={{ backgroundColor: 'rgb(219, 137, 45)', userSelect: 'none' }}
                        className="m-0 p-3 shadow border border-5 rounded-5 d-flex text-light col-12">
                        <img src={movie.poster} className='col-4 p-0' alt={`${movie.id} Poster`} />
                        <div className="col-8 d-flex flex-column justify-content-between px-5">
                            <h2 className="mx-auto">{movie.title}</h2>
                            <div className="d-flex flex-column align-items-start" >
                                <p className="text-start">Đạo diễn: {movie.director}</p>
                                <p className="text-start">Diễn viên: {movie.cast}</p>
                                <p className="text-start">Tóm tắt: {movie.story}</p>
                                <p className="text-start">Thời lượng: {movie.duration} phút</p>
                                <p className="text-start">Thể loại: {movie.genre}</p>
                                <p className="text-start">Ngôn ngữ: {movie.language}</p>
                                <p className="text-start">Khởi chiếu: {movie.openingDay}</p>
                                <p className="text-start">Đánh giá: {movie.rated}</p>
                            </div>
                            <div style={{ width: "100%", height: "1px", backgroundColor: 'rgb(81, 156, 247)' }}></div>
                            <div className="d-flex justify-content-center">
                                <a href={`${movie.trailer}`} style={{ backgroundColor: 'rgb(81, 156, 247)' }} className="btn btn-lg text-light m-5">Trailer</a>
                                <a onClick={handleOrderClick} style={{ backgroundColor: 'rgb(81, 156, 247)' }} className="btn btn-lg text-light m-5" data-bs-toggle="modal" data-bs-target="#MovieCard-modal">Đặt vé</a>
                            </div>
                        </div>
                    </div>
                </div>

            ) : (<></>)}

        </div>
    )
}

export default MovieList