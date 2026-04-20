package alerts

type SMSSender interface {
	SendSMS() error
}

type User string

func (u User) SendSMS(message string) error {
	err := ExecSend(string(u), message)
	if err != nil {
		return err
	}
	return nil
}
