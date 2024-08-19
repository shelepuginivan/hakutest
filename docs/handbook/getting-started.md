---
title: Getting Started
titleTemplate: Hakutest Handbook
description: 'Hakutest is an educational platform designed for testing, quizzes, and exams with automatic answer checking'
---

# Getting Started

![Hakutest Logo](/logo/logo.svg)

:::details TLDR

_Hakutest is an educational platform designed for testing, quizzes, and exams
with automatic answer checking. It offers advantages over analogues. Hakutest
operates by storing test files locally, generating web pages for each test, and
automatically checking student responses against expected solutions for instant
feedback and accurate grading._

---

**However, we recommend that you read the entire page.**

:::

## About Hakutest

Hakutest is an educational platform designed for testing in local network. It
allows you to test students, conduct quizzes, and even take exams. In other
words, you can use Hakutest for every task that requires automatic answer
check.

It has a number of advantages over its analogues:

-   **Security**: Your data and your students' data are stored locally, ensuring
    that no one else can access it.
-   **Efficiency**: Hakutest is an efficient system capable of checking students'
    answers in a matter of milliseconds.
-   **Customizability**: You can customize the system to suit your needs:
    internationalization, different environments, data export and more.
-   **Cross-platform**: Students can access Hakutest from any device. Server is
    available for Windows and Linux.
-   **Free**: Hakutest is a free (as in freedom) and open source software.
-   **Just works**: Hakutest is configured to work out-of-the-box.

## Motivation

Today, numerous testing systems are available for educational institutions.
However, each of them has significant drawbacks. Some of them have security
issues: they compromise answers or allow students to falsify their results.
Some of them are proprietary, so student data is stored on third-party servers.

Hakutest was created with the intention of mitigating these weaknesses. While it
is definitely not the most important tool out there (_it is neither a window
system, nor a driver, etc._), it attempts to provide a transparent and smooth
testing experience for both students and teachers. Focus on the educational
process, not on fighting software.

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
