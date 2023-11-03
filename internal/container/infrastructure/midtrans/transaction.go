package midtrans_payment

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

func BankTransferRequest(ammount int64, orderID string, bankName string) coreapi.ChargeReq {
	request := coreapi.ChargeReq{
		PaymentType: "bank_transfer",
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderID,
			GrossAmt: ammount,
		},
		BankTransfer: &coreapi.BankTransferDetails{
			Bank: midtrans.Bank(bankName),
		},
	}

	return request
}


