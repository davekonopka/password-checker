# Take-Home Project

## Password Strength Checker

We want you to create a password strength checking service. The service should provide an HTTP interface. Prepare the application as you would a production ready service with any operational features you want. Spend no more than five hours working on this project. Capture your work in a Github repository with commits as you complete features.

### Functional Requirements

The application should take in a password as a string input and return a numeric value indicating how many steps would be required to make the password strong.

A password is considered strong if the below conditions are all met:

It has at least 6 characters and at most 20 characters.
It contains at least one lowercase letter, at least one uppercase letter, and at least one digit.
It does not contain three repeating characters in a row (i.e., "Baaabb0" is weak, but "Baaba0" is strong).
Given a string password, return the minimum number of steps required to make password strong. if password is already strong, return 0.

In one step, you can:

Insert one character to password,
Delete one character from password, or
Replace one character of password with another character.
 
Example 1:

Input: password = "a"
Output: 5

Example 2:

Input: password = "aA1"
Output: 3

Example 3:

Input: password = "1337C0d3"
Output: 0 

Constraints:

1 <= password.length <= 50
password consists of letters, digits, dot '.' or exclamation mark '!'.

Source: https://leetcode.com/problems/strong-password-checker/
