import React from "react";

const ManageComment = () => {
    return (
        <>
            <main>
                <div className="header">
                    <div className="left">
                        <h3>Manage Comments</h3>
                    </div>
                    <div className="time">
                        <span>March 2023 - February 2024</span>
                    </div>
                </div>
            </main>

            <div className="db-content col">
                <div className="left">
                    <span className="top">Total comments</span>
                        <h3>10.0k</h3>

                    <span>Growth in comments on this year</span>
                </div>
            </div>
            <hr className="h" />
            <section id="testimonials">
                <div className="testimonial-box-container">
                    <div className="testimonial-box">
                        <div className="box-top">
                            <div className="profile">
                                <div className="profile-img">
                                    <img src="https://4.bp.blogspot.com/-wmrAFTC9mEk/XhF48JllnvI/AAAAAAAAAa0/TEo_xTwPyNE3uBQOzj-kyaTdHI_YAIduwCLcBGAsYHQ/s1600/Fakta-wibu-dan-otaku.jpg" style={{width: '4.5rem'}} />
                                </div>
                                <div className="name-user">
                                    <strong>Quach Duc Huy</strong>
                                    <span>@DucHuy</span>
                                    <span>28-12-2023</span>
                                </div>
                            </div>
                            <div className="reviews">
                                <i className="fa fa-star"></i>
                                <i className="fa fa-star"></i>
                                <i className="fa fa-star"></i>
                                <i className="far fa-star"></i>
                            </div>
                        </div>

                        <div className="client-comment">
                            <p>Lorem ipsum dolor sit amet consectetur, adipisicing elit. Quisquam iusto iure similique quod omnis ipsum exercitationem velit eum veritatis repellat saepe delectus laborum voluptas architecto eligendi assumenda dolores, provident voluptatem!</p>

                        </div>
                    </div>
                </div>
            </section>
        </>)
}

export default ManageComment