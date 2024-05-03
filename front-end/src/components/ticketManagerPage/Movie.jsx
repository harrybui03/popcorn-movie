import React, { useState, useEffect } from 'react';
import useAuth from "../../hooks/useAuth";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faHouse, faMagnifyingGlass, faAdd, faEdit, faTrash, faAngleRight, faUnsorted, faSortUp, faSortDown } from "@fortawesome/free-solid-svg-icons";
import LoadingSpinner from '../defaultPage/Loading'
import { set } from 'date-fns';

function Movie({ setDeleteId }) {
    const auth = useAuth();
    const [movies, setMovies] = useState(null);
    const [moviesDisplay, setMoviesDisplay] = useState(null);

    const getMovie = async () => {
        try {
            // Địa chỉ API và tham số
            const apiUrl = 'http://localhost:8080/movies/manager';

            // Gọi API bằng phương thức GET
            const response = await fetch(apiUrl, {
                method: 'GET',
                headers: {
                    'Authorization': `Bearer ${auth.accessToken}`
                }
            });
            // Kiểm tra trạng thái của response
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            else {
                // Chuyển đổi response thành JSON
                const result = await response.json();
                // Cập nhật state với dữ liệu từ API
                setMovies(result);
            }
        } catch (error) {
            console.error('Error fetching data:', error);
        }
    };

    useEffect(() => {
        getMovie();
    }, []);

    useEffect(() => {
        const a = [];
        if (movies !== null) {
            for (let i in movies[0]) {
                a.push(i);
            }
            setMoviesDisplay(movies);
        }
        setListFilterBy(a);
    }, [movies]);

    const [listFilterBy, setListFilterBy] = useState([]);

    useEffect(() => {
        if (listFilterBy.length > 0) {
            setFilterBy(listFilterBy[0]);
        }
    }, [listFilterBy])

    const [filterBy, setFilterBy] = useState(null);
    const [sortBy, setSortBy] = useState({ by: '', isDown: true });

    const handleFilterBy = (event) => {
        setFilterBy(event.target.value);
    }

    const handleSortBy = (a) => {
        setSortBy({ by: a, isDown: !sortBy?.isDown });
    }

    const [query, setQuery] = useState('');
    const handleChangeQuery = (event) => {
        setQuery(event.target.value);
    };

    useEffect(() => {
        if (query === '') {
            setMoviesDisplay(movies);
        }
    }, [query])

    useEffect(() => {
        if (moviesDisplay !== null) {
            setMoviesDisplay(moviesDisplay.sort((a, b) => {
                if (sortBy.by === 'id') {
                    return sortBy.isDown ? (a[sortBy.by] - b[sortBy.by]) : (b[sortBy.by] - a[sortBy.by]);
                }
                return sortBy.isDown ? (a[sortBy.by]?.toLowerCase().localeCompare(b[sortBy.by]?.toLowerCase())) : (b[sortBy.by]?.toLowerCase().localeCompare(a[sortBy.by]?.toLowerCase()));
            }));
        }
    }, [sortBy])

    const search = () => {
        setMoviesDisplay(movies.filter(e => {
            return (e[filterBy] + '').toLowerCase().includes(query.toLowerCase());
        }))
    }

    const [chainAction, setChainAction] = useState([{ text: 'Quản lý phim', href: '/manage/movie' }]);

    const [action, setAction] = useState('manage-movie');

    const handleChangeAction = (a, id) => {
        const newChange = chainAction;
        newChange.push({ text: a === 'insert-movie' ? 'Thêm phim mới' : 'Chỉnh sửa phim', href: '#' })
        setChainAction(newChange);
        setAction(a);
        if (a === 'edit-movie') {
            const movie = moviesDisplay.filter(e => e.id === id)[0];
            setFormData(movie);
            setSelectedImage(movie.poster);

        } else {
            setIsEditable(true);
            setFormData({
                id: '',
                title: '',
                director: '',
                cast: '',
                duration: '',
                genre: '',
                language: '',
                openingDay: '',
                rated: '',
                status: 'Upcoming',
                story: '',
                poster: '',
                trailer: '',
            });
            setSelectedImage(null);
            setSelectedFile(null);
        }
    }

    const [formData, setFormData] = useState({
        id: '',
        title: '',
        director: '',
        cast: '',
        duration: '',
        genre: '',
        language: '',
        openingDay: '',
        rated: '',
        status: 'Upcoming',
        story: '',
        poster: '',
        trailer: '',
    });

    const handleChange = (e) => {
        const { name, value } = e.target;
        setFormData((prevData) => ({
            ...prevData,
            [name]: value,
        }));
    };

    const [isEditable, setIsEditable] = useState(false);
    const handleEditToggle = () => {
        setIsEditable(true);
        setMessage({ isShow: false, text: '', success: false });
    };
    const handleBack = () => {
        setAction('manage-movie');
        const newChain = chainAction;
        newChain.pop();
        setChainAction(newChain);
        setIsEditable(false);
        setMessage({ isShow: false, text: '', success: false });
        setSelectedImage(null);
        setSelectedFile(null);
    };

    const [selectedImage, setSelectedImage] = useState(null);
    const [selectedFile, setSelectedFile] = useState(null);

    const handleChangeFormFile = (e) => {
        const file = e.target.files[0];
        if (file) {
            setSelectedFile(file);
            const reader = new FileReader();
            reader.onloadend = () => {
                setSelectedImage(reader.result);
            };
            reader.readAsDataURL(file);
        } else {
            setSelectedFile(null);
            setSelectedImage(null);
        }
    };

    const [message, setMessage] = useState({ isShow: false, text: '', success: false });

    const isObjectEmpty = (obj) => {
        for (const key in obj) {
            if (obj[key] === null || obj[key] === undefined || obj[key] === '') {
                return true;
            }
        }
        return false;
    };

    const [disabledBtn, setDisabledBtn] = useState(false);


    const addMovie = async () => {
        const movie = {
            title: formData.title,
            director: formData.director,
            cast: formData.cast,
            duration: formData.duration,
            genre: formData.genre,
            language: formData.language,
            openingDay: formData.openingDay,
            rated: formData.rated,
            status: formData.status,
            story: formData.story,
            trailer: formData.trailer,
        }
        console.log(movie);
        if (isObjectEmpty(movie) || selectedImage === null || selectedImage === '') {
            setMessage({ isShow: true, text: 'Vui lòng điền đầy đủ thông tin', success: false });
        } else {
            setDisabledBtn(true);
            setMessage({ isShow: true, text: 'Đang thêm dữ liệu, vui lòng đợi...', success: true });

            const response = await postDataWithFile(movie, selectedFile);
            getMovie();
            setFormData({
                id: '',
                title: '',
                director: '',
                cast: '',
                duration: '',
                genre: '',
                language: '',
                openingDay: '',
                rated: '',
                status: 'Upcoming',
                story: '',
                poster: '',
                trailer: '',
            });
            setSelectedImage(null);
            setSelectedFile(null);
            setMessage({ isShow: true, text: 'Thêm phim thành công', success: true });

            setDisabledBtn(false);
        }
    }

    const postDataWithFile = async (data, file) => {
        const formData = new FormData();

        formData.append('poster', file);
        const movieBlob = new Blob([JSON.stringify(data)], { type: 'application/json' });
        formData.append('movie', movieBlob);

        try {
            const response = await fetch('http://localhost:8080/movies', {
                method: 'POST',
                headers: {
                    'Authorization': `Bearer ${auth.accessToken}`,
                },
                body: formData,
            });

            if (!response.ok) {
                setMessage({ isShow: true, text: 'Vui lòng kiểm tra lại dữ liệu phải đầy đủ ý nghĩa', success: false });
                setDisabledBtn(false);
                throw new Error(`HTTP error! Status: ${response.status}`);
            }
            const result = await response.json();
            return result;
        } catch (error) {
            console.error('Error:', error);
            throw error;
        }
    };

    const editMovie = async () => {
        const movie = {
            id: formData.id,
            title: formData.title,
            director: formData.director,
            cast: formData.cast,
            duration: formData.duration,
            genre: formData.genre,
            language: formData.language,
            openingDay: formData.openingDay,
            rated: formData.rated,
            status: formData.status,
            story: formData.story,
            trailer: formData.trailer,
        }
        if (isObjectEmpty(movie) || selectedImage === null || selectedImage === '') {
            setMessage({ isShow: true, text: 'Vui lòng điền đầy đủ thông tin', success: false });
        } else {
            setDisabledBtn(true);
            setMessage({ isShow: true, text: 'Đang cập nhật dữ liệu, vui lòng đợi...', success: true });

            const response = await putDataWithFile(movie, selectedFile);
            getMovie();
            setFormData(response);
            setSelectedImage(response.poster);
            setSelectedFile(null);
            setMessage({ isShow: true, text: 'Cập nhật phim thành công', success: true });
            setDisabledBtn(false);
            setIsEditable(false);
        }
    }

    const putDataWithFile = async (data, file) => {
        const formData = new FormData();

        if (file !== null) {
            formData.append('poster', file);
        }
        const movieBlob = new Blob([JSON.stringify(data)], { type: 'application/json' });
        formData.append('movie', movieBlob);

        try {
            const response = await fetch('http://localhost:8080/movies', {
                method: 'PUT',
                headers: {
                    'Authorization': `Bearer ${auth.accessToken}`,
                },
                body: formData,
            });

            if (!response.ok) {
                setMessage({ isShow: true, text: 'Vui lòng kiểm tra lại dữ liệu phải đầy đủ ý nghĩa', success: false });
                setDisabledBtn(false);
                throw new Error(`HTTP error! Status: ${response.status}`);
            }
            const result = await response.json();
            return result;
        } catch (error) {
            console.error('Error:', error);
            throw error;
        }
    };

    const deleteMovie = (id) => {
        setDeleteId(id);
        $('#deleteMovie').modal('show');
    }

    return (
        <>
            <div className="p-4" style={{ backgroundColor: '#f0f0f0' }}>
                <div className='m-3 d-flex justify-content-start align-items-center'>
                    <div>
                        <a className='text-black' href='/manage/home'><FontAwesomeIcon icon={faHouse} /> Trang chủ</a>
                    </div>
                    {chainAction.map((item, index) => (
                        <div key={index}>
                            <FontAwesomeIcon className='mx-3' icon={faAngleRight} />
                            <a className='text-black' href={item.href}>{item.text}</a>
                        </div>
                    ))}
                </div>
                <hr />

                {action === 'manage-movie' ? (
                    <>
                        <div className='m-3 d-flex justify-content-between align-items-center'>
                            <div className='d-flex justify-content-start align-items-center'>
                                <form className='mx-3 d-flex justify-content-start align-items-center'>
                                    <label htmlFor='filter' className='mx-2'>Tìm theo</label>
                                    <select id='filter' className='p-1 rounded-2' style={{ backgroundColor: 'white', color: 'black' }} onChange={handleFilterBy}>
                                        {listFilterBy.map((item) => (
                                            <option style={{ backgroundColor: 'white', color: 'black' }} key={item} value={item}>{item}</option>
                                        ))}
                                    </select>
                                </form>
                                <div className="search-container">
                                    <div className="d-flex">
                                        <input className="form-control me-2" type="search" placeholder="Tìm kiếm" aria-label="Search" onChange={handleChangeQuery} />
                                        <button className="btn btn-outline-success" onClick={search}>
                                            <FontAwesomeIcon icon={faMagnifyingGlass} className='search-icon' />
                                        </button>
                                    </div>
                                </div>
                            </div>
                            <div>
                                <button onClick={() => handleChangeAction('insert-movie')} className='btn btn-primary'><FontAwesomeIcon icon={faAdd} /> Thêm mới</button>
                            </div>
                        </div>

                        {moviesDisplay ? (
                            <table className="table table-hover table-bordered shadow" style={{ borderRadius: '10px', overflow: 'hidden' }}>
                                <thead className="table-primary">
                                    <tr style={{ cursor: 'context-menu' }}>
                                        <th scope="col">
                                            <div onClick={() => handleSortBy('id')} className='d-flex justify-content-center align-items-center'>
                                                <div className='mx-2'>ID</div>
                                                <FontAwesomeIcon icon={sortBy.by !== 'id' ? faUnsorted : sortBy.isDown ? faSortDown : faSortUp} />
                                            </div>
                                        </th>
                                        <th scope="col" className="col-3">
                                            <div onClick={() => handleSortBy('title')} className='d-flex justify-content-center align-items-center'>
                                                <div className='mx-2'>Tên phim</div>
                                                <FontAwesomeIcon icon={sortBy.by !== 'title' ? faUnsorted : sortBy.isDown ? faSortDown : faSortUp} />
                                            </div>
                                        </th>
                                        <th scope="col" className="col-2">
                                            <div onClick={() => handleSortBy('openingDay')} className='d-flex justify-content-center align-items-center'>
                                                <div className='mx-2'>Ngày chiếu</div>
                                                <FontAwesomeIcon icon={sortBy.by !== 'openingDay' ? faUnsorted : sortBy.isDown ? faSortDown : faSortUp} />
                                            </div>
                                        </th>
                                        <th scope="col" className="col-2">
                                            <div onClick={() => handleSortBy('status')} className='d-flex justify-content-center align-items-center'>
                                                <div className='mx-2'>Trạng thái</div>
                                                <FontAwesomeIcon icon={sortBy.by !== 'status' ? faUnsorted : sortBy.isDown ? faSortDown : faSortUp} />
                                            </div>
                                        </th>
                                        <th scope="col" className="col-2">
                                            Poster
                                        </th>
                                        <th scope="col" className="col-2">Tác vụ</th>
                                    </tr>
                                </thead>
                                <tbody>
                                    {moviesDisplay.map((item) => (
                                        <tr key={item.id}>
                                            <th className="align-middle" scope="row">{item.id}</th>
                                            <td className="align-middle">{item.title}</td>
                                            <td className="align-middle">{item.openingDay}</td>
                                            <td className="align-middle">{item.status}</td>
                                            <td className="align-middle"><img style={{ height: '7rem' }} src={item.poster} alt='poster'></img></td>
                                            <td className="align-middle">
                                                <div className='d-flex justify-content-center'>
                                                    <button onClick={() => handleChangeAction('edit-movie', item.id)} className='btn btn-sm text-primary m-2' style={{ backgroundColor: '#ffffff' }}>
                                                        <FontAwesomeIcon icon={faEdit} />
                                                    </button>
                                                    <button onClick={() => deleteMovie(item.id)} className='btn btn-sm text-danger m-2' style={{ backgroundColor: '#ffffff' }}>
                                                        <FontAwesomeIcon icon={faTrash} />
                                                    </button>
                                                </div>

                                            </td>
                                        </tr>
                                    ))}
                                </tbody>
                            </table>
                        ) : (
                            <div className='d-flex justify-content-center my-5'>
                                <LoadingSpinner />
                            </div>
                        )}
                    </>
                ) : (
                    <div>
                        <div className='my-2 d-flex justify-content-center'>
                            <div className="col-6 m-2">
                                <table className="table shadow" style={{ borderRadius: '10px', overflow: 'hidden' }}>
                                    <tbody>
                                        <tr>
                                            <td className="align-middle col-2" >ID phim</td>
                                            <td className="align-middle col-2" >
                                                {action === 'insert-movie' ? (
                                                    <input
                                                        className='form-control bg-light'
                                                        name="id"
                                                        value={'Sẽ được tạo tự động'}
                                                        disabled={true}></input>
                                                ) : (
                                                    <input
                                                        className='form-control bg-light'
                                                        name="id"
                                                        value={formData.id}
                                                        disabled={true}
                                                        onChange={handleChange} required></input>
                                                )}

                                            </td>
                                            <td className="align-middle col-3" >Trạng thái</td>
                                            <td className="align-middle col-3" >
                                                <select
                                                    className='form-control bg-light'
                                                    name="status"
                                                    value={formData.status !== '' ? formData.status : 'Upcoming'}
                                                    disabled={!isEditable}
                                                    onChange={handleChange} required>
                                                    <option value="Upcoming">Upcoming</option>
                                                    <option value="Ongoing">Ongoing</option>
                                                    <option value="Over">Over</option>
                                                </select>
                                            </td>
                                        </tr>
                                        <tr>
                                            <td className="align-middle col-3" >Tên phim</td>
                                            <td colSpan={3} className="align-middle col-2" >
                                                <textarea
                                                    className='form-control bg-light'
                                                    name="title"
                                                    value={formData.title}
                                                    disabled={!isEditable}
                                                    onChange={handleChange} required rows={2}></textarea>
                                            </td>
                                        </tr>
                                        <tr>
                                            <td className="align-middle col-3" >Thời lượng</td>
                                            <td className="align-middle col-2" >
                                                <input
                                                    placeholder='Phút'
                                                    type='number'
                                                    className='form-control bg-light'
                                                    name="duration"
                                                    value={formData.duration}
                                                    disabled={!isEditable}
                                                    onChange={handleChange} required></input>
                                            </td>
                                            <td className="align-middle col-3" >Thể loại</td>
                                            <td className="align-middle col-2" >
                                                <input
                                                    className='form-control bg-light'
                                                    name="genre"
                                                    value={formData.genre}
                                                    disabled={!isEditable}
                                                    onChange={handleChange} required></input>
                                            </td>
                                        </tr>
                                        <tr>
                                            <td className="align-middle col-3" >Đánh giá</td>
                                            <td className="align-middle col-3" >
                                                <input
                                                    className='form-control bg-light'
                                                    name="rated"
                                                    value={formData.rated}
                                                    disabled={!isEditable}
                                                    onChange={handleChange} required></input>
                                            </td>
                                            <td className="align-middle col-3" >Ngôn ngữ</td>
                                            <td className="align-middle col-3" >
                                                <input
                                                    className='form-control bg-light'
                                                    name="language"
                                                    value={formData.language}
                                                    disabled={!isEditable}
                                                    onChange={handleChange} required></input>
                                            </td>
                                        </tr>
                                        <tr>
                                            <td className="align-middle col-3" >Khởi chiếu</td>
                                            <td colSpan={3} className="align-middle col-2" >
                                                <input
                                                    type='date'
                                                    className='form-control bg-light'
                                                    name="openingDay"
                                                    value={formData.openingDay}
                                                    disabled={!isEditable}
                                                    onChange={handleChange} required
                                                    style={{ color: 'black', background: 'linear-gradient(90deg, rgba(255,255,255,1) 88%, rgba(101, 190, 195,1) 88%, rgba(101, 190, 195,1) 100%)' }} ></input>
                                            </td>
                                        </tr>
                                        <tr>
                                            <td className="align-middle col-2" >Tóm tắt</td>
                                            <td colSpan={3} className="align-middle col-3" >
                                                <textarea
                                                    className='form-control bg-light'
                                                    name="story"
                                                    value={formData.story}
                                                    disabled={!isEditable}
                                                    onChange={handleChange} required rows={8}></textarea>
                                            </td>
                                        </tr>
                                    </tbody>
                                </table>
                            </div>
                            <div className='col-5 m-2'>
                                <table className="table shadow" style={{ borderRadius: '10px', overflow: 'hidden' }}>
                                    <tbody>
                                        <tr>
                                            <td className="align-middle col-1" >Đạo diễn</td>
                                            <td colSpan={3} className="align-middle col-3" >
                                                <input
                                                    className='form-control bg-light'
                                                    name="director"
                                                    value={formData.director}
                                                    disabled={!isEditable}
                                                    onChange={handleChange} required></input>
                                            </td>
                                        </tr>
                                        <tr>
                                            <td className="align-middle col-1" >Diễn viên</td>
                                            <td colSpan={3} className="align-middle col-3" >
                                                <textarea
                                                    className='form-control bg-light'
                                                    name="cast"
                                                    value={formData.cast}
                                                    disabled={!isEditable}
                                                    onChange={handleChange} required rows="3"></textarea>
                                            </td>
                                        </tr>
                                        <tr>
                                            <td className="align-middle col-1" >Trailer</td>
                                            <td colSpan={3} className="align-middle col-3" >
                                                <textarea
                                                    className='form-control bg-light'
                                                    name="trailer"
                                                    value={formData.trailer}
                                                    disabled={!isEditable}
                                                    onChange={handleChange} required rows={2}></textarea>
                                            </td>
                                        </tr>
                                        <tr>
                                            <td className="align-middle col-1" >Poster</td>
                                            <td colSpan={3} className="align-middle col-3" >
                                                <input
                                                    type='file'
                                                    className='form-control bg-light'
                                                    name="poster"
                                                    disabled={!isEditable}
                                                    onChange={handleChangeFormFile} required></input>
                                                {selectedImage && selectedImage !== '' && (
                                                    <img
                                                        src={selectedImage}
                                                        alt="Selected Poster"
                                                        style={{ maxHeight: '13.7rem', marginTop: '10px' }}
                                                    />
                                                )}
                                            </td>
                                        </tr>
                                    </tbody>
                                </table>
                            </div>
                        </div>

                        {message.text !== '' && message.isShow && (
                            <div className={message.success ? 'text-success' : 'text-danger'}>
                                {message.text}
                            </div>
                        )}
                        {action === 'edit-movie' ? (
                            <>
                                <button className='btn btn-primary m-2' disabled={disabledBtn} type="button" onClick={handleBack}>
                                    Quay về
                                </button>
                                {isEditable ? (
                                    <button className='btn btn-primary m-2' disabled={disabledBtn} type="button" onClick={editMovie}>
                                        Lưu
                                    </button>
                                ) : (
                                    <button className='btn btn-primary m-2' disabled={disabledBtn} type="button" onClick={handleEditToggle}>
                                        Chỉnh sửa
                                    </button>
                                )}
                            </>
                        ) : (
                            <>
                                <button className='btn btn-primary m-2' disabled={disabledBtn} type="button" onClick={handleBack}>
                                    Quay về
                                </button>
                                <button className='btn btn-primary m-2' disabled={disabledBtn} type="button" onClick={addMovie}>
                                    Thêm
                                </button>
                            </>
                        )}
                    </div>

                )}


            </div>
        </>
    )
}

export default Movie