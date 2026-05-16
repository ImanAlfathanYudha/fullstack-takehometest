import { useEffect, useState, useMemo } from 'react';

import { getRole, logout } from '@/services/authService';

import { DashboardTable } from './components/DashboardTable'; // <--- New import
import styles from './style.module.css';
import { Payment } from '@/entities/payment.entity';
import { getPayments } from '@/actions/payment.action';
import { SummarryGridComponent } from './components/SummaryGrid';

export const DashboardModule = () => {
  const [payments, setPayments] = useState<Payment[]>([]);
  const [loading, setLoading] = useState(true);
  const [filterStatus, setFilterStatus] = useState('');
  const role = getRole();

  const summary = useMemo(() => ({
    total: payments.length,
    success: payments.filter(p => p.status === 'completed').length,
    failed: payments.filter(p => p.status === 'failed').length,
    processing: payments.filter(p => p.status === 'processing').length,
  }), [payments]);

  useEffect(() => {
    const fetchData = async () => {
      setLoading(true);
      try {
        const data = await getPayments(filterStatus);
        setPayments(data);
      } catch (error) {
        console.error('Failed to fetch payments', error);
      } finally {
        setLoading(false);
      }
    };
    fetchData();
  }, [filterStatus]);

  return (
    <div className={styles.container}>
      <header className={styles.header}>
        <div>
          <h1>Payments Dashboard</h1>
          <span className={styles.roleBadge}>Role: {role}</span>
        </div>
        <button onClick={logout} className={styles.logoutBtn}>Logout</button>
      </header>

      <SummarryGridComponent summary={summary} />

      <div className={styles.filterSection}>
        <label htmlFor="statusFilter">Filter Status:</label>
        <select 
          id="statusFilter"
          className={styles.select}
          value={filterStatus}
          onChange={(e) => setFilterStatus(e.target.value)}
        >
          <option value="">All Payments</option>
          <option value="completed">Completed</option>
          <option value="processing">Processing</option>
          <option value="failed">Failed</option>
        </select>
      </div>

      {/* Much cleaner Table usage */}
      <DashboardTable loading={loading} payments={payments} />
    </div>
  );
};
