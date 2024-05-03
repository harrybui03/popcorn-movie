import React, { useState, useEffect } from 'react';
import useAuth from "../../hooks/useAuth";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faHouse, faAdd, faAngleRight, faMagnifyingGlass, faEdit, faTrash, faUnsorted, faSortUp, faSortDown, faSave, faBackward } from "@fortawesome/free-solid-svg-icons";
import LoadingSpinner from '../defaultPage/Loading'
import { addDays, addMinutes, format, parseISO } from 'date-fns';
import { da } from 'date-fns/locale';


function Showtime({ setDeleteIdShowtime }) {
    const auth = useAuth()

    const colorShowtime = [
        'rgb(219, 137, 45)',
        'rgb(81, 156, 247)',
        'rgb(201, 102, 247)',
        'rgb(242, 118, 143)',
        'rgb(120, 209, 111)',
        'rgb(182, 213, 103)',
        'rgb(224, 122, 97)',
        'rgb(172, 59, 51)',
        'rgb(189, 60, 88)',
        'rgb(75, 200, 173)',
    ];
    const [theater, setTheater] = useState([]);
    const [theaterChosen, setTheaterChosen] = useState(null);
    const getTheater = async () => {
        try {
            // Địa chỉ API và tham số
            const apiUrl = 'http://localhost:8080/theaters';

            // Gọi API bằng phương thức GET
            const response = await fetch(apiUrl);

            // Kiểm tra trạng thái của response
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            else {
                // Chuyển đổi response thành JSON
                const result = await response.json();
                // Cập nhật state với dữ liệu từ API
                setTheater(result);
            }
        } catch (error) {
            console.error('Error fetching data:', error);
        }
    };

    const [room, setRoom] = useState([]);
    const [roomOption, setRoomOption] = useState([]);
    const [roomChosen, setRoomChosen] = useState(null);
    const getRoom = async () => {
        try {
            // Địa chỉ API và tham số
            const apiUrl = 'http://localhost:8080/rooms';

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
                setRoom(result);
            }
        } catch (error) {
            console.error('Error fetching data:', error);
        }
    };

    const handleChangeTheater = (event) => {
        setTheaterChosen(event.target.value);
    };

    const handleChangeRoom = (event) => {
        setRoomChosen(event.target.value);
    };

    useEffect(() => {
        getTheater();
        getRoom();
        getDefaultDate();
        generateTimeList();
    }, []);

    useEffect(() => {
        if (theater.length > 0) {
            setTheaterChosen(theater[0]?.id);
        }
    }, [theater]);

    useEffect(() => {
        if (theaterChosen !== null && theaterChosen !== undefined && room.length > 0) {
            setRoomOption(room.filter(e => e.theater.id === parseInt(theaterChosen)));
        }
    }, [room]);

    useEffect(() => {
        if (theaterChosen !== null && theaterChosen !== undefined && room.length > 0) {
            setRoomOption(room.filter(e => e.theater.id === parseInt(theaterChosen)));
        }
    }, [theaterChosen]);

    useEffect(() => {
        if (roomOption.length > 0) {
            setRoomChosen(roomOption[0]?.id);
        }
    }, [roomOption]);

    const [defaultDate, setDefaultDate] = useState([]);
    const [customDate, setCustomDate] = useState([]);
    const [displayDate, setDisplayDate] = useState([]);
    const [defaultShowtime, setDefaultShowtime] = useState([]);
    const [customShowtime, setCustomShowtime] = useState([]);
    const [displayShowtime, setDisplayShowtime] = useState([]);
    const [dateFrom, setDateFrom] = useState('');
    const [dateTo, setDateTo] = useState('');

    useEffect(() => {
        if (roomChosen !== null && theaterChosen !== null && defaultDate.length > 0 && (dateFrom === '' || dateTo === '')) {
            setDisplayDate(defaultDate);
            getDefaultShowtime();
        } else {
            getCustomShowtime();
        }
    }, [roomChosen]);

    const getDefaultDate = () => {
        const today = new Date();
        today.setHours(0, 0, 0, 0); // Đặt giờ, phút, giây và mili giây về 0 để so sánh chính xác ngày

        const sevenDays = Array.from({ length: 7 }, (_, index) => {
            const currentDate = new Date(today);
            currentDate.setDate(today.getDate() + index);
            return format(currentDate, 'yyyy-MM-dd');
        });

        setDefaultDate(sevenDays);
    };

    const handleChangeDate = (e) => {
        const { name, value } = e.target;
        console.log(name, value);
        console.log(value === '');
        if (name === 'dateFrom') {
            setDateFrom(value);
        } else {
            setDateTo(value);
        }
    };

    function isValidDateRange(dateFrom, dateTo) {
        const startDate = new Date(dateFrom);
        const endDate = new Date(dateTo);

        return startDate < endDate;
    }
    useEffect(() => {
        if (dateFrom !== '' && dateTo !== '' && isValidDateRange(dateFrom, dateTo)) {
            const a = generateDateRange(dateFrom, dateTo);
            setCustomDate(a);
        } else {
            console.log('hihi');
            setCustomDate([]);
        }
    }, [dateFrom, dateTo]);

    function generateDateRange(dateFrom, dateTo) {
        const startDate = new Date(dateFrom);
        const endDate = new Date(dateTo);
        let dateList = [];

        while (startDate <= endDate) {
            dateList.push(startDate.toISOString().split('T')[0]);
            startDate.setDate(startDate.getDate() + 1);
        }
        if (dateList.length > 7) {
            dateList = dateList.slice(0, 7);
        }
        return dateList;
    }

    const [timeList, setTimeList] = useState([]);

    const generateTimeList = () => {
        const timeList = [];
        let currentTime = new Date("2000-01-01T00:00:00"); // Sử dụng một ngày cố định (đủ dài) để đảm bảo chạy qua cả 24 giờ

        while (currentTime.getDate() === 1) {
            const hours = currentTime.getHours().toString().padStart(2, "0");
            const minutes = currentTime.getMinutes().toString().padStart(2, "0");
            const timeString = `${hours}:${minutes}`;
            timeList.push(timeString);

            // Thêm 30 phút cho thời gian hiện tại
            currentTime.setMinutes(currentTime.getMinutes() + 30);
        }
        setTimeList(timeList);
    }

    const getShowTime = async (theater, room, date) => {
        try {
            // Địa chỉ API và tham số
            const apiUrl = `http://localhost:8080/showtimes?theater=${theater}&date=${date}`;

            // Gọi API bằng phương thức GET
            const response = await fetch(apiUrl);

            // Kiểm tra trạng thái của response
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            else {
                // Chuyển đổi response thành JSON
                const result = await response.json();
                const temp = result.showTimes;
                const data = temp.filter(e => e.room.id === parseInt(room));
                data.sort((a, b) => new Date(a.startAt) - new Date(b.startAt));
                return data;
            }
        } catch (error) {
            console.error('Error fetching data:', error);
        }
    }

    const getDefaultShowtime = async () => {
        const temp = [];
        for (let i = 0; i < defaultDate.length; i++) {
            const a = await getShowTime(theaterChosen, roomChosen, defaultDate[i]);
            temp.push(a);
        }
        setDefaultShowtime(temp);
        // console.log(temp);
        // console.log(timeList);
        // console.log(defaultDate);
    }
    const getCustomShowtime = async () => {
        const temp = [];
        for (let i = 0; i < customDate.length; i++) {
            const a = await getShowTime(theaterChosen, roomChosen, customDate[i]);
            temp.push(a);
        }
        setCustomShowtime(temp);
        // console.log(temp);
        // console.log(timeList);
        // console.log(defaultDate);
    }

    useEffect(() => {
        if (defaultDate.length > 0) {
            setDisplayDate(defaultDate);
        }
    }, [defaultDate]);

    useEffect(() => {
        if (roomChosen !== null && theaterChosen !== null && defaultDate.length > 0 && (dateFrom === '' || dateTo === '')) {
            setDisplayDate(defaultDate);
            getDefaultShowtime();
        } else {
            setDisplayDate(customDate);
            getCustomShowtime();
        }
    }, [customDate]);

    useEffect(() => {
        if (defaultShowtime.length > 0) {
            setDisplayShowtime(defaultShowtime);
        }
        console.log(defaultShowtime);
    }, [defaultShowtime]);

    useEffect(() => {
        if (customShowtime.length > 0) {
            setDisplayShowtime(customShowtime);
        } else {
            setDisplayShowtime(defaultShowtime);
        }
    }, [customShowtime]);

    const compareTimes = (time1, time2) => {
        // Chia chuỗi thời gian thành giờ và phút
        if (time1.length > time2.length) {
            time1 = time1.split('T')[1];
        } else {
            time2 = time2.split('T')[1];
        }
        const [hour1, minute1] = time1.split(':').map(Number);
        const [hour2, minute2] = time2.split(':').map(Number);

        // So sánh giờ và phút
        if (hour1 < hour2 || (hour1 === hour2 && minute1 < minute2)) {
            return -1;
        } else if (hour1 > hour2 || (hour1 === hour2 && minute1 > minute2)) {
            return 1;
        } else {
            return 0;
        }
    };

    let i = 0;
    let carry = 0;
    let cloneCarry = carry;
    let cloneI = i;
    let showtime2date = null;
    let cloneShowtime2date = showtime2date;
    let isNewDate = false;

    const increaseI = () => {
        i++;
    }
    const setI = (n) => {
        i = n;
    }
    const decreaseCarry = () => {
        carry--;
    }
    const setCarry = (n) => {
        carry = n;
    }
    const checkCarry = () => {
        if (carry === 0) {
            showtime2date = null;
        }
    }

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
                const data = result.filter(e => e.status !== 'Over')
                // Cập nhật state với dữ liệu từ API
                setMovies(data);
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

    const [messageAddShowtime, setMessageAddShowtime] = useState({ isShow: false, text: '', success: false, item: null });
    const [showtimeData, setShowtimeData] = useState({});
    const [timeStartAddShowtime, setTimeStartAddShowtime] = useState('');
    const [isConflict, setIsConflict] = useState(true);

    const addShowtime = (indexDate, item) => {
        setTimeStartAddShowtime(displayDate[indexDate] + 'T' + item + ':00');
        handleChangeAction('add-showtime');
    }

    function calculateEnd(start, duration) {
        var startDate = new Date(start);
        var endDate = addMinutes(startDate, duration);
        var endFormatted = format(endDate, "yyyy-MM-dd'T'HH:mm:ss");
        return endFormatted;
    }

    async function checkAvailableShowtime(end) {
        try {
            const response = await fetch(`http://localhost:8080/rooms/available?room=${parseInt(roomChosen)}&start=${timeStartAddShowtime}&end=${end}`, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${auth.accessToken}`,
                },
            });

            if (!response.ok) {
                return false;
            }
            const data = await response.json();
            if (data.message === 'Room is available') {
                return true;
            } else {
                return false;
            }
        } catch (error) {
            console.error('Error:', error);
            throw error;
        }
    }

    const checkConflict = async (item) => {
        const end = calculateEnd(timeStartAddShowtime, item.duration);
        const condition = await checkAvailableShowtime(end);
        if (condition) {
            const data = {
                movieId: item.id,
                roomId: parseInt(roomChosen),
                startAt: timeStartAddShowtime,
                endAt: end,
            };
            setIsConflict(false);
            setShowtimeData(data);
            setMessageAddShowtime({ isShow: true, text: 'Đã chọn, vui lòng lưu để cập nhật dữ liệu.', success: true, item: item });
        } else {
            setIsConflict(true);
            setShowtimeData({});
            setMessageAddShowtime({ isShow: true, text: 'Bị xung đột thời gian, vui lòng chọn phim hoặc giờ chiếu khác!', success: false, item: item });
        }
    }

    async function generateTicket(ticket) {
        try {
            const response = await fetch(`http://localhost:8080/tickets/generate-ticket/${ticket.id}?price=50000`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${auth.accessToken}`,
                },
            });

            if (!response.ok) {
                throw new Error(`HTTP error! Status: ${response.status}`);
            }
        } catch (error) {
            console.error('Error:', error);
            throw error;
        }
    }

    const postShowtime = async () => {
        try {
            const response = await fetch('http://localhost:8080/showtimes', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${auth.accessToken}`,
                },
                body: JSON.stringify(showtimeData),
            });

            if (!response.ok) {
                throw new Error(`HTTP error! Status: ${response.status}`);
            }
            const data = await response.json();
            generateTicket(data);
            setShowtimeData({});
            setTimeStartAddShowtime('');
            setIsConflict(true);
            setMessageAddShowtime({ isShow: false, text: '', success: false, item: null });
        } catch (error) {
            console.error('Error:', error);
            throw error;
        }
    };

    const [showtimeDelete, setShowtimeDelete] = useState({});
    const [messageDeleteShowtime, setMessageDeleteShowtime] = useState({ isShow: false, text: '', success: false });

    const editShowtime = (data) => {
        setShowtimeDelete(data);
        console.log(data)
        handleChangeAction('edit-showtime');
    }

    // async function deleteShowtime() {
    //     try {
    //         const response = await fetch(`http://localhost:8080/showtimes/${showtimeDelete.id}`, {
    //             method: 'DELETE',
    //             headers: {
    //                 'Content-Type': 'application/json',
    //                 'Authorization': `Bearer ${auth.accessToken}`,
    //             },
    //         });

    //         if (!response.ok) {
    //             throw new Error(`HTTP error! Status: ${response.status}`);
    //         }
    //         const data = await response.json();
    //         if (data.message === 'Some tickets of movies are booked. Can not delete this showtime') {
    //             setMessageDeleteShowtime({ isShow: true, text: 'Đã có người đặt vé, không thể xóa!', success: false });
    //         } else {
    //             setMessageDeleteShowtime({ isShow: true, text: 'Xóa lịch chiếu thành công', success: true });
    //         }
    //     } catch (error) {
    //         console.error('Error:', error);
    //         throw error;
    //     }
    // }

    const deleteTicket = async () => {
        setDeleteIdShowtime(showtimeDelete.id);
        $('#deleteShowtime').modal('show');
    };

    const [chainAction, setChainAction] = useState([{ text: 'Xếp lịch chiếu', href: '/manage/showtime' }]);
    const [action, setAction] = useState('manage-showtime');

    const handleChangeAction = (a) => {
        const newChange = chainAction;
        newChange.push({ text: a === 'add-showtime' ? 'Thêm lịch chiếu' : 'Chỉnh sửa lịch chiếu', href: '#' })
        setChainAction(newChange);
        setAction(a);
        // if (a === 'edit-movie') {
        // } else {
        // }
    }

    const handleBack = () => {
        setAction('manage-showtime');
        const newChain = chainAction;
        newChain.pop();
        setChainAction(newChain);
        setShowtimeData({});
        setShowtimeDelete({});
        setTimeStartAddShowtime('');
        setIsConflict(true);
        setMessageAddShowtime({ isShow: false, text: '', success: false, item: null });
        setMessageDeleteShowtime({ isShow: false, text: '', success: false });
        if (dateFrom !== '' && dateTo !== '') {
            getCustomShowtime();
        } else {
            getDefaultShowtime();
        }
    };

    // Hàm so sánh ngày
    function compareDate(date1, date2) {
        // Chuyển đổi chuỗi ngày thành đối tượng Date
        const datetime1 = new Date(date1);
        const datetime2 = new Date(date2);
        console.log(date1, date2);
        // So sánh ngày
        if (datetime1 < datetime2) {
            return -1;
        } else if (datetime1 > datetime2) {
            return 1;
        } else {
            return 0;
        }
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


                {action === 'manage-showtime' ? (
                    <>
                        <div className='my-3 mx-1 d-flex justify-content-between align-items-center'>
                            <form className='mx-3 d-flex justify-content-start align-items-center'>
                                <label htmlFor='chooseTheater' className='mx-2'>Chọn rạp</label>
                                <select defaultValue={theaterChosen?.toString()} id='chooseTheater' className='p-1 rounded-2' style={{ backgroundColor: 'white', color: 'black' }} onChange={handleChangeTheater}>
                                    {theater.map((item) => (
                                        <option key={item.id} value={item.id.toString()}>{item.name}</option>
                                    ))}
                                </select>
                            </form>
                            <form className='mx-3 d-flex justify-content-start align-items-center'>
                                <label htmlFor='chooseRoom' className='mx-2'>Chọn Phòng</label>
                                <select defaultValue={roomChosen?.toString()} id='chooseRoom' className='p-1 rounded-2' style={{ backgroundColor: 'white', color: 'black' }} onChange={handleChangeRoom}>
                                    {roomOption.map((item) => (
                                        <option key={item.id} value={item.id.toString()}>{item.roomNumber}</option>
                                    ))}
                                </select>
                            </form>
                            <form className='mx-3 d-flex justify-content-start align-items-center'>
                                <label htmlFor='from' className='mx-2'>Từ ngày</label>
                                <input onChange={handleChangeDate} id='from' className='p-1 rounded-2' type='date' name='dateFrom' value={dateFrom}
                                    style={{ color: 'black', background: 'linear-gradient(90deg, rgba(255,255,255,1) 80%, rgba(101, 190, 195,1) 80%, rgba(101, 190, 195,1) 100%)' }} ></input>
                            </form>
                            <form className='mx-3 d-flex justify-content-start align-items-center'>
                                <label htmlFor='to' className='mx-2'>Đến ngày</label>
                                <input onChange={handleChangeDate} id='to' className='p-1 rounded-2' type='date' name='dateTo' value={dateTo}
                                    style={{ color: 'black', background: 'linear-gradient(90deg, rgba(255,255,255,1) 80%, rgba(101, 190, 195,1) 80%, rgba(101, 190, 195,1) 100%)' }} ></input>
                            </form>
                        </div>
                        <div id='table' className='bg-light p-2 rounded-3 shadow' style={{ overflowX: 'auto', maxWidth: '100%', cursor: 'context-menu' }}>
                            <div id='header' className='d-flex'>
                                {displayDate.length > 0 ? (
                                    <>
                                        <div className='d-flex flex-column' style={{ width: '5rem' }}>
                                            <div className='p-2 fw-bold rounded-2 border border-light d-flex justify-content-center align-items-center'
                                                style={{ width: '5rem', height: '3rem', background: '#cfe2ff' }}>Ngày</div>
                                        </div>
                                        {displayDate.map((item) => (
                                            <div key={item} className='p-2 fw-bold rounded-2 border border-light d-flex justify-content-center align-items-center'
                                                style={{ width: '100%', height: '3rem', minWidth: '8rem', background: '#cfe2ff' }}>{item}</div>

                                        ))}
                                    </>
                                ) : (<></>)}
                            </div>
                            <div className='d-flex'>
                                <div id='body' className='d-flex flex-column' style={{ width: '5rem' }}>
                                    {timeList.length > 0 ? (
                                        <>
                                            {timeList.map((item) => (
                                                <div key={item} className='p-2 fw-bold rounded-2 border border-light d-flex justify-content-center align-items-center'
                                                    style={{ width: '5rem', height: '3rem', background: '#cfe2ff' }}>{item}</div>
                                            ))}
                                        </>
                                    ) : (<></>)}
                                </div>
                                {displayShowtime.length > 0 ? (
                                    <>
                                        {displayShowtime.map((date, indexDate) => (
                                            <div key={indexDate} id='body' className='d-flex flex-column' style={{ width: '100%', minWidth: '8rem' }}>
                                                {timeList.length > 0 ? (
                                                    setI(0),
                                                    checkCarry(),
                                                    isNewDate = true,
                                                    timeList.map((item, index) => {
                                                        let st;
                                                        if (carry > 0 && isNewDate) {
                                                            st = showtime2date;
                                                        } else {
                                                            st = date[i];
                                                        }
                                                        return ((carry > 0 && isNewDate) ? (
                                                            cloneCarry = carry,
                                                            decreaseCarry(),
                                                            cloneShowtime2date = showtime2date,
                                                            showtime2date = null,
                                                            isNewDate = false,
                                                            <div key={item} onClick={() => editShowtime(st)} className='p-2 fw-bold showtime rounded-2 border border-light d-flex justify-content-center align-items-center'
                                                                style={{ width: '100%', height: `${3 * (cloneCarry)}rem`, minWidth: '8rem', background: `${colorShowtime[cloneShowtime2date?.movie.id % 10]}`, color: 'white' }}>
                                                                {cloneShowtime2date?.movie.title}
                                                            </div>
                                                        ) : (carry > 0 && !isNewDate) ? (
                                                            decreaseCarry(),
                                                            isNewDate = false,
                                                            <div key={index}></div>
                                                        ) : (date.length <= 0) ? (
                                                            isNewDate = false,
                                                            <div key={item} onClick={() => { if (compareDate(displayDate[indexDate], defaultDate[0]) !== -1) addShowtime(indexDate, item) }} className={`p-2 fw-bold ${compareDate(displayDate[indexDate], defaultDate[0]) !== -1 ? 'showtime' : ''} rounded-2 border border-light d-flex justify-content-center align-items-center`}
                                                                style={{ width: '100%', height: `${3}rem`, minWidth: '8rem', background: 'white', color: 'gray' }}>
                                                                {compareDate(displayDate[indexDate], defaultDate[0]) !== -1 && <FontAwesomeIcon className='mx-3' icon={faAdd} />}
                                                            </div>
                                                        ) : (i >= date.length) ? (
                                                            isNewDate = false,
                                                            <div key={item} onClick={() => { if (compareDate(displayDate[indexDate], defaultDate[0]) !== -1) addShowtime(indexDate, item) }} className={`p-2 fw-bold ${compareDate(displayDate[indexDate], defaultDate[0]) !== -1 ? 'showtime' : ''} rounded-2 border border-light d-flex justify-content-center align-items-center`}
                                                                style={{ width: '100%', height: `${3}rem`, minWidth: '8rem', background: 'white', color: 'gray' }}>
                                                                {compareDate(displayDate[indexDate], defaultDate[0]) !== -1 && <FontAwesomeIcon className='mx-3' icon={faAdd} />}
                                                            </div>
                                                        ) : (compareTimes((index < 47 ? timeList[index + 1] + '' : '24:00'), date[i].startAt + '') <= 0) ? (
                                                            isNewDate = false,
                                                            <div key={item} onClick={() => { if (compareDate(displayDate[indexDate], defaultDate[0]) !== -1) addShowtime(indexDate, item) }} className={`p-2 fw-bold ${compareDate(displayDate[indexDate], defaultDate[0]) !== -1 ? 'showtime' : ''} rounded-2 border border-light d-flex justify-content-center align-items-center`}
                                                                style={{ width: '100%', height: `${3}rem`, minWidth: '8rem', background: 'white', color: 'gray' }}>
                                                                {compareDate(displayDate[indexDate], defaultDate[0]) !== -1 && <FontAwesomeIcon className='mx-3' icon={faAdd} />}
                                                            </div>
                                                        ) : (
                                                            isNewDate = false,
                                                            setCarry((date[i].movie.duration % 30 === 0 ? parseInt(date[i].movie.duration / 30) : parseInt(date[i].movie.duration / 30) + 1) - 1),
                                                            cloneI = i,
                                                            increaseI(),
                                                            showtime2date = index + carry > 47 ? date[cloneI] : null,
                                                            <div key={item} onClick={() => editShowtime(st)} className='p-2 fw-bold showtime rounded-2 border border-light d-flex justify-content-center align-items-center'
                                                                style={{ width: '100%', height: `${3 * (index + carry > 47 ? 47 - index + 1 : carry + 1)}rem`, minWidth: '8rem', background: `${colorShowtime[date[cloneI]?.movie.id % 10]}`, color: 'white' }}>
                                                                {date[cloneI]?.movie.title}
                                                            </div>
                                                        )
                                                        )
                                                    })
                                                ) : (<></>)}
                                            </div>
                                        ))}
                                    </>
                                ) : (<></>)}


                            </div>
                        </div>
                    </>
                ) : (action === 'add-showtime') ? (
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
                            <div className='d-flex'>
                                <button onClick={() => handleBack()} className='btn btn-primary mx-2'><FontAwesomeIcon className='mx-2' icon={faBackward} />  Quay về</button>
                                <button onClick={() => postShowtime()} className='btn btn-primary mx-2' disabled={isConflict}><FontAwesomeIcon className='mx-2' icon={faSave} />  Lưu</button>
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
                                                    <button onClick={() => checkConflict(item)} className='btn btn-sm text-primary m-2' style={{ backgroundColor: '#ffffff' }}>
                                                        <FontAwesomeIcon icon={faAdd} /> Chọn
                                                    </button>
                                                </div>
                                                {messageAddShowtime.text !== '' && messageAddShowtime.isShow && messageAddShowtime.item.id === item.id && (
                                                    <div className={messageAddShowtime.success ? 'text-success' : 'text-danger'}>
                                                        {messageAddShowtime.text}
                                                    </div>
                                                )}
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
                    <>
                        <div className='container d-flex flex-column justify-content-start align-items-center'>
                            <table className="table shadow" style={{ borderRadius: '10px', overflow: 'hidden' }}>
                                <tbody>
                                    <tr>
                                        <td className="align-middle col-2" >ID phim</td>
                                        <td className="align-middle col-2" >
                                            <input
                                                className='form-control bg-light'
                                                name="id"
                                                value={showtimeDelete.movie.id}
                                                disabled={true}
                                                required></input>
                                        </td>
                                        <td className="align-middle col-3" >Trạng thái</td>
                                        <td className="align-middle col-3" >
                                            <select
                                                className='form-control bg-light'
                                                name="status"
                                                value={showtimeDelete.movie.status}
                                                disabled={true}
                                                required>
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
                                                value={showtimeDelete.movie.title}
                                                disabled={true}
                                                required rows={2}></textarea>
                                        </td>
                                    </tr>
                                    <tr>
                                        <td className="align-middle col-3" >Rạp chiếu</td>
                                        <td className="align-middle col-2" >
                                            <input
                                                className='form-control bg-light'
                                                name="duration"
                                                value={showtimeDelete.room.theater.name}
                                                disabled={true}
                                                required></input>
                                        </td>
                                        <td className="align-middle col-3" >Bắt đầu<ul></ul></td>
                                        <td className="align-middle col-2" >
                                            <input
                                                className='form-control bg-light'
                                                name="genre"
                                                value={showtimeDelete.startAt.replace('T', '  ')}
                                                disabled={true}
                                                required></input>
                                        </td>
                                    </tr>
                                    <tr>
                                        <td className="align-middle col-3" >Phòng chiếu</td>
                                        <td className="align-middle col-3" >
                                            <input
                                                className='form-control bg-light'
                                                name="rated"
                                                value={showtimeDelete.room.roomNumber}
                                                disabled={true}
                                                required></input>
                                        </td>
                                        <td className="align-middle col-3" >Kết thúc</td>
                                        <td className="align-middle col-3" >
                                            <input
                                                className='form-control bg-light'
                                                name="language"
                                                value={showtimeDelete.endAt.replace('T', '  ')}
                                                disabled={true}
                                                required></input>
                                        </td>
                                    </tr>
                                    <tr>
                                        <td className="align-middle col-1" >Poster</td>
                                        <td colSpan={3} className="align-middle col-3" >
                                            <img
                                                src={showtimeDelete.movie.poster}
                                                alt="Selected Poster"
                                                style={{ maxHeight: '13.7rem', marginTop: '10px' }}
                                            />
                                        </td>
                                    </tr>
                                </tbody>
                            </table>

                            {messageDeleteShowtime.text !== '' && messageDeleteShowtime.isShow && (
                                <div className={messageDeleteShowtime.success ? 'text-success m-3' : 'text-danger m-3'}>
                                    {messageDeleteShowtime.text}
                                </div>
                            )}
                            <div className='d-flex'>
                                <button onClick={() => handleBack()} className='btn btn-primary mx-2'><FontAwesomeIcon className='mx-2' icon={faBackward} />  Quay về</button>
                                <button onClick={() => deleteTicket()} className='btn btn-danger mx-2' disabled={compareDate(showtimeDelete.startAt.split('T')[0], defaultDate[0]) === -1}><FontAwesomeIcon className='mx-2' icon={faTrash} />  Xóa lịch chiếu</button>
                            </div>
                        </div>
                    </>
                )}

            </div >
        </>
    )
}

export default Showtime