import React from "react";
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faFacebook, faGithub, faInstagram } from "@fortawesome/free-brands-svg-icons";
import { ReactDOM } from "react";

function Footer() {
    return (
        <div className="container-fluid my-1 footer-container pe-0 ps-0">
            <footer className="text-center text-lg-start text-dark">
                <section className="d-flex justify-content-between p-4 text-white footer-social-section">
                    <div className="me-5">
                        <span>Get connected with us on social networks:</span>
                    </div>
                    <div>
                        <a href="" className="text-white me-4">
                            <FontAwesomeIcon icon={faFacebook} />
                        </a>

                        <a href="" className="text-white me-4">
                            <FontAwesomeIcon icon={faInstagram} />
                        </a>
                    </div>
                </section>
                <section className="">
                    <div className="container text-center text-md-start mt-5">
                        <div className="row mt-3">
                            <div className="col-md-3 col-lg-4 col-xl-3 mx-auto mb-4">
                                <h6 className="text-uppercase fw-bold">Rạp chiếu phim</h6>
                                <hr className="mb-4 mt-0 d-inline-block mx-auto footer-company-name"/>
                                <p>
                                    H3DC
                                </p>
                            </div>
                            <div className="col-md-2 col-lg-2 col-xl-2 mx-auto mb-4 footer-branch-container">
                                <h6 className="text-uppercase fw-bold">Chi Nhánh</h6>
                                <hr className="mb-4 mt-0 d-inline-block mx-auto footer-branch"/>
                                <p>
                                    <a href="#!" className="text-dark">Linh Trung, Thủ Đức</a>
                                </p>
                                <p>
                                    <a href="#!" className="text-dark">Dĩ An, Bình Dương</a>
                                </p>
                                <p>
                                    <a href="#!" className="text-dark">Thành Phố Thủ Đức</a>
                                </p>

                            </div>
                            <div className="col-md-3 col-lg-2 col-xl-2 mx-auto mb-4 footer-router-container">
                                <h6 className="text-uppercase fw-bold">Ticketify</h6>
                                <hr className="mb-4 mt-0 d-inline-block mx-auto footer-router"/>
                                <p>
                                    <a href="#!" className="text-dark">Phim sắp chiếu</a>
                                </p>
                                <p>
                                    <a href="#!" className="text-dark">Phim đang chiếu</a>
                                </p>
                                <p>
                                    <a href="#!" className="text-dark">Phim mới</a>
                                </p>
                                <p>
                                    <a href="#!" className="text-dark">Lịch chiếu</a>
                                </p>
                            </div>

                            <div className="col-md-4 col-lg-3 col-xl-3 mx-auto mb-md-0 mb-4 footer-information-container">
                                <h6 className="text-uppercase fw-bold">Thông tin</h6>
                                <hr className="mb-4 mt-0 d-inline-block mx-auto footer-information"/>
                                <p> <a className="text-dark" href="">Giới thiệu</a></p>
                                <p> <a className="text-dark" href="">Liên hệ</a></p>
                            </div>
                        </div>
                    </div>
                </section>
                <div className="text-center p-3 bg-light">
                    © Copyright Ticketify
                </div>
            </footer>
        </div>
    )
}

export default Footer