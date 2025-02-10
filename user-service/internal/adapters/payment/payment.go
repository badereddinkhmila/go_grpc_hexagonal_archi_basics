package payment

import (
	order "user-service/proto/gen/go"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type PaymentAdapter struct {
	payment order.PaymentServiceClient
}

func NewPaymentAdapter(paymentServiceUrl string) (*PaymentAdapter, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(paymentServiceUrl, opts...)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := order.NewPaymentServiceClient(conn)
	return &PaymentAdapter{payment: client}, nil
}
