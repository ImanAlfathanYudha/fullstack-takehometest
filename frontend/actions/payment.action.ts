import { Payment } from "@/entities/payment.entity";
import api from "@/services/api";

export interface PaymentListResponse {
  payments: Payment[];
}

export const getPayments = async (status?: string): Promise<Payment[]> => {
  const response = await api.get<PaymentListResponse>('/dashboard/v1/payments', {
    params: {
      status: status === '' ? undefined : status,
    },
  });
  return response.data.payments;
};
