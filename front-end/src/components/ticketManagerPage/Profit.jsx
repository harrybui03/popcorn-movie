import React, { useState, useEffect } from 'react';
import { PieChart, Cell, Pie, ComposedChart, Line, Bar, XAxis, YAxis, CartesianGrid, Tooltip, Legend } from 'recharts';
import useAuth from '../../hooks/useAuth';
import { subMonths, startOfMonth, endOfMonth, format } from 'date-fns';
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faDollar } from "@fortawesome/free-solid-svg-icons";
import { da } from 'date-fns/locale';

const CombinedColumnLineChart = () => {
  const auth = useAuth();

  async function getData(last12months) {
    const dataResponse = { total: 0, arr: [] };
    for (let i of last12months) {
      try {
        // Địa chỉ API và tham số
        const apiUrl = `http://localhost:8080/transactions/monthly-revenue?year=${i.year}&month=${i.month}`;

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
          dataResponse.arr.push(result);
          dataResponse.total += result.total;
        }
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    }
    setData(dataResponse);
  }

  function getLast12Months() {
    const currentDate = new Date();
    const last12Months = [];

    for (let i = 0; i < 12; i++) {
      const startDate = startOfMonth(subMonths(currentDate, i));
      const endDate = endOfMonth(subMonths(currentDate, i));

      last12Months.push({
        year: format(startDate, 'yyyy'),
        month: format(startDate, 'MM'),
        startDate: format(startDate, 'yyyy-MM-dd'),
        endDate: format(endDate, 'yyyy-MM-dd')
      });
    }

    return last12Months;
  }

  const last12months = getLast12Months().reverse();
  const [data, setData] = useState({ total: 0, arr: [] });
  let dataTemp = {
    total: 0,
    arr: []
  }


  useEffect(() => {
    getData(last12months);
  }, []);

  const [month, setMonth] = useState(null);

  useEffect(() => {
    if (data.arr.length > 0) {
      setMonth(data.arr[0].yearMonth);
      setDataMonth(data.arr[0]);
    }
  }, [data]);

  const [dataMonth, setDataMonth] = useState(null);

  useEffect(() => {
    if (data.arr.length > 0) {
      setDataMonth(data.arr.filter(e => e.yearMonth === month)[0]);
    }
  }, [month]);

  const colors = [
    '#4e79a7', '#f28e2b', '#e15759', '#76b7b2',
    '#59a14f', '#edc948', '#b07aa1', '#ff9da7',
    '#9c755f', '#bab0ac', '#d95f02', '#2b83ba'
  ];



  const calculatePercent = (value, total) => ((value / total) * 100).toFixed(2) + '%';
  return (
    <>
      <div className="chart-container p-4" style={{ backgroundColor: 'beige' }}>
        <div className='fs-4 m-2 p-5 shadow border border-2 rounded-5' style={{ backgroundColor: 'rgb(199, 235, 248)' }}>
          <h4><FontAwesomeIcon icon={faDollar} /> Tổng doanh thu 12 tháng gần nhất</h4>
          <div className='fw-bold text-primary'>{data.total} <u>đ</u></div>
        </div>

        <div className='m-2 p-4 shadow border border-2 rounded-5 d-flex flex-column justify-content-center align-items-center' style={{ background: 'linear-gradient(356deg, rgba(226,164,114,1) 0%, rgba(242,215,232,1) 59%)' }}>
          <div className='fs-5 fw-bold'>CHI TIẾT DOANH THU 12 THÁNG GẦN NHẤT</div>
          <div>
            <ComposedChart width={950} height={400} data={data.arr} margin={{ top: 5, right: 30, left: 20, bottom: 5 }}>
              <XAxis dataKey="yearMonth" stroke="black" />
              <YAxis stroke="black" />
              <CartesianGrid strokeDasharray="3 3" stroke="black" />
              <Tooltip stroke="black" />
              <Legend stroke="black" />
              <Line type="monotone" dataKey="total" stroke="#d95f02" strokeWidth={3} />
            </ComposedChart>
          </div>
        </div>

        <div className='m-2 p-4 bg-white shadow border border-2 rounded-5 d-flex flex-column justify-content-center align-items-center'>
          <div className='fs-5 fw-bold my-2'>PHẦN TRĂM DOANH THU THEO 12 THÁNG GẦN NHẤT</div>
          <div>
            <PieChart width={900} height={400}>
              <Pie
                dataKey="total"
                data={data.arr}
                cx={400}
                cy={200}
                outerRadius={170}
                label={({ yearMonth, percent }) => `${yearMonth} : ${(percent * 100).toFixed(2)}%`}
              >
                {data.arr.map((entry, index) => (
                  <Cell key={`cell-${index}`} fill={colors[(index) % colors.length]} />
                ))}
              </Pie>
              <Tooltip formatter={(value, name, props) => [`${value} (${calculatePercent(value, props.payload.total)})`, name]} />
            </PieChart>
          </div>
        </div>

        <div className='m-2 shadow border border-2 rounded-5 d-flex flex-column justify-content-center align-items-center' style={{ position: 'relative', background: 'linear-gradient(356deg, rgba(226,164,114,1) 0%, rgba(242,215,232,1) 59%)', paddingTop: '3rem' }}>
          <div className='d-flex' style={{ position: 'absolute', top: '10px', left: '10px', zIndex: '1' }}>
            <label style={{width: '10rem'}} className='fw-bold'>Chọn tháng</label>
            <select className="form-control" onChange={(e) => setMonth(e.target.value)}>
              {data.arr.length > 0 && data.arr.map((item) => {
                return (
                  <option key={item.yearMonth} value={`${item.yearMonth}`}>{`${item.yearMonth}`}</option>
                )
              })}
            </select>
          </div>
          <div className='fs-5 fw-bold'>{`CHI TIẾT DOANH THU ${month ? month : ''}`}</div>
          <div>
            {dataMonth && (
              <ComposedChart width={950} height={400} data={dataMonth.dailyRevenues} margin={{ top: 5, right: 30, left: 20, bottom: 5 }}>
                <XAxis dataKey="date" stroke="black" />
                <YAxis stroke="black" />
                <CartesianGrid strokeDasharray="3 3" stroke="black" />
                <Tooltip stroke="black" />
                <Legend stroke="black" />
                <Bar dataKey="totalRevenue" fill="#9c755f" barSize={20} />
              </ComposedChart>
            )}
          </div>
        </div>
      </div>
    </>
  );
};

export default CombinedColumnLineChart;
