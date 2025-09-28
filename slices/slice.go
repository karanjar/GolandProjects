package slices

func GetMessageWithRetries(Primary, Secondary, Tertiary string) ([3]string, [3]int) {
	// ?
	

	Message := [3]string{Primary,Secondary,Tertiary}

	costs := [3]int{20, 12, 30}

	return Message,costs

}
