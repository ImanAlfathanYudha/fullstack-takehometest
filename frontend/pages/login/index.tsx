import Head from 'next/head';
import { LoginModule } from '@/modules/login';

export default function LoginPage() {
  return (
    <>
      <Head>
        <title>Login | Dashboard</title>
      </Head>
      <LoginModule />
    </>
  );
}
