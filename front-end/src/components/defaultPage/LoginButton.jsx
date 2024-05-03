import React, { useEffect, useState, useRef } from "react";

function LoginButton() {
    useEffect(() => {
        return () => {
            $('#btn-login').on({
                'click': function() {
                    $('.login-modal').addClass('display');
                },
                'mouseover': function() {
                    $('#btn-login').removeClass('text-light');
                    $('#btn-login').addClass('text-primary');
                    $('#btn-login').css('cursor', 'pointer');
                },
                'mouseout': function() {
                    $('#btn-login').removeClass('text-primary');
                    $('#btn-login').addClass('text-light');
                    $('#btn-login').css('cursor', 'pointer');
                }
            });
            $('.login-close-btn').on('click', function() {                
                $('.login-modal').removeClass('display');
            });
        }
    }, []);

    return (
        <div className="d-flex justify-content-end w-25 px-3"> 
            <a id="btn-login" className="text-light">Đăng nhập </a>
        </div>
    )
}

export default LoginButton