package err

import (
	"fmt"
	"strings"
)

var (
	CodeNotFound      = "resource/not_found"
	CodeProps         = "input/property_errors"
	CodeMalformed     = "input/malformed"
	CodeNoIdentifier  = "input/missing_identifier"
	CodeOneIdentifier = "input/too_many_identifiers"
	CodeInternal      = "internal"
	CodeDatastore     = "internal/datastore"
	CodeNoTemplate    = "internal/template/not_found"
	CodeExecTemplate  = "internal/template/compile"
	CodeStartup       = "internal/service_startup"
	CodeNotImpl       = "functionality/not_implemented"

	MsgProps        = "problems with input data"
	MsgPropMissing  = "required, missing"
	MsgPropInt      = "should be an integer"
	MsgPropString   = "should be a string"
	MsgNoTemplate   = "template not found"
	MsgExecTemplate = "could not compile template & data"

	DatastoreInit = func(e error) ServiceError {
		msg := fmt.Sprintf("internal service error - %s", e.Error())
		return New(msg, CodeStartup)
	}

	RepoInit = func(e error) ServiceError {
		msg := fmt.Sprintf("repo init error - %s", e.Error())
		return New(msg, CodeStartup)
	}

	ServiceInit = func(e error) ServiceError {
		msg := fmt.Sprintf("service init error - %s", e.Error())
		return New(msg, CodeStartup)
	}

	TemplateExec = func(e error) ServiceError {
		msg := fmt.Sprintf("template exec error - %s", e.Error())
		return New(msg, CodeInternal)
	}

	TemplateFind = func(e error) ServiceError {
		msg := fmt.Sprintf("template find error - %s", e.Error())
		return New(msg, CodeInternal)
	}

	ResourceNotFound = func(rSrc string, idType, id interface{}) ServiceError {
		msg := fmt.Sprintf("%s with %s [%v] not found", rSrc, idType, id)
		return New(msg, CodeNotFound)
	}

	NoResourceIdentifier = func(rSrc string, params ...string) ServiceError {
		msg := fmt.Sprintf("at least one identifying param [%s] is needed in order to fetch resource [%s]", strings.Join(params, "; "), rSrc)
		return New(msg, CodeNoIdentifier)
	}

	TooManyIdentifiers = func(rSrc string, params ...string) ServiceError {
		msg := fmt.Sprintf("only one identifying param [%s] is needed in order to fetch resource [%s]", strings.Join(params, "/"), rSrc)
		return New(msg, CodeOneIdentifier)
	}

	InputMalformed = New("input malformed - check given parameters and parameter types", CodeMalformed)

	Datastore = func(e error) ServiceError {
		msg := fmt.Sprintf("datastore error - %s", e.Error())
		return New(msg, CodeDatastore)
	}

	Internal = func(e error) ServiceError {
		msg := fmt.Sprintf("internal service error - %s", e.Error())
		return New(msg, CodeInternal)
	}

	NotImpl = New("functionality not yet implemented", CodeNotImpl)
)
