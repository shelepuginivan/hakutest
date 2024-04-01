const tasksSection = document.getElementById("tasks")
const buttonAddTask = document.getElementById("button-add-task")
const numberOfTasksInput = document.getElementById("number-of-tasks-input")

let taskIndex = Number(numberOfTasksInput.value)

const addTask = (taskIndex) => {
    const newTask = document.createElement("fieldset")

    newTask.classList.add("test-task", "task-${taskIndex}")
    newTask.innerHTML = `
        <legend>${config.get("labelTaskHeader")} ${taskIndex + 1}</legend>
        <div class="input-wrapper">
            <label for="${taskIndex}-type">
                ${config.get("labelTaskType")}
            </label>
            <select
                id="${taskIndex}-type"
                class="input-select"
                name="${taskIndex}-type"
            >
                <option value="single">
                    ${config.get("labelTaskTypeSingle")}
                </option>
                <option value="multiple">
                    ${config.get("labelTaskTypeMultiple")}
                </option>
                <option value="open">
                    ${config.get("labelTaskTypeOpen")}
                </option>
                <option value="file">
                    ${config.get("labelTaskTypeFile")}
                </option>
            </select>
        </div>
        <div class="input-wrapper">
            <label for="${taskIndex}-text">
                ${config.get("labelTaskText")}
            </label>
            <textarea
                id="${taskIndex}-text"
                class="input-text"
                name="${taskIndex}-text"
            ></textarea>
        </div>
        <div class="input-wrapper">
            <label for="${taskIndex}-answer">
                ${config.get("labelTaskAnswer")}
            </label>
            <input
                class="input-text"
                type="text"
                name="${taskIndex}-answer"
            >
        </div>
        <div class="answer-options-wrapper">
            <p>${config.get("labelTaskOptions")}</p>
            <div class="answer-options" id="${taskIndex}-options"></div>
            <button
                class="button-add-option"
                type="button"
                onclick="addOption(${taskIndex})"
            >
                ${config.get("labelTaskAddOption")}
            </button>
        </div>
        <div class="attachment-wrapper">
            <div>
                <input
                    id="${taskIndex}-has-attachment"
                    class="input-checkbox"
                    type="checkbox"
                    name="${taskIndex}-has-attachment"
                    onchange="toggleAttachment(${taskIndex}, this)"
                >
                <label for="${taskIndex}-has-attachment">
                    ${config.get("labelAddAttachment")}
                </label>
            </div>
            <div
                id="${taskIndex}-attachment"
                class="attachment"
                data-enabled="false"
            ></div>
        </div>`

    tasksSection.appendChild(newTask)
}

const addAttachment = (taskIndex) => {
    const attachment = document.getElementById(`${taskIndex}-attachment`)
    const attachmentFields = `
        <div class="input-wrapper">
            <label for="${taskIndex}-attachment-name">
                ${config.get("labelAttachmentName")}
            </label>
            <input
                id="${taskIndex}-attachment-name"
                class="input-text"
                type="text"
                name="${taskIndex}-attachment-name"
            >
        </div>
        <div class="input-wrapper">
            <label for="${taskIndex}-attachment-type">
                ${config.get("labelAttachmentType")}
            </label>
            <select
                id="${taskIndex}-attachment-type"
                class="input-select"
                name="${taskIndex}-attachment-type"
            >
                <option value="file">
                    ${config.get("labelAttachmentTypeFile")}
                </option>
                <option value="image">
                    ${config.get("labelAttachmentTypeImage")}
                </option>
                <option value="video">
                    ${config.get("labelAttachmentTypeVideo")}
                </option>
                <option value="audio">
                    ${config.get("labelAttachmentTypeAudio")}
                </option>
            </select>
        </div>
        <div class="input-wrapper">
            <label for="${taskIndex}-attachment-src">
                ${config.get("labelAttachmentSrc")}
            </label>
            <input
                id="${taskIndex}-attachment-src"
                class="input-text"
                type="text"
                name="${taskIndex}-attachment-src"
            >
        </div>`

    attachment.innerHTML = attachmentFields
    attachment.dataset.enabled = true
}

const removeAttachment = (taskIndex) => {
    const attachment = document.getElementById(`${taskIndex}-attachment`)
    attachment.innerHTML = ""
    attachment.dataset.enabled = false
}

const toggleAttachment = (taskIndex, checkbox) => {
    if (!checkbox.checked) {
        return removeAttachment(taskIndex)
    }

    addAttachment(taskIndex)
}

const addOption = (taskIndex) => {
    const taskOptions = document.getElementById(`${taskIndex}-options`)
    const optionIndex = taskOptions.children.length
    const newOptionWrapper = document.createElement("div")
    const newOptionLabel = document.createElement("label")
    const newOptionInput = document.createElement("input")

    newOptionWrapper.classList.add("input-wrapper")
    newOptionLabel.innerText = `${optionIndex + 1})`
    newOptionLabel.htmlFor = `${taskIndex}-option-${optionIndex}`
    newOptionInput.id = `${taskIndex}-option-${optionIndex}`
    newOptionInput.className = "input-text"
    newOptionInput.type = "text"
    newOptionInput.name = `${taskIndex}-options`

    newOptionWrapper.appendChild(newOptionLabel)
    newOptionWrapper.appendChild(newOptionInput)

    taskOptions.appendChild(newOptionWrapper)
}

buttonAddTask.onclick = () => {
    addTask(taskIndex)
    numberOfTasksInput.value = ++taskIndex
}
