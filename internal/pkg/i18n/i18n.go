package i18n

import (
	"embed"
	"fmt"
	"slices"

	"gopkg.in/yaml.v3"
)

//go:embed translations
var translations embed.FS

func New(lang string) *I18n {
	var i18n I18n

	if !slices.Contains(AvailableLanguages, lang) {
		lang = LanguageEn
	}

	data, err := translations.ReadFile(fmt.Sprintf("translations/%s.yaml", lang))
	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(data, &i18n); err != nil {
		panic(err)
	}

	return &i18n
}

type GtkServerI18n struct {
	Title        string `yaml:"title"`
	LabelIdle    string `yaml:"label_idle"`
	LabelRunning string `yaml:"label_running"`
	LabelError   string `yaml:"label_error"`
	ButtonStart  string `yaml:"button_start"`
	ButtonStop   string `yaml:"button_stop"`
}

type GtkEditorFormI18n struct {
	InputTitle       string `yaml:"input_title"`
	InputDescription string `yaml:"input_description"`
	InputSubject     string `yaml:"input_subject"`
	InputAuthor      string `yaml:"input_author"`
	InputTarget      string `yaml:"input_target"`
	InputInstitution string `yaml:"input_institution"`
	CheckExpiresAt   string `yaml:"check_expires_at"`
	ButtonClose      string `yaml:"button_close"`
	ButtonSave       string `yaml:"button_save"`
	LabelSuccess     string `yaml:"label_success"`
	LabelError       string `yaml:"label_error"`
}

type GtkEditorMenuI18n struct {
	Title        string `yaml:"title"`
	SidebarLabel string `yaml:"sidebar_label"`
	InputTest    string `yaml:"input_test"`
	ButtonOpen   string `yaml:"button_open"`
	ButtonCreate string `yaml:"button_create"`
}

type GtkEditorAttachmentI18n struct {
	CheckInclude           string `yaml:"check_include"`
	InputName              string `yaml:"input_name"`
	InputType              string `yaml:"input_type"`
	LabelTypeFile          string `yaml:"label_type_file"`
	LabelTypeAudio         string `yaml:"label_type_audio"`
	LabelTypeImage         string `yaml:"label_type_image"`
	LabelTypeVideo         string `yaml:"label_type_video"`
	LabelModeUrl           string `yaml:"label_mode_url"`
	LabelModeFile          string `yaml:"label_mode_file"`
	LabelModeLoaded        string `yaml:"label_mode_loaded"`
	FileDialogTitle        string `yaml:"file_dialog_title"`
	FileDialogButtonOpen   string `yaml:"file_dialog_button_open"`
	FileDialogButtonCancel string `yaml:"file_dialog_button_cancel"`
}

type GtkEditorTaskListI18n struct {
	ButtonAdd string `yaml:"button_add"`
}

type GtkEditorTaskI18n struct {
	InputType                      string                   `yaml:"input_type"`
	LabelTypeSingle                string                   `yaml:"label_type_single"`
	LabelTypeMultiple              string                   `yaml:"label_type_multiple"`
	LabelTypeOpen                  string                   `yaml:"label_type_open"`
	LabelTypeFile                  string                   `yaml:"label_type_file"`
	InputText                      string                   `yaml:"input_text"`
	InputAnswer                    string                   `yaml:"input_answer"`
	AnswerOptionsSingleLabel       string                   `yaml:"answer_options_single_label"`
	AnswerOptionsSingleButtonAdd   string                   `yaml:"answer_options_single_button_add"`
	AnswerOptionsMultipleLabel     string                   `yaml:"answer_options_multiple_label"`
	AnswerOptionsMultipleButtonAdd string                   `yaml:"answer_options_multiple_button_add"`
	Attachment                     *GtkEditorAttachmentI18n `yaml:"attachment"`
	List                           *GtkEditorTaskListI18n   `yaml:"list"`
}

type GtkEditorI18n struct {
	NewTestLabel string             `yaml:"new_test_label"`
	Form         *GtkEditorFormI18n `yaml:"form"`
	Menu         *GtkEditorMenuI18n `yaml:"menu"`
	Task         *GtkEditorTaskI18n `yaml:"task"`
}

type GtkStatsI18n struct {
	Title        string `yaml:"title"`
	LabelSuccess string `yaml:"label_success"`
	LabelError   string `yaml:"label_error"`
	InputTest    string `yaml:"input_test"`
	InputFormat  string `yaml:"input_format"`
	ButtonExport string `yaml:"button_export"`
}

type GtkConfigGeneralI18n struct {
	Title                        string `yaml:"title"`
	InputLang                    string `yaml:"input_lang"`
	InputTestsDirectory          string `yaml:"input_tests_directory"`
	InputTestsDirectoryTooltip   string `yaml:"input_tests_directory_tooltip"`
	InputResultsDirectory        string `yaml:"input_results_directory"`
	InputResultsDirectoryTooltip string `yaml:"input_results_directory_tooltip"`
	InputShowResults             string `yaml:"input_show_results"`
	InputShowResultsTooltip      string `yaml:"input_show_results_tooltip"`
	InputOverwriteResults        string `yaml:"input_overwrite_results"`
	InputOverwriteResultsTooltip string `yaml:"input_overwrite_results_tooltip"`
}

type GtkConfigServerI18n struct {
	Title                     string `yaml:"title"`
	InputPort                 string `yaml:"input_port"`
	InputPortTooltip          string `yaml:"input_port_tooltip"`
	InputMaxUploadSize        string `yaml:"input_max_upload_size"`
	InputMaxUploadSizeTooltip string `yaml:"input_max_upload_size_tooltip"`
	InputMode                 string `yaml:"input_mode"`
	InputModeTooltip          string `yaml:"input_mode_tooltip"`
}

type GtkConfigI18n struct {
	Title   string                `yaml:"title"`
	General *GtkConfigGeneralI18n `yaml:"general"`
	Server  *GtkConfigServerI18n  `yaml:"server"`
}

type GtkI18n struct {
	Server *GtkServerI18n `yaml:"server"`
	Editor *GtkEditorI18n `yaml:"editor"`
	Stats  *GtkStatsI18n  `yaml:"stats"`
	Config *GtkConfigI18n `yaml:"config"`
}

type ServerI18n struct {
	StopTitle   string `yaml:"stop_title" mapstructure:"stop_title"`
	StopTooltip string `yaml:"stop_tooltip" mapstructure:"stop_tooltip"`
}

type StatsAppI18n struct {
	LabelTest      string `yaml:"label_test" mapstructure:"label_test"`
	LabelFormat    string `yaml:"label_format" mapstructure:"label_format"`
	LabelDirectory string `yaml:"label_directory" mapstructure:"label_directory"`
	SubmitText     string `yaml:"submit_text" mapstructure:"submit_text"`
	CancelText     string `yaml:"cancel_text" mapstructure:"cancel_text"`
	SelectText     string `yaml:"select_text" mapstructure:"select_text"`
	SuccessText    string `yaml:"success_text" mapstructure:"success_text"`
	ErrorPrefix    string `yaml:"error_prefix" mapstructure:"error_prefix"`
}

type StatsExcelI18n struct {
	TestResultsSheet    string `yaml:"test_results_sheet" mapstructure:"test_results_sheet"`
	TestStatisticsSheet string `yaml:"statistics_sheet" mapstructure:"statistics_sheet"`
	HeaderStudent       string `yaml:"header_student" mapstructure:"header_student"`
	HeaderPoints        string `yaml:"header_points" mapstructure:"header_points"`
	HeaderPercentage    string `yaml:"header_percentage" mapstructure:"header_percentage"`
}

type StatsImageI18n struct {
	Title  string `yaml:"title" mapstructure:"title"`
	LabelX string `yaml:"label_x" mapstructure:"label_x"`
	LabelY string `yaml:"label_y" mapstructure:"label_y"`
}

type StatsI18n struct {
	App   StatsAppI18n   `yaml:"app" mapstructure:"app"`
	Excel StatsExcelI18n `yaml:"excel" mapstructure:"excel"`
	Image StatsImageI18n `yaml:"image" mapstructure:"image"`
}

type WebEditorI18n struct {
	Header                   string `yaml:"header" mapstructure:"header"`
	LabelTitle               string `yaml:"label_title" mapstructure:"label_title"`
	LabelDescription         string `yaml:"label_description" mapstructure:"label_description"`
	LabelSubject             string `yaml:"label_subject" mapstructure:"label_subject"`
	LabelAuthor              string `yaml:"label_author" mapstructure:"label_author"`
	LabelTarget              string `yaml:"label_target" mapstructure:"label_target"`
	LabelInstitution         string `yaml:"label_institution" mapstructure:"label_institution"`
	LabelExpiresIn           string `yaml:"label_expires_in" mapstructure:"label_expires_in"`
	LabelAddTask             string `yaml:"label_add_task" mapstructure:"label_add_task"`
	LabelTaskHeader          string `yaml:"label_task_header" mapstructure:"label_task_header"`
	LabelTaskType            string `yaml:"label_task_type" mapstructure:"label_task_type"`
	LabelTaskTypeSingle      string `yaml:"label_task_type_single" mapstructure:"label_task_type_single"`
	LabelTaskTypeMultiple    string `yaml:"label_task_type_multiple" mapstructure:"label_task_type_multiple"`
	LabelTaskTypeOpen        string `yaml:"label_task_type_open" mapstructure:"label_task_type_open"`
	LabelTaskTypeFile        string `yaml:"label_task_type_file" mapstructure:"label_task_type_file"`
	LabelTaskText            string `yaml:"label_task_text" mapstructure:"label_task_text"`
	LabelTaskAnswer          string `yaml:"label_task_answer" mapstructure:"label_task_answer"`
	LabelTaskOptions         string `yaml:"label_task_options" mapstructure:"label_task_options"`
	LabelTaskAddOption       string `yaml:"label_task_add_option" mapstructure:"label_task_add_option"`
	LabelAddAttachment       string `yaml:"label_add_attachment" mapstructure:"label_add_attachment"`
	LabelAttachmentName      string `yaml:"label_attachment_name" mapstructure:"label_attachment_name"`
	LabelAttachmentType      string `yaml:"label_attachment_type" mapstructure:"label_attachment_type"`
	LabelAttachmentTypeFile  string `yaml:"label_attachment_type_file" mapstructure:"label_attachment_type_file"`
	LabelAttachmentTypeImage string `yaml:"label_attachment_type_image" mapstructure:"label_attachment_type_image"`
	LabelAttachmentTypeVideo string `yaml:"label_attachment_type_video" mapstructure:"label_attachment_type_video"`
	LabelAttachmentTypeAudio string `yaml:"label_attachment_type_audio" mapstructure:"label_attachment_type_audio"`
	LabelAttachmentSrc       string `yaml:"label_attachment_src" mapstructure:"label_attachment_src"`
	LabelUploadTestInput     string `yaml:"label_upload_test_input" mapstructure:"label_upload_test_input"`
	LabelUploadTestButton    string `yaml:"label_upload_test_button" mapstructure:"label_upload_test_button"`
	LabelNewTest             string `yaml:"label_new_test" mapstructure:"label_new_test"`
	LabelDownloadTest        string `yaml:"label_download_test" mapstructure:"label_download_test"`
}

type WebErrorI18n struct {
	Header  string `yaml:"header" mapstructure:"header"`
	Details string `yaml:"details" mapstructure:"details"`
}

type WebExpiredI18n struct {
	Header  string `yaml:"header" mapstructure:"header"`
	Message string `yaml:"message" mapstructure:"message"`
}

type WebSearchI18n struct {
	InputPlaceholder  string `yaml:"input_placeholder" mapstructure:"input_placeholder"`
	SearchButtonLabel string `yaml:"search_button_label" mapstructure:"search_button_label"`
}

type WebSubmittedI18n struct {
	Header  string `yaml:"header" mapstructure:"header"`
	Message string `yaml:"message" mapstructure:"message"`
}

type WebTestI18n struct {
	StudentNameLabel  string `yaml:"student_name_label" mapstructure:"student_name_label"`
	OpenAnswerLabel   string `yaml:"open_answer_label" mapstructure:"open_answer_label"`
	SubmitButtonLabel string `yaml:"submit_button_label" mapstructure:"submit_button_label"`
}

type WebI18n struct {
	Editor    WebEditorI18n    `yaml:"editor" mapstructure:"editor"`
	Error     WebErrorI18n     `yaml:"error" mapstructure:"error"`
	Expired   WebExpiredI18n   `yaml:"expired" mapstructure:"expired"`
	Search    WebSearchI18n    `yaml:"search" mapstructure:"search"`
	Submitted WebSubmittedI18n `yaml:"submitted" mapstructure:"submitted"`
	Test      WebTestI18n      `yaml:"test" mapstructure:"test"`
}

type I18n struct {
	Language   string     `yaml:"lang" mapstructure:"lang"`
	Gtk        *GtkI18n   `yaml:"gtk"`
	Server     ServerI18n `yaml:"server" mapstructure:"server"`
	Statistics StatsI18n  `yaml:"stats" mapstructure:"stats"`
	Web        WebI18n    `yaml:"web" mapstructure:"web"`
}

func Default() *I18n {
	return &I18n{
		Language: "en",
		Server: ServerI18n{
			StopTitle:   "Stop Hakutest",
			StopTooltip: "Stop Hakutest server and quit",
		},
		Statistics: StatsI18n{
			App: StatsAppI18n{
				LabelTest:      "Test",
				LabelFormat:    "Format",
				LabelDirectory: "Export to",
				SubmitText:     "Export",
				CancelText:     "Cancel",
				SelectText:     "(Select one)",
				SuccessText:    "Statistics exported successfully!",
				ErrorPrefix:    "An error occurred! Detail:",
			},
			Excel: StatsExcelI18n{
				TestResultsSheet:    "Test Results",
				TestStatisticsSheet: "Test Statistics",
				HeaderStudent:       "Student",
				HeaderPoints:        "Points",
				HeaderPercentage:    "%",
			},
			Image: StatsImageI18n{
				Title:  "Student Performance",
				LabelX: "Points",
				LabelY: "Students",
			},
		},
		Web: WebI18n{
			Editor: WebEditorI18n{
				Header:                   "Test Editor",
				LabelTitle:               "Title:",
				LabelDescription:         "Description:",
				LabelSubject:             "Subject:",
				LabelAuthor:              "Author:",
				LabelTarget:              "Target audience:",
				LabelInstitution:         "Institution:",
				LabelExpiresIn:           "Expires in:",
				LabelAddTask:             "+ Add task",
				LabelTaskHeader:          "Task",
				LabelTaskType:            "Type:",
				LabelTaskTypeSingle:      "Single answer",
				LabelTaskTypeMultiple:    "Multiple answers",
				LabelTaskTypeOpen:        "Open question",
				LabelTaskTypeFile:        "Answer with file(s)",
				LabelTaskText:            "Text:",
				LabelTaskAnswer:          "Answer:",
				LabelTaskOptions:         "Answer options",
				LabelTaskAddOption:       "+ Add option",
				LabelAddAttachment:       "Add attachment",
				LabelAttachmentName:      "Name:",
				LabelAttachmentType:      "Type:",
				LabelAttachmentTypeFile:  "File",
				LabelAttachmentTypeImage: "Image",
				LabelAttachmentTypeVideo: "Video",
				LabelAttachmentTypeAudio: "Audio",
				LabelAttachmentSrc:       "Source (URL):",
				LabelUploadTestInput:     "Upload test file",
				LabelUploadTestButton:    "Upload and edit",
				LabelNewTest:             "Create new test",
				LabelDownloadTest:        "Download test",
			},
			Error: WebErrorI18n{
				Header:  "An error occurred!",
				Details: "Details",
			},
			Expired: WebExpiredI18n{
				Header:  "Test expired!",
				Message: "This test is no longer available",
			},
			Search: WebSearchI18n{
				InputPlaceholder:  "Search for a test",
				SearchButtonLabel: "Search",
			},
			Submitted: WebSubmittedI18n{
				Header:  "Submitted!",
				Message: "The test results are not displayed according to the system settings",
			},
			Test: WebTestI18n{
				StudentNameLabel:  "Your name:",
				OpenAnswerLabel:   "Answer:",
				SubmitButtonLabel: "Submit",
			},
		},
	}
}
