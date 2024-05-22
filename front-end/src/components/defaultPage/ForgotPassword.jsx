import React from 'react';
import { useForm } from 'react-hook-form';
import Footer from './Footer';
import HomepageHeader from './Header';
import { useForgotPassword } from './hook/useMutation';

const ForgotPassword = () => {
  const { register, handleSubmit, formState: { errors } } = useForm();
  const { isSuccess, onForgot } = useForgotPassword();

  const onSubmit = (data) => {
    onForgot(data.email);
  };

  return (
    <>
      <HomepageHeader />
      <div style={styles.container}>
        <h2>Forgot Password</h2>
        <p style={styles.infoText}>
          Quên mật khẩu? Vui lòng nhập tên đăng nhập hoặc địa chỉ email. Bạn sẽ nhận được một liên kết tạo mật khẩu mới qua email.
        </p>
        <form onSubmit={handleSubmit(onSubmit)} style={styles.form}>
          <label htmlFor="email" style={styles.label}>
            Enter your email address:
          </label>
          <input
            type="email"
            id="email"
            name="email"
            {...register('email', { required: 'Email is required' })}
            style={styles.input}
          />
          {errors.email && <p style={styles.error}>{errors.email.message}</p>}
          <button type="submit" style={styles.button}>
            Send Reset Link
          </button>
        </form>
        {isSuccess && (
          <p style={styles.message}>
            Một thư email khôi phục mật khẩu đã được gửi cho địa chỉ email tài khoản của bạn.
          </p>
        )}
      </div>
      <Footer />
    </>
  );
};

const styles = {
  container: {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    justifyContent: 'center',
    height: '100vh',
    backgroundColor: '#f0f0f0',
    padding: '20px',
  },
  infoText: {
    maxWidth: '400px',
    textAlign: 'center',
    marginBottom: '20px',
    fontSize: '14px',
    color: '#555',
  },
  form: {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    width: '100%',
    maxWidth: '400px',
    backgroundColor: '#fff',
    padding: '20px',
    borderRadius: '8px',
    boxShadow: '0 0 10px rgba(0, 0, 0, 0.1)',
  },
  label: {
    marginBottom: '10px',
    fontSize: '16px',
    fontWeight: 'bold',
  },
  input: {
    width: '100%',
    padding: '10px',
    marginBottom: '20px',
    borderRadius: '4px',
    border: '1px solid #ccc',
    fontSize: '16px',
  },
  button: {
    width: '100%',
    padding: '10px',
    backgroundColor: '#007bff',
    color: '#fff',
    border: 'none',
    borderRadius: '4px',
    fontSize: '16px',
    cursor: 'pointer',
  },
  message: {
    marginTop: '20px',
    fontSize: '14px',
    color: 'green',
  },
  error: {
    color: 'red',
    fontSize: '12px',
  },
};

export default ForgotPassword;
