start:
	go run ./...

test:
	go test -v -cover ./...

generate_mock:
	@ mockery --dir=usecase --name=TodoUsecaseInterface --filename=todo_mock.go --output=usecase/mocks --outpkg=todousecasemock
	@ mockery --dir=repository/mysql --name=TodoRepositoryInterface --filename=todo_mock.go --output=repository/mysql/mocks --outpkg=todorepositorymock
