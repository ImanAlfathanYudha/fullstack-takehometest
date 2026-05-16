import { useEffect, useMemo, useState } from 'react';
import { getPayments } from '@/actions/payment.action';
import { getRole, logout } from '@/services/authService';
import styles from './dashboard.module.css';
import { Payment } from '@/entities/payment.entity';
import { formatCurrency, formatDate } from '@/utils/format';

export const DashboardModule = () => {
    const [payments, setPayments] = useState<Payment[]>([]);
    const [loading, setLoading] = useState(true);
    const [filterStatus, setFilterStatus] = useState('');
    const role = getRole();

    const summary = useMemo(() => {
        return {
            total: payments.length,
            success: payments.filter(p => p.status === 'completed').length,
            failed: payments.filter(p => p.status === 'failed').length,
            processing: payments.filter(p => p.status === 'processing').length,
        };
    }, [payments]);

    useEffect(() => {
        const fetchData = async () => {
            setLoading(true); // Show loading every time we filter
            try {
                // Pass the current filter to our action
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

    const getBadgeClass = (status: string) => {
        if (status === 'completed') return styles.badgeCompleted;
        if (status === 'processing') return styles.badgeProcessing;
        if (status === 'failed') return styles.badgeFailed;
        return '';
    };

    return (
        <div className={styles.container}>
            <header className={styles.header}>
                <div>
                    <h1>Payments Dashboard</h1>
                    <span className={styles.roleBadge}>Role: {role}</span>
                </div>
                <button onClick={logout} className={styles.logoutBtn}>Logout</button>
            </header>

            {/* Summary Cards */}
            <div className={styles.summaryGrid}>
                <div className={styles.card}>
                    <div className={styles.cardTitle}>Total Payments</div>
                    <div className={styles.cardValue}>{summary.total}</div>
                </div>
                <div className={`${styles.card} ${styles.cardSuccess}`}>
                    <div className={styles.cardTitle}>Successful</div>
                    <div className={styles.cardValue}>{summary.success}</div>
                </div>
                <div className={`${styles.card} ${styles.cardProcessing}`}>
                    <div className={styles.cardTitle}>Processing</div>
                    <div className={styles.cardValue}>{summary.processing}</div>
                </div>
                <div className={`${styles.card} ${styles.cardFailed}`}>
                    <div className={styles.cardTitle}>Failed</div>
                    <div className={styles.cardValue}>{summary.failed}</div>
                </div>
            </div>

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

            {/* Main Table */}
            <div className={styles.tableWrapper}>
                <table className={styles.table}>
                    <thead>
                        <tr>
                            <th>ID</th>
                            <th>Merchant</th>
                            <th style={{ textAlign: 'right' }}>Amount</th>
                            <th style={{ textAlign: 'center' }}>Status</th>
                            <th>Date</th>
                        </tr>
                    </thead>
                    <tbody>
                        {!loading && payments.map((p) => (
                            <tr key={p.id}>
                                <td style={{ color: '#9ca3af' }}>#{p.id}</td>
                                <td style={{ fontWeight: 600 }}>{p.merchant}</td>
                                <td style={{ textAlign: 'right', fontWeight: 600 }}>
                                    {formatCurrency(p.amount)}
                                </td>
                                <td style={{ textAlign: 'center' }}>
                                    <span className={`${styles.badge} ${getBadgeClass(p.status)}`}>
                                        {p.status}
                                    </span>
                                </td>
                                <td>{formatDate(p.created_at)}</td>
                            </tr>
                        ))}
                    </tbody>
                </table>

                {loading && <div className={styles.emptyState}>Updating list...</div>}
                {!loading && payments.length === 0 && (
                    <div className={styles.emptyState}>No payments found for this status.</div>
                )}
            </div>
        </div>
    );
};