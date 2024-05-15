import React, { useEffect, useState } from "react";
import axios from "axios"
import useAuth from "../../hooks/useAuth";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faFilter, faHouse, faAngleRight, faAdd } from "@fortawesome/free-solid-svg-icons";
import 'bootstrap/dist/css/bootstrap.min.css'
import { useGetAllUsers } from "./hook/useQuery";
import { useCreateUser, useLockUser } from "./hook/useMutation";
import { useNavigate } from "react-router-dom";

const ManageUser = () => {
    const auth = useAuth();
    const navigate = useNavigate()
    const [formData, setFormData] = useState({
        email: '',
        password: '',
        confirmPassword: '',
        displayName: '',
        role: 'CUSTOMER'
    });
    const [userEdit, setUserEdit] = useState({});
    const usersData = useGetAllUsers()
    const users = usersData?.data??[]
    const [search, setSearch] = useState('');
    const [selectedOption, setSelectedOption] = useState('id');
    const [selectedStatus, setSelectedStatus] = useState('active');
    const [searchPage, setSearchPage] = useState('');
    const [currentPage, setCurrentPage] = useState(1);
    const recordsPerPage = 10;
    const lastIndex = currentPage * recordsPerPage;
    const firstIndex = lastIndex - recordsPerPage;
    const records = users.slice(firstIndex, lastIndex);
    const numberOfPages = Math.ceil(users.length / recordsPerPage);
    const maxPageDisplay = 5;

    const onError = () =>{
        console.log('Insert fail')
    }

    const onSuccess = () => {
        window.location.href = '/admin/manage-users'
    }


    const {handleCreateUser , messageInsertUser , isPending} = useCreateUser({onSuccess , onError})

    function createUser() {
        handleCreateUser(formData)
    }

    const {handleLockUser } = useLockUser()
    function handleUpdateLockUser(user){
        const isLocked = !user.isLocked
        handleLockUser({isLocked , id: user.id})
    }




    const getPages = () => {
        const startPage = Math.max(currentPage - Math.floor(maxPageDisplay / 2), 1);
        const endPage = Math.min(startPage + maxPageDisplay - 1, numberOfPages);
        return Array.from({ length: endPage - startPage + 1 }, (_, index) => startPage + index);
    };

    const handleClickPrev = () => {
        if (currentPage > 1)
            setCurrentPage(currentPage - 1);
    }

    const handleClickNext = () => {
        if (currentPage < numberOfPages)
            setCurrentPage(currentPage + 1);
    }

    const changeCurrentPage = (page) => {
        setCurrentPage(page)
    }

    const handleSearchPage = () => {
        setCurrentPage(searchPage);
    }

    const handleOptionChange = (e) => {
        const newOption = e.target.value;
        setSelectedOption(newOption);
    };

    const handleStatusChange = (e) => {
        const newOption = e.target.value;
        setSelectedStatus(newOption);
    };

    // useEffect(() => {
    //     getUsers();
    // }, [])

    useEffect(() => {
        setCurrentPage(1);
    }, [search, users, recordsPerPage]);


    const [chainAction, setChainAction] = useState([{ text: 'Quản lý tài khoản', href: '/admin/manage-users' }]);

    const [action, setAction] = useState('manage-users');

    const handleChangeAction = (a) => {
        const newChange = chainAction;
        newChange.push({ text: a === 'insert-user' ? 'Thêm tài khoản' : 'Chỉnh sửa tài khoản', href: '#' })
        setChainAction(newChange);
        setAction(a);
    }

  
    const handleChange = (e) => {
        const { name, value } = e.target;
        setFormData((prevData) => ({
            ...prevData,
            [name]: value,
        }));
    };



    const handleBack = () => {
        setAction('manage-users');
        const newChain = chainAction;
        newChain.pop();
        setChainAction(newChain);
        setFormData({
            email: '',
            password: '',
            confirmPassword: '',
            displayName: '',
            role: 'CUSTOMER'
        });
        setMessageInsertUser({ isShow: false, text: '', success: false });
    };

    // const handleLockUser = async (user) => {
    //     try {
    //         const response = await fetch(`http://localhost:8080/users?user=${user.id}`, {
    //             method: 'PUT',
    //             headers: {
    //                 'Content-Type': 'application/json',
    //                 'Authorization': `Bearer ${auth.accessToken}`
    //             },
    //             body: JSON.stringify({ locked: !user.locked }),
    //         });

    //         if (!response.ok) {
    //             throw new Error(`HTTP error! Status: ${response.status}`);
    //         }            
    //         const data = await response.json();
    //     } catch (error) {
    //         console.error('Error during sign-in:', error);
    //     }
    // }


    return (
        <div className="p-4" style={{ backgroundColor: '#f0f0f0' }}>
            <div className='m-3 d-flex justify-content-start align-items-center'>
                <div>
                    <a className='text-black' href='/admin/home'><FontAwesomeIcon icon={faHouse} /> Trang chủ</a>
                </div>
                {chainAction.map((item, index) => (
                    <div key={index}>
                        <FontAwesomeIcon className='mx-3' icon={faAngleRight} />
                        <a className='text-black' href={item.href}>{item.text}</a>
                    </div>
                ))}
            </div>
            <hr />


            {action === 'manage-users' ? (
                <>
                    <div className="table-filter">
                        <div className="row">
                            <div className="d-flex justify-content-between px-5">
                                <div className="d-flex justify-content-start">
                                    <div className="filter-group d-flex">
                                        <label>Tìm theo</label>
                                        <select className="form-control" onChange={handleOptionChange}>
                                            <option value="id">ID</option>
                                            <option value="displayName">Username</option>
                                            <option value="email">Email</option>
                                            <option value="role">Role</option>
                                        </select>
                                    </div>
                                    <div className="filter-group d-flex">
                                        <input type="text" className="form-control" placeholder="Tìm kiếm" onChange={(e) => setSearch(e.target.value)} />
                                    </div>

                                    <div className="filter-group d-flex">
                                        <label>Status</label>
                                        <select className="form-control" onChange={handleStatusChange}>
                                            <option value="active">Active</option>
                                            <option value="inactive">Inactive</option>
                                        </select>
                                    </div>
                                </div>

                                <button onClick={() => handleChangeAction('insert-user')} className="btn btn-primary d-flex justify-content-center align-items-center"><FontAwesomeIcon className="mx-2" icon={faAdd} /> Thêm</button>
                            </div>
                        </div>
                    </div>

                    <div className="board shadow border border-2">
                        <table width="100%" className="table table-hover">
                            <thead>
                                <tr className="table table-primary">
                                    <th className="text-center">ID</th>
                                    <th className="text-center">Username</th>
                                    <th className="text-center">Email</th>
                                    <th className="text-center">Status</th>
                                    <th className="text-center">Role</th>
                                    <th className="text-center"></th>
                                </tr>
                            </thead>

                            <tbody>
                                {
                                    users && users.filter((itemStatus) => {
                                        if (selectedStatus === 'active') {
                                            return itemStatus.isLocked === false;
                                        }

                                        else if (selectedStatus === 'inactive') {

                                            return itemStatus.isLocked === true;
                                        }
                                        return itemStatus;
                                    }).filter((item) => {
                                        let lowerCaseSearch = search.toLowerCase();
                                        if (selectedOption === 'id') {
                                            lowerCaseSearch = search;
                                            return lowerCaseSearch === '' ? item : item[`${selectedOption}`].toString().match(lowerCaseSearch)
                                        }
                                        return lowerCaseSearch === '' ? item : item[`${selectedOption}`].toLowerCase().includes(lowerCaseSearch)
                                    }).slice(firstIndex, lastIndex).map((user, index) => (
                                        <tr key={index}>
                                            <th className="pb-4">
                                                {user.id}
                                            </th>
                                            <td className="px-auto">
                                                <h5 className="mt-1">
                                                    {user.displayName}
                                                </h5>
                                            </td>

                                            <td className="people-des">
                                                <p>{user.email}</p>
                                            </td>

                                            <td className="active"><p>{user.isLocked === false ? "Active" : "Inactive"}</p></td>

                                            <td className="role">
                                                <p>{user.role.replace("ROLE_", "")}</p>
                                            </td>

                                            <td className="edit">
                                                <a onClick={() => setUserEdit(user)} className="me-3" type="button" data-bs-toggle="modal" data-bs-target="#LockModal">{user.locked ? 'Unlock' : 'Lock'}</a>
                                            </td>
                                        </tr>
                                    ))}
                            </tbody>
                        </table>
                        <br />
                        <nav aria-label="Page navigation example my-5">
                            <ul className="pagination justify-content-center">
                                <li className="page-item m-1">
                                    <button className="page-link bg-light" href="#" aria-label="Previous" onClick={handleClickPrev}>
                                        <span aria-hidden="true">&laquo;</span>
                                    </button>
                                </li>
                                {
                                    getPages().map((page, index) => (
                                        <li className={`page-item m-1 ${currentPage === page ? "active" : ""}`} key={index}><button className="page-link bg-light text-primary" href="#" onClick={() => changeCurrentPage(page)}>{page}</button></li>
                                    ))
                                }

                                <li className="page-item m-1">
                                    <button className="page-link bg-light" href="#" aria-label="Next" onClick={handleClickNext}>
                                        <span aria-hidden="true">&raquo;</span>
                                    </button>
                                </li>
                                <div className="mx-5 d-flex align-items-center">
                                    <input
                                        type="text"
                                        className="bg-white text-black px-2"
                                        onInput={(e) => setSearchPage(parseInt(e.target.value))}
                                        style={{ width: '3rem', height: '2.5rem', borderRadius: '10px' }}
                                    />
                                    <button onClick={handleSearchPage} className="ms-3 btn btn-outline-primary">Go to page</button>
                                </div>

                            </ul>
                        </nav>

                        <div className="modal fade" id="LockModal" tabIndex="-1" aria-labelledby="exampleModalLabel" aria-hidden="true">
                            <div className="modal-dialog">
                                <div className="modal-content">
                                    <div className="modal-header">
                                        <h1 className="modal-title fs-5" id="exampleModalLabel">Khóa tài khoản</h1>
                                        <button type="button" className="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                                    </div>
                                    <div className="modal-body">
                                        {userEdit.locked ? 'Bạn có chắc muốn mở khóa tài khoản này?' : 'Bạn có chắc muốn khóa tài khoản này?'}
                                    </div>
                                    <div className="modal-footer">
                                        <button type="button" className="btn btn-secondary" data-bs-dismiss="modal">Hủy</button>
                                        <button type="button" className="btn btn-primary" data-bs-dismiss="modal" onClick={() => handleUpdateLockUser(userEdit)}>{userEdit.locked ? 'Mở khóa' : 'Khóa'}</button>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </>
            ) : (
                <div className="container px-5">
                    <div className="mx-auto px-5 shadow border border-black rounded-5" style={{ maxWidth: '40rem', backgroundColor: 'beige' }}>
                        <div className="group group-login-signup my-3">
                            <label htmlFor="email" className="label mx-2">Email</label>
                            <input id="email" onChange={handleChange} name="email" type="text" className="form-control bg-light py-2" placeholder="Nhập email" required />
                        </div>
                        <div className="group group-login-signup my-3">
                            <label htmlFor="user" className="label mx-2">Username</label>
                            <input id="user" onChange={handleChange} name="displayName" type="text" className="form-control bg-light py-2" placeholder="Nhập tên đăng nhập" required />
                        </div>
                        <div className="group group-login-signup my-3">
                            <label htmlFor="pass" className="label mx-2">Password</label>
                            <input id="pass" onChange={handleChange} name="password" type="password" className="form-control bg-light py-2" data-type="password" placeholder="Nhập mật khẩu" required />
                        </div>
                        <div className="group group-login-signup my-3">
                            <label htmlFor="confirmPass" className="label mx-2">Confirm Password</label>
                            <input id="confirmPass" onChange={handleChange} name="confirmPassword" type="password" className="form-control bg-light py-2" data-type="password" placeholder="Nhập lại mật khẩu" required />
                        </div>
                        <div className="group group-login-signup my-3">
                            <label htmlFor="role" className="label mx-2">Role</label>
                            <select
                                id="role"
                                className='form-control bg-light'
                                name="role"
                                onChange={handleChange} required>
                                <option value="CUSTOMER">CUSTOMER</option>
                                <option value="TICKET_MANAGER">TICKET_MANAGER</option>
                                <option value="STAFF">STAFF</option>
                            </select>
                        </div>

                        {messageInsertUser.text !== '' && messageInsertUser.isShow && (
                            <div className={messageInsertUser.success ? 'text-success' : 'text-danger'}>
                                {messageInsertUser.text}
                            </div>
                        )}

                        <div className="group group-login-signup my-4">
                            <button onClick={handleBack} className="btn btn-primary m-2">Quay về</button>
                            {!isPending && <button onClick={createUser} className="btn btn-primary m-2">Tạo tài khoản</button>}
                        </div>
                    </div>
                </div>
            )}

        </div>
    )
}

export default ManageUser