import { startTransition, useState, useTransition } from 'react';
import { useRouter } from 'next/router';
import styles from './style.module.css';
import { login } from '@/services/authService';

export const LoginModule = () => {
    const router = useRouter();
    const [error, setError] = useState('');
    const [isPending, startTransition] = useTransition();

    const handleLoginAction = (formData: FormData) => {
        const email = formData.get('email') as string;
        const password = formData.get('password') as string;
        setError('');
        startTransition(async () => {
            try {
                await login(email, password);
                router.push('/dashboard');
            } catch (err: any) {
                setError(err.message || 'Login failed');
            }
        });
    };

    return (
        <div className={styles.container}>
            <form action={handleLoginAction} className={styles.card}>
                <h1 className={styles.title}>Welcome Back</h1>

                {error && <div className={styles.errorMessage}>{error}</div>}
                <div className={styles.inputGroup}>
                    <label className={styles.label}>Email</label>
                    <input
                        name="email"
                        className={styles.input}
                        type="email"
                        required
                    />
                </div>
                <div className={styles.inputGroup}>
                    <label className={styles.label}>Password</label>
                    <input
                        name="password"
                        className={styles.input}
                        type="password"
                        required
                    />
                </div>
                <button className={styles.button} type="submit" disabled={isPending}>
                    {isPending ? 'Logging in...' : 'Login'}
                </button>
            </form>
        </div>
    );

};
