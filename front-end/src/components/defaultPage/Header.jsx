import React from 'react'
import SearchButton from './SearchButton'
import LoginButton from './LoginButton'

function HomepageHeader() {
    return (
        <nav className="navbar navbar-expand-lg navbar-light " style={{backgroundColor: '#2b1a1a'}}>
            <div className="container-fluid d-flex justify-content-center">
                <a className="navbar-brand" href="/"><img src="/Logo.png" style={{height: '3rem'}} alt='logo'></img></a>
                <button className="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNavDropdown" aria-controls="navbarNavDropdown" aria-expanded="false" aria-label="Toggle navigation">
                    <span className="navbar-toggler-icon"></span>
                </button>
                <div className="collapse navbar-collapse" id="navbarNavDropdown">
                    <ul className="navbar-nav nav-bar">
                        <li className="nav-item">
                            <a className="nav-link text-light active" aria-current="page" href="/">Trang chủ</a>
                        </li>                        
                    </ul>
                </div>
                <SearchButton />
                <LoginButton />
            </div>
        </nav>
    )
}

export default HomepageHeader