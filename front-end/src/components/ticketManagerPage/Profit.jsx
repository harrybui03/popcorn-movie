import React, { useState, useEffect } from 'react';
import { PieChart, Cell, Pie, ComposedChart, Line, Bar, XAxis, YAxis, CartesianGrid, Tooltip, Legend } from 'recharts';
import useAuth from '../../hooks/useAuth';
import { subMonths, startOfMonth, endOfMonth, format } from 'date-fns';
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faDollar } from "@fortawesome/free-solid-svg-icons";
import { da } from 'date-fns/locale';
import { useGetRevenue } from './hook/useQuery';

const CombinedColumnLineChart = () => {
  const dataRes = useGetRevenue(2024)
  const data = dataRes?.data??{}
  
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
      </div>
    </>
  );
};

export default CombinedColumnLineChart;
