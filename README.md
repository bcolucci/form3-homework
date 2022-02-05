
# Form3 Account API homework

Written by **Brice Colucci**.

I did not have much time so I keep things simple. The only packages I used is one to generate UUID and testify.

## Tests

When you start the containers, you'll see **unit** and **integration** tests:

``` 
% docker-compose up
[...]
form3-homework-accountapi_client_unit-1         | === RUN   TestAccountCreateSuccess
form3-homework-accountapi_client_unit-1         | --- PASS: TestAccountCreateSuccess (0.00s)
form3-homework-accountapi_client_unit-1         | === RUN   TestAccountCreateReqErr
form3-homework-accountapi_client_unit-1         | --- PASS: TestAccountCreateReqErr (0.00s)
form3-homework-accountapi_client_unit-1         | === RUN   TestAccountCreateInvalidStatus
form3-homework-accountapi_client_unit-1         | --- PASS: TestAccountCreateInvalidStatus (0.00s)
form3-homework-accountapi_client_unit-1         | === RUN   TestAccountFetchSuccess
form3-homework-accountapi_client_unit-1         | --- PASS: TestAccountFetchSuccess (0.00s)
form3-homework-accountapi_client_unit-1         | === RUN   TestAccountFetchReqErr
form3-homework-accountapi_client_unit-1         | --- PASS: TestAccountFetchReqErr (0.00s)
form3-homework-accountapi_client_unit-1         | === RUN   TestAccountFetchInvalidStatus
form3-homework-accountapi_client_unit-1         | --- PASS: TestAccountFetchInvalidStatus (0.00s)
form3-homework-accountapi_client_unit-1         | === RUN   TestAccountDeleteSuccess
form3-homework-accountapi_client_unit-1         | --- PASS: TestAccountDeleteSuccess (0.00s)
form3-homework-accountapi_client_unit-1         | === RUN   TestAccountDeleteReqErr
form3-homework-accountapi_client_unit-1         | --- PASS: TestAccountDeleteReqErr (0.00s)
form3-homework-accountapi_client_unit-1         | === RUN   TestAccountDeleteNotFound
form3-homework-accountapi_client_unit-1         | --- PASS: TestAccountDeleteNotFound (0.00s)
form3-homework-accountapi_client_unit-1         | === RUN   TestAccountDeleteConflict
form3-homework-accountapi_client_unit-1         | --- PASS: TestAccountDeleteConflict (0.00s)
form3-homework-accountapi_client_unit-1         | PASS
form3-homework-accountapi_client_unit-1         | ok    github.com/bcolucci/form3-homework/account      (cached)
[...]
form3-homework-accountapi_client_integration-1  | === RUN   TestCreate
form3-homework-accountapi_client_integration-1  |     account_api_test.go:30: creating account ...
form3-homework-accountapi_client_integration-1  |     account_api_test.go:35: account 3c17e1d2-23c2-42b5-b6bf-e2dd222d9935 created
form3-homework-accountapi_client_integration-1  |     account_api_test.go:41: fetching account 3c17e1d2-23c2-42b5-b6bf-e2dd222d9935 ...
form3-homework-accountapi_client_integration-1  |     account_api_test.go:52: deleting account 3c17e1d2-23c2-42b5-b6bf-e2dd222d9935 ...
form3-homework-accountapi_client_integration-1  | --- PASS: TestCreate (0.02s)
form3-homework-accountapi_client_integration-1  | PASS
form3-homework-accountapi_client_integration-1  | ok    github.com/bcolucci/form3-homework/integration_test     (cached)
```
