<div align="center">

<img src="./assets/logo.svg" alt="Hakutest" width="680">
<h3>Modern and efficient educational testing</h3>

<br>

<a href="https://hakutest.org" target="_blank"><img src="https://img.shields.io/badge/website-hakutest.org-008000?style=for-the-badge" alt="Website: hakutest.org"></a>
<a href="https://www.codefactor.io/repository/github/shelepuginivan/hakutest/overview" target="_blank"><img src="https://www.codefactor.io/repository/github/shelepuginivan/hakutest/badge?style=for-the-badge" alt="CodeFactor report"></a>
<a href="" target="_blank"><img src="https://goreportcard.com/badge/github.com/shelepuginivan/hakutest?style=for-the-badge"></a>
<a href="https://github.com/shelepuginivan/hakutest/releases/latest"><img src="https://img.shields.io/github/v/release/shelepuginivan/hakutest?style=for-the-badge&color=67B458" alt="Latest Hakutest release"></a>

<br>

**[<kbd> <br> Installation <br> </kbd>](https://hakutest.org/handbook/installation.html)**
**[<kbd> <br> Handbook <br> </kbd>](https://hakutest.org/handbook/getting-started.html)**


</div>

---

## About

**Hakutest** is an educational platform designed for testing in local network. It
allows you to test students, conduct quizzes, and even take exams. In other
words, you can use Hakutest for every task that requires automatic answer
check.

## Installation

Please refer to the [Installation](https://hakutest.org/handbook/installation.html) page on the website.
Alternatively, check out the [Releases](https://github.com/shelepuginivan/hakutest/releases) page.

If you want to compile Hakutest from source, see
[BUILDING.md](https://github.com/shelepuginivan/hakutest/blob/main/BUILDING.md).

## Why Hakutest?

Hakutest offers a fair number of advantages over its analogues:

1. **Efficiency** &mdash; Hakutest is an efficient system capable of checking answers in a matter of milliseconds.
2. **Customizability** &mdash; You can customize the system to suit your needs: internationalization, data export. etc.
3. **Easy to use** &mdash; Hakutest has a clean and accessible interface, uses a familiar markup format (Markdown), etc.
4. **Cross-platform** &mdash; Run Hakutest on Windows, Linux, in a graphical session or on a server.
5. **Security** &mdash; Hakutest uses a reliable policy-based security model.
6. **Privacy** &mdash; You own your data, period. Hakutest respects your privacy.
7. **Freedom** &mdash; Hakutest is a free (as in freedom) and open source software.

## How it works

Hakutest runs on the local network (but it can also be run on a global network
as well). Each test is represented by a JSON file that is stored locally on
your device. When the Hakutest server is running, students can access the test
through the browser. The platform dynamically generates a web page for each
test, displaying the questions and any additional content specified in the test
file.

Once students have completed the test and submitted their answers, Hakutest
automatically checks their responses against the expected solutions. This
automated answer checking process provides instant feedback to students. The
platform evaluates each answer based on the predefined criteria set in the test
file, allowing for accurate and efficient grading.

By utilizing this approach, Hakutest ensures that the testing process is secure
and reliable. Since the test files are stored locally, the platform maintains
data privacy and prevents unauthorized access to sensitive information.
