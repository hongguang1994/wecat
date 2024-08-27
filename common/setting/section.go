package setting

import "time"

type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	IsHttps      bool
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	SSL          SSLSettingS
}

type SSLSettingS struct {
	Key string
	Pem string
}

type AppSettingS struct {
	DefaultPageSize      int
	MaxPageSize          int
	LogSavePath          string
	LogFileName          string
	LogFileExt           string
	UploadSavePath       string
	UploadServerUrl      string
	UploadImageMaxSize   int
	UploadImageAllowExts []string
}

type DatabaseSettingS struct {
	DBType       string
	UserName     string
	Passworld    string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type JWTSettingS struct {
	Secret string
	Issuer string
	Expire time.Duration
}

type EmailSettingS struct {
	Host     string
	Port     int
	UserName string
	Password string
	IsSSL    bool
	From     string
	To       []string
}

type LogSettingS struct {
	Compress      bool
	ConsoleStdout bool
	FileStdout    bool
	Level         string
	LocalTime     bool
	MaxAge        int
	MaxBackups    int
	MaxSize       int
	Path          string
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
