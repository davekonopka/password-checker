# 07 A new test failure

## Prompt 

That helped but I'm still getting a failing test. Can you look at the test output and provide a fix?

```
--- FAIL: TestCheckPasswordStrength (0.00s)
    --- FAIL: TestCheckPasswordStrength/Test_5:_Exceeding_Maximum_Length (0.00s)
        main_test.go:42: Expected 36 steps, but got 34 steps
FAIL
exit status 1
FAIL	github.com/davekonopka/password-checker	0.290s
```

## Response 

Thanks for bringing this up. The test "Test_5: Exceeding Maximum Length" is failing because it expects the result to be 36, but the actual output is 34.

Looking at the test case:

```go
{
    name:     "Test 5: Exceeding Maximum Length",
    password: "abcABC123abcABC123abcABC123abcABC123abcABC123abcABC123",
    steps:    36,
},
```

The password length is 54 characters long, which is 34 characters over the limit (20 characters). So, the function is correctly returning 34 steps to remove the excess characters. I believe there's an error in the test case itself - the expected output should be 34, not 36.

Here's the corrected test case:

```go
{
    name:     "Test 5: Exceeding Maximum Length",
    password: "abcABC123abcABC123abcABC123abcABC123abcABC123abcABC123",
    steps:    34,
},
```

Please replace the test case in your tests and try again. Your tests should all pass with this change.
