package domain

const (
	awaiting = "awaiting"
	done     = "done"
)

func Status() map[string]string {
	return map[string]string{
		awaiting: "awaiting",
		done:     "done",
	}
}
