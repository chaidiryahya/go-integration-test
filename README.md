# go-integration-test
Integration test using golang. The main objective is to compare ACTUAL response of API/GQL with the EXPECTED one. The app will tell you if all test is succeed or there's an error from the test. It's usefull to get early failure detection of your app before the end user.

### Prequisites
- golang 1.18

### How to run
- clone repository `git clone https://github.com/chaidiryahya/go-integration-test`
- enter to directory then run app using one of the commands :
    - `make run` to run API and GQL test case(s)
    - `make run-api-test` to run API test only
    - `make run-gql-test` to run GQL test only
- You'll see the integration test result like below
```
2023/01/21 20:06:51 [Integration Test] Starting
2023/01/21 20:06:51 [Integration Test] Testing ALL Module(s)
2023/01/21 20:06:51 [Integration Test] Found 3 API test case file(s)
2023/01/21 20:06:52 [API] Get avatar info - Failed Case => Failed. Reason : Expected Response not equal with Current Response
2023/01/21 20:06:53 [API] Get avatar info - Success Get Info => Success
2023/01/21 20:06:53 [API] Get hot coffee list - Success Get Hot coffee list => Success
2023/01/21 20:06:53 [Integration Test] Found 1 GQL test case file(s)
2023/01/21 20:06:54 [GQL] Query Country - Success Get Country Info => Success
2023/01/21 20:06:54 [Integration Test] End
2023/01/21 20:06:54 [Integration Test] Total Process Time : 2.218164074 seconds
```

All tests are expected to be successful, but there's one failed test case from the sample to show you if one of the case has been failed.
Sample failed : `[API] Get avatar info - Failed Case => Failed. Reason : Expected Response not equal with Current Response`

You can add another API test case in `test_case/api` directory and GQL test case in `test_case/gql`.

### Next update
- Run in http mode, so the app can do integration testing on demand
- Run in cron mode, the app will run periodically to run the testcase