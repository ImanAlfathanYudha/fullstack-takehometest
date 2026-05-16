import { SummaryData } from '@/entities/payment.entity';
import styles from '../../style.module.css';

interface SummaryGridProps {
    summary: SummaryData;
}
export const SummarryGridComponent = ({ summary }: SummaryGridProps) => {
    return (
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
    )
}