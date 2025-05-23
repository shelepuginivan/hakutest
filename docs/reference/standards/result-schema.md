---
title: Result JSON schema
description: "Hakutest result JSON schema is available at https://hakutest.org/standards/result.schema.json"
---

:::info

Hakutest Result JSON schema is available at https://hakutest.org/standards/result.schema.json

:::

---
# Result

- [1. Property `Result > student`](#student)
- [2. Property `Result > submitted_at`](#submitted_at)
- [3. Property `Result > answers`](#answers)
  - [3.1. Result > answers > answers items](#answers_items)
    - [3.1.1. Property `Result > answers > answers items > type`](#answers_items_type)
    - [3.1.2. Property `Result > answers > answers items > value`](#answers_items_value)
    - [3.1.3. Property `Result > answers > answers items > correct`](#answers_items_correct)
- [4. Property `Result > points`](#points)
- [5. Property `Result > total`](#total)
- [6. Property `Result > percentage`](#percentage)

**Title:** Result

|                           |                  |
| ------------------------- | ---------------- |
| **Type**                  | `object`         |
| **Required**              | No               |
| **Additional properties** | Any type allowed |

**Description:** Hakutest Result JSON file schema.

| Property                         | Pattern | Type            | Deprecated | Definition | Title/Description                                                             |
| -------------------------------- | ------- | --------------- | ---------- | ---------- | ----------------------------------------------------------------------------- |
| + [student](#student )           | No      | string          | No         | -          | The student who submitted the solution.                                       |
| + [submitted_at](#submitted_at ) | No      | string          | No         | -          | Solution submission time                                                      |
| + [answers](#answers )           | No      | array of object | No         | -          | Answer given by the student for each task present in test.                    |
| + [points](#points )             | No      | integer         | No         | -          | Number of correct answers given by the student.                               |
| + [total](#total )               | No      | integer         | No         | -          | Total number of tasks present in the test. Same as the total scorable points. |
| + [percentage](#percentage )     | No      | integer         | No         | -          | Percentage scored by student. Defined as (100 * points / total).              |

## 1. Property `Result > student` {#student}

|              |          |
| ------------ | -------- |
| **Type**     | `string` |
| **Required** | Yes      |

**Description:** The student who submitted the solution.

## 2. Property `Result > submitted_at` {#submitted_at}

|              |             |
| ------------ | ----------- |
| **Type**     | `string`    |
| **Required** | Yes         |
| **Format**   | `date-time` |

**Description:** Solution submission time

## 3. Property `Result > answers` {#answers}

|              |                   |
| ------------ | ----------------- |
| **Type**     | `array of object` |
| **Required** | Yes               |

**Description:** Answer given by the student for each task present in test.

|                      | Array restrictions |
| -------------------- | ------------------ |
| **Min items**        | N/A                |
| **Max items**        | N/A                |
| **Items unicity**    | False              |
| **Additional items** | False              |
| **Tuple validation** | See below          |

| Each item of this array must be | Description |
| ------------------------------- | ----------- |
| [answers items](#answers_items) | -           |

### 3.1. Result > answers > answers items {#answers_items}

|                           |                  |
| ------------------------- | ---------------- |
| **Type**                  | `object`         |
| **Required**              | No               |
| **Additional properties** | Any type allowed |

| Property                             | Pattern | Type             | Deprecated | Definition | Title/Description                                                                                                                                                                                                                                  |
| ------------------------------------ | ------- | ---------------- | ---------- | ---------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| + [type](#answers_items_type )       | No      | enum (of string) | No         | -          | Type of the task.                                                                                                                                                                                                                                  |
| + [value](#answers_items_value )     | No      | string           | No         | -          | The actual answer given by the student. For "single" - index of the chosen option (zero-indexed). For "multiple" - comma-separated indices of chosen options (zero-indexed, lexically ascending order). For "open" and "detailed" - answer string. |
| + [correct](#answers_items_correct ) | No      | boolean          | No         | -          | Whether the answer is correct.                                                                                                                                                                                                                     |

#### 3.1.1. Property `Result > answers > answers items > type` {#answers_items_type}

|              |                    |
| ------------ | ------------------ |
| **Type**     | `enum (of string)` |
| **Required** | Yes                |

**Description:** Type of the task.

Must be one of:
* "single"
* "multiple"
* "open"
* "detailed"

#### 3.1.2. Property `Result > answers > answers items > value` {#answers_items_value}

|              |          |
| ------------ | -------- |
| **Type**     | `string` |
| **Required** | Yes      |

**Description:** The actual answer given by the student. For "single" - index of the chosen option (zero-indexed). For "multiple" - comma-separated indices of chosen options (zero-indexed, lexically ascending order). For "open" and "detailed" - answer string.

#### 3.1.3. Property `Result > answers > answers items > correct` {#answers_items_correct}

|              |           |
| ------------ | --------- |
| **Type**     | `boolean` |
| **Required** | Yes       |

**Description:** Whether the answer is correct.

## 4. Property `Result > points` {#points}

|              |           |
| ------------ | --------- |
| **Type**     | `integer` |
| **Required** | Yes       |

**Description:** Number of correct answers given by the student.

## 5. Property `Result > total` {#total}

|              |           |
| ------------ | --------- |
| **Type**     | `integer` |
| **Required** | Yes       |

**Description:** Total number of tasks present in the test. Same as the total scorable points.

## 6. Property `Result > percentage` {#percentage}

|              |           |
| ------------ | --------- |
| **Type**     | `integer` |
| **Required** | Yes       |

**Description:** Percentage scored by student. Defined as (100 * points / total).

----------------------------------------------------------------------------------------------------------------------------
Generated using [json-schema-for-humans](https://github.com/coveooss/json-schema-for-humans) on 2025-04-16 at 21:38:51 +0300
