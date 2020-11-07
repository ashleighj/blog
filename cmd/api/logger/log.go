package logger

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"runtime"
	"strings"
)

type ctxKey string

const (
	prefError = "[ ERROR ] "
	prefWarn  = "[ WARN  ] "
	prefInfo  = "[ INFO  ] "
	prefFatal = "[ FATAL ] "

	keyReqID  ctxKey = "x-request-id"
	keyAddr   ctxKey = "remote-address"
	keySesson ctxKey = "session-id"

	callDepth = 3
)

var (
	infoLog  = log.New(os.Stdout, prefInfo, log.Ldate|log.Ltime)
	warnLog  = log.New(os.Stdout, prefWarn, log.Ldate|log.Ltime)
	errLog   = log.New(os.Stderr, prefError, log.Ldate|log.Ltime)
	fatalLog = log.New(os.Stderr, prefFatal, log.Ldate|log.Ltime)
)

// Info logs info-level logs
func Info(ctx context.Context, vars ...interface{}) {
	lg(ctx, infoLog, vars)
}

// Infof logs formatted info-level logs
func Infof(ctx context.Context, fmtStr string, vars ...interface{}) {
	msg := fmt.Sprintf(fmtStr, vars...)
	lg(ctx, infoLog, msg)
}

// Warn logs warning-level logs
func Warn(ctx context.Context, vars ...interface{}) {
	lg(ctx, warnLog, vars)
}

// Error logs error-level logs
func Error(ctx context.Context, vars ...interface{}) {
	lg(ctx, errLog, vars)
}

// Errorf logs formatted error-level logs
func Errorf(ctx context.Context, fmtStr string, vars ...interface{}) {
	msg := fmt.Sprintf(fmtStr, vars...)
	lg(ctx, errLog, msg)
}

// Fatal logs fatal logs
func Fatal(ctx context.Context, vars ...interface{}) {
	lg(ctx, fatalLog, vars)
}

// Fatalf logs formatted fatal logs
func Fatalf(ctx context.Context, fmtStr string, vars ...interface{}) {
	msg := fmt.Sprintf(fmtStr, vars...)
	lg(ctx, fatalLog, msg)
}

func lg(ctx context.Context, l *log.Logger, vars ...interface{}) {
	fileLn, fn := getLoc()

	strs := []string{}
	for _, v := range vars {
		str := strings.Trim(fmt.Sprintf("%v", v), "[]")
		strs = append(strs, str)
	}
	msg := strings.Join(strs, ", ")

	l.Println(ctxDataStr(ctx), fmt.Sprintf(
		`location="%s" fn="%s" msg="%s"`, fileLn, fn, msg))
}

func getLoc() (fileLine string, fn string) {
	workDir, _ := os.Getwd()

	pc, file, line, _ := runtime.Caller(callDepth)
	file = strings.TrimPrefix(file, workDir+"/")

	caller := runtime.FuncForPC(pc)
	if caller != nil {
		r := regexp.MustCompile(`[^\/]+$`)
		fn = r.FindString(caller.Name())
	}

	return fmt.Sprintf("%s:%d", file, line), fn
}

func ctxDataStr(ctx context.Context) (lg string) {
	rID := ctx.Value(keyReqID)
	if rID != nil {
		lg += fmt.Sprintf(`req_id="%s"`, rID.(string))
	}

	rAddr := ctx.Value(keyReqID)
	if rAddr != nil {
		lg += fmt.Sprintf(`r_addr="%s"`, rAddr.(string))
	}

	session := ctx.Value(keyReqID)
	if session != nil {
		lg += fmt.Sprintf(`session="%s"`, session.(string))
	}

	return
}
