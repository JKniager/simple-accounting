package account

type AccountId = int

type AccountType uint32

const (
	AccountTypeSavings AccountType = iota
	AccountTypeExpense
)

var (
	accountTypeNameMap = map[AccountType]string{
		AccountTypeSavings: "AccountTypeSavings",
		AccountTypeExpense: "AccountTypeExpense",
	}
)

func (t AccountType) Name() string {
	return accountTypeNameMap[t]
}

func (t AccountType) Val() uint32 {
	return uint32(t)
}

type Account interface {
	Credit(float64) error
	Debit(float64) error
	Id() AccountId
	Name() string
	Balance() float64
}

type SavingsAcct struct {
	id   AccountId
	name string
	blnc float64
}

func NewSavingsAcct(id AccountId, name string, initBalance float64) *SavingsAcct {
	return &SavingsAcct{
		id:   id,
		name: name,
		blnc: initBalance,
	}
}

func (a *SavingsAcct) Id() AccountId {
	return a.id
}

func (a *SavingsAcct) Name() string {
	return a.name
}

func (a *SavingsAcct) Balance() float64 {
	return a.blnc
}

func (a *SavingsAcct) Credit(amt float64) error {
	a.blnc -= amt
	return nil
}

func (a *SavingsAcct) Debit(amt float64) error {
	a.blnc += amt
	return nil
}

type ExpenseAcct struct {
	id   AccountId
	name string
	blnc float64
}

func NewExpenseAcct(id AccountId, name string, initBalance float64) *ExpenseAcct {
	return &ExpenseAcct{
		id:   id,
		name: name,
		blnc: initBalance,
	}
}

func (a *ExpenseAcct) Id() AccountId {
	return a.id
}

func (a *ExpenseAcct) Name() string {
	return a.name
}

func (a *ExpenseAcct) Balance() float64 {
	return a.blnc
}

func (a *ExpenseAcct) Credit(amt float64) error {
	a.blnc -= amt
	return nil
}

func (a *ExpenseAcct) Debit(amt float64) error {
	a.blnc += amt
	return nil
}
