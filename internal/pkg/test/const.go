package test

// The Type of the Attachment can be one of these types.
const (
	AttachmentAudio = "audio" // Audio attachment.
	AttachmentFile  = "file"  // File attachment.
	AttachmentImage = "image" // Image attachment.
	AttachmentVideo = "video" // Video attachment.
)

// The Type of the Task can be one of these types.
const (
	TaskSingle   = "single"   // Single-choice task.
	TaskMultiple = "multiple" // Multiple-choice task.
	TaskOpen     = "open"     // Open-ended task.
)
