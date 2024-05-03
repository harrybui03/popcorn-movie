import React from 'react';


function Home() {

    return (
        <div className="jumbotron" style={{ backgroundColor: '#f0f0f0', height: '70vh', paddingTop: '7rem'}}>
            <h1 className="display-4">Quản Lý Phim</h1>
            <p className="lead">Đây là trang quản lý phim của website.</p>
            <hr className="my-4" />
            <p>Vui lòng chọn rạp trước khi thực hiện các thao tác quản lý.</p>
            <div>
                <a href={`/manage/movie`} className='btn btn-primary m-3'>Quản lý phim</a>
                <a href={`/manage/showtime`} className='btn btn-primary m-3'>Xếp lịch chiếu</a>
                <a href={`/manage/profit`} className='btn btn-primary m-3'>Xem doanh thu</a>
            </div>
        </div>
    )
}

export default Home