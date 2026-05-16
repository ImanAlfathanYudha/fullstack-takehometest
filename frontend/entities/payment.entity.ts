export interface Payment {
    id: string;
    merchant: string;
    status: 'completed' | 'processing' | 'failed' | string;
    amount: string;
    created_at: string;
}

export interface SummaryData {
  total: number;
  success: number;
  failed: number;
  processing: number;
}