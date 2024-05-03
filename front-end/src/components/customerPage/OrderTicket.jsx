import React, { useState, useEffect } from 'react';
import Header from './Header'
import Footer from '../defaultPage/Footer'
import LogOut from './LogOut'
import ChangePW from './ChangePW'
import { useParams } from 'react-router-dom';
import { Form } from 'react-bootstrap';
import LoadingSpinner from '../defaultPage/Loading'
import { addDays, format, parseISO } from 'date-fns';
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faXmark } from "@fortawesome/free-solid-svg-icons";
import './styles.css'
import useAuth from "../../hooks/useAuth";

function CustomerDashBoard(params) {
  const { id } = useParams();
  const auth = useAuth();

  const [step, setStep] = useState('choose-time');
  const [prevStep, setPrevStep] = useState(null);
  const [movie, setMovie] = useState(null);
  const [theater, setTheater] = useState([]);
  const [date, setDate] = useState([]);

  const [showtimes, setShowTimes] = useState(null);

  const getMovie = async () => {
    try {
      // Địa chỉ API và tham số
      const apiUrl1 = 'http://localhost:8080/movies?status=Ongoing';
      const apiUrl2 = 'http://localhost:8080/movies?status=Upcoming';

      // Gọi API bằng phương thức GET
      const response1 = await fetch(apiUrl1);
      const response2 = await fetch(apiUrl2);

      // Kiểm tra trạng thái của response
      if (!response1.ok || !response2.ok) {
        throw new Error('Network response was not ok');
      }
      else {
        // Chuyển đổi response thành JSON
        const result1 = await response1.json();
        const result2 = await response2.json();
        const data = result1.concat(result2).filter(e => e.id === parseInt(id));
        // Cập nhật state với dữ liệu từ API
        setMovie(data[0]);
      }
    } catch (error) {
      console.error('Error fetching data:', error);
    }
  };
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
  const getDate = () => {
    const today = new Date();
    const fiveDays = Array.from({ length: 5 }, (_, index) => {
      const currentDate = addDays(today, index);
      return format(currentDate, 'dd/MM/yyyy');
    });
    setDate(fiveDays);
  };
  const getShowTime = async (idMovie, idTheater, date) => {
    try {
      // Địa chỉ API và tham số
      const apiUrl = `http://localhost:8080/showtimes?theater=${idTheater}&movie=${idMovie}`;

      // Gọi API bằng phương thức GET
      const response = await fetch(apiUrl);

      // Kiểm tra trạng thái của response
      if (!response.ok) {
        throw new Error('Network response was not ok');
      }
      else {
        // Chuyển đổi response thành JSON
        const result = await response.json();
        const data = result.filter(e => {
          const dateString = e.startAt;
          const dateTime = parseISO(dateString);

          const datePart = format(dateTime, 'dd/MM/yyyy');
          return datePart === date;
        });
        data.sort((a, b) => {
          const timeA = parseISO(a.startAt);
          const timeB = parseISO(b.startAt);
          return timeA - timeB;
        });
        setShowTimes(data);
      }
    } catch (error) {
      console.error('Error fetching data:', error);
    }
  }
  const [theaterChosen, setTheaterChosen] = useState(theater[0]?.id);
  const [dateChosen, setDateChosen] = useState(date[0]);

  const [theaterChosenName, setTheaterChosenName] = useState(theater[0]);

  useEffect(() => {
    getMovie();
    getTheater();
    getDate();
  }, []);

  useEffect(() => {
    setTheaterChosen(theater[0]?.id);
    setDateChosen(date[0]);
    setTheaterChosenName((theater.filter(e => e.id === theater[0]?.id))[0]?.name);
  }, [movie, theater, date]);


  useEffect(() => {
    // console.log(theaterChosen);
    // console.log(dateChosen);
    if (theaterChosen) {
      getShowTime(id, theaterChosen, dateChosen);
    }
  }, [dateChosen, theaterChosen]);

  const handleChangeTheater = (event) => {
    setTheaterChosen(event.target.value);
    setTheaterChosenName((theater.filter(e => e.id === parseInt(event.target.value)))[0]?.name);
  };
  const handleChangeDate = (event) => {
    setDateChosen(event.target.value);
  };

  const [showtimeChosen, setShowTimeChosen] = useState(null);
  const [tickets, setTickets] = useState(null);

  const getTickets = async (idShowtime) => {
    try {
      // Địa chỉ API và tham số
      const apiUrl = `http://localhost:8080/tickets?showtime=${idShowtime}`;

      // Gọi API bằng phương thức GET
      const response = await fetch(apiUrl);

      // Kiểm tra trạng thái của response
      if (!response.ok) {
        throw new Error('Network response was not ok');
      }
      else {
        // Chuyển đổi response thành JSON
        const result = await response.json();
        result.sort((a, b) => {
          const aPrefix = a.seat.seatNumber.substring(0, 1);
          const bPrefix = b.seat.seatNumber.substring(0, 1);
          if (aPrefix === bPrefix) {
            const aNumber = parseInt(a.seat.seatNumber.substring(1), 10);
            const bNumber = parseInt(b.seat.seatNumber.substring(1), 10);
            return aNumber - bNumber;
          }
          return aPrefix.localeCompare(bPrefix);
        });
        setTickets(result);
      }
    } catch (error) {
      console.error('Error fetching data:', error);
    }
  }

  const handleChooseSeat = (showtime) => {
    setStep('choose-seat');
    setShowTimeChosen(showtime);
  }

  useEffect(() => {
    if (showtimeChosen !== null) {
      getTickets(showtimeChosen?.id);
    }
  }, [showtimeChosen]);

  const handleDisabledTime = (time) => {
    const currentTime = new Date();
    const targetTime = new Date(time);
    // Tính toán và kiểm tra khoảng thời gian
    const timeDifference = (targetTime.getTime() + 30 * 60 * 1000) - currentTime.getTime();
    return timeDifference < 0;
  }

  const handleBack = () => {
    setStep(prevStep);
  }
  const handleBuyFood = () => {
    setStep('choose-food');
  }
  const handlePayment = () => {
    setStep('choose-payment');
  }
  const handleCloseBooking = () => {
    setStep('choose-time');
  }

  useEffect(() => {
    if (step !== 'choose-payment') {
      setIsCheckedCondition(false);
      setPaymentMethod(false);
    }
    if (step === 'choose-time') {
      setShowTimeChosen(null);
      setPrevStep(null);
      setSeatNumbers('');
      setSeats([]);
      setTotal(0);
    } else if (step === 'choose-seat') {
      setPrevStep('choose-time');
    } else if (step === 'choose-food') {
      setPrevStep('choose-seat');
    } else if (step === 'choose-payment') {
      if (prevStep === 'choose-seat') {
        setPrevStep('choose-food');
      } else {
        setPrevStep('choose-seat');
      }
    }
  }, [step]);

  const [seats, setSeats] = useState([]);
  const [total, setTotal] = useState(0);
  const [seatNumbers, setSeatNumbers] = useState('');

  const [foods, setFoods] = useState([]);
  const [feeFoods, setfeeFoods] = useState(0);

  const handleSelectSeats = (ticket) => {
    const seatIndex = seats.indexOf(ticket);
    if (seatIndex === -1) {
      // Ghế chưa được chọn, thêm vào mảng
      setSeats([...seats, ticket]);
    } else {
      // Ghế đã được chọn, loại bỏ khỏi mảng
      const newSeats = [...seats];
      newSeats.splice(seatIndex, 1);
      setSeats(newSeats);
    }
  }

  useEffect(() => {
    console.log(seats);
    let sum = 0;
    let str = '';
    seats.forEach(e => {
      sum += e.price;
      str += e.seat.seatNumber + ', ';
    })
    setSeatNumbers(str.substring(0, str.length - 2));
    setTotal(sum);
  }, [seats])

  const [isCheckedCondition, setIsCheckedCondition] = useState(false);
  const handleCheckCondition = () => {
    setIsCheckedCondition(!isCheckedCondition);
  }

  const [paymentMethod, setPaymentMethod] = useState(false);
  const handlePaymentMethod = () => {
    setPaymentMethod(!paymentMethod);
  }

  const [message, setMessage] = useState({ isShow: false, text: '', status: '' });
  const [disabledBtnPayment, setDisabledBtnPayment] = useState(false);

  const handleCheckoutInit = async () => {
    if (!isCheckedCondition || !paymentMethod) {
      alert("Vui lòng chọn chấp nhận các điều khoản, điều kiện và phương thức thanh toán.");
    } else {
      setDisabledBtnPayment(true);
      setMessage({ isShow: true, text: 'Đang xử lý dữ liệu, vui lòng đợi...', status: 'loading' });
      const response = await postCheckoutInit();
      setDisabledBtnPayment(false);
      // setMessage({ isShow: false, text: '', status: '' });
      if (response && response.status === 'Success') {
        console.log(seats, foods);
        let seatId = [];
        // let foodId = [];
        for (let i of seats) {
          seatId.push(i.id);
        }
        // for (let i of foods) {
        //   Thêm food here
        //   {
        //     "foodId": 0,
        //     "quantity": 0
        //   }
        //   foodId.push()
        // }
        const item = {
          "ticketIds": seatId,
          "orderLineDtoList": []
        };
        localStorage.setItem('listItem', JSON.stringify(item));
        window.location.href = response.redirectUrl;
      }
    }
  }

  const postCheckoutInit = async () => {
    try {
      const response = await fetch(`http://localhost:8080/checkout/init?fee=${total}`, {
        method: 'POST',
        headers: {
          'Authorization': `Bearer ${auth.accessToken}`,
        }
      });

      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }
      const result = await response.json();
      return result;
    } catch (error) {
      console.error('Error:', error);
      throw error;
    }
  }

  return (
    <>
      <Header {...params} ></Header>
      {step === 'choose-time' && movie ? (
        <>
          <Form className='d-flex justify-content-center mt-5'>
            <Form.Group controlId="theater" className='mx-5'>
              <Form.Label>Rạp</Form.Label>
              <Form.Control as="select" className="" onChange={handleChangeTheater}>
                {theater.map((item) => (
                  <option key={item.id} value={item.id}>{item.name}</option>
                ))}
              </Form.Control>
            </Form.Group>
            <Form.Group controlId="date" className='mx-5'>
              <Form.Label>Lịch chiếu</Form.Label>
              <Form.Control as="select" className="" onChange={handleChangeDate}>
                {date.map((item) => (
                  <option key={item} value={item}>{item}</option>
                ))}
              </Form.Control>
            </Form.Group>
          </Form>
          <div className="container mx-auto my-5">
            <div style={{ backgroundColor: 'rgb(219, 137, 45)', userSelect: 'none' }}
              className="m-0 p-3 shadow border border-5 rounded-5 d-flex text-light col-12">
              <img src={movie.poster} className='col-4 p-0' alt={`${movie.id} Poster`} />
              <div className="col-8 d-flex flex-column justify-content-start px-5">
                <h1 className="mx-auto my-4">{movie.title}</h1>
                <div className="d-flex flex-column align-items-start my-4" >
                  <h3>{theaterChosenName}</h3>
                  <h3>{dateChosen}</h3>
                </div>
                {!showtimes ? (
                  <div className='d-flex justify-content-center my-5'>
                    <LoadingSpinner />
                  </div>
                ) : showtimes.length === 0 ? (
                  <h5>Chưa có suất chiếu</h5>
                ) : (
                  < div className='flex-wrap d-flex justify-content-start'>
                    {showtimes.map((item) => (
                      <button onClick={() => handleChooseSeat(item)} key={item.id} className='btn btn-success col-2 text-light m-2' disabled={handleDisabledTime(item.startAt)}>{format(parseISO(item.startAt), 'HH:mm:ss').substring(0, 5)}</button>
                    ))}
                  </div>
                )}
              </div>
            </div>
          </div>
        </>
      ) : step === 'choose-seat' && showtimeChosen ? (
        <div >
          <div className='d-flex flex-column justify-content-start align-items-center p-5' style={{ position: 'relative', overflow: 'hidden', margin: '1px', background: 'linear-gradient(176deg, rgba(251,255,200,1) 42%, rgba(149,255,169,1) 100%)' }}>
            <button onClick={handleCloseBooking} type="button" className="btn btn-lg btn-secondary rounded-circle position-absolute top-0 end-0" style={{ marginRight: '1rem', marginTop: '1rem' }}>
              <FontAwesomeIcon icon={faXmark} />
            </button>
            <div className='d-flex justify-content-start align-items-center' >
              <div className='m-3 '>Tên phim</div>
              <h4 className='text-success text-start fw-bold'>{showtimeChosen.movie.title}</h4>
            </div>
            <div className='d-flex justify-content-center'>
              <div >
                <div id='time' className='d-flex justify-content-center' style={{ marginBottom: '1px' }}>
                  <div className='p-3' style={{ marginRight: '1px', backgroundColor: 'rgb(219, 137, 45)', width: '10rem' }}>
                    <div>Suất chiếu</div>
                    <h4 className='fw-bold'>
                      {format(parseISO(showtimeChosen.startAt), 'HH:mm:ss').substring(0, 5)}
                    </h4>
                  </div>
                  <div className='p-3' style={{ backgroundColor: 'rgb(219, 137, 45)', width: '20rem' }}>
                    <div>
                      Ngày chiếu
                    </div>
                    <h4 className='fw-bold'>
                      {format(parseISO(showtimeChosen.startAt), 'dd/MM/yyyy')}
                    </h4>
                  </div>
                </div>
                <div id='seat' className='d-flex justify-content-start'>
                  <div style={{ width: '10rem', backgroundColor: 'rgb(201, 102, 247)', marginRight: '1px' }} className='p-3 d-flex justify-content-center align-items-center'>Số ghế</div>
                  <div style={{ width: '20rem', backgroundColor: 'rgb(201, 102, 247)' }} className='fw-bold p-3'>
                    <h4>
                      {seatNumbers}
                    </h4>
                  </div>
                </div>
              </div>
              <div className='p-3 d-flex flex-column justify-content-center align-items-center' style={{ backgroundColor: 'rgb(242, 118, 143)', marginLeft: '1px', width: '15rem' }}>
                <div>Tổng tiền</div>
                <h4 className='fw-bold'>{total} <u>đ</u></h4>
              </div>
            </div>
          </div>
          <h5 className='d-flex justify-content-center align-items-center p-3' style={{ margin: '1px', background: 'rgb(81, 156, 247)' }}>
            {showtimeChosen.room.theater.name.toUpperCase()}
          </h5>
          <div className='container text-light'>
            <h5 className='m-5 p-3 bg-secondary border border-5 rounded-5'>MÀN HÌNH</h5>
          </div>
          <div className='p-5' style={{ userSelect: 'none' }}>
            {tickets ? (
              tickets.map((item, index) => (
                <React.Fragment key={item.id}>
                  <button
                    className={`btn seat 
                      ${item.booked ? 'seat-booked' : seats.indexOf(item) === -1 ? 'seat-vacant' : 'seat-selecting'}
                      ${item.seat.category === 'Couple' ? 'seat-couple' : 'seat-standard'}
                      
                      `}
                    disabled={item.booked}
                    onClick={() => handleSelectSeats(item)}
                  >
                    {item.seat.seatNumber}
                  </button>
                  {((item.seat.category !== 'Couple' && index % 10 === 4) || (item.seat.category === 'Couple' && (index % 10 === 2 || index % 10 === 7))) && <div style={{ width: '8rem', display: 'inline-block' }}></div>}
                  {(index % 10 === 9 || (item.seat.category === 'Couple' && index % 10 === 4)) && <br />}
                </React.Fragment>
              ))
            ) : (
              <div className='d-flex justify-content-center my-5'>
                <LoadingSpinner />
              </div>
            )}
          </div>
          <div className='container d-flex justify-content-around my-5'>
            <div>
              <button className={`btn m-3 seat-note seat-note-standard-vacant`} disabled={true}></button><label>Ghế thường</label>
            </div>
            <div>
              <button className={`btn m-3 seat-note seat-note-couple-vacant`} disabled={true}></button><label>Ghế đôi</label>
            </div>
            <div>
              <button className={`btn m-3 seat-note seat-note-standard-booked`} disabled={true}></button><label>Ghế đã chọn</label>
            </div>
            <div>
              <button className={`btn m-3 seat-note seat-note-standard-selecting`} disabled={true}></button><label>Ghế đang chọn</label>
            </div>
          </div>
          <div className='container'>
            <button onClick={handleBack} className='btn btn-lg text-light fw-bold m-2' style={{ background: 'orange' }}>Quay lại</button>
            <button onClick={handleBuyFood} className='btn btn-lg text-light fw-bold m-2' style={{ background: 'orange' }}>Mua đồ ăn</button>
            <button onClick={handlePayment} className='btn btn-lg text-light fw-bold m-2' disabled={seats.length <= 0} style={{ background: 'orange' }}>Thanh toán</button>
          </div>
        </div >

      ) : step === 'choose-food' ? (
        <div>
          <div className='d-flex flex-column justify-content-start align-items-center p-5' style={{ position: 'relative', overflow: 'hidden', margin: '1px', background: 'linear-gradient(176deg, rgba(251,255,200,1) 42%, rgba(149,255,169,1) 100%)' }}>
            <button onClick={handleCloseBooking} type="button" className="btn btn-lg btn-secondary rounded-circle position-absolute top-0 end-0" style={{ marginRight: '1rem', marginTop: '1rem' }}>
              <FontAwesomeIcon icon={faXmark} />
            </button>
            <div className='d-flex justify-content-start align-items-center' >
              <div className='m-3 '>Tên phim</div>
              <h4 className='text-success text-start fw-bold'>{showtimeChosen.movie.title}</h4>
            </div>
            <div className='d-flex justify-content-center'>
              <div >
                <div id='time' className='d-flex justify-content-center' style={{ marginBottom: '1px' }}>
                  <div className='p-3' style={{ marginRight: '1px', backgroundColor: 'rgb(219, 137, 45)', width: '10rem' }}>
                    <div>Suất chiếu</div>
                    <h4 className='fw-bold'>
                      {format(parseISO(showtimeChosen.startAt), 'HH:mm:ss').substring(0, 5)}
                    </h4>
                  </div>
                  <div className='p-3' style={{ backgroundColor: 'rgb(219, 137, 45)', width: '20rem' }}>
                    <div>
                      Ngày chiếu
                    </div>
                    <h4 className='fw-bold'>
                      {format(parseISO(showtimeChosen.startAt), 'dd/MM/yyyy')}
                    </h4>
                  </div>
                </div>
                <div id='seat' className='d-flex justify-content-start'>
                  <div style={{ width: '10rem', backgroundColor: 'rgb(201, 102, 247)', marginRight: '1px' }} className='p-3 d-flex justify-content-center align-items-center'>Số ghế</div>
                  <div style={{ width: '20rem', backgroundColor: 'rgb(201, 102, 247)' }} className='fw-bold p-3'>
                    <h4>
                      {seatNumbers}
                    </h4>
                  </div>
                </div>
              </div>
              <div className='p-3 d-flex flex-column justify-content-center align-items-center' style={{ backgroundColor: 'rgb(242, 118, 143)', marginLeft: '1px', width: '15rem' }}>
                <div>Tổng tiền</div>
                <h4 className='fw-bold'>{total} <u>đ</u></h4>
              </div>
            </div>
          </div>
          <h5 className='d-flex justify-content-center align-items-center p-3' style={{ margin: '1px', background: 'rgb(81, 156, 247)' }}>
            {showtimeChosen.room.theater.name.toUpperCase()}
          </h5>
          <div>Choose food</div>
          <div className='container'>
            <button onClick={handleBack} className='btn btn-lg text-light fw-bold m-2' style={{ background: 'orange' }}>Quay lại</button>
            <button onClick={handlePayment} className='btn btn-lg text-light fw-bold m-2' disabled={seats.length <= 0} style={{ background: 'orange' }}>Thanh toán</button>
          </div>
        </div>
      ) : step === 'choose-payment' ? (
        <div>
          <div className='d-flex flex-column justify-content-start align-items-center p-5' style={{ position: 'relative', overflow: 'hidden', margin: '1px', background: 'linear-gradient(176deg, rgba(251,255,200,1) 42%, rgba(149,255,169,1) 100%)' }}>
            <button onClick={handleCloseBooking} type="button" className="btn btn-lg btn-secondary rounded-circle position-absolute top-0 end-0" style={{ marginRight: '1rem', marginTop: '1rem' }}>
              <FontAwesomeIcon icon={faXmark} />
            </button>
            <div className='d-flex justify-content-start align-items-center' >
              <div className='m-3 '>Tên phim</div>
              <h4 className='text-success text-start fw-bold'>{showtimeChosen.movie.title}</h4>
            </div>
            <div className='d-flex justify-content-center'>
              <div >
                <div id='time' className='d-flex justify-content-center' style={{ marginBottom: '1px' }}>
                  <div className='p-3' style={{ marginRight: '1px', backgroundColor: 'rgb(219, 137, 45)', width: '10rem' }}>
                    <div>Suất chiếu</div>
                    <h4 className='fw-bold'>
                      {format(parseISO(showtimeChosen.startAt), 'HH:mm:ss').substring(0, 5)}
                    </h4>
                  </div>
                  <div className='p-3' style={{ backgroundColor: 'rgb(219, 137, 45)', width: '20rem' }}>
                    <div>
                      Ngày chiếu
                    </div>
                    <h4 className='fw-bold'>
                      {format(parseISO(showtimeChosen.startAt), 'dd/MM/yyyy')}
                    </h4>
                  </div>
                </div>
                <div id='seat' className='d-flex justify-content-start'>
                  <div style={{ width: '10rem', backgroundColor: 'rgb(201, 102, 247)', marginRight: '1px' }} className='p-3 d-flex justify-content-center align-items-center'>Số ghế</div>
                  <div style={{ width: '20rem', backgroundColor: 'rgb(201, 102, 247)' }} className='fw-bold p-3'>
                    <h4>
                      {seatNumbers}
                    </h4>
                  </div>
                </div>
              </div>
              <div className='p-3 d-flex flex-column justify-content-center align-items-center' style={{ backgroundColor: 'rgb(242, 118, 143)', marginLeft: '1px', width: '15rem' }}>
                <div>Tổng tiền</div>
                <h4 className='fw-bold'>{total} <u>đ</u></h4>
              </div>
            </div>
          </div>
          <h5 className='d-flex justify-content-center align-items-center p-3' style={{ margin: '1px', background: 'rgb(81, 156, 247)' }}>
            {showtimeChosen.room.theater.name.toUpperCase()}
          </h5>
          <div className="container mx-auto my-5" style={{ paddingLeft: '5rem', paddingRight: '5rem' }}>
            <div style={{ backgroundColor: '#f5f5f5' }}
              className="m-0 p-5 shadow border border-5 rounded-5 d-flex flex-column">
              <h4>Cảm ơn quý khách đã đến với H3DC!</h4>
              <h4>Xin quý khách vui lòng kiểm tra lại thông tin đặt vé</h4>
              <div className='d-flex'>
                <img src={movie.poster} className='col-4 m-3 p-0' alt={`${movie.id} Poster`} style={{ width: '25%' }} />
                <div className="col-8 d-flex flex-column justify-content-start px-5">
                  <div className="d-flex flex-column align-items-start my-5" >
                    <h5>Phim: {movie.title}</h5>
                    <h5>Rạp: {theaterChosenName}</h5>
                    <h5>Phòng: {showtimeChosen.room.roomNumber}</h5>
                    <h5>Ngày chiếu: {dateChosen}</h5>
                    <h5>Suất chiếu: {format(parseISO(showtimeChosen.startAt), 'HH:mm:ss').substring(0, 5)}</h5>
                  </div>
                </div>
              </div>
              <div className='d-flex p-2'>
                <div className='p-3 d-flex justify-content-center align-items-center fs-5' style={{ width: '20%', backgroundColor: '#e3e3e3', margin: '1px' }}>Ghế</div>
                <div className='d-flex flex-column' style={{ width: '80%' }}>
                  {seats.length > 0 ? (
                    seats.map((item) => (
                      <div key={item.id} className='d-flex justify-content-between' style={{ backgroundColor: '#e3e3e3', margin: '1px' }}>
                        <div className='p-3 fs-5'>{item.seat.seatNumber}</div>
                        <div className='p-3 fs-5'>{item.price} <u>đ</u></div>
                      </div>
                    ))
                  ) : (
                    <div className='d-flex justify-content-center my-5'>
                      <LoadingSpinner />
                    </div>
                  )}
                </div>
              </div>
              <div className='d-flex p-2'>
                <div className='p-3 d-flex justify-content-center align-items-center fs-5' style={{ width: '20%', backgroundColor: '#e3e3e3', margin: '1px' }}>Đồ ăn</div>
                <div className='d-flex flex-column' style={{ width: '80%' }}>
                  {foods.length > 0 ? (
                    foods.map((item) => (
                      <div key={item.id} className='d-flex justify-content-between' style={{ backgroundColor: '#e3e3e3', margin: '1px' }}>
                        {/* <div className='p-3 fs-5'>{item.seat.seatNumber}</div>
                        <div className='p-3 fs-5'>{item.price} <u>đ</u></div> */}
                      </div>
                    ))
                  ) : (
                    <div className='d-flex justify-content-between' style={{ backgroundColor: '#e3e3e3', margin: '1px' }}>
                      <div className='p-3 fs-5'></div>
                      <div className='p-3 fs-5'>0 <u>đ</u></div>
                    </div>
                  )}
                </div>
              </div>
              <div className='d-flex justify-content-between p-3' >
                <div className='fw-bold fs-3'>Tổng số tiền</div>
                <div className='fw-bold fs-3'>{total} <u>đ</u></div>
              </div>
              <div className='d-flex flex-column justify-content-start align-items-start p-4'>
                <div className='fs-5'>Quý khách vui lòng nhập mã giảm giá (nếu có)</div>
                <div>
                  <input type='text' placeholder='Nhập mã giảm giá' className='bg-light text-black p-2 m-2' style={{ width: '10rem' }} />
                  <button className='btn btn-primary'>Gửi</button>
                </div>
              </div>
              <div className='d-flex justify-content-between p-4'>
                <div className='fw-bold fs-3'>Số tiền giảm</div>
                <div className='fw-bold fs-3'>0 <u>đ</u></div>
              </div>
              <hr />
              <div className='d-flex justify-content-between p-4'>
                <div className='fw-bold fs-3'>Số tiền cần thanh toán</div>
                <div className='fw-bold fs-3'>{total} <u>đ</u></div>
              </div>
            </div>
          </div>
          <div className="container mx-auto my-5" style={{ paddingLeft: '5rem', paddingRight: '5rem' }}>
            <div style={{ backgroundColor: '#f5f5f5' }}
              className="m-0 p-5 shadow border border-5 rounded-5 d-flex flex-column">
              <div className='fs-2 fw-bold mb-2'>Điều khoản chung</div>
              <div className="overflow-auto p-4 rounded-5" style={{ maxHeight: '25rem', textAlign: 'start', border: 'solid 1px black' }}>
                <h5>Việc bạn sử dụng website này đồng nghĩa với việc bạn đồng ý với những thỏa thuận dưới đây.</h5>
                <h5>Nếu bạn không đồng ý, xin vui lòng không sử dụng website.</h5>
                <div>
                  <strong>1. Trách nhiệm của người sử dụng:</strong>
                  <p>Khi truy cập vào trang web này, bạn đồng ý chấp nhận mọi rủi ro. H3DC và các bên đối tác khác không chịu trách nhiệm về bất kỳ tổn thất nào do những hậu quả trực tiếp, tình cờ hay gián tiếp; những thất thoát, chi phí (bao gồm chi phí pháp lý, chi phí tư vấn hoặc các khoản chi tiêu khác) có thể phát sinh trực tiếp hoặc gián tiếp do việc truy cập trang web hoặc khi tải dữ liệu về máy; những tổn hại gặp phải do virus, hành động phá hoại trực tiếp hay gián tiếp của hệ thống máy tính khác, đường dây điện thoại, phần cứng, phần mềm, lỗi chương trình, hoặc bất kì các lỗi nào khác; đường truyền dẫn của máy tính hoặc nối kết mạng bị chậm…</p>
                </div>

                <div>
                  <strong>2. Về nội dung trên trang web:</strong>
                  <p>Tất cả những thông tin ở đây được cung cấp cho bạn một cách trung thực như bản thân sự việc. H3DC và các bên liên quan không bảo đảm, hay có bất kỳ tuyên bố nào liên quan đến tính chính xác, tin cậy của việc sử dụng hay kết quả của việc sử dụng nội dung trên trang web này. Nột dung trên website được cung cấp vì lợi ích của cộng đồng và có tính phi thương mại. Các cá nhân và tổ chức không được phếp sử dụng nội dung trên website này với mục đích thương mại mà không có sự ưng thuận của H3DC bằng văn bản. Mặc dù H3DC luôn cố gắng cập nhật thường xuyên các nội dung tại trang web, nhưng chúng tôi không bảo đảm rằng các thông tin đó là mới nhất, chính xác hay đầy đủ. Tất cả các nội dung website có thể được thay đổi bất kỳ lúc nào.</p>
                </div>
                <div>
                  <strong>3. Về bản quyền:</strong>
                  <p>H3DC là chủ bản quyền của trang web này. Việc chỉnh sửa trang, nội dung, và sắp xếp thuộc về thẩm quyền của H3DC. Sự chỉnh sửa, thay đổi, phân phối hoặc tái sử dụng những nội dung trong trang này vì bất kì mục đích nào khác được xem như vi phạm quyền lợi hợp pháp của H3DC.</p>
                </div>
                <div>
                  <strong>4. Về việc sử dụng thông tin:</strong>
                  <p>Chúng tôi sẽ không sử dụng thông tin cá nhân của bạn trên website này nếu không được phép. Nếu bạn đồng ý cung cấp thông tin cá nhân, bạn sẽ được bảo vệ. Thông tin của bạn sẽ được sử dụng với mục đích, liên lạc với bạn để thông báo các thông tin cập nhật của H3DC như lịch chiếu phim, khuyến mại qua email hoặc bưu điện. Thông tin cá nhân của bạn sẽ không được gửi cho bất kỳ ai sử dụng ngoài trang web H3DC, ngoại trừ những mở rộng cần thiết để bạn có thể tham gia vào trang web (những nhà cung cấp dịch vụ, đối tác, các công ty quảng cáo) và yêu cầu cung cấp bởi luật pháp. Nếu chúng tôi chia sẻ thông tin cá nhân của bạn cho các nhà cung cấp dịch vụ, công ty quảng cáo, các công ty đối tác liên quan, thì chúng tôi cũng yêu cầu họ bảo vệ thông tin cá nhân của bạn như cách chúng tôi thực hiện.</p>
                </div>
                <div>
                  <strong>5. Vể việc tải dữ liệu:</strong>
                  <p>Nếu bạn tải về máy những phần mềm từ trang này, thì phần mềm và các dữ liệu tải sẽ thuộc bản quyền của H3DC và cho phép bạn sử dụng. Bạn không được sở hữu những phầm mềm đã tải và H3DC không nhượng quyền cho bạn. Bạn cũng không được phép bán, phân phối lại, hay bẻ khóa phần mềm…</p>
                </div>
                <div>
                  <strong>6. Thay đổi nội dung:</strong>
                  <p>H3DC giữ quyền thay đổi, chỉnh sửa và loại bỏ những thông tin hợp pháp vào bất kỳ thời điểm nào vì bất kỳ lý do nào.</p>
                </div>
                <div>
                  <strong>7. Liên kết với các trang khác:</strong>
                  <p>Mặc dù trang web này có thể được liên kết với những trang khác, H3DC không trực tiếp hoặc gián tiếp tán thành, tổ chức, tài trợ, đứng sau hoặc sát nhập với những trang đó, trừ phi điều này được nêu ra rõ ràng. Khi truy cập vào trang web bạn phải hiểu và chấp nhận rằng H3DC không thể kiểm soát tất cả những trang liên kết với trang H3DC và cũng không chịu trách nhiệm cho nội dung của những trang liên kết.</p>
                </div>
                <div>
                  <strong>8. Đưa thông tin lên trang web:</strong>
                  <p>Bạn không được đưa lên, hoặc chuyển tải lên trang web tất cả những hình ảnh, từ ngữ khiêu dâm, thô tục, xúc phạm, phỉ báng, bôi nhọ, đe dọa, những thông tin không hợp pháp hoặc những thông tin có thể đưa đến việc vi phạm pháp luật, trách nhiệm pháp lý. H3DC và tất cả các bên có liên quan đến việc xây dựng và quản lý trang web không chịu trách nhiệm hoặc có nghĩa vụ pháp lý đối với những phát sinh từ nội dung do bạn tải lên trang web.</p>
                </div>
                <div>
                  <strong>9. Luật áp dụng:</strong>
                  <p>Mọi hoạt động phát sinh từ trang web có thể sẽ được phân tích và đánh giá theo luật pháp Việt Nam và toà án Tp. Hồ Chí Minh. Và bạn phải đồng ý tuân theo các điều khoản riêng của các toà án này.</p>
                </div>
























              </div>
              <div className='my-5 d-flex justify-content-start align-items-center'>
                <input className='m-2' id='check_terms_condition' type='checkbox' checked={isCheckedCondition} onChange={handleCheckCondition} />
                <label className='fw-bold fs-6' htmlFor='check_terms_condition'>Tôi đồng ý với điều khoản trên và bảo đảm mua vé xem phim này theo đúng độ tuổi quy định</label>
              </div>
              <div className='fs-5 my-2' style={{ textAlign: 'start' }}>Chọn hình thức thanh toán</div>
              <div className='my-1 d-flex justify-content-start align-items-center'>
                <input onChange={handlePaymentMethod} className='radio-button mx-3' type="checkbox" id="option1" checked={paymentMethod}></input>
                <label htmlFor="option1">Paypal</label>
                <br />
              </div>
              <img onClick={handlePaymentMethod} className='mx-5' src="https://th.bing.com/th/id/OIP.wBKSzdf1HTUgx1Ax_EecKwHaHa?rs=1&pid=ImgDetMain" style={{ width: '3rem' }} alt="logo Paypal" />

            </div>
          </div>

          {message.isShow && (
            <div className='d-flex flex-column justify-content-center align-items-center my-4'>
              <div>{message.text}</div>
              <LoadingSpinner />
            </div>
          )}
          <div className='container'>
            <button onClick={handleBack} className='btn btn-lg text-light fw-bold m-2' style={{ background: 'orange' }}>Quay lại</button>
            <button onClick={handleCheckoutInit} className='btn btn-lg text-light fw-bold m-2' style={{ background: 'orange' }} disabled={seats.length <= 0 || disabledBtnPayment}>Thanh toán</button>
          </div>
        </div>
      ) : (
        <div className='d-flex justify-content-center my-5'>
          <LoadingSpinner />
        </div>
      )
      }

      <LogOut></LogOut>
      <ChangePW {...params}></ChangePW>
      <Footer></Footer>
    </>
  )
}

export default CustomerDashBoard