package transaction

type TxInput struct {
	PreviousTx []byte
	ScriptSig  []byte
}
