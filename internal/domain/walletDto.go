package domain

import (
	"github.com/joey1123455/beds-api/internal/common"
	"github.com/joey1123455/beds-api/internal/validator"
)

type CreateWalletRequest struct {
	UserID     string `json:"user_id"`
	CurrencyID string `json:"currency_id"`
}

func (r *CreateWalletRequest) Validate(v *validator.Validator) bool {
	v.Check(!common.IsEmptyOrSpaces(r.UserID), "user_id", "user id must be provided")
	v.Check(!common.IsEmptyOrSpaces(r.CurrencyID), "currency_id", "currencyid must be provided")
	return v.Valid()
}

type DepositRequest struct {
	WalletID    string `json:"wallet_id"`
	Amount      string `json:"amount"`
	Description string `json:"description"`
}

func (r *DepositRequest) Validate(v *validator.Validator) bool {
	v.Check(!common.IsEmptyOrSpaces(r.WalletID), "wallet_id", "wallet id must be provided")
	v.Check(!common.IsEmptyOrSpaces(r.Amount), "amount", "amount must be provided")
	return v.Valid()
}

type WithdrawRequest struct {
	WalletID    string `json:"wallet_id"`
	Amount      string `json:"amount"`
	Description string `json:"description"`
	AccountNo   string `json:"account_no"`
	BankCode    string `json:"bank_code"`
}

func (r *WithdrawRequest) Validate(v *validator.Validator) bool {
	v.Check(!common.IsEmptyOrSpaces(r.WalletID), "wallet_id", "wallet id must be provided")
	v.Check(!common.IsEmptyOrSpaces(r.Amount), "amount", "amount must be provided")
	v.Check(!common.IsEmptyOrSpaces(r.AccountNo), "account_no", "account no must be provided")
	v.Check(!common.IsEmptyOrSpaces(r.BankCode), "bank_code", "bank code must be provided")
	return v.Valid()
}
