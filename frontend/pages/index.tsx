
import Head from "next/head";

export default function Home() {
  return (
    <>

      <Head>
        <title>Durianpay Dashboard</title>
        <meta name="description" content="Durianpay Take-home Test Dashboard" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <div style={{
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
        justifyContent: 'center',
        height: '100vh',
        fontFamily: 'system-ui, sans-serif'
      }}>
        <h1 style={{ fontSize: '3rem', marginBottom: '1rem' }}>Durianpay Dashboard</h1>
        <p style={{ fontSize: '1.2rem', color: '#666' }}>
          Welcome to the payment dashboard. Please login to continue.
        </p>
        <button
          onClick={() => window.location.href = '/login'}
          style={{
            marginTop: '2rem',
            padding: '0.8rem 2rem',
            fontSize: '1.1rem',
            backgroundColor: '#000',
            color: '#fff',
            border: 'none',
            borderRadius: '8px',
            cursor: 'pointer'
          }}
        >
          Go to Login
        </button>
      </div>

    </>
  );
}
