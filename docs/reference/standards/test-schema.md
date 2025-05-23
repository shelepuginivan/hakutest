---
title: Test JSON schema
description: "Hakutest test JSON schema is available at https://hakutest.org/standards/test.schema.json"
---

:::info

Hakutest Test JSON schema is available at https://hakutest.org/standards/test.schema.json

:::

---
# Test

- [1. Property `Test > title`](#title)
- [2. Property `Test > description`](#description)
- [3. Property `Test > subject`](#subject)
- [4. Property `Test > author`](#author)
- [5. Property `Test > target`](#target)
- [6. Property `Test > institution`](#institution)
- [7. Property `Test > createdAt`](#createdAt)
- [8. Property `Test > expiresAt`](#expiresAt)
- [9. Property `Test > shuffleTasks`](#shuffleTasks)
- [10. Property `Test > tasks`](#tasks)
  - [10.1. Test > tasks > tasks items](#tasks_items)
    - [10.1.1. Property `Test > tasks > tasks items > type`](#tasks_items_type)
    - [10.1.2. Property `Test > tasks > tasks items > text`](#tasks_items_text)
    - [10.1.3. Property `Test > tasks > tasks items > options`](#tasks_items_options)
      - [10.1.3.1. Test > tasks > tasks items > options > options items](#tasks_items_options_items)
    - [10.1.4. Property `Test > tasks > tasks items > answer`](#tasks_items_answer)

**Title:** Test

|                           |                  |
| ------------------------- | ---------------- |
| **Type**                  | `object`         |
| **Required**              | No               |
| **Additional properties** | Any type allowed |

**Description:** Hakutest Test JSON file schema.

| Property                         | Pattern | Type            | Deprecated | Definition | Title/Description                                               |
| -------------------------------- | ------- | --------------- | ---------- | ---------- | --------------------------------------------------------------- |
| + [title](#title )               | No      | string          | No         | -          | Title of the test.                                              |
| - [description](#description )   | No      | string          | No         | -          | Description of the test.                                        |
| - [subject](#subject )           | No      | string          | No         | -          | Subject or topic of the test.                                   |
| - [author](#author )             | No      | string          | No         | -          | Author (creator) of the test.                                   |
| - [target](#target )             | No      | string          | No         | -          | Target audience of the test.                                    |
| - [institution](#institution )   | No      | string          | No         | -          | Educational institution where the test is conducted.            |
| - [createdAt](#createdAt )       | No      | string          | No         | -          | Test creation time.                                             |
| - [expiresAt](#expiresAt )       | No      | string          | No         | -          | Test expiration time.                                           |
| - [shuffleTasks](#shuffleTasks ) | No      | boolean         | No         | -          | Whether tasks should be shuffled and displayed in random order. |
| + [tasks](#tasks )               | No      | array of object | No         | -          | Tasks of the test.                                              |

## 1. Property `Test > title` {#title}

|              |          |
| ------------ | -------- |
| **Type**     | `string` |
| **Required** | Yes      |

**Description:** Title of the test.

## 2. Property `Test > description` {#description}

|              |          |
| ------------ | -------- |
| **Type**     | `string` |
| **Required** | No       |

**Description:** Description of the test.

## 3. Property `Test > subject` {#subject}

|              |          |
| ------------ | -------- |
| **Type**     | `string` |
| **Required** | No       |

**Description:** Subject or topic of the test.

## 4. Property `Test > author` {#author}

|              |          |
| ------------ | -------- |
| **Type**     | `string` |
| **Required** | No       |

**Description:** Author (creator) of the test.

## 5. Property `Test > target` {#target}

|              |          |
| ------------ | -------- |
| **Type**     | `string` |
| **Required** | No       |

**Description:** Target audience of the test.

## 6. Property `Test > institution` {#institution}

|              |          |
| ------------ | -------- |
| **Type**     | `string` |
| **Required** | No       |

**Description:** Educational institution where the test is conducted.

## 7. Property `Test > createdAt` {#createdAt}

|              |             |
| ------------ | ----------- |
| **Type**     | `string`    |
| **Required** | No          |
| **Format**   | `date-time` |

**Description:** Test creation time.

## 8. Property `Test > expiresAt` {#expiresAt}

|              |             |
| ------------ | ----------- |
| **Type**     | `string`    |
| **Required** | No          |
| **Format**   | `date-time` |

**Description:** Test expiration time.

## 9. Property `Test > shuffleTasks` {#shuffleTasks}

|              |           |
| ------------ | --------- |
| **Type**     | `boolean` |
| **Required** | No        |

**Description:** Whether tasks should be shuffled and displayed in random order.

## 10. Property `Test > tasks` {#tasks}

|              |                   |
| ------------ | ----------------- |
| **Type**     | `array of object` |
| **Required** | Yes               |

**Description:** Tasks of the test.

|                      | Array restrictions |
| -------------------- | ------------------ |
| **Min items**        | N/A                |
| **Max items**        | N/A                |
| **Items unicity**    | False              |
| **Additional items** | False              |
| **Tuple validation** | See below          |

| Each item of this array must be | Description |
| ------------------------------- | ----------- |
| [tasks items](#tasks_items)     | -           |

### 10.1. Test > tasks > tasks items {#tasks_items}

|                           |                  |
| ------------------------- | ---------------- |
| **Type**                  | `object`         |
| **Required**              | No               |
| **Additional properties** | Any type allowed |

| Property                           | Pattern | Type                    | Deprecated | Definition | Title/Description                                                                                                                                                                                                                                                   |
| ---------------------------------- | ------- | ----------------------- | ---------- | ---------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| + [type](#tasks_items_type )       | No      | enum (of string)        | No         | -          | Type of the task.                                                                                                                                                                                                                                                   |
| + [text](#tasks_items_text )       | No      | string                  | No         | -          | Text of the task. Can be multi-line. Markdown is supported.                                                                                                                                                                                                         |
| - [options](#tasks_items_options ) | No      | array of string or null | No         | -          | Answer options of the task. null for "open" and "detailed" task.                                                                                                                                                                                                    |
| + [answer](#tasks_items_answer )   | No      | string                  | No         | -          | Correct answer of the task. For "single" - index of correct option (zero-indexed). For "multiple" - comma-separated indices of correct answer options (zero-indexed, lexically ascending order). For "open" - correct answer string. For "detailed" - empty string. |

#### 10.1.1. Property `Test > tasks > tasks items > type` {#tasks_items_type}

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

#### 10.1.2. Property `Test > tasks > tasks items > text` {#tasks_items_text}

|              |          |
| ------------ | -------- |
| **Type**     | `string` |
| **Required** | Yes      |

**Description:** Text of the task. Can be multi-line. Markdown is supported.

#### 10.1.3. Property `Test > tasks > tasks items > options` {#tasks_items_options}

|              |                           |
| ------------ | ------------------------- |
| **Type**     | `array of string or null` |
| **Required** | No                        |

**Description:** Answer options of the task. null for "open" and "detailed" task.

|                      | Array restrictions |
| -------------------- | ------------------ |
| **Min items**        | N/A                |
| **Max items**        | N/A                |
| **Items unicity**    | False              |
| **Additional items** | False              |
| **Tuple validation** | See below          |

| Each item of this array must be             | Description |
| ------------------------------------------- | ----------- |
| [options items](#tasks_items_options_items) | -           |

##### 10.1.3.1. Test > tasks > tasks items > options > options items {#tasks_items_options_items}

|              |          |
| ------------ | -------- |
| **Type**     | `string` |
| **Required** | No       |

#### 10.1.4. Property `Test > tasks > tasks items > answer` {#tasks_items_answer}

|              |          |
| ------------ | -------- |
| **Type**     | `string` |
| **Required** | Yes      |

**Description:** Correct answer of the task. For "single" - index of correct option (zero-indexed). For "multiple" - comma-separated indices of correct answer options (zero-indexed, lexically ascending order). For "open" - correct answer string. For "detailed" - empty string.

----------------------------------------------------------------------------------------------------------------------------
Generated using [json-schema-for-humans](https://github.com/coveooss/json-schema-for-humans) on 2025-04-16 at 21:38:51 +0300
