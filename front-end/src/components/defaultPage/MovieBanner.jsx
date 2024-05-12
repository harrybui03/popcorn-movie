import { React, useEffect, useState } from "react";
import { Carousel, Container, Row, Card, CarouselItem, Col } from 'react-bootstrap';
import { useNavigate } from "react-router-dom";
import LoadingSpinner from './Loading'
import MovieCardCarousel from "./MovieCardCarousel";
import "./styles.css"
import { useGetAllMovies } from "./hook/useQuery";

function MovieBanner(username) {
    const navigate = useNavigate();
    const [isOnGoing, setIsOnGoing] = useState(true);
    const [loading, setLoading] = useState(true);
    const isLogin = username.username ? true : false;
    const { data , isLoading} = useGetAllMovies(null);

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
        }
    };

    useEffect(() => {
        let items = document.querySelectorAll('.carousel .carousel-item')

        items.forEach((el) => {
            const minPerSlide = 3
            let next = el.nextElementSibling
            for (var i = 1; i < minPerSlide; i++) {
                if (!next) {
                    // wrap carousel by using first child
                    next = items[0]
                }
                let cloneChild = next.cloneNode(true)
                el.appendChild(cloneChild.children[0])
                next = next.nextElementSibling
            }
        })
    }, []);

    const itemsPerSlide = 3;

    const renderCarouselItems = () => {
        const carouselItems = [];

        for (let i = 0; i < data.length; i += itemsPerSlide) {
            // divide data into 3 item per slide
            const slideItems = data.slice(i, i + itemsPerSlide).map((item, index) => (
                <Col key={index} md={4}>
                    <MovieCardCarousel
                        isLogin={isLogin}
                        index={i + index + 1}
                        movie={item}
                        onClick={() => handleCardClick(item.id)}
                    />

                    


                </Col>
            ));

            carouselItems.push(
                <Carousel.Item key={i}>
                    {slideItems}
                </Carousel.Item>
            );
        }

        return carouselItems;
    };



    return (
        <Container className="text-center">
            <Row className="mx-auto my-auto justify-content-center">
                {(!detailMovie && data) ? (
                    <Carousel id="recipeCarousel" className="slide me-auto ms-auto" interval={3000} slide={true} role="listbox">
                        {isLoading &&
                            <div className='d-flex justify-content-center my-5'>
                                <LoadingSpinner />
                            </div>}
                        {renderCarouselItems()}
                    </Carousel>
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
                                    <a onClick={handleOrderClick} style={{ backgroundColor: 'rgb(81, 156, 247)' }} className="btn btn-lg text-light m-5" data-bs-toggle="modal" data-bs-target="#MovieList-modal">Đặt vé</a>
                                    {!isLogin && (<div className="modal fade" id="MovieList-modal" tabIndex="-1" aria-labelledby="exampleModalLabel">
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
                                    </div>)}

                                </div>
                            </div>
                        </div>
                    </div>

                ) : (<></>)}

            </Row>
        </Container>
    )
}

export default MovieBanner