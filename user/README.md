## Commands

mockgen --source=user.go --destination=user_mock_test.go --package=user_test

mockgen --source=user_handler.go --destination=user_handler_mock_test.go --package=user_test

go test -tags=integration