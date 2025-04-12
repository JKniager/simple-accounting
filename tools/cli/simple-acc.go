package main

import (
	"simple_accounting/internal/accounting/ledger"
	"simple_accounting/tools/cli/cmd"
)

type mainStruct struct {
	ldgr ledger.Ledger
}

func main() {
	ms := mainStruct{
		ldgr: *ledger.NewLedger(),
	}
	cmd.Execute()
}
