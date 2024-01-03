package transaction

type TxInput struct {
	PreviousTx []byte
	Index      int
	ScriptSig  []byte
}
