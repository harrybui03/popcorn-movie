import React from 'react';
import { useForm } from 'react-hook-form';
import HomepageHeader from './Header';
import Footer from './Footer';
import { useResetPassword } from './hook/useMutation';

const ResetPassword = () => {
    const params = new URLSearchParams(window.location.search);
    const tokenParam = params.get('token');
  const { register, handleSubmit, watch, formState: { errors } } = useForm();
  const {onReset} = useResetPassword()
  const onSubmit = data => {
    onReset({newPassword:data.newPassword , confirmNewPassword:data.confirmPassword , token:tokenParam})
  };

  // Watch the password input to ensure passwords match
  const newPassword = watch('newPassword', '');

  return (
    <>
      <HomepageHeader />
      <form onSubmit={handleSubmit(onSubmit)} id="reset-password-form" style={formStyle}>
        <h5 style={h5Style}>Quên mật khẩu</h5>
        <h3 style={h3Style}>Nhập mật khẩu mới bên dưới.</h3>
        <div style={inputContainerStyle}>
          <label htmlFor="newPassword" style={labelStyle}>New Password</label>
          <input
            id="newPassword"
            type="password"
            {...register('newPassword', { required: 'New password is required' })}
            style={inputStyle}
          />
          {errors.newPassword && <p style={errorStyle}>{errors.newPassword.message}</p>}
        </div>

        <div style={inputContainerStyle}>
          <label htmlFor="confirmPassword" style={labelStyle}>Confirm Password</label>
          <input
            id="confirmPassword"
            type="password"
            {...register('confirmPassword', {
              required: 'Please confirm your new password',
              validate: value =>
                value === newPassword || 'The passwords do not match'
            })}
            style={inputStyle}
          />
          {errors.confirmPassword && <p style={errorStyle}>{errors.confirmPassword.message}</p>}
        </div>

        <button type="submit" style={buttonStyle}>Reset Password</button>
      </form>
      <Footer />
    </>
  );
};

const formStyle = {
    maxWidth: '400px',
    margin: '100px auto',
    padding: '2em', // Increased padding to 2em
    backgroundColor: '#f9f9f9',
    borderRadius: '8px',
    boxShadow: '0 0 10px rgba(0, 0, 0, 0.1)',
  };

const h5Style = {
  textAlign: 'center',
  color: '#333',
  marginBottom: '0.5em',
};

const h3Style = {
  textAlign: 'center',
  color: '#555',
  marginBottom: '1em',
};

const inputContainerStyle = {
  marginBottom: '1em',
};

const labelStyle = {
  display: 'block',
  marginBottom: '0.5em',
  color: '#333',
};

const inputStyle = {
  width: '100%',
  padding: '0.5em',
  fontSize: '1em',
  borderRadius: '4px',
  border: '1px solid #ccc',
};

const errorStyle = {
  color: 'red',
  fontSize: '0.875em',
};

const buttonStyle = {
  width: '100%',
  padding: '0.75em',
  fontSize: '1em',
  color: '#fff',
  backgroundColor: '#007BFF',
  border: 'none',
  borderRadius: '4px',
  cursor: 'pointer',
};

// Add a hover effect for the button
buttonStyle[':hover'] = {
  backgroundColor: '#0056b3',
};



export default ResetPassword;
