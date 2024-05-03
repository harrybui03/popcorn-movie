import React, { useState, useEffect } from 'react';
import './styles.css';
import useAuth from '../../hooks/useAuth';
import { redirect } from 'react-router-dom';
import { useNavigate } from 'react-router-dom';

const MovieCard = ({ isLogin, index, movie, onClick }) => {
    const navigate = useNavigate();
    const styles = ['rgb(219, 137, 45)', 'rgb(81, 156, 247)', 'rgb(201, 102, 247)', 'rgb(242, 118, 143)'];
    const handleOrderClick = () => {
        if (isLogin) {
            navigate(`/movies/${movie.id}`);
        } else {
            $('#requireLogin').modal('show');
        }
    };

    return (
        <div style={{ backgroundColor: styles[index % 4], userSelect: 'none' }} id={`${movie.id}`}
            className="movie-card col-md-6 p-3 shadow border border-5 rounded-5 d-flex justify-content-around text-light"
            onClick={() => {
                onClick(index)
            }}>
            <img src={movie.poster} alt={`${movie.id} Poster`} style={{ height: "20rem" }} />
            <div className="d-flex flex-column justify-content-between" style={{ width: "15rem" }}>
                <h5 className="mx-auto">{movie.title}</h5>
                <div>
                    <p className="mx-auto">Đạo diễn: {movie.director}</p>
                    <p className="mx-auto">Diễn viên: {movie.cast}</p>
                </div>
                <div style={{ width: "100%", height: "1px", backgroundColor: styles[(index + 1) % 4] }}></div>
                <div className="d-flex justify-content-around">
                    <a href={`${movie.trailer}`}  onClick={(e) => {e.stopPropagation()}} style={{ backgroundColor: styles[(index + 1) % 4] }} className="btn text-light">Trailer</a>
                    <a onClick={(e) => {
                        e.stopPropagation();
                        handleOrderClick();
                    }}
                        style={{ backgroundColor: styles[(index + 1) % 4] }} className="btn text-light">Đặt vé</a>

                </div>
            </div>
        </div>
    );
};

export default MovieCard;
