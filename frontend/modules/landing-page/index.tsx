import { useRouter } from 'next/router';
import styles from './landing.module.css';

export const LandingPageModule = () => {
    const router = useRouter();

    return (
        <div className={styles.container}>
            <h1 className={styles.title}>Payment Dashboard</h1>
            <p className={styles.subtitle}>
                A payment dashboard built with Next.js and TypeScript.
            </p>
            <button
                className={styles.button}
                onClick={() => router.push('/login')}
            >
                Go to Login
            </button>
        </div>
    );
};
