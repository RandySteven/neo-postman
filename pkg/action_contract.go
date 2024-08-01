package pkg

type (
	GraphqlActionContract interface {
	}

	YAMLActionContract interface {
		ReadYAML() ([]byte, error)
	}

	RedisActionContract interface {
		ConnectToRedis() error
	}

	PostgresActionContract interface{}

	JiraActionContract interface {
		CreateJiraTicket()
	}
)
