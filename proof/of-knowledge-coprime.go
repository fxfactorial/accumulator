package proof

type Membership interface {
}

type CoprimeProver interface {
	Prove()
}

type CoprimeVerifier interface {
	Verify()
}

func NewCoprime() {

}
