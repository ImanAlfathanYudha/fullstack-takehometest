import Head from 'next/head';
import ProtectedRoute from '@/components/protectedRoute';
import { DashboardModule } from '@/modules/dashboard';

export default function DashboardPage() {
  return (
    <ProtectedRoute>
      <Head>
        <title>Dashboard | Durianpay</title>
      </Head>
      <DashboardModule />
    </ProtectedRoute>
  );
}
