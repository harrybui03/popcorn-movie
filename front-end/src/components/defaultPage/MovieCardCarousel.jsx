import React, { useState, useEffect } from 'react';
import useAuth from '../../hooks/useAuth';
import { redirect } from 'react-router-dom';
import { useNavigate } from 'react-router-dom';
import { CarouselItem, Card, Col, Row } from 'react-bootstrap';
import './styles.css';

const MovieCardCarousel = ({ isLogin, index, movie, onClick }) => {
    const navigate = useNavigate();
    const handleOrderClick = () => {
        if (isLogin) {
            navigate(`/movies/${movie.id}`);
        }
    };
    // useEffect(() => {

    // }, [isLogin]);
    return (
        <Card onClick={() => onClick(index)} className='carousel-item-card m-2'>
            <Card.Img src={movie.poster} className="img-fluid h-100 w-100" alt={`${movie.id} Poster`} />
            {/* <Card.ImgOverlay>Slide {index}</Card.ImgOverlay> */}
        </Card>
        /* {!isLogin && (<div className="modal fade" id={`MovieCard-modal`} aria-labelledby="exampleModalLabel" onClick={(e) => {e.stopPropagation()}}>
            <div className="modal-dialog">
                <div className="modal-content">
                    <div className="modal-header">
                        <h1 className="modal-title fs-5" id="exampleModalLabel">Thông báo</h1>
                        <button type="button" className="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                    </div>
                    <div className="modal-body">
                        Vui lòng đăng nhập trước khi thực hiện thao tác này!
                    </div>
                    <div className="modal-footer">
                        <button type="button" className="btn btn-primary" data-bs-dismiss="modal" onClick={() => {
                            $('.login-modal').addClass('display');
                            $('.login-close-btn').on('click', function () {
                                $('.login-modal').removeClass('display');
                            });
 
                        }}>Đăng nhập</button>
                    </div>
                </div>
            </div>
        </div>
        )} */
    );
};

export default MovieCardCarousel;