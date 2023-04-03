package mail

type Mail interface {
	Verify(email, code string) error
}
