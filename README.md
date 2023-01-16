# go-integration-test
Integration test using golang. The main objective is to compare ACTUAL response of API with the EXPECTED one. The app will tell you if all test is succeed or there's an error from the test. It's usefull to get early failure detection of your app.

For next update, i'll try to add for graphql cases.

### Prequisites
- golang 1.18

### How to run
- clone repository `git clone https://github.com/chaidiryahya/go-integration-test`
- enter to directory then run app using `make run`
- You'll see the integration test result like below
```
[Integration Test] Starting
[Integration Test] Found 3 scenario file(s)
Get avatar info - Failed Case => Failed. Reason : Expected Response not equal with Current Response
Get avatar info - Success Get Info => Success
Get hot coffee list - Success Get Hot coffee list => Success
[Integration Test] End
```

You can add another API scenario in `scenario` directory.
