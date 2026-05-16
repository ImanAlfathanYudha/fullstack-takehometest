import { Payment } from '@/entities/payment.entity';
import { formatCurrency, formatDate } from '@/utils/format';
import styles from '../../style.module.css';

interface DashboardTableProps {
  loading: boolean;
  payments: Payment[];
}

export const DashboardTable = ({ loading, payments }: DashboardTableProps) => {
  const getBadgeClass = (status: string) => {
    if (status === 'completed') return styles.badgeCompleted;
    if (status === 'processing') return styles.badgeProcessing;
    if (status === 'failed') return styles.badgeFailed;
    return '';
  };

  return (
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
  );
};
